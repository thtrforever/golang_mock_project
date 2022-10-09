package customer_grpc_handler

import (
	"context"
	"database/sql"
	"fmt"
	db "thindv/golang-mock/db/sqlc"
	"thindv/golang-mock/pb"
	"thindv/golang-mock/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CustomerGrpcHandler) UpdateCustomer(c context.Context, in *pb.GrpcUpdateCustomerRequest) (*pb.GrpcUpdateCustomerResponse, error) {
	fmt.Println("Grpc update customer")

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

	if !(len(existedCustomer) > 0) {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Customer not found"),
		)
	}

	//Update new customer
	existedCustomerItem := existedCustomer[0]

	updateCustomerParam := db.UpdateCustomerParams{
		ID:          existedCustomerItem.ID,
		Username:    existedCustomerItem.Username,
		FirstName:   sql.NullString{String: in.FirstName, Valid: true},
		LastName:    sql.NullString{String: in.LastName, Valid: true},
		Email:       sql.NullString{String: in.Email, Valid: true},
		PhoneNumber: sql.NullString{String: in.PhoneNumber, Valid: true},
		Address:     sql.NullString{String: in.Address, Valid: true},
	}

	response, err := h.store.Queries.UpdateCustomer(c, updateCustomerParam)
	fmt.Printf("Updated for customer %v", response.Username)

	return &pb.GrpcUpdateCustomerResponse{
		Status: "OK",
	}, nil
}
