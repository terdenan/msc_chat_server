package chat

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/terdenan/msc_chat_server/pkg/chat_v1"
)

func (s *ChatServer) SendMessage(context.Context, *desc.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
