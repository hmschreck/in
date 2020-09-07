package in

import (
	"errors"
	"fmt"
	"reflect"
)

// takes in a slice or array of a given type as `input`
func CheckOne(input interface{}, check interface{}) (isIn bool, err error) {
	inputType := reflect.TypeOf(input)
	inputKind := inputType.Kind()
	if !(inputKind == reflect.Slice || inputKind == reflect.Array) {
		return false, errors.New("input is not a valid slice or array")
	}
	inputElementType := inputType.Elem()
	checkType := reflect.TypeOf(check)
	if inputElementType != checkType {
		errString := fmt.Sprintf("types do not match, list type is %v, check type is %v", inputElementType, checkType)
		return false, errors.New(errString)
	}
	isIn = false
	inputValue := reflect.ValueOf(input)
	checkValue := reflect.ValueOf(check)
	for i := 0; i < inputValue.Len(); i++ {
		inputValueReal := reflect.ValueOf(inputValue.Index(i).Interface())
		// I know.  It's horrible.  But I verify above that the types ARE the same.  So.  It should be fine.
		if fmt.Sprintf("%v", inputValueReal) == fmt.Sprintf("%v", checkValue) {
			isIn = true
			break
		}
	}
	return
}
