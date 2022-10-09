package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	db "thindv/golang-mock/db/sqlc"
	flight_grpc_handler "thindv/golang-mock/grpc/flight-grpc/handler"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Start flight-service-grpc")
	config, err := utils.LoadGrpcConfig("./grpc/flight-grpc/")

	if nil != err {
		log.Fatal("Cannot load config ", err)
		return
	}

	//init database connection
	conn, err := sql.Open(config.Database.DriverName, utils.GetDatabaseSourceNameForGrpc(config.Database))
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}
	store := db.InitalStore(conn)

	//init GRPC service
	grpcServer := grpc.NewServer()
	grpcHandler, err := flight_grpc_handler.InitFlightGrpcHandler(config, store)
	if err != nil {
		log.Fatal("Can not initial grpc service:", err)
	}

	reflection.Register(grpcServer)
	pb.RegisterFlightGrpcServer(grpcServer, grpcHandler)

	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", config.Server.Port))
	if err != nil {
		panic(err)
	}
	log.Printf("Listening at port: %v\n", config.Server.Port)

	grpcServer.Serve(listen)
}
