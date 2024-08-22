package helpers

import (
	"prodtrack-api/pkg/errs"
	"reflect"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(data interface{}) errs.ErrorResponse {
	_, err := govalidator.ValidateStruct(data)
	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}

	err = validateStructValue(data)
	if err != nil {
		return err.(errs.ErrorResponse)
	}

	return nil
}

func validateStructValue(data interface{}) errs.ErrorResponse {
	reflectValue := reflect.ValueOf(data)
	reflectType := reflect.TypeOf(data)

	for i := 0; i < reflectValue.NumField(); i++ {
		fieldValue := reflectValue.Field(i)
		fieldType := reflectType.Field(i)

		switch fieldValue.Kind() {
		case reflect.Float64:
			if fieldValue.Float() <= 0 {
				errMsg := fieldType.Name + " must be greater than 0"
				return errs.NewBadRequestError(errMsg)
			}
		}
	}

	return nil
}
