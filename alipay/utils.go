package alipay

import (
	"fmt"
	"reflect"
)

func SetField(obj interface{}, field string, value string) error {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("expected a pointer")
	}
	v = v.Elem()

	f := v.FieldByName(field)
	if !f.IsValid() {
		return fmt.Errorf("no such field: %s in obj", field)
	}
	if !f.CanSet() {
		return fmt.Errorf("cannot set field %s", field)
	}

	f.SetString(value)
	return nil
}
