//go:build tools

package tools

import (
	_ "github.com/gin-gonic/gin"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/sirupsen/logrus"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
