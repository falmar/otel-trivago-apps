package kithelper

type ErrInvalidArgument struct {
	Message string
}

func (e *ErrInvalidArgument) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "missing argument"
}
