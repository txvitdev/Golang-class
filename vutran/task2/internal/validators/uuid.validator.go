package validator

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

var uuidValidator validator.Func = func(fl validator.FieldLevel) bool {
	field := fl.Field().String()

	_, err := uuid.Parse(field)

	if err != nil {
		return false
	}

	return true
}

func RegisterCustomValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("uuid", uuidValidator)
	}
}
