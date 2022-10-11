package booking_api_handler

import core_api "thindv/golang-mock/core/api"

type BookingApiHandler struct {
}

func InitBookingApiHandler() BookingApiHandler {
	return BookingApiHandler{}
}

var (
	BAD_REQUEST_ERROR    = core_api.InitApiError(400, "Bad Request")
	BUSINESS_LOGIC_ERROR = core_api.InitApiError(500, "Business Logic Error")
)
