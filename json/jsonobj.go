
package json_util

import (
	io "io"
)

type JsonObj map[string]interface{}

func DecodeJsonObj(data []byte)(obj JsonObj, err error){
	obj = make(JsonObj)
	err = DecodeJson(data, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func DecodeJsonObjStr(data string)(obj JsonObj, err error){
	obj = make(JsonObj)
	err = DecodeJsonStr(data, &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func ReadJsonObj(r io.Reader)(obj JsonObj, err error){
	obj = make(JsonObj)
	err = ReadJson(r, &obj)
	return
}

func (obj JsonObj)Has(key string)(ok bool){
	_, ok = obj[key]
	return
}

func (obj JsonObj)GetOk(key string)(v interface{}, ok bool){
	if ind >= obj.Len() {
		return nil, false
	}
	return obj[key], true
}

func (obj JsonObj)GetDefault(key string, d interface{})(v interface{}){
	v, ok := obj.GetOk(key)
	if !ok { return d }
	return
}

func (obj JsonObj)Get(key string)(v interface{}){
	return obj.GetDefault(key, nil)
}

func (obj JsonObj)GetBoolOk(key string)(v bool, ok bool){
	v, ok = obj[key].(bool)
	return
}

func (obj JsonObj)GetByteOk(key string)(v byte, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (byte)(v0) }else{
		v, ok = obj[key].(byte)
	}
	return
}

func (obj JsonObj)GetIntOk(key string)(v int, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (int)(v0) }else{
		v, ok = obj[key].(int)
	}
	return
}

func (obj JsonObj)GetUIntOk(key string)(v uint, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (uint)(v0) }else{
		v, ok = obj[key].(uint)
	}
	return
}

func (obj JsonObj)GetInt8Ok(key string)(v int8, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (int8)(v0) }else{
		v, ok = obj[key].(int8)
	}
	return
}

func (obj JsonObj)GetInt16Ok(key string)(v int16, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (int16)(v0) }else{
		v, ok = obj[key].(int16)
	}
	return
}

func (obj JsonObj)GetInt32Ok(key string)(v int32, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (int32)(v0) }else{
		v, ok = obj[key].(int32)
	}
	return
}

func (obj JsonObj)GetInt64Ok(key string)(v int64, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (int64)(v0) }else{
		v, ok = obj[key].(int64)
	}
	return
}

func (obj JsonObj)GetUInt8Ok(key string)(v uint8, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (uint8)(v0) }else{
		v, ok = obj[key].(uint8)
	}
	return
}

func (obj JsonObj)GetUInt16Ok(key string)(v uint16, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (uint16)(v0) }else{
		v, ok = obj[key].(uint16)
	}
	return
}

func (obj JsonObj)GetUInt32Ok(key string)(v uint32, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (uint32)(v0) }else{
		v, ok = obj[key].(uint32)
	}
	return
}

func (obj JsonObj)GetUInt64Ok(key string)(v uint64, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (uint64)(v0) }else{
		v, ok = obj[key].(uint64)
	}
	return
}

func (obj JsonObj)GetFloat32Ok(key string)(v float32, ok bool){
	v0, ok := obj[key].(float64)
	if ok { v = (float32)(v0) }else{
		v, ok = obj[key].(float32)
	}
	return
}

func (obj JsonObj)GetFloat64Ok(key string)(v float64, ok bool){
	v, ok = obj[key].(float64)
	return
}

func (obj JsonObj)GetStringOk(key string)(v string, ok bool){
	v, ok = obj[key].(string)
	return
}

func (obj JsonObj)GetArrayOk(key string)(v JsonArr, ok bool){
	v0, ok := obj[key].([]interface{})
	if ok {
		v = (JsonArr)(v0)
	}
	return
}

func (obj JsonObj)GetBytesOk(key string)(v []byte, ok bool){
	a, ok := obj.GetArrayOk(key)
	if !ok { return }
	v = make([]byte, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetByteOk(i)
		if !ok { return }
	}
	return
}

func (obj JsonObj)GetStringsOk(key string)(v []string, ok bool){
	a, ok := obj.GetArrayOk(key)
	if !ok { return }
	v = make([]string, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetStringOk(i)
		if !ok { return }
	}
	return
}

func (obj JsonObj)GetArrays(key string)(v []JsonArr, ok bool){
	a, ok := obj.GetArrayOk(key)
	if !ok { return }
	v = make([]JsonArr, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetArrayOk(i)
		if !ok { return }
	}
	return
}

func (obj JsonObj)GetObjOk(key string)(v JsonObj, ok bool){
	v0, ok := obj[key].(map[string]interface{})
	if ok {
		v = (JsonObj)(v0)
	}
	return
}

