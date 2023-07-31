package main

import (
	"context"
	"encoding/json"
	"fmt"
	"latihan_grpc/common/proto"
	"log"

	"google.golang.org/grpc"
)

func serviceUser() proto.UsersClient {
	port := ":9090"
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	return proto.NewUsersClient(conn)
}

func main() {

	user := serviceUser()
	res, err := user.ShowAll(context.Background(), &proto.User{})
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("sebelum di parse")
	fmt.Println(res)
	fmt.Println("sesudah di parse")
	b, _ := json.Marshal(res.Data)
	fmt.Println(string(b))
}
