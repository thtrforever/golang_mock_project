package main

import (
	"fmt"
	"log"
	customer_api_handler "thindv/golang-mock/customer-service/handler"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Start module customer service API")
	config, err := utils.LoadApiConfig("./customer-service/")
	if err != nil {
		log.Fatal("Cannot load config with path:", err)
	}

	//config grpc
	grpcClient, err := grpc.Dial(fmt.Sprintf("%v:%v", config.Grpc.CustomerGrpc.Host, config.Grpc.CustomerGrpc.Port), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Cannot create grpc client %v", err)
		panic(err)
	}

	grpcCustomerClient := pb.NewCustomerGrpcClient(grpcClient)
	customerApiHandler := customer_api_handler.InitCustomerApiHandler(grpcCustomerClient)

	//config API router
	engine := gin.Default()
	routerGroup := engine.Group("/api/customer")

	routerGroup.POST("/create-customer", customerApiHandler.CreateCustomer)
	routerGroup.PUT("/update-customer", customerApiHandler.UpdateCustomer)

	err = engine.Run(fmt.Sprintf("127.0.0.1:%v", config.Server.Port))
	if err != nil {
		log.Fatal("Can not start server:", err)
	}
}
