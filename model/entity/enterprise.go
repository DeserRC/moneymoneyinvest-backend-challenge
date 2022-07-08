package entity

import (
	"gopkg.in/validator.v2"
)

type Enterprise struct {
	Cnpj      string `json:"cnpj" validate:"regexp=^([\\d]{2}\\.[\\d]{3}\\.[\\d]{3}\\/[\\d]{4}\\-[\\d]{2})|([\\d]{14})$"`
	Invoicing int    `json:"invoicing"`
	Address   string `json:"address" validate:"nonzero"`
	Requester Client `json:"requester" validate:"nonnil"`
}

func ValidateEnterprise(enterprise *Enterprise) error {
	return validator.Validate(enterprise)
}
