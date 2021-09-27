package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rlgino/go-grcp-example/internal/net/grpc"
	googlegrpc "google.golang.org/grpc"
)

func main() {
	addr := fmt.Sprintf("%s:%s", "localhost", "3333")
	conn, err := googlegrpc.Dial(addr, googlegrpc.WithInsecure())
	if err != nil {
		log.Fatalf("impossible connect: %v", err)
	}
	client := grpc.NewWishListServiceClient(conn)

	// here we can start using the client:
	name := "pepe"
	w := &grpc.WishList{
		Name:   name,
		Status: grpc.WishList_ACTIVE,
	}

	ctx := context.Background()
	res, err := client.Create(ctx, &grpc.CreateWishListReq{WishList: w})
	fmt.Println(res)
	fmt.Println(err)

	item := &grpc.AddItemReq{
		Item: &grpc.Item{
			WishListId: res.WishListId,
			Id: "item-1",
			Price: 20.0,
		},
	}
	resAddItem, err := client.Add(ctx, item)
	fmt.Println(resAddItem)
	fmt.Println(err)
	item2 := &grpc.AddItemReq{
		Item: &grpc.Item{
			WishListId: res.WishListId,
			Id: "item-2",
			Price: 20.0,
		},
	}
	resAddItem2, err := client.Add(ctx, item2)
	fmt.Println(resAddItem2)
	fmt.Println(err)

	list, err := client.List(ctx, &grpc.ListWishListReq{WishListId: "pepe"})
	fmt.Println(list)
	fmt.Println(err)
}
