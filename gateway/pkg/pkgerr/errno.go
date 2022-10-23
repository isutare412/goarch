package pkgerr

type Errno int32

// Internal Error
const (
	ErrnoInternal Errno = iota
	ErrnoRepository
)

// Bad Request
const (
	ErrnoBadRequest Errno = 1000 + iota
	ErrnoEmptyField
)

// Confict
const (
	ErrnoConflict Errno = 2000 + iota
	ErrnoValueAlreadyExists
)

// Not Found
const (
	ErrnoNotFound Errno = 3000 + iota
)
