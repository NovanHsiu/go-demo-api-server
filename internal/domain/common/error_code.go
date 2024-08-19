package common

import "net/http"

type ErrorCode struct {
	StatusCode int
}

/*
	General error codes
*/

var ErrorCodeInternalProcess = ErrorCode{
	StatusCode: http.StatusInternalServerError,
}

/*
	Authentication and Authorization error codes
*/

var ErrorCodeAuthPermissionDenied = ErrorCode{
	StatusCode: http.StatusForbidden,
}

var ErrorCodeAuthNotAuthenticated = ErrorCode{
	StatusCode: http.StatusUnauthorized,
}

/*
	Resource-related error codes
*/

var ErrorCodeResourceNotFound = ErrorCode{
	StatusCode: http.StatusNotFound,
}

/*
	Parameter-related error codes
*/

var ErrorCodeParameterInvalid = ErrorCode{
	StatusCode: http.StatusBadRequest,
}

/*
	Remote server-related error codes
*/

var ErrorCodeRemoteProcess = ErrorCode{
	StatusCode: http.StatusBadGateway,
}
