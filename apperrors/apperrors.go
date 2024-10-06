package apperrors

import (
	"errors"
	"log"
	"net/http"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (re RequestError) Error() string {
	log.Println(re.Err.Error())
	return re.Err.Error()
}

var ErrNotFound = RequestError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("not found"),
}

var ErrBadRequest = RequestError{
	StatusCode: http.StatusBadRequest,
	Err:        errors.New("bad request data"),
}

var ErrBadApiResponse = RequestError{
	StatusCode: http.StatusInternalServerError,
	Err:        errors.New("bad api response"),
}

func GetErrorAndStatus(err error) (int, string) {
	switch re := err.(type) {
	case RequestError:
		return re.StatusCode, re.Error()
	default:
		message := "internal server error"
		log.Println(message)
		return http.StatusInternalServerError, message
	}
}
