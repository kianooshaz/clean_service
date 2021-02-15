package errors

import (
	"encoding/json"
	"github.com/kianooshaz/clean_service/contract"
	"net/http"
)

type Section string

type serviceError struct {
	Section Section       `json:"section"`
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func (e *serviceError) GetSection() string {
	return string(e.Section)
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

func (e *serviceError) AppendCause(causes interface{}) contract.IServiceError {
	e.Causes = append(e.Causes, causes)
	return e
}

func (e *serviceError) IsEqual(err contract.IServiceError) bool {
	if e.GetStatus() != err.GetStatus() {
		return false
	} else if e.GetMessage() != err.GetMessage() {
		return false
	} else if e.GetError() != err.GetError() {
		return false
	} else if e.GetSection() != err.GetSection() {
		return false
	}
	return true
}

func NewServiceError(section Section, message string, status int, error string, causes []interface{}) contract.IServiceError {
	return &serviceError{
		Section: section,
		Message: message,
		Status:  status,
		Error:   error,
		Causes:  causes,
	}
}

func NewServiceErrorFromByte(bytes []byte) (error contract.IServiceError, ok bool) {
	if err := json.Unmarshal(bytes, &error); err != nil {
		return nil, false
	}
	return error, true
}

func NewBadRequestError(section Section, message string) contract.IServiceError {
	return &serviceError{
		Section: section,
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFound(section Section, message string) contract.IServiceError {
	return &serviceError{
		Section: section,
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewUnauthorizedError(section Section, message string) contract.IServiceError {
	return &serviceError{
		Section: section,
		Message: message,
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

func NewInternalServerError(section Section, message string, err error) contract.IServiceError {
	result := &serviceError{
		Section: section,
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		result.AppendCause(err.Error())
	}
	return result
}
