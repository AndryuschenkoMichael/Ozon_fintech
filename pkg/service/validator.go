package service

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

const (
	validationError = "validation error"
)

type Valid struct{}

func NewValid() *Valid {
	return &Valid{}
}

func (v *Valid) ValidateLink(link string) error {
	if utf8.RuneCountInString(link) != LengthLink {
		return errors.New(validationError)
	}

	expr := fmt.Sprintf("^[a-zA-z0-9_]{%d}", utf8.RuneCountInString(link))
	validLink := regexp.MustCompile(expr)

	if !validLink.MatchString(link) {
		return errors.New(validationError)
	}

	return nil
}
