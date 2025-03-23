package mktextr

type ITaskManager interface {
	Enqueue(Task) (TaskHandle, error)
}

type TaskManager struct {
	textureStorage ITextureStorage
}

func NewTaskManager(textureStorage ITextureStorage) *TaskManager {
	return &TaskManager{textureStorage}
}

func (t TaskManager) Enqueue(task Task) (TaskHandle, error) {
	
	return EmptyTaskHandle(), nil
}
