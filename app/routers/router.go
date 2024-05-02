package routers

import (
	"dgut-icourse/app/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(handlers.AuthHandler())
	// auth路由组
	auth := r.Group("/auth")
	auth.POST("/login", handlers.LoginHandler())
	// user路由组
	user := r.Group("/user")
	user.GET("/me", handlers.GetMe())
	// object路由组
	object := r.Group("/object")
	object.GET("/:oid", handlers.GetObjectHandler())
	object.GET("/type/:otype", handlers.GetObjectsByTypeHandler())
	object.POST("/create", handlers.CreateObjectHandler())
	// comment路由组
	comment := r.Group("/comment")
	comment.GET("/:cid", handlers.GetCommentHandler())
	comment.GET("/object/:oid", handlers.GetCommentsByObjIDHandler())
	comment.GET("/last", handlers.GetCommentsByTimeHandler())
	comment.POST("/create", handlers.CreateCommentHandler())
}
