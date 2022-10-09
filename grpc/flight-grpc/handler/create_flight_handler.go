package flight_grpc_handler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	db "thindv/golang-mock/db/sqlc"
	"thindv/golang-mock/pb"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *FlightGrpcHandler) CreateFlight(ctx context.Context, in *pb.GrpcCreateFlightRequest) (*pb.GrpcCreateFlightResponse, error) {
	log.Fatal("Create Flight by Grpc handler")

	//searching existed Flight
	searchFlightQueryData := db.GetFlightsByFieldsParams{
		Offset:                 0,
		Limit:                  10,
		OriginAirportCode:      in.OriginAirportCode,
		DestinationAirportCode: in.DestinationAirportCode,
		DepartureDateTime:      time.Time(in.DepartureDateTime.AsTime()),
		BookingClass:           in.BookingClass,
		FlightNumber:           in.FlightNumber,
	}

	existedFlight, err := h.store.Queries.GetFlightsByFields(ctx, searchFlightQueryData)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Server internal error|%v", err))
	}

	if len(existedFlight) > 0 {
		return nil, status.Error(codes.Internal, "This flight has existed")
	}

	//insert new flight
	createBy := "thindv"
	createdDate := time.Now()
	createFlightData := db.CreateFlightParams{
		OriginAirportCode:      strings.ToUpper(in.OriginAirportCode),
		DestinationAirportCode: strings.ToUpper(in.DestinationAirportCode),
		DepartureDateTime:      time.Time(in.DepartureDateTime.AsTime()),
		BookingClass:           strings.ToUpper(in.BookingClass),
		FlightNumber:           strings.ToUpper(in.FlightNumber),
		FlightDuration:         in.FlightDuration,
		Enabled:                true,
		SeatAvailable:          in.SeatRemaining,
		CurrencyCode:           in.CurrencyCode,
		Price:                  in.Price,
		CreateBy:               sql.NullString{String: createBy, Valid: true},
		CreateDate:             sql.NullTime{Time: createdDate, Valid: true},
	}

	flight, err := h.store.Queries.CreateFlight(ctx, createFlightData)
	if nil != err {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Create flight error|%v", err))
	}

	outData := &pb.GrpcCreateFlightResponse{
		Id:       flight.ID,
		CreateBy: createBy,
		CreateAt: timestamppb.New(createdDate),
	}
	return outData, nil
}
