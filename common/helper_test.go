/**
 * @Author: 木荣
 * @Date: 2020/10/28 10:19 上午
 * @File: helper_test
 * @Package: common
 * @Description: 辅助工具单元测试
 */
package common

import "testing"

func TestJoinString(t *testing.T) {
	var s []string = []string{"s", "b", "d"}
	var end = JoinString(s)
	if end == "sbd" {
		t.Log("TestJoinString Success")
	} else {
		t.Error("TestJoinString Error")
	}
}
