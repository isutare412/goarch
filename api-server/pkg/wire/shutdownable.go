package wire

import "context"

type shutdownable interface {
	Shutdown(context.Context) error
}

type namedShutdownable struct {
	shutdownable
	name string
}

func asShutdownable(name string, s shutdownable) namedShutdownable {
	return namedShutdownable{name: name, shutdownable: s}
}
