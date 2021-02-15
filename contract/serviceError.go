package contract

type IServiceError interface {
	GetMessage() string
	GetStatus() int
	GetError() string
	GetSection() string
	GetCauses() []interface{}
	AppendCause(causes interface{}) IServiceError
	IsEqual(iServiceError IServiceError) bool
}
