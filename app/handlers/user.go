package handlers

import (
	"dgut-icourse/ent"
	"github.com/gin-gonic/gin"
)

// GetMe
// @Description: 获取当前用户信息
// @return gin.HandlerFunc:
func GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := Response{0, "success", nil, nil}
		// 从上下文中获取用户信息
		if user, ok := c.Get("user"); !ok || user == nil {
			resp.Code = 401
			resp.Msg = "unauthorized"
		} else {
			user := user.(*ent.User)
			resp.Data = gin.H{
				"id":        user.ID,
				"nick_name": user.NickName,
				"avatar":    user.Icon,
			}
		}
		ResponseOK(c, resp)
		return
	}
}
