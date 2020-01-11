package errors

/*
	定义响应错误
*/
type ResponseError struct {
	Code       int
	Message    string
	StatusCode int
	ERR        error
}
