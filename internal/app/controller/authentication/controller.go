package authentication

import (
	"github.com/lutracorp/aonyx/internal/pkg/validator"
	"github.com/matthewhartstonge/argon2"
)

// Controller represents authentication controller.
type Controller struct {
	Argon     *argon2.Config
	Validator *validator.Validator
}

// NewController creates new instance of authentication Controller.
func NewController(arc *argon2.Config) *Controller {
	return &Controller{
		Argon:     arc,
		Validator: validator.NewValidator(),
	}
}
