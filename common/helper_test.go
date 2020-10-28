/**
 * @Author: 木荣
 * @Date: 2020/10/28 10:19 上午
 * @File: helper_test
 * @Package: common
 * @Description: 辅助工具单元测试
 */
package common

import (
	"testing"
)

func TestJoinString(t *testing.T) {
	var end = JoinString("s", "b", "d")
	if end == "sbd" {
		t.Log("TestJoinString Success")
	} else {
		t.Error("TestJoinString Error")
	}
}

func TestStructConvertMap(t *testing.T) {
	type testStruct struct {
		A string `json:"a"`
		B int    `json:"b"`
		C string `json:"c"`
	}
	s := testStruct{
		A: "test",
		B: 123,
		C: "test123",
	}

	target := map[string]interface{}{
		"a": "test",
		"b": 123,
		"c": "test123",
	}

	result := StructConvertMap(s)
	for key, value := range target {
		if result[key] != value {
			t.Error("TestStructConvertMap Error")
			return
		}
	}
	t.Log("TestStructConvertMap Success")
}
