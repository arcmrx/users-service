package main

import (
	"log"

	"github.com/arcmrx/users-service/internal/database"
	"github.com/arcmrx/users-service/internal/user"
	"github.com/arcmrx/users-service/internal/transport/grpc"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repository := user.NewRepository(db)
	service := user.NewService(repository)

	if err := grpc.RunGRPC(service); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}