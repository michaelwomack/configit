package configit

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func Load(target interface{}) error {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || targetValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("invalid argument: target must be a pointer to a struct")
	}
	val := targetValue.Elem()
	return populateStruct(val)
}

func populateStruct(val reflect.Value) error {
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if err := populateField(field, val.Field(i)); err != nil {
			return fmt.Errorf("failed to populate target: %v", err)
		}
	}
	return nil
}

func populateField(field reflect.StructField, value reflect.Value) error {
	envName := field.Tag.Get("env")

	envValue, _ := os.LookupEnv(envName)

	ParseError := fmt.Errorf("failed to parse %s as %v", envValue, value.Kind())

	switch value.Kind() {
	case reflect.Bool:
		parsed, err := strconv.ParseBool(envValue)
		if err != nil {
			return ParseError
		}
		value.SetBool(parsed)
	case reflect.String:
		value.SetString(envValue)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		parsed, err := strconv.ParseInt(envValue, 10, 64)
		if err != nil {
			return ParseError
		}
		value.SetInt(parsed)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		parsed, err := strconv.ParseUint(envValue, 10, 64)
		if err != nil {
			return ParseError
		}
		value.SetUint(parsed)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		parsed, err := strconv.ParseFloat(envValue, 64)
		if err != nil {
			return ParseError
		}
		value.SetFloat(parsed)
	case reflect.Struct:
		return populateStruct(value)
	}
	return nil
}
