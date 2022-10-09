package flight_grpc_handler

import (
	db "thindv/golang-mock/db/sqlc"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"
)

type FlightGrpcHandler struct {
	pb.UnimplementedFlightGrpcServer
	config utils.GrpcConfig
	store  *db.Store
}

func InitFlightGrpcHandler(config utils.GrpcConfig, store *db.Store) (*FlightGrpcHandler, error) {
	return &FlightGrpcHandler{
		config: config,
		store:  store,
	}, nil
}
