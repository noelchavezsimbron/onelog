package encoder

// EncodeObject encodes an object to JSON
func (enc *Encoder) EncodeObject(v JsonObject) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	_, err := enc.encodeObject(v)
	if err != nil {
		enc.err = err
		return err
	}
	_, err = enc.Write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}

// EncodeObjectKeys encodes an object to JSON
func (enc *Encoder) EncodeObjectKeys(v JsonObject, keys []string) error {
	if enc.isPooled == 1 {
		panic(InvalidUsagePooledEncoderError("Invalid usage of pooled encoder"))
	}
	enc.hasKeys = true
	enc.keys = keys
	_, err := enc.encodeObject(v)
	if err != nil {
		enc.err = err
		return err
	}
	_, err = enc.Write()
	if err != nil {
		enc.err = err
		return err
	}
	return nil
}

func (enc *Encoder) encodeObject(v JsonObject) ([]byte, error) {
	enc.grow(512)
	enc.writeByte('{')
	if v!=nil {
		v.MarshalJSONObject(enc)
	}
	if enc.hasKeys {
		enc.hasKeys = false
		enc.keys = nil
	}
	enc.writeByte('}')
	return enc.buf, enc.err
}

// AddObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObject(v JsonObject) {
	enc.Object(v)
}

// AddObjectOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectOmitEmpty(v JsonObject) {
	enc.ObjectOmitEmpty(v)
}

// AddObjectNullEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectNullEmpty(v JsonObject) {
	enc.ObjectNullEmpty(v)
}

// AddObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectKey(key string, v JsonObject) {
	enc.ObjectKey(key, v)
}

// AddObjectKeyOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectKeyOmitEmpty(key string, v JsonObject) {
	enc.ObjectKeyOmitEmpty(key, v)
}

// AddObjectKeyNullEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) AddObjectKeyNullEmpty(key string, v JsonObject) {
	enc.ObjectKeyNullEmpty(key, v)
}

// JsonObject adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) Object(v JsonObject) {
	if v==nil {
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
	enc.writeByte('{')
	v.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectWithKeys adds an object to be encoded, must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject. It will only encode the keys in keys.
func (enc *Encoder) ObjectWithKeys(v JsonObject, keys []string) {
	if v==nil {
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
	enc.writeByte('{')
	var origKeys = enc.keys
	var origHasKeys = enc.hasKeys
	enc.hasKeys = true
	enc.keys = keys
	v.MarshalJSONObject(enc)
	enc.hasKeys = origHasKeys
	enc.keys = origKeys
	enc.writeByte('}')
}

// ObjectOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectOmitEmpty(v JsonObject) {
	if v==nil {
		return
	}
	enc.grow(2)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	enc.writeByte('{')
	v.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectNullEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectNullEmpty(v JsonObject) {
	enc.grow(2)
	r := enc.getPreviousRune()
	if r != '[' {
		enc.writeByte(',')
	}
	if v==nil {
		enc.writeBytes(nullBytes)
		return
	}
	enc.writeByte('{')
	v.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectKey adds a struct to be encoded, must be used inside an object as it will encode a key
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectKey(key string, value JsonObject) {
	if enc.hasKeys {
		if !enc.keyExists(key) {
			return
		}
	}
	if value==nil {
		enc.grow(2 + len(key))
		r := enc.getPreviousRune()
		if r != '{' {
			enc.writeByte(',')
		}
		enc.writeByte('"')
		enc.writeStringEscape(key)
		enc.writeBytes(objKeyObj)
		enc.writeByte('}')
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	value.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectKeyWithKeys adds a struct to be encoded, must be used inside an object as it will encode a key.
// Value must implement MarshalerJSONObject. It will only encode the keys in keys.
func (enc *Encoder) ObjectKeyWithKeys(key string, value JsonObject, keys []string) {
	if enc.hasKeys {
		if !enc.keyExists(key) {
			return
		}
	}
	if value==nil {
		enc.grow(2 + len(key))
		r := enc.getPreviousRune()
		if r != '{' {
			enc.writeByte(',')
		}
		enc.writeByte('"')
		enc.writeStringEscape(key)
		enc.writeBytes(objKeyObj)
		enc.writeByte('}')
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	var origKeys = enc.keys
	var origHasKeys = enc.hasKeys
	enc.hasKeys = true
	enc.keys = keys
	value.MarshalJSONObject(enc)
	enc.hasKeys = origHasKeys
	enc.keys = origKeys
	enc.writeByte('}')
}

// ObjectKeyOmitEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectKeyOmitEmpty(key string, value JsonObject) {
	if enc.hasKeys {
		if !enc.keyExists(key) {
			return
		}
	}
	if value==nil  {
		return
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKeyObj)
	value.MarshalJSONObject(enc)
	enc.writeByte('}')
}

// ObjectKeyNullEmpty adds an object to be encoded or skips it if IsNil returns true.
// Must be used inside a slice or array encoding (does not encode a key)
// value must implement MarshalerJSONObject
func (enc *Encoder) ObjectKeyNullEmpty(key string, value JsonObject) {
	if enc.hasKeys {
		if !enc.keyExists(key) {
			return
		}
	}
	enc.grow(5 + len(key))
	r := enc.getPreviousRune()
	if r != '{' {
		enc.writeByte(',')
	}
	enc.writeByte('"')
	enc.writeStringEscape(key)
	enc.writeBytes(objKey)
	if value==nil {
		enc.writeBytes(nullBytes)
		return
	}
	enc.writeByte('{')
	value.MarshalJSONObject(enc)
	enc.writeByte('}')
}
