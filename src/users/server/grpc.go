package server

import (
	"net"

	log "cloud/users/logger"
	userspb "cloud/users/proto/users"
	"cloud/users/service"

	"google.golang.org/grpc"
)

type GrpcConfig struct {
	Address     string
	UsersServer *service.UsersServer
	listener    net.Listener

	server *grpc.Server
}

func NewGrpcConfig(address string, userServer *service.UsersServer) *GrpcConfig {
	return &GrpcConfig{
		Address:     address,
		UsersServer: userServer,
	}
}

func (g *GrpcConfig) RunGrpcServer() error {
	listen, err := net.Listen("tcp", g.Address)
	if err != nil {
		log.Entry.WithError(err).WithField("address", g.Address).Errorln("Failed to bind target port")
		return err
	}

	g.server = grpc.NewServer()
	userspb.RegisterUsersServer(g.server, g.UsersServer)
	g.server.Serve(listen)
	return nil
}
