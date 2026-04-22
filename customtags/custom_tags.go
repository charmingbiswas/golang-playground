package customtags

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Example struct for validations
// type User struct {
// 	Username string `validate:"required,min=3,max=20"`
// 	Email    string `validate:"required,email"`
// 	Age      int    `validate:"min=18,max=100"`
// 	Bio      string `validate:"max=500"`
// }

type ValidationError struct {
	Field   string
	Message string
}

// ValidationError is now a type of error since it implements the error package
func (e *ValidationError) Error() string {
	return fmt.Sprintf("field '%s': %s", e.Field, e.Message)
}

func Validate(s any) []ValidationError {
	var errors []ValidationError

	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		return []ValidationError{
			{Field: "root", Message: "input must be a struct"},
		}
	}

	for index := range typ.NumField() {
		field := typ.Field(index)
		fieldVal := val.Field(index)

		tag, ok := field.Tag.Lookup("validate")
		if !ok {
			continue
		}

		rules := strings.Split(tag, ",")

		for _, rule := range rules {
			rule := strings.TrimSpace(rule)
			errs := applyRule(field.Name, fieldVal, rule)
			errors = append(errors, errs...)
		}
	}

	return errors
}

func applyRule(fieldName string, fieldVal reflect.Value, rule string) []ValidationError {
	var errors []ValidationError

	parts := strings.SplitN(rule, "=", 2)
	ruleName := parts[0]
	var ruleParam string
	if len(parts) == 2 {
		ruleParam = parts[1]
	}
	switch ruleName {
	case "required":
		if isEmpty(fieldVal) {
			errors = append(errors, ValidationError{
				Field:   fieldName,
				Message: "this field is required, cannot be empty",
			})
		}
	case "min":
		n, _ := strconv.Atoi(ruleParam)
		switch fieldVal.Kind() {
		case reflect.String:
			if len(fieldVal.String()) < n {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("length must be at least %d characters", n),
				})
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldVal.Int() < int64(n) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("must be atleast %d", n),
				})
			}
		}
	case "max":
		n, _ := strconv.Atoi(ruleParam)
		switch fieldVal.Kind() {
		case reflect.String:
			if len(fieldVal.String()) > n {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("must be at most %d characters", n),
				})
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if fieldVal.Int() > int64(n) {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: fmt.Sprintf("must be <= %d", n),
				})
			}
		}
	case "email":
		if fieldVal.Kind() == reflect.String {
			if !strings.Contains(fieldVal.String(), "@") || !strings.Contains(fieldVal.String(), ".") {
				errors = append(errors, ValidationError{
					Field:   fieldName,
					Message: "must be a valid email address",
				})
			}
		}
	}

	return errors
}

func isEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int() == 0
	case reflect.Slice, reflect.Map:
		return val.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return val.IsNil()
	}
	return false
}
