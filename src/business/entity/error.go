package entity

// UnauthorizedError

type UnauthorizedError struct {
	message string
}

func (error UnauthorizedError) Error() string {
	return error.message
}

func NewUnauthorizedError(error string) error {
	return &UnauthorizedError{
		message: error,
	}
}

// NotFoundError

type NotFoundError struct {
	message string
}

func (error NotFoundError) Error() string {
	return error.message
}

func NewNotFoundError(error string) error {
	return &NotFoundError{
		message: error,
	}
}

// DuplicateEntryError

type DuplicateEntryError struct {
	message string
}

func (error *DuplicateEntryError) Error() string {
	return error.message
}

func NewDuplicateEntryError(error string) error {
	return &DuplicateEntryError{
		message: error,
	}
}

// BadGateWayError

type BadGateWayError struct {
	message string
}

func (error *BadGateWayError) Error() string {
	return error.message
}

func NewBadGateWayError(error string) error {
	return &BadGateWayError{
		message: error,
	}
}

// BadRequestError

type BadRequestError struct {
	message string
}

func (error *BadRequestError) Error() string {
	return error.message
}

func NewBadRequestError(error string) error {
	return &BadRequestError{
		message: error,
	}
}
