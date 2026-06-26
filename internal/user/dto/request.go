package dto

type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"omitempty,oneof=driver admin"`
}

type UpdateUserRequest struct {
	Name  string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email string `json:"email,omitempty" validate:"omitempty,email"`
	Role  string `json:"role,omitempty" validate:"omitempty,oneof=driver admin"`
}
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required"`
}
// UpdateUserRequest represents a request to update a user
