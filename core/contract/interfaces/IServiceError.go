package interfaces

type IServiceError interface {
	GetMessage() string
	GetStatus() int
	GetError() string
	GetCauses() []interface{}
	AppendCause(causes interface{}) IServiceError
}
