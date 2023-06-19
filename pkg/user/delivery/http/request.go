package http

type UserCreateRequest struct {
	UserID string `json:"user_id"`
}

type UserCreateResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ApiKey   string `json:"api_key"`
}
