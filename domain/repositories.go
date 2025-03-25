package domain

import (
	"context"
	"fmt"
)

var (
	ErrInsertIdsMismatch         = fmt.Errorf("inserted id doesn't match with the one returned by a database")
	TextureRefNotFoundErr        = fmt.Errorf("texture not found")
	RenderImageTaskNotFound      = fmt.Errorf("rendering task not found")
	ErrTextureSetNotFoundInCache = fmt.Errorf("TextureSet not found in cache")
)

type ITextureSetRepository interface {
	InsertTextureSet(ctx context.Context, m TextureSet) (string, error)
	GetTextureSetByAddress(ctx context.Context, address TextureAddress) (TextureSet, error)
	GetTextureSetById(ctx context.Context, id string) (TextureSet, error)
	UpdateTextureSet(ctx context.Context, set TextureSet) error
}

type ITaskRepository interface {
	GetTextureRenderingTask(ctx context.Context, taskId string) (RenderImageTask, error)
	InsertTextureRenderingTasks(ctx context.Context, task []RenderImageTask) error
	DeleteTask(ctx context.Context, taskId string) error
	GetTaskQueue(ctx context.Context) ([]string, error)
	DeleteTaskFromQueue(ctx context.Context, taskId string) error
}
