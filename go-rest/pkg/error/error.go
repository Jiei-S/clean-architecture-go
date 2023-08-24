package error

type ErrorLevel int8

const (
	_ ErrorLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

type ErrorCode int

const (
	_ ErrorCode = iota
	CodeBadRequest
	CodeNotFound
	CodeDuplicate
	CodeInternalServerError
)

type ApplicationError struct {
	message string
	level   ErrorLevel
	code    ErrorCode
}

func (e *ApplicationError) Error() string {
	return e.message
}

func (e *ApplicationError) Level() ErrorLevel {
	return e.level
}

func (e *ApplicationError) Code() ErrorCode {
	return e.code
}

func NewApplicationError(message string, level ErrorLevel, code ErrorCode) *ApplicationError {
	return &ApplicationError{
		message: message,
		level:   level,
		code:    code,
	}
}
