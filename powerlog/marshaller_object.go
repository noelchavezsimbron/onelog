package powerlog

import (
	"fmt"
	"reflect"
	"strings"
)

const jsonTag = "json"

type objectJsonMarshaller struct {
	obj interface{}
}

func newObjectJsonMarshaller(obj interface{}) *objectJsonMarshaller {
	return &objectJsonMarshaller{obj: obj}
}

func (marshaller objectJsonMarshaller) MarshalJSONObject(enc IEncoder) {
	obj := marshaller.obj

	if obj == nil {
		return
	}

	sType := reflect.TypeOf(obj)
	sValue := reflect.ValueOf(obj)

	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}

	if sType.Kind() == reflect.Struct {

		for i := 0; i < sType.NumField(); i++ {
			field := sType.Field(i)
			value := sValue.Field(i)
			key := marshaller.getFieldName(field)
			marshaller.valueKey(key, enc, value)
		}

	}

	if sType.Kind() == reflect.Map {

		for _, keyItem := range sValue.MapKeys() {
			key := fmt.Sprintf("%s", keyItem.Interface())
			value := sValue.MapIndex(keyItem)
			marshaller.valueKey(key, enc, value)
		}
	}

	if sType.Kind() == reflect.Interface {
		newObjectJsonMarshaller(sValue.Interface()).MarshalJSONObject(enc)
	}

}

func (marshaller objectJsonMarshaller) valueKey(keyName string, enc IEncoder, value reflect.Value) {

	if value.Interface() == nil {
		return
	}

	if fieldEncoder, ok := fieldEncoders[value.Kind()]; ok {
		fieldEncoder(keyName, value, enc)
		return
	}

	defaultEncoder(keyName, value, enc)
}

func (marshaller objectJsonMarshaller) getFieldName(field reflect.StructField) string {

	if jsonTagValue, ok := field.Tag.Lookup(jsonTag); ok {
		return strings.Split(jsonTagValue, ",")[0]
	}
	return field.Name
}
