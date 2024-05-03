package handlers

import (
	"dgut-icourse/app/services"
	"dgut-icourse/ent"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CreateCommentHandler
// @Description: 创建评论
// @return gin.HandlerFunc
func CreateCommentHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentReq ent.Comment
		var commentResp *ent.Comment
		resp := Response{0, "success", nil, nil}
		// 获取字段
		if resp.err = c.ShouldBindJSON(&commentReq); resp.err != nil {
			resp.Code = 10001
			resp.Msg = "invalid json"
			ResponseBadRequest(c, resp)
			return
		} else if commentResp, resp.err = services.CreateComment(c, commentReq); resp.err != nil {
			resp.Code = 10002
			resp.Msg = "create comment failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = commentResp
		ResponseOK(c, resp)
		return
	}
}

// GetCommentHandler
// @Description: 获取评论
// @return gin.HandlerFunc
func GetCommentHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentResp *ent.Comment
		resp := Response{0, "success", nil, nil}
		cid := c.Param("cid")
		if commentResp, resp.err = services.GetCommentByUUID(c, cid); resp.err != nil {
			resp.Code = 10003
			resp.Msg = "get comment failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = commentResp
		ResponseOK(c, resp)
		return
	}
}

// GetCommentsByObjIDHandler
// @Description: 通过OID获取评论
// @return gin.HandlerFunc
func GetCommentsByObjIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentsResp []*ent.Comment
		resp := Response{0, "success", nil, nil}
		oid := c.Param("oid")
		if commentsResp, resp.err = services.GetCommentsByObjectUUID(c, oid); resp.err != nil {
			resp.Code = 10003
			resp.Msg = "get comments failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = commentsResp
		ResponseOK(c, resp)
		return
	}
}

func GetCommentsByTimeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var commentsResp []*ent.Comment
		var offset, limit int
		resp := Response{0, "success", nil, nil}
		if offset, resp.err = strconv.Atoi(c.Query("offset")); resp.err != nil {
			resp.Code = 10010
			resp.Msg = "invalid offset"
			ResponseBadRequest(c, resp)
			return
		}
		if limit, resp.err = strconv.Atoi(c.Query("limit")); resp.err != nil {
			resp.Code = 10011
			resp.Msg = "invalid limit"
			ResponseBadRequest(c, resp)
			return
		}
		if commentsResp, resp.err = services.GetCommentsByTime(c, offset, limit); resp.err != nil {
			resp.Code = 10003
			resp.Msg = "get comments failed, " + resp.err.Error()
			ResponseBadRequest(c, resp)
			return
		}
		resp.Data = commentsResp
		ResponseOK(c, resp)
		return
	}
}
