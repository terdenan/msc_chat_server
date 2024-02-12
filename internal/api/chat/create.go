package chat

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"

	desc "github.com/terdenan/msc_chat_server/pkg/chat_v1"
)

func (s *ChatServer) Create(context.Context, *desc.CreateChatRequest) (*desc.CreateChatResponse, error) {
	return &desc.CreateChatResponse{
		Id: gofakeit.Int64(),
	}, nil
}
