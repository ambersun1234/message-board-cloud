package main

import (
	"fmt"
	"os"

	log "cloud/users/logger"
	"cloud/users/server"
	"cloud/users/service/users"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	local         = "127.0.0.1"
	host          = "0.0.0.0"
	grpcPort      = 6666
	serverPort    = 7777
	grpcAddress   = fmt.Sprintf("%s:%d", host, grpcPort)
	serverAddress = fmt.Sprintf("%s:%d", host, serverPort)
)

func main() {
	// logger
	config := log.LoggerConfig{
		Formatter: &logrus.JSONFormatter{},
		Data:      logrus.Fields{"service": "users"},
		Level:     logrus.DebugLevel,
		Output:    os.Stdout,
		Caller:    true,
	}
	log.New(config)

	// gRPC server
	grpcConfig := server.NewGrpcConfig(grpcAddress, &users.UsersServer{})
	go grpcConfig.RunGrpcServer()
	log.Entry.Infoln("gRPC server listening on", grpcAddress)

	// gateway
	gatewayConfig, err := server.NewGatewayConfig(grpcAddress)
	if err != nil {
		panic(err)
	}

	// Gin server
	routingTable := map[string]string{
		"GET": "/api/users/:user_id",
	}
	ginConfig := server.NewGinConfig(
		serverAddress,
		log.GinLogger(),
		gin.Recovery(),
		routingTable,
		gin.WrapH(gatewayConfig.ServerMux),
	)
	ginConfig.RunGinServer()
}
