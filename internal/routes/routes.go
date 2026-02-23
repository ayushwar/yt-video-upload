package routes

import (
	"github.com/gin-gonic/gin"
	"yt-video-platform/internal/controllers"
)

func Register(r *gin.Engine, upload *controllers.UploadHandler, list *controllers.ListHandler) {
	api := r.Group("/api/videos")
	{
		api.POST("/upload", upload.Upload)
		api.GET("", list.List)
	}
}