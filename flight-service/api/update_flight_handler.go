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

func (h *FlightHandler) UpdateFlightHandler(c *gin.Context) {
	fmt.Println("Update flight API")
	request := api_flight_payload.UpdateFlightRequest{}
	err := c.ShouldBindJSON(&request)

	if nil != err {
		fmt.Printf("parsing request exception %v", err)
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(
			BAD_REQUEST_CODE, err.Error()))
		return
	}

	grpcUpdateFlightRequest := &pb.GrpcUpdateFlightRequest{
		Id:                     request.Identifier.Id,
		OriginAirportCode:      request.OriginAirportCode,
		DestinationAirportCode: request.DestinationAirportCode,
		DepartureDateTime:      timestamppb.New(time.Time(request.DepartureDateTime)),
		BookingClass:           request.BookingClass,
		FlightNumber:           request.FlightNumber,
		FlightDuration:         request.FlightDuration,
		SeatRemaining:          request.SeatRemaining,
		CurrencyCode:           request.CurrencyCode,
		Price:                  request.Price,
		Enabled:                request.Enabled,
	}

	//call to grpc service
	grpcUpdateFlightResponse, err := h.GrpcClient.UpdateFlight(c.Request.Context(), grpcUpdateFlightRequest)
	if err != nil {
		c.JSON(http.StatusOK, core_api.CreateCustomApiErrorResponse(BAD_REQUEST_CODE, err.Error(), nil))
		return
	}

	response := api_flight_payload.UpdateFlightResponse{
		RequestData: request,
		UpdateBy:    grpcUpdateFlightResponse.UpdateBy,
		UpdateAt:    grpcUpdateFlightResponse.UpdateAt.AsTime(),
	}
	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}
