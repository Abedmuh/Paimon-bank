package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateTransferProof(fl validator.FieldLevel) bool {
	// Kode validasi yang disesuaikan sesuai kebutuhan Anda
	match, _ := regexp.MatchString(`^(http|https)://[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(?:/[^/#?]+)+\.(?:jpg|jpeg|png)$`, fl.Field().String())
	return match
}