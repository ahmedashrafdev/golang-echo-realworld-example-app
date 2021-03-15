package model

type GetAccountRequest struct {
	Code int    `json:"Code" validate:"required"`
	Name string `json:"Name" validate:"required"`
	Type int    `json:"Type" validate:"required"`
}

type Account struct {
	Serial      int
	AccountCode int
	AccountName string
}
