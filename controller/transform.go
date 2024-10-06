package controller

import (
	"fmt"
	"net/http"
	"songs/apperrors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func transformToInt64(val string, name string, errs *[]error) int64 {
	if val == "" {
		return 0
	}
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		*errs = append(*errs, fmt.Errorf("cannot get number from %v ", name))
		return 0
	}
	return num
}

func transformToInt(val string, name string, errs *[]error) int {
	return int(transformToInt64(val, name, errs))
}

func validateDTO[T any](c *gin.Context, data *T) error {
	if err := c.BindJSON(&data); err != nil {
		return apperrors.ErrBadRequest
	}

	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		return apperrors.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        fmt.Errorf("validation error: %s", errors),
		}
	}
	return nil
}