func (obj JsonObj)GetObjsOk(key string)(v []JsonObj, ok bool){
	a, ok := obj.GetArrayOk(key)
	if !ok { return }
	v = make([]JsonObj, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetObjOk(i)
		if !ok { return }
	}
	return
}

func (obj JsonObj)GetStringMap(key string)(v map[string]string, ok bool){
	v, ok = obj[key].(map[string]string)
	if ok { return }
	v0, ok := obj.GetObjOk(key)
	if !ok { return }
	v = make(map[string]string)
	for k, _ := range v0 {
		v[k], ok = v0.GetStringOk(k)
		if !ok { return }
	}
	return
}


func (obj JsonObj)GetBoolDefault(key string, d bool)(v bool){
	v, ok := obj.GetBoolOk(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetByteDefault(key string, d byte)(v byte){
	v, ok := obj.GetByteOk(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetIntDefault(key string, d int)(v int){
	v, ok := obj.GetIntOk(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetUIntDefault(key string, d uint)(v uint){
	v, ok := obj.GetUIntOk(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetInt8Default(key string, d int8)(v int8){
	v, ok := obj.GetInt8Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetInt16Default(key string, d int16)(v int16){
	v, ok := obj.GetInt16Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetInt32Default(key string, d int32)(v int32){
	v, ok := obj.GetInt32Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetInt64Default(key string, d int64)(v int64){
	v, ok := obj.GetInt64Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetUInt8Default(key string, d uint8)(v uint8){
	v, ok := obj.GetUInt8Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetUInt16Default(key string, d uint16)(v uint16){
	v, ok := obj.GetUInt16Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetUInt32Default(key string, d uint32)(v uint32){
	v, ok := obj.GetUInt32Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetUInt64Default(key string, d uint64)(v uint64){
	v, ok := obj.GetUInt64Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetFloat32Default(key string, d float32)(v float32){
	v, ok := obj.GetFloat32Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetFloat64Default(key string, d float64)(v float64){
	v, ok := obj.GetFloat64Ok(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetStringDefault(key string, d string)(v string){
	v, ok := obj.GetStringOk(key)
	if !ok { return d }
	return
}

func (obj JsonObj)GetBool(key string)(v bool){
	return obj.GetBoolDefault(key, false)
}

func (obj JsonObj)GetByte(key string)(v byte){
	return obj.GetByteDefault(key, 0)
}

func (obj JsonObj)GetInt(key string)(v int){
	return obj.GetIntDefault(key, 0)
}

func (obj JsonObj)GetUInt(key string)(v uint){
	return obj.GetUIntDefault(key, 0)
}

func (obj JsonObj)GetInt8(key string)(v int8){
	return obj.GetInt8Default(key, 0)
}

func (obj JsonObj)GetInt16(key string)(v int16){
	return obj.GetInt16Default(key, 0)
}

func (obj JsonObj)GetInt32(key string)(v int32){
	return obj.GetInt32Default(key, 0)
}

func (obj JsonObj)GetInt64(key string)(v int64){
	return obj.GetInt64Default(key, 0)
}

func (obj JsonObj)GetUInt8(key string)(v uint8){
	return obj.GetUInt8Default(key, 0)
}

func (obj JsonObj)GetUInt16(key string)(v uint16){
	return obj.GetUInt16Default(key, 0)
}

func (obj JsonObj)GetUInt32(key string)(v uint32){
	return obj.GetUInt32Default(key, 0)
}

func (obj JsonObj)GetUInt64(key string)(v uint64){
	return obj.GetUInt64Default(key, 0)
}

func (obj JsonObj)GetFloat32(key string)(v float32){
	return obj.GetFloat32Default(key, 0.0)
}

func (obj JsonObj)GetFloat64(key string)(v float64){
	return obj.GetFloat64Default(key, 0.0)
}

func (obj JsonObj)GetString(key string)(v string){
	return obj.GetStringDefault(key, "")
}

func (obj JsonObj)GetArray(key string)(v JsonArr){
	v, ok := obj.GetArrayOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)GetBytes(key string)(v []byte){
	v, ok := obj.GetBytesOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)GetStrings(key string)(v []string){
	v, ok := obj.GetStringsOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)GetArrays(key string)(v []JsonArr){
	v, ok := obj.GetArraysOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)GetObj(key string)(v JsonObj){
	v, ok := obj.GetObjOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)GetObjs(key string)(v []JsonObj){
	v, ok := obj.GetObjsOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)GetStringMap(key string)(v map[string]string){
	v, ok := obj.GetStringMapOk(key)
	if !ok { v = nil }
	return
}

func (obj JsonObj)Size()(int){
	return len(obj)
}

func (obj JsonObj)Bytes()([]byte){
	return EncodeJson(obj)
}

func (obj JsonObj)String()(string){
	return EncodeJsonStr(obj)
}
