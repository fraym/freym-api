package management

import (
	"fmt"

	"github.com/fraym/freym-api/go/migrations/config"
	"github.com/fraym/freym-api/go/proto/migrations/managementpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewClient(config *config.Config) (managementpb.ServiceClient, func() error, error) {
	address := config.Address
	if len(address) == 0 {
		return nil, nil, fmt.Errorf("cannot use empty value for MIGRATIONS_CLIENT_ADDRESS")
	}

	credentialsOptions := grpc.WithTransportCredentials(insecure.NewCredentials())
	keepaliveOptions := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                config.KeepaliveInterval,
		Timeout:             config.KeepaliveTimeout,
		PermitWithoutStream: true,
	})

	conn, err := grpc.NewClient(address, credentialsOptions, keepaliveOptions)
	if err != nil {
		return nil, nil, err
	}

	return managementpb.NewServiceClient(conn), func() error {
		return conn.Close()
	}, nil
}
