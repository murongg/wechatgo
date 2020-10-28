/**
 * @Author: 木荣
 * @Date: 2020/10/28 10:10 上午
 * @File: http_client_test
 * @Package: common
 * @Description: http 单元测试
 */
package common

import (
	"testing"
)

func TestGet(t *testing.T) {
	h := NewHttpSend(GetUrlBuild("http://www.baidu.com", map[string]string{"name": "xiaochuan"}))
	_, err := h.Get()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func TestPost(t *testing.T) {
	h := NewHttpSend("http://www.baidu.com")
	h.SetBody(map[string]string{"name": "xiaochuan"})
	_, err := h.Post()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func TestJson(t *testing.T) {
	h := NewHttpSend("http://www.baidu.com")
	h.SetSendType("JSON")
	h.SetBody(map[string]string{"name": "xiaochuan"})
	_, err := h.Post()
	if err != nil {
		t.Error("请求错误:", err)
	} else {
		t.Log("正常返回")
	}
}

func BenchmarkGET(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := NewHttpSend(GetUrlBuild("http://www.baidu.com", map[string]string{"name": "xiaochuan"}))
		_, err := h.Get()
		if err != nil {
			b.Error("请求错误:", err)
		} else {
			b.Log("正常返回")
		}
	}
}
