package main

import (
	"fmt"
	"log"
	flight_api_handler "thindv/golang-mock/flight-service/api"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start flight-service")
	config, err := utils.LoadApiConfig("./flight-service/")
	if err != nil {
		fmt.Println("Cannot load config")
		log.Fatal("Cannot load config with path:", err)
	}

	//config grpc
	grpcClient, err := grpc.Dial(fmt.Sprintf("%v:%v", config.Grpc.FlightGrpc.Host, config.Grpc.FlightGrpc.Port), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	grpcFlightClient := pb.NewFlightGrpcClient(grpcClient)
	flightApiHandler := flight_api_handler.InitFlightHandler(grpcFlightClient)

	//config API router
	engine := gin.Default()
	routerGroup := engine.Group("/api/flight")

	routerGroup.GET("/ping", flightApiHandler.PingFlightHandler)
	routerGroup.POST("/create-flight", flightApiHandler.CreateFlightHandler)
	routerGroup.PUT("/update-flight", flightApiHandler.UpdateFlightHandler)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", config.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
