package common

type Error interface {
	Error() string
}

type DomainError struct {
	code         ErrorCode
	ErrorMessage string // error message for developer
	Description  string // error message for user
	Err          error  // native error
}

// NewError return Error defined by code. parameter messages, first message be assigned to ErrorMessage, second message be assigned to Description.
func NewError(code ErrorCode, err error, messages ...string) Error {
	domainError := DomainError{
		code: code,
		Err:  err,
	}
	for i := range messages {
		if i == 0 {
			domainError.ErrorMessage = messages[0]
		} else if i == 1 {
			domainError.Description = messages[1]
		} else {
			break
		}
	}
	if domainError.code.StatusCode == 0 {
		domainError.code.StatusCode = 500
		domainError.ErrorMessage = "unknown error"
		domainError.Description = domainError.ErrorMessage
	}
	if domainError.ErrorMessage == "" {
		if domainError.Err != nil {
			domainError.ErrorMessage = err.Error()
		}
	}
	if domainError.Description == "" {
		domainError.Description = domainError.ErrorMessage
	}
	return domainError
}

func (e DomainError) StatusCode() int {
	return e.code.StatusCode
}

func (e DomainError) Error() string {
	if e.ErrorMessage != "" {
		return e.ErrorMessage
	}
	return e.Err.Error()
}
