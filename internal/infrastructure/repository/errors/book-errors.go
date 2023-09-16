package repository

type BookError struct {
	s string
}

func (e *BookError) Error() string {
	return e.s
}

func NoRowsAffectedError() *BookError {
	return &BookError{
		s: "no rows affected",
	}
}

func MultipleRowsAffectedError() *BookError {
	return &BookError{
		s: "multiple rows affected",
	}
}
