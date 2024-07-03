package domain

type AuthUserRequest struct {
	ID    uint64 `json:"id"`
	Token string `json:"token"`
}

type AuthUserResponse struct {
	ID uint64 `json:"id"`
}
