package domain

import (
	"context"
	"errors"
)

type ITextureSetCache interface {
	GetTextureSetByAddress(ctx context.Context, address TextureAddress) (TextureSet, error)
	SetTextureSet(ctx context.Context, textureSet TextureSet) (string, error)
	DelTextureSet(ctx context.Context, textureSet TextureSet) error
}

type ITaskManager interface {
	CreateTextureSetTasks(ctx context.Context, texRef TextureSet) error
	CompleteTask(ctx context.Context, taskId string) (RenderImageTask, error)
	GetTaskQueue(ctx context.Context) ([]string, error)
}

type TaskManager struct {
	taskRepository ITaskRepository
}

func (t TaskManager) GetTaskQueue(ctx context.Context) ([]string, error) {
	return t.taskRepository.GetTaskQueue(ctx)
}

func (t TaskManager) CreateTextureSetTasks(ctx context.Context, textureSet TextureSet) error {
	var tasks []RenderImageTask
	if bm, ok := textureSet.TextureSet[TextureTypeBaseMap]; ok {
		if bm.State != TextureReady {
			tasks = append(tasks, NewRenderBaseMapImageTask(textureSet.Address))
		}
	}

	if bm, ok := textureSet.TextureSet[TextureTypeContourMap]; ok {
		if bm.State != TextureReady {
			tasks = append(tasks, NewRenderContourMapImageTask(textureSet.Address))
		}
	}

	err := t.taskRepository.InsertTextureRenderingTasks(ctx, tasks)
	return err
}

func (t TaskManager) CompleteTask(ctx context.Context, taskId string) (RenderImageTask, error) {
	tRef, err := t.GetTaskById(ctx, taskId)

	if err != nil {
		return RenderImageTask{}, err
	}

	return tRef, t.taskRepository.DeleteTask(ctx, taskId)
}

func (t TaskManager) GetTaskById(ctx context.Context, taskId string) (RenderImageTask, error) {
	r, err := t.taskRepository.GetTextureRenderingTask(ctx, taskId)
	if err != nil {
		if errors.Is(err, RenderImageTaskNotFound) {
			delErr := t.taskRepository.DeleteTaskFromQueue(ctx, taskId)
			if delErr != nil {
				// TODO: Log!
			}

			return RenderImageTask{}, err
		}
		return RenderImageTask{}, err
	}

	return r, nil
}

func NewTaskManager(taskRepository ITaskRepository) ITaskManager {
	return &TaskManager{taskRepository}
}
