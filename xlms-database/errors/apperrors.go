package cutomeerror

func ErrBadRequest(msg string, err error) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: 400,
		Err:        err,
	}
}

func ErrNotFound(msg string, err error) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: 404,
		Err:        err,
	}
}

func ErrInternal(err error) *AppError {
	return &AppError{
		Message:    "Internal server error",
		StatusCode: 500,
		Err:        err,
	}
}

func ErrDuplication(err error) *AppError {
	return &AppError{
		Message:    "Filed already exists",
		StatusCode: 409,
		Err:        err,
	}
}

func ErrForbidden(err error) *AppError {
	return &AppError{
		Message:    "Unauthorized access",
		StatusCode: 403,
		Err:        err,
	}
}
