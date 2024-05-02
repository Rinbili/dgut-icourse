package handlers

import (
	"dgut-icourse/app/services"
	"dgut-icourse/app/utils"
	"entgo.io/ent/privacy"
	"github.com/gin-gonic/gin"
)

// AuthHandler
// @Description: 中间件：验证token
// @return gin.HandlerFunc:
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		if tokenString := c.Request.Header.Get("Authorization"); len(tokenString) != 0 {
			tokenString = tokenString[7:]
			// 解析token
			if claims, err := utils.ParseToken(tokenString); err == nil {
				// 验证token是否有效
				if u, err := services.GetUserByUUID(privacy.DecisionContext(c, privacy.Allow), claims.UId); err == nil && u.UpdatedAt <= claims.IssuedAt.Time.Unix() {
					//token有效可用
					c.Set("user", u)
					c.Next()
					return
				}
			}
		}
		c.Set("user", nil)
		c.Next()
	}
}

// LoginHandler
// @Description: 登录
// @return gin.HandlerFunc:
func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取code
		if code := c.Request.Header.Get("Authorization"); len(code) != 0 {
			// 获取openid
			if openid, err := utils.GetWechatUserInfo(code); err == nil {
				// 获取或创建用户
				if u, err := services.GetOrCreateUserByOpenID(c, openid); err == nil {
					// 生成token
					if token, err := utils.GetToken(u.ID.String()); err == nil {
						resp := Response{
							Code: 0,
							Msg:  "login success",
							Data: gin.H{
								"token": token,
							},
						}
						ResponseOK(c, resp)
						return
					}
				}
			}
		}
		ResponseUnauthorized(c, Response{
			Code: 401,
			Msg:  "login failed",
		})
	}
}
