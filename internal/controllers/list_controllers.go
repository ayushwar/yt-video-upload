package controllers


import (
	"net/http"

	"github.com/gin-gonic/gin"
	"yt-video-platform/internal/repository"
)

type ListHandler struct {
	Repo *repository.VideoRepository
}

func (h *ListHandler) List(c *gin.Context) {
	videos, _ := h.Repo.List()

	var result []gin.H
	for _, v := range videos {
		result = append(result, gin.H{
			"title": v.Title,
			"embed_url": "https://www.youtube.com/embed/" + v.YouTubeVideoID,
		})
	}

	c.JSON(http.StatusOK, result)
}