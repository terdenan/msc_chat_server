package chat

import (
	desc "github.com/terdenan/msc_chat_server/pkg/chat_v1"
)

type ChatServer struct {
	desc.UnimplementedChatV1Server
}

func NewChatServer() *ChatServer {
	return &ChatServer{}
}
