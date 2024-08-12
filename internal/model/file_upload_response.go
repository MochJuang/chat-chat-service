package model

import "time"

type CreateFileUploadResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	FileURL    string    `json:"file_url"`
	UploadedAt time.Time `json:"uploaded_at"`
}
