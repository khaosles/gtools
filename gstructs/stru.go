package gstructs

import (
	"encoding/json"
	"reflect"
)

/*
   @File: stru.go
   @Author: khaosles
   @Time: 2023/3/5 12:26
   @Desc:
*/

// Assignment 把结构体1的相同字段赋值给结构体2
func Assignment[T1, T2 any](elem1 T1, elem2 *T2) {
	// 获取结构体类型和值
	t1 := reflect.TypeOf(elem1)
	v1 := reflect.ValueOf(elem1)
	t2 := reflect.TypeOf(*elem2)
	v2 := reflect.ValueOf(elem2).Elem()

	// 遍历1结构体的字段
	for i := 0; i < t1.NumField(); i++ {
		// 获取字段名称和值
		fieldName := t1.Field(i).Name
		fieldValue := v1.Field(i).Interface()
		// 判断2结构体是否有同名字段
		if _, ok := t2.FieldByName(fieldName); ok {
			// 根据字段名称设置2结构体相应字段的值
			v2.FieldByName(fieldName).Set(reflect.ValueOf(fieldValue))
		}
	}
}

// Struct2Map 将结构体转换为 Map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		if value.Kind() == reflect.Slice {
			bytes, _ := json.Marshal(value.Interface())
			data[field.Name] = string(bytes)
		} else {
			data[field.Name] = value.Interface()
		}
	}
	return data
}
