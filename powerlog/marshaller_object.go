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
			enc.InterfaceKey(key, value.Interface())
		}

	}

	if sType.Kind() == reflect.Map {

		for _, keyItem := range sValue.MapKeys() {
			key := fmt.Sprintf("%s", keyItem.Interface())
			value := sValue.MapIndex(keyItem)
			enc.InterfaceKey(key, value.Interface())
		}
	}

	if sType.Kind() == reflect.Interface {
		newObjectJsonMarshaller(sValue.Interface()).MarshalJSONObject(enc)
	}

}

func (marshaller objectJsonMarshaller) getFieldName(field reflect.StructField) string {

	if jsonTagValue, ok := field.Tag.Lookup(jsonTag); ok {
		return strings.Split(jsonTagValue, ",")[0]
	}
	return field.Name
}
