package clients

import (
	"context"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// TaskBuilder provides a fluent interface for constructing and executing tasks.
type TaskBuilder struct {
	client  *TaskClient
	fileID  string
	config  interface{}
	lastErr error
}

// WithVideoConfig sets the configuration for a video conversion task.
func (b *TaskBuilder) WithVideoConfig(config models.VideoConversionConfig) services.TaskBuilder {
	if b.lastErr != nil {
		return b
	}
	config.Type = "video"
	b.config = config
	return b
}

// WithImageConfig sets the configuration for an image conversion task.
func (b *TaskBuilder) WithImageConfig(config models.ImageConversionConfig) services.TaskBuilder {
	if b.lastErr != nil {
		return b
	}
	config.Type = "image"
	b.config = config
	return b
}

// WithMusicConfig sets the configuration for a music conversion task.
func (b *TaskBuilder) WithMusicConfig(config models.MusicConversionConfig) services.TaskBuilder {
	if b.lastErr != nil {
		return b
	}
	config.Type = "music"
	b.config = config
	return b
}

// Execute finalizes the task construction and sends it to the API for processing.
func (b *TaskBuilder) Execute(ctx context.Context) (*models.Task, error) {
	if b.lastErr != nil {
		return nil, b.lastErr
	}
	req := models.CreateTaskRequest{
		FileID: b.fileID,
		Config: b.config,
	}
	return b.client.CreateTask(ctx, req)
}
