package main

import (
	"github.com/gin-gonic/gin"
	"yt-video-platform/internal/config"
	"yt-video-platform/internal/database"
	"yt-video-platform/internal/controllers"
	"yt-video-platform/internal/repository"
	"yt-video-platform/internal/routes"
)

func main() {
	cfg := config.Load()

	db := database.Connect(cfg.DatabaseURL)
	repo := &repository.VideoRepository{DB: db}

	upload := &controllers.UploadHandler{
		Repo: repo,
	}
	upload.Config.ClientID = cfg.YouTubeClientID
	upload.Config.ClientSecret = cfg.YouTubeClientSecret
	upload.Config.RefreshToken = cfg.YouTubeRefreshToken

	list := &controllers.ListHandler{Repo: repo}

	r := gin.Default()
	routes.Register(r, upload, list)

	r.Run(":" + cfg.Port)
}