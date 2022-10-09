package api

import (
	"fmt"
	"net/http"
	core_api "thindv/golang-mock/core/api"
	api_flight_payload "thindv/golang-mock/flight-service/payload"
	"thindv/golang-mock/pb"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FlightHandler struct {
	GrpcClient pb.FlightGrpcClient
}

func InitFlightHandler(grpcClient pb.FlightGrpcClient) FlightHandler {
	return FlightHandler{
		GrpcClient: grpcClient,
	}
}

func (h *FlightHandler) PingFlightHandler(c *gin.Context) {
	grpcPingFlightGrpcRequest := &pb.GrpcPingFlightGrpcRequest{
		Status: "",
	}

	grpcPingFlightServiceResponse, err := h.GrpcClient.PingFlightGrpc(c.Request.Context(), grpcPingFlightGrpcRequest)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateCustomApiErrorResponse(BAD_REQUEST_CODE, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(grpcPingFlightServiceResponse.GetStatus()))
}

func (f *FlightHandler) CreateFlightHandler(c *gin.Context) {
	fmt.Println("Create flight API")
	request := CreateFlightRequest{}

	err := c.ShouldBindJSON(&request)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}

	grpcCreateFlightRequest := &pb.GrpcCreateFlightRequest{
		OriginAirportCode:      request.OriginAirportCode,
		DestinationAirportCode: request.DestinationAirportCode,
		DepartureDateTime:      timestamppb.New(time.Time(request.DepartureDateTime)),
		BookingClass:           request.BookingClass,
		FlightNumber:           request.FlightNumber,
		FlightDuration:         request.FlightDuration,
		SeatRemaining:          request.SeatRemaining,
		CurrencyCode:           request.CurrencyCode,
		Price:                  request.Price,
	}

	grpcCreateFlightResponse, err := f.GrpcClient.CreateFlight(c, grpcCreateFlightRequest)

	response := api_flight_payload.CreateFlightResponse{
		Id:       grpcCreateFlightResponse.Id,
		CreateBy: grpcCreateFlightResponse.CreateBy,
		CreateAt: grpcCreateFlightResponse.CreateAt.AsTime(),
	}

	fmt.Println("Create flight API - DONE")
	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}
