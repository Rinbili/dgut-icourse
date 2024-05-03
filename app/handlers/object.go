package handlers

import (
	"dgut-icourse/app/services"
	"dgut-icourse/ent"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateObjectHandler
// @Description: 创建对象
// @return gin.HandlerFunc
func CreateObjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var objReq ent.Object
		var objResp *ent.Object
		resp := Response{0, "success", nil, nil}
		// 获取字段
		if resp.err = c.ShouldBindJSON(&objReq); resp.err != nil {
			resp.Code = 10001
			resp.Msg = "invalid json"
			ResponseBadRequest(c, resp)
			return
		} else if objResp, resp.err = services.CreateObject(c, objReq); resp.err != nil {
			resp.Code = 10002
			resp.Msg = "create object failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = objResp
		ResponseOK(c, resp)
		return
	}
}

// GetObjectHandler
// @Description: 获取对象
// @return gin.HandlerFunc
func GetObjectHandler() gin.HandlerFunc {
	var objResp *ent.Object
	resp := Response{0, "success", nil, nil}
	return func(c *gin.Context) {
		oid := c.Param("oid")
		if objResp, resp.err = services.GetObjectByUUID(c, oid); resp.err != nil {
			resp.Code = 10003
			resp.Msg = "get object failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = objResp
		ResponseOK(c, resp)
		return
	}
}

func GetObjectsByTypeHandler() gin.HandlerFunc {
	var objsResp []*ent.Object
	var otype int
	resp := Response{0, "success", nil, nil}
	return func(c *gin.Context) {
		if otype, resp.err = strconv.Atoi(c.Param("otype")); resp.err != nil {
			resp.Code = 10005
			resp.Msg = "invalid otype"
			ResponseBadRequest(c, resp)
			return
		}
		if objsResp, resp.err = services.GetObjectsByType(c, otype); resp.err != nil {
			resp.Code = 10003
			resp.Msg = "get objects failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = objsResp
		ResponseOK(c, resp)
		return
	}
}
