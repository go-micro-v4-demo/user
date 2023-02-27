package main

import (
	"github.com/go-micro-v4-demo/user/handler"
	pb "github.com/go-micro-v4-demo/user/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	regs "go-micro.dev/v4/registry"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	// Create service
	reg := regs.NewMemoryRegistry() //没存
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address("0.0.0.0:8080"),
		micro.Registry(reg),
	)

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), new(handler.User)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
