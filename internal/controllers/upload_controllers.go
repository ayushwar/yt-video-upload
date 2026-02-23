package controllers


import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"yt-video-platform/internal/models"
	"yt-video-platform/internal/repository"
	"yt-video-platform/internal/services"
)

type UploadHandler struct {
	Repo   *repository.VideoRepository
	Config struct {
		ClientID     string
		ClientSecret string
		RefreshToken string
	}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	title := c.PostForm("title")
	desc := c.PostForm("description")

	file, _ := c.FormFile("video")
	path := filepath.Join("uploads", file.Filename)
	c.SaveUploadedFile(file, path)

	videoID, err := services.UploadToYouTube(
		path, title, desc,
		h.Config.ClientID,
		h.Config.ClientSecret,
		h.Config.RefreshToken,
	)

	os.Remove(path)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.Repo.Create(&models.Video{
		Title:          title,
		Description:    desc,
		YouTubeVideoID: videoID,
	})

	c.JSON(200, gin.H{
		"embed_url": "https://www.youtube.com/embed/" + videoID,
	})
}