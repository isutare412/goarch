package pkgerr

import "fmt"

type Known struct {
	Errno  Errno
	Simple error
	Origin error
}

func (k Known) Error() string {
	if k.Origin == nil {
		if k.Simple == nil {
			return fmt.Sprintf("errno: %d", k.Errno)
		}
		return k.Simple.Error()
	}

	if k.Simple == nil {
		return k.Origin.Error()
	}
	return fmt.Sprintf("%s: %s", k.Simple.Error(), k.Origin.Error())
}

func (k Known) SimpleError() string {
	if k.Simple == nil {
		if k.Origin == nil {
			return fmt.Sprintf("errno: %d", k.Errno)
		}
		return k.Origin.Error()
	}
	return k.Simple.Error()
}
