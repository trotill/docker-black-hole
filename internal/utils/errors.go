package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HttpError struct {
	Validation  error  `json:"validation,omitempty"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func ErrorResponse(ctx *gin.Context, code int, err HttpError) {
	result := map[string]interface{}{}
	if err.Validation != nil {
		var errors []string
		for _, e := range err.Validation.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("Field '%s' failed validation: %s", e.Field(), e.Tag()))
		}

		result["validation"] = errors
	}
	result["code"] = err.Code
	result["description"] = err.Description
	ctx.JSON(code, result)
}
