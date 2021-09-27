package platform

import (
	"context"
	"fmt"
	"log"
	"net"

	googlegrpc "google.golang.org/grpc"

	"github.com/rlgino/go-grcp-example/internal/net/grpc"
	"github.com/rlgino/go-grcp-example/internal/platform/server"
)

type grpcServer struct {
	config server.Config
}

func (g *grpcServer) Serve() error {
	addr := fmt.Sprintf("%s:%s", g.config.Host, g.config.Port)
	listener, err := net.Listen(g.config.Protocol, addr)
	if err != nil {
		return err
	}

	srv := googlegrpc.NewServer()
	serviceServer := NewWishListServer()
	grpc.RegisterWishListServiceServer(srv, serviceServer)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}

func NewServer(config server.Config) server.Server {
	return &grpcServer{config: config}
}

type wishListHandler struct {
	itemList []*grpc.Item
}

func NewWishListServer() grpc.WishListServiceServer {
	return &wishListHandler{}
}

func (server *wishListHandler) Create(_ context.Context, req *grpc.CreateWishListReq) (*grpc.CreateWishListResp, error) {
	log.Println("Recive wishlist name", req.WishList.Name)
	resp := grpc.CreateWishListResp{
		WishListId: "101",
	}
	return &resp, nil
}

func (server *wishListHandler) Add(ctx context.Context, newItem *grpc.AddItemReq) (*grpc.AddItemResp, error) {
	server.itemList = append(server.itemList, newItem.Item)
	resp := grpc.AddItemResp{
		ItemId: fmt.Sprintf("%d", len(server.itemList)),
	}
	return &resp, nil
}

func (server *wishListHandler) List(_ context.Context, req *grpc.ListWishListReq) (*grpc.ListWishListResp, error) {
	resp := grpc.ListWishListResp{
		Items: server.itemList,
	}
	return &resp, nil
}
