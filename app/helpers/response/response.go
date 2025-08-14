package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	common "event-reporting/app/utils/common/src"
	constants "event-reporting/app/utils/constants"
)

func SendSuccessResponse[T any](c *gin.Context, data T) {
	common.CreateJSONResponse(c, constants.ResponseOK, http.StatusOK, "", data)
}

func SendBadRequestResponse[T comparable](c *gin.Context, message T) {
	common.CreateJSONResponse(c, constants.ResponseBad, http.StatusBadRequest, fmt.Sprintf("%v", message), nil)
}

func SendErrorResponse[T comparable](c *gin.Context, statusCode int, message T) {
	common.CreateJSONResponse(c, constants.ResponseError, statusCode, fmt.Sprintf("%v", message), nil)
}

func SendConflictResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseConflict, http.StatusConflict, message, nil)
}

func SendForbiddenResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseForbidden, http.StatusForbidden, message, nil)
}

func SendUnauthorizedResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseUnauthorized, http.StatusUnauthorized, message, nil)
}

func SendTooManyRequestsResponse(c *gin.Context, message string) {
	common.CreateJSONResponse(c, constants.ResponseTooManyRequests, http.StatusTooManyRequests, message, nil)
}
