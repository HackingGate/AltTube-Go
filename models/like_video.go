package models

// LikeVideo represents a record of a user liking a video
type LikeVideo struct {
	Entity
	ID      uint   `json:"id"`
	UserID  string `json:"user_id"`
	VideoID string `json:"video_id"`
}

// LikeVideoRequest represents the payload for liking a video
type LikeVideoRequest struct {
	VideoID string `json:"video_id" binding:"required"`
}

// LikeVideoResponse represents the response for a liked video
type LikeVideoResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ThumbnailURL string `json:"thumbnail_url"`
}
