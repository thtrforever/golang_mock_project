package customer_grpc_handler

import (
	db "thindv/golang-mock/db/sqlc"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"
)

type CustomerGrpcHandler struct {
	pb.UnimplementedCustomerGrpcServer
	config utils.GrpcConfig
	store  *db.Store
}

func InitCustomerGrpcHandler(config utils.GrpcConfig, store *db.Store) (*CustomerGrpcHandler, error) {
	return &CustomerGrpcHandler{
		config: config,
		store:  store,
	}, nil
}
