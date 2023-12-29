package validator

import (
	"slices"
	"strings"
	"unicode/utf8"
)

type Validator struct {
	FieldErrors map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

func (v *Validator) Add(field, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}

	if _, ok := v.FieldErrors[field]; !ok {
		v.FieldErrors[field] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.Add(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func MaxStringLength(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

func PermitteValues[T comparable](value T, options ...T) bool {
	return slices.Contains(options, value)
}
