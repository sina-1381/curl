package validation

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

func Simple(ver validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range ver {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}
	return errs
}

func CheckValidate(input interface{}, c *gin.Context) {
	if err := c.ShouldBindJSON(input); err != nil {
		var ver validator.ValidationErrors
		if errors.As(err, &ver) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": Simple(ver)})
			panic("validation error")
		}
	}
}
