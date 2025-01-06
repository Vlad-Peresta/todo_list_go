package schemas

import "time"

// Todo struct for the request HTTP body
type TodoRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Active      bool      `json:"active"`
	StatusID    uint      `json:"status_id"`
	UserID      uint      `json:"user_id" binding:"required"`
}

// Todo struct for the HTTP response
type TodoResponse struct {
	TodoRequest
	ID uint `json:"id"`
}

type AuthInputData struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
