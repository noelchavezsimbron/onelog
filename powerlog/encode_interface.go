package powerlog

import "reflect"

func (enc *Encoder) AddInterfaceKey(key string, v interface{}) {
	enc.InterfaceKey(key, v)
}

func (enc *Encoder) InterfaceKey(key string, v interface{}) {

	value := reflect.ValueOf(v)
	if v == nil || value.Interface() == nil {
		return
	}

	if fieldEncoder, ok := fieldEncoders[value.Kind()]; ok {
		fieldEncoder(key, value, enc)
		return
	}
	defaultEncoder(key, value, enc)
}

func (enc *Encoder) AddInterface(v interface{}) {
	enc.Interface(v)
}

func (enc *Encoder) Interface(v interface{}) {
	if v == nil {
		enc.grow(2)
		r := enc.getPreviousRune()
		if r != '{' && r != '[' {
			enc.writeByte(',')
		}
		enc.writeByte('{')
		enc.writeByte('}')
		return
	}
	enc.grow(4)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}

	vm := newObjectJsonMarshaller(v)
	enc.writeByte('{')
	vm.MarshalJSONObject(enc)
	enc.writeByte('}')
}
