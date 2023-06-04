package http

type ManagementUserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"is_admin"`
}

type ManagementUserCreateResponse struct {
	Username string `json:"username"`
	ApiKey   string `json:"api_key"`
}
