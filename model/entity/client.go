package entity

import "gopkg.in/validator.v2"

type Client struct {
	Name  string `json:"name" validate:"nonzero"`
	Cpf   string `json:"cpf" validate:"regexp=^([\\d]{3}\\.[\\d]{3}\\.[\\d]{3}\\-[\\d]{2})|([\\d]{11})$"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func ValidateClient(client *Client) error {
	return validator.Validate(client)
}
