package domain

import "errors"

var (
	ErrNotFound          = errors.New("resource not found")
	ErrDuplicate         = errors.New("resource duplicate")
	ErrInvalidIDParam    = errors.New("invalid id parameter")
	ErrInvalidPageParam  = errors.New("invalid page parameter")
	ErrInvalidLimitParam = errors.New("invalid limit parameter")
)
