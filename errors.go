package ilcd

import "errors"

var (
	// ErrDataSetNotFound indicates that a data set could not be found
	ErrDataSetNotFound = errors.New("data set not found")
)
