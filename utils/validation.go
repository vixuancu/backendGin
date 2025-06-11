package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func ValidationRequied(fieldName, value string) error {
	if value == "" {
		return fmt.Errorf("%s bat buoc phai nhap", fieldName)
	}
	return nil
}
func ValidationStringLength(fieldName, value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return fmt.Errorf("%s phai co do dai tu %d toi %d", fieldName, min, max)
	}
	return nil
}
func ValidationRegex(fieldName, value, errorMessage string, regex *regexp.Regexp) error {
	if !regex.MatchString(value) {
		return fmt.Errorf("%s: %s ", fieldName, errorMessage)
	}
	return nil
}

func HandleValidationError(err error) gin.H {
	if validation, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validation {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn %s", e.Field(), e.Param())
			case "lt":
				errors[e.Field()] = fmt.Sprintf("%s phải nhỏ hơn %s", e.Field(), e.Param())
			case "gte":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", e.Field(), e.Param())
			case "lte":
				errors[e.Field()] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", e.Field(), e.Param())
			case "uuid":
				errors[e.Field()] = fmt.Sprintf("%s phải là một UUID hợp lệ", e.Field())
			case "slug":
				errors[e.Field()] = fmt.Sprintf("%s chỉ chữ cái thường,số,dấu - .", e.Field())
			case "min":
				errors[e.Field()] = fmt.Sprintf("%s phải lớn hơn %s", e.Field(), e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("%s phải nhỏ hơn %s", e.Field(), e.Param())
			case "oneof":
				allowdValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[e.Field()] = fmt.Sprintf("%s phải là 1 trong các giá trị %s", e.Field(), allowdValues)
			case "required":
				errors[e.Field()] = fmt.Sprintf("Trường %s bắt buộc phải nhập", e.Field())
			case "search":
				errors[e.Field()] = fmt.Sprintf("Trường %s không nhập được các kí tự đặc biệt", e.Field())
			case "email":
				errors[e.Field()] = fmt.Sprintf("Trường %s phải đúng định dạng", e.Field())
			case "datetime":
				errors[e.Field()] = fmt.Sprintf("Trường %s phải đúng định dạng YYYY-MM-DD", e.Field())
			}
		}
		return gin.H{"error": errors}
	}
	return gin.H{
		"error": "Validation failed- yêu cầu không hợp lệ " + err.Error(),
	}
}
func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("không thể lấy validator từ binding")
	}
	var slugRegex = regexp.MustCompile("^[a-z0-9]+(?:[-.][a-z0-9]+)*$")
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return slugRegex.MatchString(value)
	})
	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return searchRegex.MatchString(value)
	})
	return nil
}
