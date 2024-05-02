package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"status"`
	Msg  string `json:"msg"`
	err  error
	Data interface{} `json:"data"`
}

func ResponseOK(c *gin.Context, response Response) {
	if response.err != nil {
		response.Msg += response.err.Error()
	}
	c.JSON(200, response)
}

func ResponseUnauthorized(c *gin.Context, response Response) {
	if response.err != nil {
		response.Msg += response.err.Error()
	}
	c.JSON(401, response)
}

func ResponseBadRequest(c *gin.Context, response Response) {
	if response.err != nil {
		response.Msg += response.err.Error()
	}
	c.JSON(400, response)
}
