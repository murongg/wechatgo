/**
 * @Author: 木荣
 * @Date: 2020/10/28 10:19 上午
 * @File: helper
 * @Package: common
 * @Description: 辅助工具
 */
package common

import "strings"

func JoinString(s []string) string {
	var r string
	var build strings.Builder
	for _, value := range s {
		build.WriteString(value)
	}
	r = build.String()
	return r
}
