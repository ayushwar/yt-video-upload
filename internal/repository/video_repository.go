package repository

import (
	"database/sql"
	"yt-video-platform/internal/models"
)

type VideoRepository struct {
	DB *sql.DB
}

func (r *VideoRepository) Create(video *models.Video) error {
	query := `
	INSERT INTO videos (title, description, youtube_video_id)
	VALUES ($1, $2, $3)
	`
	_, err := r.DB.Exec(query, video.Title, video.Description, video.YouTubeVideoID)
	return err
}

func (r *VideoRepository) List() ([]models.Video, error) {
	rows, err := r.DB.Query(`
		SELECT id, title, description, youtube_video_id, created_at
		FROM videos ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []models.Video
	for rows.Next() {
		var v models.Video
		rows.Scan(&v.ID, &v.Title, &v.Description, &v.YouTubeVideoID, &v.CreatedAt)
		videos = append(videos, v)
	}
	return videos, nil
}