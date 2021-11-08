
package json_util

import (
	io "io"
)

type JsonArr []interface{}

func DecodeJsonArr(data []byte)(arr JsonArr, err error){
	arr = make(JsonArr, 0)
	err = DecodeJson(data, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func DecodeJsonArrStr(data string)(arr JsonArr, err error){
	arr = make(JsonArr, 0)
	err = DecodeJsonStr(data, &arr)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

func ReadJsonArr(r io.Reader)(arr JsonArr, err error){
	arr = make(JsonArr, 0)
	err = ReadJson(r, &arr)
	return
}

func (arr JsonArr)Append(ins ...interface{})(JsonArr){
	return arr = append(arr, ins...)
}

func (arr JsonArr)Insert(ind int, ins ...interface{})(JsonArr){
	return append(arr[0:ind], append(ins, arr[ind:len(arr)])...)
}

func (arr JsonArr)Delete(ind int)(JsonArr){
	return append(arr[0:ind], arr[ind + 1:len(arr)]...)
}

func (arr JsonArr)GetOk(ind int)(v interface{}, ok bool){
	if ind >= arr.Len() {
		return nil, false
	}
	return arr[ind], true
}

func (arr JsonArr)GetDefault(ind int, d interface{})(v interface{}){
	v, ok := arr.GetOk(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)Get(ind int)(v interface{}){
	return arr.GetDefault(ind, nil)
}

func (arr JsonArr)GetBoolOk(ind int)(v bool, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v, ok = arr[ind].(bool)
	return
}

func (arr JsonArr)GetByteOk(ind int)(v byte, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (byte)(v0) }else{
		v, ok = arr[ind].(byte)
	}
	return
}

func (arr JsonArr)GetIntOk(ind int)(v int, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (int)(v0) }else{
		v, ok = arr[ind].(int)
	}
	return
}

func (arr JsonArr)GetUIntOk(ind int)(v uint, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (uint)(v0) }else{
		v, ok = arr[ind].(uint)
	}
	return
}

func (arr JsonArr)GetInt8Ok(ind int)(v int8, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (int8)(v0) }else{
		v, ok = arr[ind].(int8)
	}
	return
}

func (arr JsonArr)GetInt16Ok(ind int)(v int16, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (int16)(v0) }else{
		v, ok = arr[ind].(int16)
	}
	return
}

func (arr JsonArr)GetInt32Ok(ind int)(v int32, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (int32)(v0) }else{
		v, ok = arr[ind].(int32)
	}
	return
}

func (arr JsonArr)GetInt64Ok(ind int)(v int64, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (int64)(v0) }else{
		v, ok = arr[ind].(int64)
	}
	return
}

func (arr JsonArr)GetUInt8Ok(ind int)(v uint8, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (uint8)(v0) }else{
		v, ok = arr[ind].(uint8)
	}
	return
}

func (arr JsonArr)GetUInt16Ok(ind int)(v uint16, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (uint16)(v0) }else{
		v, ok = arr[ind].(uint16)
	}
	return
}

func (arr JsonArr)GetUInt32Ok(ind int)(v uint32, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (uint32)(v0) }else{
		v, ok = arr[ind].(uint32)
	}
	return
}

func (arr JsonArr)GetUInt64Ok(ind int)(v uint64, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (uint64)(v0) }else{
		v, ok = arr[ind].(uint64)
	}
	return
}

func (arr JsonArr)GetFloat32Ok(ind int)(v float32, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(float64)
	if ok { v = (float32)(v0) }else{
		v, ok = arr[ind].(float32)
	}
	return
}

func (arr JsonArr)GetFloat64Ok(ind int)(v float64, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v, ok = arr[ind].(float64)
	return
}

func (arr JsonArr)GetStringOk(ind int)(v string, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v, ok = arr[ind].(string)
	return
}

func (arr JsonArr)GetArrayOk(ind int)(v JsonArr, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].([]interface{})
	if ok {
		v = (JsonArr)(v0)
	}
	return
}

func (arr JsonArr)GetBytesOk(ind int)(v []byte, ok bool){
	a, ok := arr.GetArrayOk(ind)
	if !ok { return }
	v = make([]byte, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetByteOk(i)
		if !ok { return }
	}
	return
}

func (arr JsonArr)GetStringsOk(ind int)(v []string, ok bool){
	a, ok := arr.GetArrayOk(ind)
	if !ok { return }
	v = make([]string, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetStringOk(i)
		if !ok { return }
	}
	return
}

func (arr JsonArr)GetArrays(ind int)(v []JsonArr, ok bool){
	a, ok := arr.GetArrayOk(ind)
	if !ok { return }
	v = make([]JsonArr, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetArrayOk(i)
		if !ok { return }
	}
	return
}

func (arr JsonArr)GetObjOk(ind int)(v JsonObj, ok bool){
	if ind >= arr.Len() {
		ok = false
		return
	}
	v0, ok := arr[ind].(map[string]interface{})
	if ok {
		v = (JsonObj)(v0)
	}
	return
}

func (arr JsonArr)GetObjsOk(ind int)(v []JsonObj, ok bool){
	a, ok := arr.GetArrayOk(ind)
	if !ok { return }
	v = make([]JsonObj, a.Len())
	for i, _ := range a {
		v[i], ok = a.GetObjOk(i)
		if !ok { return }
	}
	return
}

