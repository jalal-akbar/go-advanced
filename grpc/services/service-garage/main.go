package servicegarage

import (
	"context"
	"grpc/common/config"
	"grpc/common/model"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.GarageListBuyer

func init() {
	localStorage = new(model.GarageListBuyer)
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (GaragesServer) Add(_ context.Context, param *model.GarageAndUserId) (*emptypb.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(model.GarageList)
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(emptypb.Empty), nil
}

func (GaragesServer) List(_ context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}

func main() {
	srv := grpc.NewServer()
	var grgSrv GaragesServer
	model.RegisterGaragesServer(srv, grgSrv)

	log.Println("Starting RPC server at", config.ServiceGaragePort)

	l, err := net.Listen("tcp", config.ServiceGaragePort)
	if err != nil {
		log.Fatalf("could not listen %s: %v", config.ServiceGaragePort, err)
	}

	log.Fatal(srv.Serve(l))
}
