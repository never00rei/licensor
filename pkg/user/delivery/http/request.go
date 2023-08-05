package http

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserCreateResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	ApiKey   string `json:"api_key"`
}
