package clients

import (
	"context"
	"fmt"

	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/models"
	"github.com/Q300Z/go_sdk_qalpuch_api/pkg/services"
)

// TaskBuilder implements services.TaskBuilder.
type TaskBuilder struct {
	client *TaskClient
	fileID string
	config interface{}
}

// NewTaskBuilder creates a new TaskBuilder.
func NewTaskBuilder(client *TaskClient, fileID string) services.TaskBuilder {
	return &TaskBuilder{client: client, fileID: fileID}
}

// WithVideoConfig sets the video conversion configuration.
func (b *TaskBuilder) WithVideoConfig(config models.VideoConversionConfig) services.TaskBuilder {
	b.config = config
	return b
}

// WithImageConfig sets the image conversion configuration.
func (b *TaskBuilder) WithImageConfig(config models.ImageConversionConfig) services.TaskBuilder {
	b.config = config
	return b
}

// WithAudioConfig sets the audio conversion configuration.
func (b *TaskBuilder) WithAudioConfig(config models.AudioConversionConfig) services.TaskBuilder {
	b.config = config
	return b
}

// Execute creates the task with the configured parameters.
func (b *TaskBuilder) Execute(ctx context.Context) (*models.Task, error) {
	if b.config == nil {
		return nil, fmt.Errorf("task configuration is incomplete. Please call one of the With...Config methods")
	}

	req := models.CreateTaskRequest{
		FileID: b.fileID,
		Config: &b.config,
	}

	return b.client.CreateTask(ctx, req)
}
