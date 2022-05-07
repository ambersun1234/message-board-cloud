package server

import (
	"context"

	log "cloud/users/logger"
	userspb "cloud/users/proto/users"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayConfig struct {
	Address   string
	ServerMux *runtime.ServeMux
}

func NewGatewayConfig(address string) (*GatewayConfig, error) {
	/* we don't need reverse proxy of grpc-gateway(we use gin as web server)
	 * but still need it's translate of HTTP/2 to HTTP/1.1
	 */
	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	mux := runtime.NewServeMux()
	if err := userspb.RegisterUsersHandlerFromEndpoint(
		context.Background(), mux,
		address, dialOptions,
	); err != nil {
		log.Entry.WithError(err).Errorln("Error register users handler")
		return nil, err
	}

	config := &GatewayConfig{
		Address:   address,
		ServerMux: mux,
	}
	return config, nil
}
