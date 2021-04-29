package models

type ResponseNotif struct {
	Code    int    `json:"code", form:"code"`
	Message string `json:"message", form:"message"`
	Status  string `json:"status", form:"status"`
}
