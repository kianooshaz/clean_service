package errors

import (
	"encoding/json"
	"github.com/kianooshaz/clean_service/core/contract/interfaces"
	"net/http"
)

type serviceError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func (e *serviceError) GetMessage() string {
	return e.Message
}

func (e *serviceError) GetStatus() int {
	return e.Status
}

func (e *serviceError) GetError() string {
	return e.Error
}

func (e *serviceError) GetCauses() []interface{} {
	return e.Causes
}

func (e *serviceError) AppendCause(causes interface{}) interfaces.IServiceError {
	e.Causes = append(e.Causes, causes)
	return e
}

func NewServiceError(message string, status int, error string, causes []interface{}) interfaces.IServiceError {
	return &serviceError{
		Message: message,
		Status:  status,
		Error:   error,
		Causes:  causes,
	}
}

func NewServiceErrorFromByte(bytes []byte) (error interfaces.IServiceError, ok bool) {
	if err := json.Unmarshal(bytes, &error); err != nil {
		return nil, false
	}
	return error, true
}

func NewBadRequestError(message string) interfaces.IServiceError {
	return &serviceError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFound(message string) interfaces.IServiceError {
	return &serviceError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedError(message string) interfaces.IServiceError {
	return &serviceError{
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewInternalServerError(message string, err error) interfaces.IServiceError {
	result := &serviceError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.AppendCause(err.Error())
	}
	return result
}
