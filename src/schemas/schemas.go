package schemas

import "time"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// Todo struct for the request HTTP body
type TodoRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Active      bool      `json:"active"`
	StatusID    uint      `json:"status_id"`
	UserID      uint      `json:"user_id"`
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
