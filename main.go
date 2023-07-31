package main

import (
	"latihan_grpc/common/proto"
	"latihan_grpc/db"
	"latihan_grpc/modules/repository"
	"latihan_grpc/server"
	"log"
	"net"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	DB := db.NewDBPostges()
	repoUser := repository.NewRepoUserImpl(DB)
	userHandler := server.NewHandlerUser(repoUser)

	srv := grpc.NewServer()

	proto.RegisterUsersServer(srv, userHandler)

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("server failed to run at port 9090 ")
	}
	err = srv.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
