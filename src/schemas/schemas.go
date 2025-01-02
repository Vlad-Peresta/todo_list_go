package schemas

// Todo struct for the request HTTP body
type TodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
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
