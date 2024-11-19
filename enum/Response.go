package enum

type ResponseCode int

const (
	Success       ResponseCode = 200
	Created       ResponseCode = 201
	Accepted      ResponseCode = 202
	Deleted       ResponseCode = 204
	Fail          ResponseCode = 400
	Unauthorized  ResponseCode = 401
	Forbidden     ResponseCode = 403
	NotFound      ResponseCode = 404
	Validate      ResponseCode = 422
	ManyRequest   ResponseCode = 429
	Error         ResponseCode = 500
)