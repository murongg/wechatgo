/**
 * @Author: 木荣
 * @Date: 2020/10/28 10:19 上午
 * @File: helper
 * @Package: common
 * @Description: 辅助工具
 */
package common

import (
	"reflect"
	"strings"
)

// 拼接字符串
func JoinString(s ...string) string {
	var r string
	var build strings.Builder
	for _, value := range s {
		build.WriteString(value)
	}
	r = build.String()
	return r
}

// 结构体转map
func StructConvertMap(target interface{}) map[string]interface{} {
	t := reflect.TypeOf(target)
	v := reflect.ValueOf(target)
	var result = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tagName := t.Field(i).Tag.Get("json")
		if tagName != "" && tagName != "-" {
			result[tagName] = v.Field(i).Interface()
		}
	}
	return result
}
