package powerlog

import "reflect"

type arrayJsonMarshaller struct {
	slice []interface{}
}

func newArrayJsonMarshaller(obj interface{}) *arrayJsonMarshaller {
	marshaller := &arrayJsonMarshaller{}
	marshaller.fill(obj)
	return marshaller
}

func (json arrayJsonMarshaller) MarshalJSONArray(enc IEncoder) {

	for _, v := range json.slice {

		sType := reflect.TypeOf(v)
		sValue := reflect.ValueOf(v)

		if sType.Kind() == reflect.Ptr {
			sType = sType.Elem()
			sValue = sValue.Elem()
		}

		if fieldEncoder, ok := elementEncoders[sValue.Kind()]; ok {
			fieldEncoder(enc, sValue)
		}
	}

}

func (json arrayJsonMarshaller) IsNil() bool {
	return json.slice == nil || len(json.slice) == 0
}

func (json *arrayJsonMarshaller) fill(obj interface{}) {

	if obj == nil {
		return
	}

	sType := reflect.TypeOf(obj)
	sValue := reflect.ValueOf(obj)

	if sType.Kind() == reflect.Ptr {
		sType = sType.Elem()
		sValue = sValue.Elem()
	}

	if sType.Kind() == reflect.Slice {
		if data, ok := obj.([]interface{}); ok {
			json.slice = data
			return
		}

		json.slice = make([]interface{}, sValue.Len())
		for i := 0; i < sValue.Len(); i++ {
			json.slice[i] = sValue.Index(i).Interface()
		}
	}

	return
}
