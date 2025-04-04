package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Error list.
var (
	ErrInternalDB     = errors.New("internal database error")
	ErrInternalServer = errors.New("internal server error")
	//customer error
	ErrFailedGetListCustomer    = errors.New("failed to get list customer")
	ErrFailedGetCustomer        = errors.New("failed to get customer")
	ErrFailedCreateCustomer     = errors.New("failed to create customer")
	ErrFailedUpdateCustomer     = errors.New("failed to update customer")
	ErrFailedDeleteCustomer     = errors.New("failed to delete customer")
	ErrFailedActivateCustomer   = errors.New("failed to activate customer")
	ErrFailedDeactivateCustomer = errors.New("failed to deactivate customer")
)

// ErrUniqueField is error for unique/duplicate field.
func ErrUniqueField(field string, value string) error {
	return fmt.Errorf("data with %s = %s already exist", field, value)
}

// ErrRequiredField is error for missing field.
func ErrRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

// ErrGTField is error for greater than field.
func ErrRangeField(str, value interface{}, value2 interface{}) error {
	return fmt.Errorf("field %s must be between %v and %v", str, value, value2)
}

// ErrGTField is error for greater than field.
func ErrGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

// ErrGTEField is error for greater than or equal field.
func ErrGTEField(str, value string) error {
	return fmt.Errorf("field %s must be greater than or equal %s", str, value)
}

// ErrLTField is error for lower than field.
func ErrLTField(str, value string) error {
	return fmt.Errorf("field %s must be lower than %s", str, value)
}

// ErrLTEField is error for lower than or equal field.
func ErrLTEField(str, value string) error {
	return fmt.Errorf("field %s must be lower than or equal %s", str, value)
}

// ErrLenField is error for length field.
func ErrLenField(str, value string) error {
	return fmt.Errorf("field %s length must be %s", str, value)
}

// ErrISO3166Alpha2Field is error for ISO 3166-1 alpha-2 field.
func ErrISO3166Alpha2Field(str string) error {
	return fmt.Errorf("field %s must be in ISO 3166-1 alpha-2 format", str)
}

// ErrEmailField is error for email field.
func ErrEmailField(str string) error {
	return fmt.Errorf("field %s must be in email format", str)
}

// ErrURLField is error for url field.
func ErrURLField(str string) error {
	return fmt.Errorf("field %s must be in URL format", str)
}

// ErrInvalidFormatField is error for invalid format field.
func ErrInvalidFormatField(str string) error {
	return fmt.Errorf("invalid format field %s", str)
}

// ErrOneOfField is error for oneof field.
func ErrOneOfField(str, value string) error {
	return fmt.Errorf("field %s must be one of %s", str, strings.Join(strings.Split(value, " "), "/"))
}

// ErrNumericField is error for numeric field.
func ErrNumericField(str string) error {
	return fmt.Errorf("field %s must contain number only", str)
}

// ErrAlphaField is error for alpha field.
func ErrAlphaField(str string) error {
	return fmt.Errorf("field %s must contain letter only", str)
}

// ErrMinPrice is error for min price.
func ErrMinPrice(price float64) error {
	return fmt.Errorf("minimum price is %d", int(price))
}

// ErrMaxPrice is error for max price.
func ErrMaxPrice(price float64) error {
	return fmt.Errorf("maximum price is %d", int(price))
}

// ErrMinAmount is error for min amount.
func ErrMinAmount(amount float64) error {
	return fmt.Errorf("minimum amount is %d", int(amount))
}

// ErrMaxAmount is error for max amount.
func ErrMaxAmount(amount float64) error {
	return fmt.Errorf("maximum amount is %d", int(amount))
}

// ErrDatetimeField is error for datetime field.
func ErrDatetimeField(str string) error {
	return fmt.Errorf("%s must be in date format (dd/mm/yyyy)", str)
}

// Err generate token.
func ErrGenerateToken(str string) error {
	return fmt.Errorf("generate token: <%s>", str)
}

// ErrPhoneNumber is error for email field.
func ErrPhoneField(str string) error {
	return fmt.Errorf("field %s must be in phone format", str)
}

// CodeFromGRPC to convert grpc response code
// to http response code.
func CodeFromGRPC(err error) int {
	if err == nil {
		return http.StatusOK
	}

	st, ok := status.FromError(err)
	if !ok {
		return http.StatusInternalServerError
	}

	switch st.Code() {
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.NotFound:
		return http.StatusNotFound
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.Internal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

// MessageFromGRPC to convert grpc response message
// to http response message.
func MessageFromGRPC(err error) error {
	if err == nil {
		return nil
	}

	st, ok := status.FromError(err)
	if !ok {
		return err
	}

	return errors.New(st.Message())
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
