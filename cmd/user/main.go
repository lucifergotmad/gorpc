package main

import (
	"database/sql"
	"log"
	"net"

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

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=myproject sslmode=disable")
	if err != nil {
		log.Fatal("Failed to open database: ", err)
		return
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(*userRepository)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, user.NewHandler(userService))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
