package constants

type ResponseType string

const (
	ResponseOK              ResponseType = "OK"
	ResponseBad             ResponseType = "BAD"
	ResponseError           ResponseType = "ERROR"
	ResponseConflict        ResponseType = "CONFLICT"
	ResponseForbidden       ResponseType = "FORBIDDEN"
	ResponseUnauthorized    ResponseType = "UNAUTHORIZED"
	ResponseTooManyRequests ResponseType = "TOO_MANY_REQUESTS"
)
type Status string

const (
	Unread    Status = "unread"
	Read      Status = "read"
	Sent      Status = "sent"
	Delivered Status = "delivered"
	Received  Status = "received"
)
const (
	True  string = "true"
	False string = "false"
)
