package domain

import (
	"context"
	"fmt"
)

var TaskAlreadyExistsError = fmt.Errorf("TaskAlreadyExistsError")

type StoredTextureMeta struct {
	Uri    string
	FileId string
}

type ITaskRepository interface {
	GetTextureRenderingTask(ctx context.Context, taskId string) (RenderImageTask, error)
	InsertTextureRenderingTask(ctx context.Context, task RenderImageTask) error
	DeleteTask(ctx context.Context, taskId string) error
}

type ITextureStorage interface {
	StoreTexture(rawData []byte) (StoredTextureMeta, error)
	GetTextureById(textureId string) ([]byte, error)
	DeleteTextureById(textureId string) error
}
