package domain

type RegisterUserRequest struct {
	Email      string `json:"email" validate:"required,email,max=30"`
	Password   string `json:"password" validate:"required,min=8"`
	Name       string `json:"name" validate:"required,max=100"`
	Address    string `json:"address" validate:"required"`
	PostalCode uint32 `json:"postal_code" validate:"required,numeric"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email,max=30"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserResponse struct {
	ID uint64 `json:"id"`
}
