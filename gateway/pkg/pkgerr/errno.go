package pkgerr

type Errno int32

// Internal Error
const (
	ErrnoInternal Errno = iota
	ErrnoRepository
)

// Bad Request
const (
	ErrnoEmptyField Errno = 1000 + iota
)

// Confict
const (
	ErrnoValueAlreadyExists Errno = 2000 + iota
)
