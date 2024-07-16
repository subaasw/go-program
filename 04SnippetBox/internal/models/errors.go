package models

import "errors"

var (
	ErrNoRows             = errors.New("models: no matching record found")
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)
