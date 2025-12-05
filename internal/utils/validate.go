package utils

import (
	"regexp"
)

func ValidateNumber(telefono string) bool {
	ok, _ := regexp.MatchString(`^3\d{9}$`,telefono)
	return ok
}
