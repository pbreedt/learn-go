package errors

type CustomError string

const (
	ErrUserNotFound CustomError = "User not found"
	ErrDatabase     CustomError = "Database error"
	ErrInvalidInput CustomError = "Invalid input"
)

func (e CustomError) Error() string {
	return string(e)
}
