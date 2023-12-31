package main

import (
	"context"
	"fmt"
	"github.com/bulutcan99/grpc_weather/api/grpc_server"
	"github.com/bulutcan99/grpc_weather/internal/query"
	config_mongodb "github.com/bulutcan99/grpc_weather/pkg/config/mongodb"
	config_rabbitmq "github.com/bulutcan99/grpc_weather/pkg/config/rabbitmq"
	config_redis "github.com/bulutcan99/grpc_weather/pkg/config/redis"
	"github.com/bulutcan99/grpc_weather/pkg/env"
	"github.com/bulutcan99/grpc_weather/pkg/logger"
	"github.com/bulutcan99/grpc_weather/proto/pb"
	"github.com/bulutcan99/grpc_weather/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	grpcServer *grpc.Server
	ServerPort int
	Services   *service.Services
	Mongo      *config_mongodb.Mongo
	Redis      *config_redis.Redis
	RabbitMQ   *config_rabbitmq.RabbitMQ
	Logger     *zap.Logger
	Env        *env.ENV
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	Logger.Info("gRPC method called: " + info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	Logger.Info("gRPC stream method called: " + info.FullMethod)
	return handler(srv, ss)
}

func Init() {
	Env = env.ParseEnv()
	ServerPort = Env.ServerPort
	Logger = logger.InitLogger(Env.LogLevel)
	Mongo = config_mongodb.NewConnetion()
	Redis = config_redis.NewRedisConnection()
	RabbitMQ = config_rabbitmq.NewRabbitMQConnection()
	grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	reflection.Register(grpcServer)
	userRepo := query.NewUserRepositry(Mongo, Env.UserCollection)
	weatherRepo := query.NewWeatherRepository(Mongo, Env.WeatherCollection)

	userService := service.NewUserService(userRepo)
	weatherService := service.NewWeatherService(weatherRepo, userRepo)

	Services = service.RegisterServices(userService, weatherService)
	weatherServer := grpc_server.NewWeatherServer(Services)
	pb.RegisterUserServiceServer(grpcServer, weatherServer)
	pb.RegisterWeatherServiceServer(grpcServer, weatherServer)
}

func main() {
	Init()
	defer Logger.Sync()
	defer Mongo.Close()
	defer Redis.Close()
	defer RabbitMQ.Close()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	zap.S().Info("gRPC Weather Server starting on: ", Env.ServerPort)
	grpcPort := fmt.Sprintf(":%d", ServerPort)
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		zap.S().Fatalf("Failed to listen: %v", err)
	}

	go func() {
		zap.S().Info("GRPC Server starting...")
		if err := grpcServer.Serve(lis); err != nil {
			zap.S().Fatalf("Failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	zap.S().Info("Signal received, shutting down")
	stop()
	grpcServer.Stop()
	zap.S().Info("Shutting down server...")
	zap.S().Info("Server gracefully stopped")
}
