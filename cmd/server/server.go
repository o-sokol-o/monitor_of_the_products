package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/o-sokol-o/as/internal/config"
	"github.com/o-sokol-o/as/internal/repository"
	"github.com/o-sokol-o/as/internal/service"
	"github.com/o-sokol-o/as/internal/transport/grpcserver"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
	})
	opts.ApplyURI(cfg.DB.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	db := dbClient.Database(cfg.DB.Database)

	repo := repository.NewRepository(db)
	service := service.NewServices(repo)
	server := grpcserver.NewGrpcServer(service)

	fmt.Println("SERVER STARTED", time.Now())

	if err := server.Run(cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
