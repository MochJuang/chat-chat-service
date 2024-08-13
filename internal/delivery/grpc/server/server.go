package server

import (
	"chat-service/internal/config"
	grpcservice "chat-service/internal/delivery/grpc"
	"chat-service/internal/repository/postgresql"
	"chat-service/internal/service"
	"github.com/MochJuang/chat-grpc/service/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func SetupGrpc(cfg config.Config) {
	// repository
	messageRepo := postgresql.NewMessageRepository(cfg.DB)
	conversationRepo := postgresql.NewConversationRepository(cfg.DB)

	// Services
	messageService := service.NewMessageService(messageRepo, conversationRepo)
	conversationService := service.NewConversationService(conversationRepo)

	grpcServer := grpc.NewServer()
	chatService := grpcservice.NewChatService(conversationService, messageService)
	chat.RegisterChatServiceServer(grpcServer, chatService)

	// Start gRPC server
	lis, err := net.Listen("tcp", cfg.GrpcServer)
	if err != nil {
		log.Fatalf("Failed to listen on: %s %v", cfg.GrpcServer, err)
	}
	log.Printf("gRPC server listening on %s", cfg.GrpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
