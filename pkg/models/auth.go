package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"user"`
}

type LoginResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    LoginResponseData `json:"data"`
}

// AuthWorkerResponseData is returned when a worker registers.
type AuthWorkerResponseData struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
type AuthWorkerResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    AuthWorkerResponseData `json:"data"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse is now the same as LoginResponse as the API returns the same structure.
type RegisterResponse = LoginResponse

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type RefreshResponseData struct {
	Token string `json:"token"`
}

type RefreshResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    RefreshResponseData `json:"data"`
}
