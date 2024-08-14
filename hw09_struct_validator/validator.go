package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

type validateHandlerFunc func(validationKey, validationValue, fieldName string, fieldValue interface{}) error

func (v ValidationErrors) Error() string {
	errString := strings.Builder{}

	for _, err := range v {
		errString.WriteString(fmt.Sprintf("Field: %s - Error: %v\n", err.Field, err.Err))
	}

	return errString.String()
}

var (
	ErrInterfaceType       = errors.New("interface is not struct")
	ErrInterfaceConversion = errors.New("invalid type for conversion")

	ErrStringLength = errors.New("length of string not equals number in tag")
	ErrStringRegexp = errors.New("string not equals regexp in tag")
	ErrStringIn     = errors.New("string is not contains in subset of tag")

	ErrNumberMax = errors.New("number is bigger than max")
	ErrNumberMin = errors.New("number is less than min")
	ErrNumberIn  = errors.New("number is not contains in subset of tag")

	ErrEmptyTag = errors.New("the tag cannot be empty")
)

func Validate(v interface{}) error {
	// Place your code here.
	var (
		validateTagName  = "validate"
		validationErrors ValidationErrors
		wg               sync.WaitGroup
		mu               sync.Mutex
	)

	structure := reflect.ValueOf(v)

	if structure.Kind() != reflect.Struct {
		validationErrors = append(validationErrors, ValidationError{
			Field: structure.Type().Name(),
			Err:   ErrInterfaceType,
		})

		return validationErrors
	}

	for i := 0; i < structure.NumField(); i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			field := structure.Type().Field(i)
			tagValue, ok := field.Tag.Lookup(validateTagName)
			if !ok {
				return
			}

			switch field.Type.Kind() {
			case reflect.String:
				fieldValue := structure.Field(i).String()

				if err := validateHandler(tagValue, field.Name, fieldValue, validateString); err != nil {
					mu.Lock()
					validationErrors = errHandler(err, validationErrors, field.Name)
					mu.Unlock()
				}

			case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int:
				fieldValue := structure.Field(i).Int()

				if err := validateHandler(tagValue, field.Name, fieldValue, validateNumber); err != nil {
					mu.Lock()
					validationErrors = errHandler(err, validationErrors, field.Name)
					mu.Unlock()
				}

			case reflect.Slice:
				fieldValue := structure.Field(i).Interface()

				if err := validateHandler(tagValue, field.Name, fieldValue, validateSlice); err != nil {
					mu.Lock()
					validationErrors = errHandler(err, validationErrors, field.Name)
					mu.Unlock()
				}

			case reflect.Struct:
				fieldValue := structure.Field(i).Interface()
				err := Validate(fieldValue)

				mu.Lock()
				validationErrors = errHandler(err, validationErrors, field.Name)
				mu.Unlock()
			case reflect.Invalid, reflect.Bool, reflect.Uint, reflect.Uint8, reflect.Uint16,
				reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64,
				reflect.Complex128, reflect.Array, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
				reflect.Pointer, reflect.UnsafePointer:
				fallthrough
			default:
				return
			}
		}(i)
	}

	wg.Wait()

	return validationErrors
}

func validateHandler(allConstraints, fieldName string, fieldValue interface{},
	validateHandler validateHandlerFunc,
) error {
	constraintSeparator := "|"
	allValidationErrors := make(ValidationErrors, 0)

	if allConstraints == "" || fieldValue == nil {
		return nil
	}

	constraints := strings.Split(allConstraints, constraintSeparator)

	for _, constraint := range constraints {
		validationKey, validationValue, err := getValidationPair(constraint)
		if err != nil {
			allValidationErrors = append(allValidationErrors, ValidationError{
				Field: fieldName,
				Err:   err,
			})
			continue
		}

		if err = validateHandler(validationKey, validationValue, fieldName, fieldValue); err != nil {
			allValidationErrors = errHandler(err, allValidationErrors, fieldName)
		}
	}

	return allValidationErrors
}

func getValidationPair(constraint string) (string, string, error) {
	keyValueSeparator := ":"
	splitConstraint := strings.Split(constraint, keyValueSeparator)

	if len(splitConstraint) == 1 {
		return "", "", ErrEmptyTag
	} else if len(splitConstraint) == 2 && splitConstraint[1] == "" {
		return "", "", ErrEmptyTag
	}

	key := splitConstraint[0]
	value := splitConstraint[1]

	return key, value, nil
}

