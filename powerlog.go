// Package powerlog is a fast, low allocation and modular JSON logger.
//
// Basic usage:
// 	import "github.com/noelchavezsimbron/powerlog/log"
//
//	log.Info("hello world !") // {"level":"info","message":"hello world !", "time":1494567715}
//
// You can create your own logger:
//	import "github.com/noelchavezsimbron/powerlog"
//
//	var logger = onelog.New(os.Stdout, onelog.ALL)
//
//	func main() {
//		logger.Info("hello world !") // {"level":"info","message":"hello world !"}
//	}
package powerlog

type encoderFunc func(IEncoder)

type EmbeddedJSON []byte

type JsonObject interface {
	MarshalJSONObject(enc IEncoder)
}
type JsonArray interface {
	MarshalJSONArray(enc IEncoder)
	IsNil() bool
}

type ObjectBuilder encoderFunc

func (f ObjectBuilder) MarshalJSONObject(enc IEncoder) {
	f(enc)
}

type IEncoder interface {
	AppendBytes(b []byte)
	AppendByte(b byte)
	Buf() []byte
	Write() (int, error)
	Release()
	EncodeString(s string) error
	AppendString(v string)
	AddString(v string)
	AddStringOmitEmpty(v string)
	AddStringNullEmpty(v string)
	AddStringKey(key, v string)
	AddStringKeyOmitEmpty(key, v string)
	AddStringKeyNullEmpty(key, v string)
	String(v string)
	StringOmitEmpty(v string)
	StringNullEmpty(v string)
	StringKey(key, v string)
	StringKeyOmitEmpty(key, v string)
	StringKeyNullEmpty(key, v string)
	EncodeInt(n int) error
	EncodeInt64(n int64) error
	AddInt(v int)
	AddIntOmitEmpty(v int)
	AddIntNullEmpty(v int)
	Int(v int)
	IntOmitEmpty(v int)
	IntNullEmpty(v int)
	AddIntKey(key string, v int)
	AddIntKeyOmitEmpty(key string, v int)
	AddIntKeyNullEmpty(key string, v int)
	IntKey(key string, v int)
	IntKeyOmitEmpty(key string, v int)
	IntKeyNullEmpty(key string, v int)
	AddInt64(v int64)
	AddInt64OmitEmpty(v int64)
	AddInt64NullEmpty(v int64)
	Int64(v int64)
	Int64OmitEmpty(v int64)
	Int64NullEmpty(v int64)
	AddInt64Key(key string, v int64)
	AddInt64KeyOmitEmpty(key string, v int64)
	AddInt64KeyNullEmpty(key string, v int64)
	Int64Key(key string, v int64)
	Int64KeyOmitEmpty(key string, v int64)
	Int64KeyNullEmpty(key string, v int64)
	AddInt32(v int32)
	AddInt32OmitEmpty(v int32)
	AddInt32NullEmpty(v int32)
	Int32(v int32)
	Int32OmitEmpty(v int32)
	Int32NullEmpty(v int32)
	AddInt32Key(key string, v int32)
	AddInt32KeyOmitEmpty(key string, v int32)
	Int32Key(key string, v int32)
	Int32KeyOmitEmpty(key string, v int32)
	Int32KeyNullEmpty(key string, v int32)
	AddInt16(v int16)
	AddInt16OmitEmpty(v int16)
	Int16(v int16)
	Int16OmitEmpty(v int16)
	Int16NullEmpty(v int16)
	AddInt16Key(key string, v int16)
	AddInt16KeyOmitEmpty(key string, v int16)
	AddInt16KeyNullEmpty(key string, v int16)
	Int16Key(key string, v int16)
	Int16KeyOmitEmpty(key string, v int16)
	Int16KeyNullEmpty(key string, v int16)
	AddInt8(v int8)
	AddInt8OmitEmpty(v int8)
	AddInt8NullEmpty(v int8)
	Int8(v int8)
	Int8OmitEmpty(v int8)
	Int8NullEmpty(v int8)
	AddInt8Key(key string, v int8)
	AddInt8KeyOmitEmpty(key string, v int8)
	AddInt8KeyNullEmpty(key string, v int8)
	Int8Key(key string, v int8)
	Int8KeyOmitEmpty(key string, v int8)
	Int8KeyNullEmpty(key string, v int8)
	EncodeFloat(n float64) error
	EncodeFloat32(n float32) error
	AddFloat(v float64)
	AddFloatOmitEmpty(v float64)
	AddFloatNullEmpty(v float64)
	Float(v float64)
	FloatOmitEmpty(v float64)
	FloatNullEmpty(v float64)
	AddFloatKey(key string, v float64)
	AddFloatKeyOmitEmpty(key string, v float64)
	AddFloatKeyNullEmpty(key string, v float64)
	FloatKey(key string, v float64)
	FloatKeyOmitEmpty(key string, v float64)
	FloatKeyNullEmpty(key string, v float64)
	AddFloat64(v float64)
	AddFloat64OmitEmpty(v float64)
	Float64(v float64)
	Float64OmitEmpty(v float64)
	Float64NullEmpty(v float64)
	AddFloat64Key(key string, v float64)
	AddFloat64KeyOmitEmpty(key string, v float64)
	Float64Key(key string, value float64)
	Float64KeyOmitEmpty(key string, v float64)
	Float64KeyNullEmpty(key string, v float64)
	AddFloat32(v float32)
	AddFloat32OmitEmpty(v float32)
	AddFloat32NullEmpty(v float32)
	Float32(v float32)
	Float32OmitEmpty(v float32)
	Float32NullEmpty(v float32)
	AddFloat32Key(key string, v float32)
	AddFloat32KeyOmitEmpty(key string, v float32)
	AddFloat32KeyNullEmpty(key string, v float32)
	Float32Key(key string, v float32)
	Float32KeyOmitEmpty(key string, v float32)
	Float32KeyNullEmpty(key string, v float32)
	EncodeBool(v bool) error
	AddBool(v bool)
	AddBoolOmitEmpty(v bool)
	AddBoolNullEmpty(v bool)
	AddBoolKey(key string, v bool)
	AddBoolKeyOmitEmpty(key string, v bool)
	AddBoolKeyNullEmpty(key string, v bool)
	Bool(v bool)
	BoolOmitEmpty(v bool)
	BoolNullEmpty(v bool)
	BoolKey(key string, value bool)
	BoolKeyOmitEmpty(key string, v bool)
	BoolKeyNullEmpty(key string, v bool)
	EncodeObject(v JsonObject) error
	EncodeObjectKeys(v JsonObject, keys []string) error
	AddObject(v JsonObject)
	AddObjectOmitEmpty(v JsonObject)
	AddObjectNullEmpty(v JsonObject)
	AddObjectKey(key string, v JsonObject)
	AddObjectKeyOmitEmpty(key string, v JsonObject)
	AddObjectKeyNullEmpty(key string, v JsonObject)
	Object(v JsonObject)
	ObjectWithKeys(v JsonObject, keys []string)
	ObjectOmitEmpty(v JsonObject)
	ObjectNullEmpty(v JsonObject)
	ObjectKey(key string, value JsonObject)
	ObjectKeyWithKeys(key string, value JsonObject, keys []string)
	ObjectKeyOmitEmpty(key string, value JsonObject)
	ObjectKeyNullEmpty(key string, value JsonObject)
	EncodeArray(v JsonArray) error
	AddArray(v JsonArray)
	AddArrayOmitEmpty(v JsonArray)
	AddArrayNullEmpty(v JsonArray)
	AddArrayKey(key string, v JsonArray)
	AddArrayKeyOmitEmpty(key string, v JsonArray)
	AddArrayKeyNullEmpty(key string, v JsonArray)
	Array(v JsonArray)
	ArrayOmitEmpty(v JsonArray)
	ArrayNullEmpty(v JsonArray)
	ArrayKey(key string, v JsonArray)
	ArrayKeyOmitEmpty(key string, v JsonArray)
	ArrayKeyNullEmpty(key string, v JsonArray)
	EncodeEmbeddedJSON(v *EmbeddedJSON) error
	encodeEmbeddedJSON(v *EmbeddedJSON) ([]byte, error)
	AddEmbeddedJSON(v *EmbeddedJSON)
	AddEmbeddedJSONOmitEmpty(v *EmbeddedJSON)
	AddEmbeddedJSONKey(key string, v *EmbeddedJSON)
	AddEmbeddedJSONKeyOmitEmpty(key string, v *EmbeddedJSON)
	AddInterfaceKey(key string, v interface{})
	InterfaceKey(key string, v interface{})
	AddInterface(v interface{})
	Interface(v interface{})
}
