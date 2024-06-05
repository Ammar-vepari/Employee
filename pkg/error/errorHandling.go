package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type serviceError struct {
	errorType ErrorResponse
	errorCode int
}

func (e serviceError) Error() string {
	return e.errorType.Message
}

func (e serviceError) getErrorCode() int {
	return e.errorCode
}

func (e serviceError) getErrorType() string {
	return e.errorType.ErrorType
}

type ErrorResponse struct {
	ErrorType string
	Message   string
}

func createResponseError(err error) (ErrorResponse, int) {
	srvError, ok := err.(serviceError)
	if !ok {
		srvError = newErrorObject(http.StatusInternalServerError, ErrorResponse{ErrorType: "Unhandelled Error", Message: "Unexpected error"})
	}

	responseError := ErrorResponse{ErrorType: srvError.getErrorType(), Message: srvError.Error()}
	return responseError, srvError.getErrorCode()
}

func SendError(w http.ResponseWriter, err error) {
	responseError, resErrorCode := createResponseError(err)
	responseErrors := make([]ErrorResponse, 1)
	responseErrors[0] = responseError

	data, err := json.Marshal(responseErrors)
	if err != nil {
		http.Error(w, "UnexpectedError", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(resErrorCode)
	fmt.Fprintln(w, string(data))
}

func newErrorObject(errorCode int, errorType ErrorResponse) serviceError {
	return serviceError{
		errorCode: errorCode,
		errorType: errorType,
	}
}

func NewApplicationError(errorCode int, errorMessage ErrorResponse) error {
	return newErrorObject(errorCode, errorMessage)
}
