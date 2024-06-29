package constant

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	UsernameRules = []validation.Rule{
		validation.Required,
		validation.Match(regexp.MustCompile(`^[a-zA-Z0-9]*$`)),
	}
)
