package powerlog

import (
	"reflect"
)

var fieldEncoders map[reflect.Kind]fieldEncoder
var elementEncoders map[reflect.Kind]elementEncoder

type fieldEncoder func(keyName string, value reflect.Value, enc IEncoder)

var defaultEncoder fieldEncoder

type elementEncoder func(enc IEncoder, value reflect.Value)

func init() {
	fieldEncoders = make(map[reflect.Kind]fieldEncoder)
	fieldEncoders[reflect.String] = encoderString
	fieldEncoders[reflect.Int] = encoderInt
	fieldEncoders[reflect.Int8] = encoderInt
	fieldEncoders[reflect.Int16] = encoderInt
	fieldEncoders[reflect.Int32] = encoderInt
	fieldEncoders[reflect.Int64] = encoderInt
	fieldEncoders[reflect.Float32] = encoderFloat
	fieldEncoders[reflect.Float64] = encoderFloat
	fieldEncoders[reflect.Bool] = encoderBool
	fieldEncoders[reflect.Struct] = encoderStruct
	fieldEncoders[reflect.Map] = encoderMap
	fieldEncoders[reflect.Slice] = encoderSlice
	fieldEncoders[reflect.Interface] = encoderInterface

	defaultEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
		enc.ObjectKey(keyName, newObjectJsonMarshaller(value.Interface()))
	}

	//----------------

	elementEncoders = make(map[reflect.Kind]elementEncoder)
	elementEncoders[reflect.String] = elementEncoderString
	elementEncoders[reflect.Int] = elementEncoderInt
	elementEncoders[reflect.Int8] = elementEncoderInt
	elementEncoders[reflect.Int16] = elementEncoderInt
	elementEncoders[reflect.Int32] = elementEncoderInt
	elementEncoders[reflect.Int64] = elementEncoderInt
	elementEncoders[reflect.Float32] = elementEncoderFloat
	elementEncoders[reflect.Float64] = elementEncoderFloat
	elementEncoders[reflect.Bool] = elementEncoderBool
	elementEncoders[reflect.Struct] = elementEncoderStruct
	elementEncoders[reflect.Map] = elementEncoderMap
	elementEncoders[reflect.Slice] = elementEncoderSlice
	elementEncoders[reflect.Interface] = elementEncoderInterface

}

var encoderString fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.StringKey(keyName, value.String())
}

var encoderInt fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.Int64Key(keyName, value.Int())
}

var encoderFloat fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.Float64Key(keyName, value.Float())
}

var encoderBool fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.BoolKey(keyName, value.Bool())
}

var encoderStruct fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.ObjectKey(keyName, newObjectJsonMarshaller(value.Interface()))
}

var encoderMap fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.ObjectKey(keyName, newObjectJsonMarshaller(value.Interface()))
}

var encoderSlice fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.ArrayKey(keyName, newArrayJsonMarshaller(value.Interface()))
}

var encoderInterface fieldEncoder = func(keyName string, value reflect.Value, enc IEncoder) {
	enc.InterfaceKey(keyName, value.Interface())
}

// ----------------------------

var elementEncoderString elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.AddString(value.String())
}

var elementEncoderInt elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.AddInt64(value.Int())
}

var elementEncoderFloat elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.AddFloat(value.Float())
}

var elementEncoderBool elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.AddBool(value.Bool())
}

var elementEncoderStruct elementEncoder = func(enc IEncoder, value reflect.Value) {

	enc.AddObject(newObjectJsonMarshaller(value.Interface()))
}

var elementEncoderMap elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.AddObject(newObjectJsonMarshaller(value.Interface()))
}

var elementEncoderSlice elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.Array(newArrayJsonMarshaller(value.Interface()))
}

var elementEncoderInterface elementEncoder = func(enc IEncoder, value reflect.Value) {
	enc.AddInterface(value.Interface())
}
