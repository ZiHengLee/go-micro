package main

import (
	"cartApi/handler"
	pb "cartApi/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("cartapi"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterCartApiHandler(srv.Server(), new(handler.CartApi))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
