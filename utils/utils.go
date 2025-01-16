package utils

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}

func SetMapFieldToStruct(myStruct interface{}, fieldName string, fieldValue interface{}) error {
	structValue := reflect.ValueOf(myStruct).Elem()
	structFieldValue := structValue.FieldByName(fieldName)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("field: %s does not exist in struct", fieldName)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("cannot set %s field value", fieldName)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(fieldValue)
	if structFieldType != val.Type() {
		return errors.New("provided value type didn't match struct field type")
	}

	structFieldValue.Set(val)
	return nil
}

func FillStructFromMap(myStruct interface{}, myMap map[string]interface{}) error {
	for k, v := range myMap {
		err := SetMapFieldToStruct(myStruct, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func BoolToCheckboxValue(val bool) string {
	if val {
		return "on"
	}
	return ""
}
func CheckboxValueToBool(val string) bool {
	return val == "on"
}
func GeneratePassword() (string, error) {
	pw, err := password.Generate(16, 5, 5, false, false)
	if err != nil {
		return "", err
	}
	return pw, nil
}

func GenerateEncryptedPassword() (string, error) {
	pw, err := password.Generate(32, 5, 5, false, false)
	if err != nil {
		return "", err
	}
	encPw, err := EncryptPassword(pw)
	if err != nil {
		return "", err
	}
	return encPw, nil
}
