package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

type EchoValidator struct {
	Validator *validator.Validate
}

func (v *EchoValidator) Validate(i any) error {
	return v.Validator.Struct(i)
}

func RegisterValidator(e *echo.Echo) {
	e.Validator = &EchoValidator{Validator: validator.New()}
}
