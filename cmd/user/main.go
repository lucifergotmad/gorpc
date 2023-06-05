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
	log.Println("Starting user-service...")
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	HOST := os.Getenv("DB_HOST")
	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s", HOST, USER, PASSWORD, DBNAME, PORT)

	log.Printf("Connecting to DB with: %s", dsn)
	
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
