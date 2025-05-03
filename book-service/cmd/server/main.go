package main

import (
	"log"
	"net"

	pb "github.com/Prrost/protoFinalAP2/book-service/book"
	"github.com/Prrost/protoFinalAP2/book-service/internal/config"
	"github.com/Prrost/protoFinalAP2/book-service/internal/model"
	"github.com/Prrost/protoFinalAP2/book-service/internal/repository"
	"github.com/Prrost/protoFinalAP2/book-service/internal/service"

	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// миграция схемы
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}

	repo := repository.NewBookRepo(db)
	srv := service.NewBookServiceServer(repo)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBookServiceServer(grpcServer, srv)

	log.Println("BookService gRPC listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
