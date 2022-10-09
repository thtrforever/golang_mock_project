package customer_grpc_handler

import (
	"context"
	"database/sql"
	"fmt"
	db "thindv/golang-mock/db/sqlc"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *CustomerGrpcHandler) CreateCustomer(c context.Context, in *pb.GrpcCreateCustomerRequest) (*pb.GrpcCreateCustomerResponse, error) {
	fmt.Println("Grpc Create customer")

	//validation data
	if utils.IsEmptyString(in.Username) {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Username is mandantory field"),
		)
	}

	//check existed customer
	existedCustomer, err := h.store.Queries.GetCustomerInfoByUsername(c, in.GetUsername())
	if nil != err {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Cannot look up existed customer| %v", err))
	}

	if len(existedCustomer) > 0 {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Username is existed"),
		)
	}

	//Insert new customer
	isActive := true
	createdAt := time.Now()
	createCustomerParam := db.CreateCustomerParams{
		Username:    in.Username,
		Password:    sql.NullString{String: in.Password, Valid: true},
		FirstName:   sql.NullString{String: in.FirstName, Valid: true},
		LastName:    sql.NullString{String: in.LastName, Valid: true},
		Email:       sql.NullString{String: in.Email, Valid: true},
		PhoneNumber: sql.NullString{String: in.PhoneNumber, Valid: true},
		Address:     sql.NullString{String: in.Address, Valid: true},
		Active:      sql.NullBool{Bool: isActive, Valid: true},
		CreateAt:    sql.NullTime{Time: createdAt, Valid: true},
	}

	response, err := h.store.Queries.CreateCustomer(c, createCustomerParam)

	return &pb.GrpcCreateCustomerResponse{
		Status:   "OK",
		Username: response.Username,
		Active:   response.Active.Bool,
		CreateAt: timestamppb.New(createdAt),
	}, nil
}
