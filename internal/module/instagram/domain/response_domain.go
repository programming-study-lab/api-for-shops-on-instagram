package domain

import "errors"

var ErrUserSuspended = errors.New("domain: user profile is suspended")

type Response struct {
	Status  bool
	Message string
	Data    []map[string]any
}
