package serviceuser

import (
	"context"
	"grpc/common/config"
	"grpc/common/model"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (UsersServer) Register(_ context.Context, param *model.User) (*emptypb.Empty, error) {
	user := param

	localStorage.List = append(localStorage.List, user)

	log.Println("Register User", user.String())
	return new(emptypb.Empty), nil
}

func (UsersServer) List(context.Context, *emptypb.Empty) (*model.UserList, error) {
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	var usrSrv UsersServer
	model.RegisterUsersServer(srv, usrSrv)

	log.Println("Starting RPC server at", config.ServiceUserPort)

	l, err := net.Listen("tcp", config.ServiceUserPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServiceUserPort, err)
	}
	log.Fatal(srv.Serve(l))
}