func errHandler(err error, validationErrors ValidationErrors, fieldName string) ValidationErrors {
	var valErr ValidationErrors

	if errors.As(err, &valErr) {
		validationErrors = append(validationErrors, valErr...)
	} else {
		validationErrors = append(validationErrors, ValidationError{
			Field: fieldName,
			Err:   err,
		})
	}

	return validationErrors
}

func validateString(validationKey, validationValue, fieldName string, fieldValue interface{}) error {
	validateStringErrors := make(ValidationErrors, 0)
	str, ok := fieldValue.(string)
	if !ok {
		validateStringErrors = append(validateStringErrors, ValidationError{
			Field: fieldName,
			Err:   ErrInterfaceConversion,
		})
		return validateStringErrors
	}

	switch validationKey {
	case "len":
		lengthNumber, err := strconv.ParseInt(validationValue, 10, 64)
		if err != nil {
			validateStringErrors = append(validateStringErrors, ValidationError{
				Field: fieldName,
				Err:   err,
			})
			return validateStringErrors
		}

		if len(str) != int(lengthNumber) {
			validateStringErrors = append(validateStringErrors, ValidationError{
				Field: fieldName,
				Err:   ErrStringLength,
			})
		}

	case "regexp":
		reg, err := regexp.Compile(validationValue)
		if err != nil {
			validateStringErrors = append(validateStringErrors, ValidationError{
				Field: fieldName,
				Err:   err,
			})
		}

		if !reg.MatchString(str) {
			validateStringErrors = append(validateStringErrors, ValidationError{
				Field: fieldName,
				Err:   ErrStringRegexp,
			})
		}

	case "in":
		if !strings.Contains(validationValue, str) {
			validateStringErrors = append(validateStringErrors, ValidationError{
				Field: fieldName,
				Err:   ErrStringIn,
			})
		}
	}

	return validateStringErrors
}

func validateNumber(validationKey, validationValue, fieldName string, fieldValue interface{}) error {
	var validateNumberErrors ValidationErrors
	number, ok := fieldValue.(int64)
	if !ok {
		validateNumberErrors = append(validateNumberErrors, ValidationError{
			Field: fieldName,
			Err:   ErrInterfaceConversion,
		})
		return validateNumberErrors
	}

	switch validationKey {
	case "max":
		maxConstraint, err := strconv.ParseInt(validationValue, 10, 64)
		if err != nil {
			validateNumberErrors = append(validateNumberErrors, ValidationError{
				Field: fieldName,
				Err:   err,
			})
			return validateNumberErrors
		}
		if number > maxConstraint {
			validateNumberErrors = append(validateNumberErrors, ValidationError{
				Field: fieldName,
				Err:   ErrNumberMax,
			})
		}

	case "min":
		minConstraint, err := strconv.ParseInt(validationValue, 10, 64)
		if err != nil {
			validateNumberErrors = append(validateNumberErrors, ValidationError{
				Field: fieldName,
				Err:   err,
			})
			return validateNumberErrors
		}

		if number < minConstraint {
			validateNumberErrors = append(validateNumberErrors, ValidationError{
				Field: fieldName,
				Err:   ErrNumberMin,
			})
		}

	case "in":
		if !strings.Contains(validationValue, strconv.Itoa(int(number))) {
			validateNumberErrors = append(validateNumberErrors, ValidationError{
				Field: fieldName,
				Err:   ErrNumberIn,
			})
		}
	}

	return validateNumberErrors
}

func validateSlice(validationKey, validationValue, fieldName string, fieldValue interface{}) error {
	validateErrors := make(ValidationErrors, 0)

	switch field := fieldValue.(type) {
	case []int64:
		for _, elem := range field {
			err := validateNumber(validationKey, validationValue, fieldName, elem)
			var valErr ValidationErrors

			if errors.As(err, &valErr) {
				validateErrors = append(validateErrors, valErr...)
			} else {
				valErr = append(valErr, ValidationError{
					Field: fieldName,
					Err:   err,
				})
			}
		}
	case []string:
		for _, elem := range field {
			err := validateString(validationKey, validationValue, fieldName, elem)
			var valErr ValidationErrors

			if errors.As(err, &valErr) {
				validateErrors = append(validateErrors, valErr...)
			} else {
				valErr = append(valErr, ValidationError{
					Field: fieldName,
					Err:   err,
				})
			}
		}
	default:
		validateErrors = append(validateErrors, ValidationError{
			Field: fieldName,
			Err:   ErrInterfaceConversion,
		})
	}

	return validateErrors
}
