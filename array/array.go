// array/map/struct助手
package array

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/spf13/cast"
)

type (
	GetIder interface {
		GetId() int
	}
)

// GetIdIndex 获取[]GetIder切片里所有ID => index的对应关系
func GetIdIndex(slice interface{}) map[int]int {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic("parameters can only be slices")
	}
	m := make(map[int]int)
	for i := 0; i < value.Len(); i++ {
		v, ok := value.Index(i).Interface().(GetIder)
		if !ok {
			panic("The slice element must implement the GetIder interface")
		}
		m[v.GetId()] = i
	}
	return m
}

// GetIds 获取[]GetIder切片里所有ID的集合
func GetIds(slice interface{}) []int {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic("parameters can only be slices")
	}
	m := make([]int, 0)
	for i := 0; i < value.Len(); i++ {
		v, ok := value.Index(i).Interface().(GetIder)
		if !ok {
			panic("The slice element must implement the GetIder interface")
		}
		m = append(m, v.GetId())
	}
	return m
}

// Keys 获取一个Map/struct的键
func Keys(out interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(out))
	valueType := value.Type()
	if value.Kind() == reflect.Map {
		keys := value.MapKeys()
		length := len(keys)
		resultSlice := reflect.MakeSlice(reflect.SliceOf(valueType.Key()), length, length)
		for i, key := range keys {
			resultSlice.Index(i).Set(key)
		}
		return resultSlice.Interface()
	}
	if value.Kind() == reflect.Struct {
		length := value.NumField()
		resultSlice := make([]string, length)
		for i := 0; i < length; i++ {
			resultSlice[i] = valueType.Field(i).Name
		}
		return resultSlice
	}
	panic(fmt.Sprintf("Type %s is not supported by Keys", valueType.String()))
}

// Values 获取map/struct的值
func Values(out interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(out))
	valueType := value.Type()
	if value.Kind() == reflect.Map {
		keys := value.MapKeys()
		length := len(keys)
		resultSlice := reflect.MakeSlice(reflect.SliceOf(valueType.Elem()), length, length)
		for i, key := range keys {
			resultSlice.Index(i).Set(value.MapIndex(key))
		}
		return resultSlice.Interface()
	}
	if value.Kind() == reflect.Struct {
		length := value.NumField()
		resultSlice := make([]interface{}, length)
		for i := 0; i < length; i++ {
			resultSlice[i] = value.Field(i).Interface()
		}
		return resultSlice
	}
	panic(fmt.Sprintf("Type %s is not supported by Keys", valueType.String()))
}

func redirectValue(value reflect.Value) reflect.Value {
	for {
		if !value.IsValid() || value.Kind() != reflect.Ptr {
			return value
		}
		res := reflect.Indirect(value)
		if res.Kind() == reflect.Ptr && value.Pointer() == res.Pointer() {
			return value
		}
		value = res
	}
}

func MapToString(params map[string]interface{}) (returnStr string) {
	keys := make([]string, 0, len(params))
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Sort(sort.StringSlice(keys))
	var buf bytes.Buffer
	for _, k := range keys {
		val := params[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(cast.ToString(val))
	}
	return buf.String()
}
