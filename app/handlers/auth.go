package handlers

import (
	"dgut-icourse/app/services"
	"dgut-icourse/app/utils"
	"dgut-icourse/ent"
	"entgo.io/ent/privacy"
	"github.com/gin-gonic/gin"
)

// AuthHandler
// @Description: 中间件：验证token
// @return gin.HandlerFunc:
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		if tokenString := c.Request.Header.Get("Authorization"); len(tokenString) > 7 {
			tokenString = tokenString[7:]
			// 解析token
			if claims, err := utils.ParseToken(tokenString); err == nil {
				// 验证token是否有效
				if u, err := services.GetUserByUUID(privacy.DecisionContext(c, privacy.Allow), claims.UId); err == nil {
					//token有效可用
					c.Set("user", u)
					c.Next()
					return
				}
			}
		}
		c.Set("user", (*ent.User)(nil))
		c.Next()
	}
}

// LoginHandler
// @Description: 登录
// @return gin.HandlerFunc:
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code, openid, token string
		var u *ent.User
		resp := Response{0, "success", nil, nil}
		// 获取code
		if code = c.Request.Header.Get("Authorization"); len(code) != 0 {
			// 获取openid
			if openid, resp.err = utils.GetWechatUserInfo(code); resp.err == nil {
				// 获取或创建用户
				if u, resp.err = services.GetOrCreateUserByOpenID(c, openid); resp.err == nil {
					// 生成token
					if token, resp.err = utils.GetToken(u.ID.String()); resp.err == nil {
						resp.Data = gin.H{
							"token": token,
						}
						ResponseOK(c, resp)
						return
					}
				}
			}
		}
		resp.Code = 40000
		resp.Msg = "login failed. "
		ResponseUnauthorized(c, resp)
	}
}
