package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	log "users/logger"
	userspb "users/proto/users"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = "0.0.0.0"
	grpcPort = 6666
	gatewayPort = 7777
	grpcAddress = fmt.Sprintf("%s:%d", host, grpcPort)
	gatewayAddress = fmt.Sprintf("%s:%d", host, gatewayPort)
)

type UsersServer struct {
	userspb.UnimplementedUsersServer
}

func (u *UsersServer) GetUser(ctx context.Context, req *userspb.UserRequest) (*userspb.User, error) {
	log.Entry.WithField("body", req).Infoln("Start GetUser request")
	
	return &userspb.User{
		UserId: req.UserId,
		UserName: "test_user_name",
		FirstName: "test_first_name",
		LastName: "test_last_name",
		Email: "test@test.com",
	}, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// logger config
	loggerConfig := log.LoggerConfig{
		Formatter: &logrus.JSONFormatter{},
		Data: logrus.Fields{"service": "users"},
		Level: logrus.InfoLevel,
		Output: os.Stdout,
		Caller: true,
	}
	log.New(loggerConfig)

	// gRPC server
	lis, err := net.Listen("tcp", grpcAddress)
	checkError(err)

	server := grpc.NewServer()
	userspb.RegisterUsersServer(server, &UsersServer{})
	go server.Serve(lis)
	log.Entry.WithField("address", grpcAddress).Infoln("gRPC server start")

	// gateway server
	mux := runtime.NewServeMux()
	err = userspb.RegisterUsersHandlerFromEndpoint(
		context.Background(), 
		mux,
		grpcAddress,
		[]grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)
	checkError(err)
	log.Entry.WithField("address", gatewayAddress).Infoln("gateway server start")
	http.ListenAndServe(gatewayAddress, mux)
}