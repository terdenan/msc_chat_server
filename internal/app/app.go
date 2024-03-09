package app

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/terdenan/msc_chat_server/internal/config"
	desc "github.com/terdenan/msc_chat_server/pkg/chat_v1"
)

type App struct {
	grpcServer      *grpc.Server
	serviceProvider ServiceProvider
}

func NewApp() (*App, error) {
	a := &App{}

	err := a.initDeps()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runGRPCServer()
}

func (a *App) initDeps() error {
	inits := []func() error{
		a.initEnv,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, init := range inits {
		err := init()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initEnv() error {
	config.Load(".env")
	return nil
}

func (a *App) initServiceProvider() error {
	a.serviceProvider = NewServiceProvider()
	return nil
}

func (a *App) initGRPCServer() error {
	a.grpcServer = grpc.NewServer()

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatServer())

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	return a.grpcServer.Serve(lis)
}
