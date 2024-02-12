package app

import (
	"log"

	"github.com/terdenan/msc_chat_server/internal/api/chat"
	"github.com/terdenan/msc_chat_server/internal/config"
)

type ServiceProvider struct {
	grpcConfig config.GRPCConfig
	chatServer *chat.ChatServer
}

func NewServiceProvider() ServiceProvider {
	return ServiceProvider{}
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *ServiceProvider) ChatServer() *chat.ChatServer {
	if s.chatServer == nil {
		s.chatServer = chat.NewChatServer()
	}

	return s.chatServer
}
