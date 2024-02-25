package utils

import (
	"io"
	"log"
	"reflect"
)

func CloseElement(closable io.Closer) {
	err := closable.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// Mapper Use: give a pointer to the source type, and a pointer to a new target type to populate
func Mapper[K any, T any](inPointer K, outPointer T) (T, bool) {
	if reflect.TypeOf(inPointer).Kind() != reflect.Pointer || reflect.TypeOf(outPointer).Kind() != reflect.Pointer {
		return outPointer, false
	}
	in := reflect.ValueOf(inPointer).Elem()
	out := reflect.ValueOf(outPointer).Elem()

	if in.Kind() != reflect.Struct || out.Kind() != reflect.Struct {
		return outPointer, false
	}
	for i := 0; i < in.NumField(); i++ {
		fieldVal := out.FieldByName(in.Type().Field(i).Name)
		if fieldVal.IsValid() && fieldVal.CanSet() && fieldVal.Type() == in.Field(i).Type() {
			fieldVal.Set(in.Field(i))
		}
	}
	return outPointer, true
}
