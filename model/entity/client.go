package entity

import "gopkg.in/validator.v2"

type Client struct {
	Name  string `json:"name" validate:"nonzero"`
	Cpf   string `json:"cpf" validate:"regexp=^([\\d]{3}\\.[\\d]{3}\\.[\\d]{3}\\-[\\d]{2}|[\\d]{11})$"`
	Phone string `json:"phone" validate:"regexp=^([0]?[\\d]{2})?[\\s]?([\\d]{8}|[\\d]{9})$"`
	Email string `json:"email" validate:"regexp=^[\\w\\d]+@[\\w\\d]+\\.[\\w\\d]+$"`
}

func ValidateClient(client *Client) error {
	return validator.Validate(client)
}
