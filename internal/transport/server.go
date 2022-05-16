package transport

import "context"

type Server interface {
	Run(cxt context.Context) error
}
