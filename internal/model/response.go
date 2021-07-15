package model

type EmployeeResponse struct {
	ServerStatus int        `json:"server_status"`
	Message      string     `json:"message"`
	Employee     []Employee `json:"employee"`
}

type StatusMessageResponse struct {
	ServerStatus int    `json:"server_status"`
	Message      string `json:"message"`
}
