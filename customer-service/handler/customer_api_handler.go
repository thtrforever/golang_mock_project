package customer_api_handler

import (
	"fmt"
	"net/http"
	core_api "thindv/golang-mock/core/api"
	customer_api_payload "thindv/golang-mock/customer-service/payload"
	"thindv/golang-mock/pb"

	"github.com/gin-gonic/gin"
)

type CustomerApiHandler struct {
	GrpcClient pb.CustomerGrpcClient
}

func InitCustomerApiHandler(grpcClient pb.CustomerGrpcClient) CustomerApiHandler {
	return CustomerApiHandler{
		GrpcClient: grpcClient,
	}
}

//create new customer
func (h *CustomerApiHandler) CreateCustomer(c *gin.Context) {
	request := customer_api_payload.CreateCustomerRequest{}

	err := c.ShouldBindJSON(&request)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}

	grpcCreateCustomerReq := &pb.GrpcCreateCustomerRequest{
		Username:    request.Username,
		Password:    request.Password,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	grpcCreateCustomerRes, err := h.GrpcClient.CreateCustomer(c, grpcCreateCustomerReq)

	if err != nil {
		c.JSON(http.StatusOK, core_api.CreateCustomApiErrorResponse(BAD_REQUEST_CODE, err.Error(), nil))
		return
	}

	response := customer_api_payload.CreateCustomerResponse{
		Status:    grpcCreateCustomerRes.Status,
		Username:  grpcCreateCustomerRes.Username,
		CreatedAt: grpcCreateCustomerRes.CreateAt.String(),
	}

	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}

//Update customer
func (h *CustomerApiHandler) UpdateCustomer(c *gin.Context) {
	fmt.Print("Update customer API")
	request := customer_api_payload.UpdateCustomerRequest{}

	err := c.ShouldBindJSON(&request)
	if nil != err {
		c.JSON(http.StatusOK, core_api.CreateApiErrorResponse(BAD_REQUEST_CODE, err.Error()))
	}

	grpcUpdateCustomerRequest := &pb.GrpcUpdateCustomerRequest{
		Username:    request.Username,
		Password:    "",
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Address:     request.Address,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
	}
	grpcUpdateCustomerRes, err := h.GrpcClient.UpdateCustomer(c, grpcUpdateCustomerRequest)

	if err != nil {
		c.JSON(http.StatusOK, core_api.CreateCustomApiErrorResponse(BAD_REQUEST_CODE, err.Error(), nil))
		return
	}

	response := customer_api_payload.UpdateCustomerResponse{
		Status: grpcUpdateCustomerRes.Status,
	}

	c.JSON(http.StatusOK, core_api.CreateSuccessResponse(response))
}
