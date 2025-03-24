package domain

import (
	"context"
	"errors"
)

type ITaskManager interface {
	CreateTextureRenderingTask(ctx context.Context, worldId string, x int, y int) (string, error)
	CompleteTask(ctx context.Context, taskId string) (RenderImageTask, error)
	GetTaskById(ctx context.Context, taskId string) (RenderImageTask, error)
}

type TaskManager struct {
	taskRepository ITaskRepository
}

func (t TaskManager) CompleteTask(ctx context.Context, taskId string) (RenderImageTask, error) {
	tRef, err := t.GetTaskById(ctx, taskId)

	if err != nil {
		return RenderImageTask{}, err
	}

	return tRef, t.taskRepository.DeleteTask(ctx, taskId)
}

func (t TaskManager) CreateTextureRenderingTask(ctx context.Context, worldId string, x int, y int) (string, error) {

	id := TextureTaskId(worldId, x, y)

	tsk, err := t.taskRepository.GetTextureRenderingTask(ctx, id)

	if errors.Is(err, TaskAlreadyExistsError) {
		return "", nil
	}

	if errors.Is(err, RenderImageTaskNotFound) {
		iErr := t.taskRepository.InsertTextureRenderingTask(ctx, RenderImageTask{
			Task: Task{
				Id: id,
			},
			WorldId: worldId,
			X:       x,
			Y:       y,
		})
		if iErr != nil {
			return "", iErr
		}
		return id, nil
	}

	return tsk.Id, err
}

func (t TaskManager) GetTaskById(ctx context.Context, taskId string) (RenderImageTask, error) {
	return t.taskRepository.GetTextureRenderingTask(ctx, taskId)
}

func NewTaskManager(taskRepository ITaskRepository) ITaskManager {
	return &TaskManager{taskRepository}
}
