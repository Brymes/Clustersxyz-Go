package clusters

import (
	"errors"
)

var (
	ErrorNotFound         = errors.New("resource not found")
	ClustersInternalError = errors.New("clusters had trouble responding to your request")
	UnrecoverableError    = errors.New("internal error")
)
