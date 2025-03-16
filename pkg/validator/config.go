package validator

import (
	"fmt"
	"regexp"

	cons "github.com/aaalik/anton-users/internal/constant"
	"github.com/aaalik/anton-users/internal/service"
	validation "github.com/go-ozzo/ozzo-validation"
)

func init() {
	RegisterValidator("login", func(value interface{}) error {
		request, ok := value.(*service.RequestLogin)
		if !ok {
			return fmt.Errorf("invalid type, %T is not RequestLogin", value)
		}

		errs := validation.Errors{
			"username": validation.Validate(request.Username, cons.UsernameRules...),
			"password": validation.Validate(request.Password, validation.Required),
		}

		return errs.Filter()
	})

	RegisterValidator("refresh_token", func(value interface{}) error {
		request, ok := value.(*service.RequestRefreshToken)
		if !ok {
			return fmt.Errorf("invalid type, %T is not RequestRefreshToken", value)
		}

		errs := validation.Errors{
			"refresh_token": validation.Validate(request.RefreshToken, validation.Required),
		}

		return errs.Filter()
	})

	RegisterValidator("register", func(value interface{}) error {
		request, ok := value.(*service.RequestRegister)
		if !ok {
			return fmt.Errorf("invalid type, %T is not RequestRegister", value)
		}

		errs := validation.Errors{
			"username": validation.Validate(request.Username, cons.UsernameRules...),
			"password": validation.Validate(request.Password, validation.Required),
			"name":     validation.Validate(request.Name, validation.Required),
		}

		return errs.Filter()
	})

	RegisterValidator("create_user", func(value interface{}) error {
		request, ok := value.(*service.RequestCreateUser)
		if !ok {
			return fmt.Errorf("invalid type, %T is not RequestCreateUser", value)
		}

		errs := validation.Errors{
			"username": validation.Validate(request.Username, cons.UsernameRules...),
			"password": validation.Validate(request.Password, validation.Required),
			"name":     validation.Validate(request.Name, validation.Required),
			"dob":      validation.Validate(request.Dob, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]{4}[-][0-9]{2}[-][0-9]{2}$`))),
			"gender":   validation.Validate(request.Gender, validation.Required),
		}

		return errs.Filter()
	})

	RegisterValidator("update_user", func(value interface{}) error {
		request, ok := value.(*service.RequestUpdateUser)
		if !ok {
			return fmt.Errorf("invalid type, %T is not RequestCreateUser", value)
		}

		errs := validation.Errors{
			"id":     validation.Validate(request.Id, validation.Required),
			"name":   validation.Validate(request.Name, validation.Required),
			"dob":    validation.Validate(request.Dob, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]{4}[-][0-9]{2}[-][0-9]{2}$`))),
			"gender": validation.Validate(request.Gender, validation.Required),
		}

		return errs.Filter()
	})
}
