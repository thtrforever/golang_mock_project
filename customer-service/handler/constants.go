package customer_api_handler

import core_api "thindv/golang-mock/core/api"

var (
	BAD_REQUEST_CODE      = core_api.InitApiError(400, "Bad request")
	BAD_REQUEST_ERROR     = core_api.InitApiError(400, "Bad request")
	INTERNAL_SERVER_ERROR = core_api.InitApiError(500, "Server error")
)
