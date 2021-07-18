package services

type RequestStatus string

const (
	Success  RequestStatus = "Success"
	NotFound RequestStatus = "NotFound"
	Error    RequestStatus = "Error"
)
