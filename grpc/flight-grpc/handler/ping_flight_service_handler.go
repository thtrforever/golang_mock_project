package flight_grpc_handler

import (
	"context"
	"thindv/golang-mock/pb"
)

func (h *FlightGrpcHandler) PingFlightGrpc(
	c context.Context,
	in *pb.GrpcPingFlightGrpcRequest) (*pb.GrpcPingFlightGrpcResponse, error) {
	out := &pb.GrpcPingFlightGrpcResponse{
		Status: "OK",
	}
	return out, nil
}
