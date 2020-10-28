/**
 * @Author: 木荣
 * @Date: 2020/10/28 10:30 上午
 * @File: mini
 * @Package: constant
 * @Description: 微信小程序 url
 */
package constant

import "github.com/murongg/wechatgo/common"

var (
	CODE_TO_SESSION   = common.JoinString([]string{BASE_URL, "/sns/jscode2session"}) // 登录凭证校验
	GET_PAID_UNION_ID = common.JoinString([]string{BASE_URL, "/wxa/getpaidunionid"}) // 获取该用户的 UnionId
	GET_ACCESS_TOKEN  = common.JoinString([]string{BASE_URL, "/cgi-bin/token"})      // 获取小程序全局唯一后台接口调用凭据
)
