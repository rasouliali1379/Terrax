package errorium

type Error struct {
	code int
	err  string
}

func New(err string, code int) error {
	return &Error{err: err, code: code}
}

func (e Error) Error() string {
	return e.err
}

func (e Error) Code() int {
	return e.code
}

func Is(err error, target int) bool {
	if err, ok := err.(*Error); ok {
		if err.Code() == target {
			return true
		}
	}

	return false
}
