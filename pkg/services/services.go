package services

import (
	"context"
	"go_sdk_qalpuch_api/pkg/models"
)

type AuthService interface {
	Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error)
	Register(ctx context.Context, req models.RegisterRequest) (*models.RegisterResponse, error)
	Logout(ctx context.Context, req models.LogoutRequest) error
	ChangePassword(ctx context.Context, req models.ChangePasswordRequest) error
	RefreshToken(ctx context.Context, req models.RefreshTokenRequest) (*models.RefreshResponse, error)
}

type UserService interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUser(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, id int, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
	DeleteCurrentUser(ctx context.Context) error
	CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error)
	SearchUsers(ctx context.Context, query string) ([]models.User, error)
}

type FileService interface {
	UploadFile(ctx context.Context, name string, file []byte) (*models.File, error)
	GetFileMetadata(ctx context.Context, cuid string) (*models.File, error)
	DownloadFile(ctx context.Context, cuid string) ([]byte, error)
	ListUserFiles(ctx context.Context) ([]models.File, error)
	DeleteFile(ctx context.Context, cuid string) error
	RenameFile(ctx context.Context, cuid string, newName string) (*models.File, error)
}

type TaskService interface {
	GetUserTasks(ctx context.Context) ([]models.Task, error)
	CreateTask(ctx context.Context, req models.CreateTaskRequest) (*models.Task, error)
	DeleteTask(ctx context.Context, cuid string) error
	GetPendingTask(ctx context.Context) (*models.Task, error)
	UpdateTaskStatus(ctx context.Context, cuid string, req models.UpdateTaskStatusRequest) error
	UploadTaskResult(ctx context.Context, cuid string, file []byte) error
	Build(fileID string) TaskBuilder
}

// TaskBuilder defines the fluent interface for building a task.
type TaskBuilder interface {
	WithVideoConfig(config models.VideoConversionConfig) TaskBuilder
	WithImageConfig(config models.ImageConversionConfig) TaskBuilder
	WithMusicConfig(config models.MusicConversionConfig) TaskBuilder
	Execute(ctx context.Context) (*models.Task, error)
}

type WorkerService interface {
	GetWorkers(ctx context.Context) ([]models.Worker, error)
	GetWorker(ctx context.Context, cuid string) (*models.Worker, error)
	CreateWorker(ctx context.Context, name string, capabilities []string) (*models.Worker, error)
	DeleteWorker(ctx context.Context, cuid string) error
	RegisterWorker(ctx context.Context, token string) (*models.AuthWorkerResponse, error)
	RefreshAuth(ctx context.Context, refreshToken string) (*models.AuthWorkerResponse, error)
}
