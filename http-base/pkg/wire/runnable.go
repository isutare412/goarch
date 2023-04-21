package wire

type runnable interface {
	Run() <-chan error
}

type namedRunnable struct {
	runnable
	name string
}

func asRunnable(name string, r runnable) namedRunnable {
	return namedRunnable{name: name, runnable: r}
}
