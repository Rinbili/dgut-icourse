package handlers

import (
	"dgut-icourse/app/services"
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
			resp.Code = 40001
			resp.Msg = "unauthorized. "
			ResponseUnauthorized(c, resp)
			return
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

// UpdateMeHandler
// @Description: 更新当前用户信息
// @return gin.HandlerFunc:
func UpdateMeHandler() gin.HandlerFunc {
	var userReq ent.User
	var userResp *ent.User
	resp := Response{0, "success", nil, nil}
	return func(c *gin.Context) {
		if resp.err = c.ShouldBindJSON(&userReq); resp.err != nil {
			resp.Code = 10001
			resp.Msg = "invalid json"
			ResponseBadRequest(c, resp)
			return
		} else if userResp, resp.err = services.UpdateMe(c, userReq); resp.err != nil {
			resp.Code = 10002
			resp.Msg = "update user failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = userResp
		ResponseOK(c, resp)
		return
	}
}
