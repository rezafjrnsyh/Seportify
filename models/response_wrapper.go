package models

type ResponseWrapper struct {
	Success bool
	Message string
	Data    interface{}
}