func (arr JsonArr)GetStringMap(ind int)(v map[string]string, ok bool){
	v, ok = arr[ind].(map[string]string)
	if ok { return }
	v0, ok := arr.GetObjOk(ind)
	if !ok { return }
	v = make(map[string]string)
	for k, _ := range v0 {
		v[k], ok = v0.GetStringOk(k)
		if !ok { return }
	}
	return
}


func (arr JsonArr)GetBoolDefault(ind int, d bool)(v bool){
	v, ok := arr.GetBoolOk(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetByteDefault(ind int, d byte)(v byte){
	v, ok := arr.GetByteOk(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetIntDefault(ind int, d int)(v int){
	v, ok := arr.GetIntOk(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetUIntDefault(ind int, d uint)(v uint){
	v, ok := arr.GetUIntOk(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetInt8Default(ind int, d int8)(v int8){
	v, ok := arr.GetInt8Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetInt16Default(ind int, d int16)(v int16){
	v, ok := arr.GetInt16Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetInt32Default(ind int, d int32)(v int32){
	v, ok := arr.GetInt32Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetInt64Default(ind int, d int64)(v int64){
	v, ok := arr.GetInt64Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetUInt8Default(ind int, d uint8)(v uint8){
	v, ok := arr.GetUInt8Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetUInt16Default(ind int, d uint16)(v uint16){
	v, ok := arr.GetUInt16Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetUInt32Default(ind int, d uint32)(v uint32){
	v, ok := arr.GetUInt32Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetUInt64Default(ind int, d uint64)(v uint64){
	v, ok := arr.GetUInt64Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetFloat32Default(ind int, d float32)(v float32){
	v, ok := arr.GetFloat32Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetFloat64Default(ind int, d float64)(v float64){
	v, ok := arr.GetFloat64Ok(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetStringDefault(ind int, d string)(v string){
	v, ok := arr.GetStringOk(ind)
	if !ok { return d }
	return
}

func (arr JsonArr)GetBool(ind int)(v bool){
	return arr.GetBoolDefault(ind, false)
}

func (arr JsonArr)GetByte(ind int)(v byte){
	return arr.GetByteDefault(ind, 0)
}

func (arr JsonArr)GetInt(ind int)(v int){
	return arr.GetIntDefault(ind, 0)
}

func (arr JsonArr)GetUInt(ind int)(v uint){
	return arr.GetUIntDefault(ind, 0)
}

func (arr JsonArr)GetInt8(ind int)(v int8){
	return arr.GetInt8Default(ind, 0)
}

func (arr JsonArr)GetInt16(ind int)(v int16){
	return arr.GetInt16Default(ind, 0)
}

func (arr JsonArr)GetInt32(ind int)(v int32){
	return arr.GetInt32Default(ind, 0)
}

func (arr JsonArr)GetInt64(ind int)(v int64){
	return arr.GetInt64Default(ind, 0)
}

func (arr JsonArr)GetUInt8(ind int)(v uint8){
	return arr.GetUInt8Default(ind, 0)
}

func (arr JsonArr)GetUInt16(ind int)(v uint16){
	return arr.GetUInt16Default(ind, 0)
}

func (arr JsonArr)GetUInt32(ind int)(v uint32){
	return arr.GetUInt32Default(ind, 0)
}

func (arr JsonArr)GetUInt64(ind int)(v uint64){
	return arr.GetUInt64Default(ind, 0)
}

func (arr JsonArr)GetFloat32(ind int)(v float32){
	return arr.GetFloat32Default(ind, 0.0)
}

func (arr JsonArr)GetFloat64(ind int)(v float64){
	return arr.GetFloat64Default(ind, 0.0)
}

func (arr JsonArr)GetString(ind int)(v string){
	return arr.GetStringDefault(ind, "")
}

func (arr JsonArr)GetArray(ind int)(v JsonArr){
	v, ok := arr.GetArrayOk(ind)
	if !ok { v = nil }
	return
}

func (arr JsonArr)GetBytes(ind int)(v []byte){
	v, ok := arr.GetBytesOk(ind)
	if !ok { v = nil }
	return
}

func (arr JsonArr)GetStrings(ind int)(v []string){
	v, ok := arr.GetStringsOk(ind)
	if !ok { v = nil }
	return
}

func (arr JsonArr)GetArrays(ind int)(v []JsonArr){
	v, ok := arr.GetArraysOk(ind)
	if !ok { v = nil }
	return
}

func (arr JsonArr)GetObj(ind int)(v JsonObj){
	v, ok := arr.GetObjOk(ind)
	if !ok { v = nil }
	return
}

func (arr JsonArr)GetObjs(ind int)(v []JsonObj){
	v, ok := arr.GetObjsOk(ind)
	if !ok { v = nil }
	return
}

func (arr JsonArr)GetStringMap(ind int)(v map[string]string){
	v, ok := arr.GetStringMapOk(ind)
	if !ok { v = nil }
	return
}

func (obj JsonArr)Bytes()([]byte){
	return EncodeJson(obj)
}

func (obj JsonArr)String()(string){
	return EncodeJsonStr(obj)
}

func (obj JsonArr)Len()(int){
	return len(obj)
}

func (obj JsonArr)Cap()(int){
	return cap(obj)
}


