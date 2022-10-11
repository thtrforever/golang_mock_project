package booking_grpc_handler

import "context"

type BookingGrpcHandler struct {
}

func InitBookingGrpcHandler(c context.Context) (*BookingGrpcHandler, error) {
	return &BookingGrpcHandler{}, nil
}
