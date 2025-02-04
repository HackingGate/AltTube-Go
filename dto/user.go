package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateEmailRequest struct {
	Email string `json:"email"`
}
