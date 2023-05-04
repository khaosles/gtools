package gstru

import (
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

func StructToMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)
	objType := objValue.Type()
	// 如果传入的不是结构体指针，则直接返回空map
	if objType.Kind() != reflect.Ptr || objType.Elem().Kind() != reflect.Struct {
		return map[string]interface{}{}
	}
	data := make(map[string]interface{})
	for i := 0; i < objValue.Elem().NumField(); i++ {
		field := objType.Elem().Field(i)
		value := objValue.Elem().Field(i)
		// 如果字段是空值，则跳过
		if reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) {
			continue
		}
		data[field.Name] = value.Interface()
	}
	return data
}
