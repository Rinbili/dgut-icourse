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
	user.PUT("/me", handlers.UpdateMeHandler())
	// object路由组
	object := r.Group("/object")
	object.GET("/:oid", handlers.GetObjectHandler())
	object.GET("/type/:otype", handlers.GetObjectsByTypeHandler())
	object.GET("/last", handlers.GetObjectsHandler())
	object.POST("/create", handlers.CreateObjectHandler())
	// comment路由组
	comment := r.Group("/comment")
	comment.GET("/:cid", handlers.GetCommentHandler())
	comment.GET("/my", handlers.GetMyCommentsHandler())
	comment.GET("/comment/:cid", handlers.GetCommentsByCommentUUIDHandler())
	comment.GET("/object/:oid", handlers.GetCommentsByObjIDHandler())
	comment.GET("/last", handlers.GetCommentsByTimeHandler())
	comment.PUT("/update", handlers.UpdateCommentHandler())
	comment.POST("/create", handlers.CreateCommentHandler())
	comment.GET("/:cid/like", handlers.GetCommentLikeCountAndDidILikeHandler())
	comment.POST("/:cid/like", handlers.LikeCommentHandler())
	comment.POST("/:cid/unlike", handlers.UnlikeCommentHandler())
}
