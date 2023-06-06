package pkgerr

// Simple hides detailed error chain by calling SimpleError. It is used to
// response only simple messages to client.
type Simple struct {
	Origin error
	Simple error
}

func (s Simple) Error() string {
	if s.Origin != nil {
		return s.Origin.Error()
	}
	if s.Simple != nil {
		return s.Simple.Error()
	}
	return "known error uninitialized"
}

func (s Simple) SimpleError() string {
	if s.Simple != nil {
		return s.Simple.Error()
	}
	if s.Origin != nil {
		return s.Origin.Error()
	}
	return "known error uninitialized"
}

func (s Simple) Unwrap() []error {
	if s.Origin != nil {
		return []error{s.Origin}
	}
	if s.Simple != nil {
		return []error{s.Simple}
	}
	return nil
}
