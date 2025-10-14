package models

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	} `json:"data"`
}

// AuthWorkerResponse is returned when a worker registers.
// It contains only the access token.
type AuthWorkerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn"`
	} `json:"data"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

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

type RefreshResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
		ExpiresIn    int    `json:"expiresIn"`
	} `json:"data"`
}
