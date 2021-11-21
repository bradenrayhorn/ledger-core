package core

import "google.golang.org/grpc"

type Config struct {
	HttpPort string

	PgHost       string
	PgPort       string
	PgUsername   string
	PgPassword   string
	PgDatabase   string
	PgParameters string

	GrpcConn *grpc.ClientConn
}
