package server

import (
	log "cloud/users/logger"

	"github.com/gin-gonic/gin"
)

type GinConfig struct {
	Address  string
	Logger   gin.HandlerFunc
	Recovery gin.HandlerFunc
	Routing  map[string]string
	Handler  gin.HandlerFunc

	server *gin.Engine
}

func NewGinConfig(address string, logger gin.HandlerFunc, recovery gin.HandlerFunc, routingTable map[string]string, handler gin.HandlerFunc) *GinConfig {
	return &GinConfig{
		Address:  address,
		Logger:   logger,
		Recovery: recovery,
		Routing:  routingTable,
		Handler:  handler,
	}
}

func (g *GinConfig) RunGinServer() error {
	gin.SetMode(gin.ReleaseMode)

	g.server = gin.New()
	g.server.Use(g.Logger)
	g.server.Use(g.Recovery)

	for method, route := range g.Routing {
		g.server.Handle(method, route, g.Handler)
	}

	log.Entry.Infoln("Gin server listening on", g.Address)
	if err := g.server.Run(g.Address); err != nil {
		log.Entry.WithError(err).Errorln("Error to start server")
		return err
	}

	return nil
}
