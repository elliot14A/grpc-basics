package custom_errors

type CustomError struct {
	Err     error
	Context CustomErrorType
}

type CustomErrorType string

const (
	NotFound              CustomErrorType = "not found"
	InvalidDatabaseClient CustomErrorType = "invalid database client"
)

// Error implements error
func (c *CustomError) Error() string {
	return c.Err.Error()
}
