package testutils

import (
	"context"
	"log"
	"net"

	"github.com/bradenrayhorn/ledger-core/internal/mocks"
	"github.com/bradenrayhorn/ledger-protos/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type MockGrpcServer struct {
	SessionService *mocks.SessionAuthenticatorServer

	Conn     *grpc.ClientConn
	listener *bufconn.Listener
	server   *grpc.Server
}

func (s *MockGrpcServer) Close() {
	s.listener.Close()
	s.server.Stop()
}

func SetupGrpcConn() *MockGrpcServer {
	mockServer := &MockGrpcServer{
		SessionService: new(mocks.SessionAuthenticatorServer),

		listener: bufconn.Listen(1024 * 1024),
		server:   grpc.NewServer(),
	}

	session.RegisterSessionAuthenticatorServer(mockServer.server, mockServer.SessionService)

	go func() {
		if err := mockServer.server.Serve(mockServer.listener); err != nil {
			log.Fatalf("gRPC server exited: %v", err)
		}
	}()

	conn, _ := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return mockServer.listener.Dial() }), grpc.WithInsecure())
	mockServer.Conn = conn

	return mockServer
}
