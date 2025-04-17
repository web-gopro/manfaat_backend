package models

type Id struct {
	Id string
}

type IsExists struct {
	TableName  string
	ClomunName string
	ExpValue   string
}

type IsExistsResp struct {
	IsExists bool
	Status   string
}

type Claim struct {
	UserId   string
	UserRole string
}

type GetById struct {
	Id string `json:"id"`
}

type GetList struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}
