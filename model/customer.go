package model

type (
	//Customer type represents customer
	Customer struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Mobile uint   `json:"mobile"`
	}
)
