package utils

// Any accepts any type
type Any interface{}

//APIResponse is a representation of a success response
type APIResponse struct {
	Data    Any
	Success bool
	Message string
	Error   error
}
