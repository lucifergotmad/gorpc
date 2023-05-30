package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "gorpc/api/user"
	"gorpc/pkg/user"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database: ", err)
		return
	}
	
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
