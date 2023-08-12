package types

type ErrStayNotFound struct {
	Message string
}

func (e *ErrStayNotFound) Error() string {
	if e.Message == "" {
		return "stay not found"
	}

	return e.Message
}
