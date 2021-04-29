package models

import "gorm.io/gorm"

type UserCompany struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}

type UserCompanyResponse struct {
	Code    int           `json:"code", form:"code"`
	Message string        `json:"message", form:"message"`
	Data    []UserCompany `json:"data", form:"data"`
	Status  string        `json:"status", form:"status"`
}

type UserCompanyRequest struct {
	Name      string `json:"name", form:"name"`
	CompanyID int    `json:"companyId", form:"companyId"`
}

type UserResponseCompanySingle struct {
	Code    int         `json:"code", form:"code"`
	Message string      `json:"message", form:"message"`
	Data    UserCompany `json:"data", form:"data"`
	Status  string      `json:"status", form:"status"`
}
