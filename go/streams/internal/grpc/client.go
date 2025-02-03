package grpc

import (
	"fmt"
	"math"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
	"github.com/fraym/freym-api/go/streams/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewClient(config *config.Config) (managementpb.ServiceClient, func() error, error) {
	address := config.Address
	if len(address) == 0 {
		return nil, nil, fmt.Errorf("cannot use empty value for STREAMS_CLIENT_ADDRESS")
	}

	credentialsOptions := grpc.WithTransportCredentials(insecure.NewCredentials())
	keepaliveOptions := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                config.KeepaliveInterval,
		Timeout:             config.KeepaliveTimeout,
		PermitWithoutStream: true,
	})
	callOptions := grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32),
	)

	conn, err := grpc.NewClient(address, credentialsOptions, keepaliveOptions, callOptions)
	if err != nil {
		return nil, nil, err
	}

	return managementpb.NewServiceClient(conn), func() error {
		return conn.Close()
	}, nil
}
