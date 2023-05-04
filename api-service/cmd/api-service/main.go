package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/HeadGardener/api-service/configs"
	"github.com/HeadGardener/api-service/internal/app/handlers"
	"github.com/HeadGardener/api-service/internal/app/handlers/auth"
	user_service "github.com/HeadGardener/api-service/internal/app/services/user"
	"github.com/HeadGardener/api-service/internal/pkg/server"
	"go.uber.org/zap"
	"log"
	"os/signal"
	"syscall"
	"time"
)

var confPath = flag.String("conf-path", "./configs/.env", "path to config env")

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(fmt.Sprintf("error whilr creating logger: %s", err.Error()))
	}

	logger.Info("initializing router...")
	router := handlers.NewRouter()

	serviceConfig, err := configs.NewServiceConfig(*confPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to read config file, error: %s", err.Error()))
	}

	logger.Info("initializing user service...")
	userService := user_service.NewUserService(serviceConfig.UserServiceURL, "users")
	authHandler := auth.NewAuthHandler(userService)
	authHandler.InitRoutes(router)

	serverConfig, err := configs.NewServerConfig(*confPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("unable to read config file, error: %s", err.Error()))
	}

	srv := &server.Server{}
	go func() {
		if err := srv.Run(serverConfig.Port, router); err != nil {
			logger.Error(fmt.Sprintf("error occurring while running server, err:%s", err.Error()))
		}
	}()

	logger.Info("server start working")
	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(fmt.Sprintf("server forced to shutdown: %s", err.Error()))
	}

	logger.Info("server exiting")
}
