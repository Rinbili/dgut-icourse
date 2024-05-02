package utils

import (
	"dgut-icourse/config"
	"encoding/json"
	"errors"
	"net/http"
)

// SessionResult 登录凭证校验响应结构
type SessionResult struct {
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误信息
}

// GetWechatUserInfo
// @Description: 获取微信用户OpenID
// @param s: string
// @return string: string
// @return error: error
func GetWechatUserInfo(s string) (string, error) {
	if resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" + config.Config.Wechat.AppID + "&secret=" + config.Config.Wechat.AppSecret + "&js_code=" + s + "&grant_type=authorization_code"); err != nil {
		return "", err
	} else {
		var result SessionResult
		if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return "", err
		} else {
			if len(result.OpenID) != 0 {
				return result.OpenID, nil
			} else {
				// TODO: 返回错误信息
				return "", errors.New(result.ErrMsg)
			}
		}
	}
}
