package encoder

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

const jsonTag = "json"

type objectJsonMarshaller struct {
	obj interface{}
}

func NewObjectJsonMarshaller(obj interface{}) *objectJsonMarshaller {
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
			key := marshaller.getFieldName(field)
			value := marshaller.getValueInterface(sValue, i)
			enc.InterfaceKey(key, value)
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
		NewObjectJsonMarshaller(sValue.Interface()).MarshalJSONObject(enc)
	}

}

func (marshaller objectJsonMarshaller) getFieldName(field reflect.StructField) string {

	if jsonTagValue, ok := field.Tag.Lookup(jsonTag); ok {
		return strings.Split(jsonTagValue, ",")[0]
	}
	return field.Name
}

func (marshaller objectJsonMarshaller) getValueInterface(objValue reflect.Value, fieldIndex int) (valueInterface interface{}) {
	defer func() {
		if err := recover(); err != nil {
			objPointer := reflect.New(objValue.Type()).Elem()
			objPointer.Set(objValue)

			fieldValue := objPointer.Field(fieldIndex)
			valueInterface = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem().Interface()
		}
	}()

	fieldValue := objValue.Field(fieldIndex)
	valueInterface = fieldValue.Interface()

	return valueInterface
}
