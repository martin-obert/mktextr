package domain

import (
	"github.com/google/uuid"
)

type Task struct {
}

type RenderImageTask struct {
	Task
	WorldId string
	X       int
	Y       int
}

type TaskHandle struct {
	Id string
}

func EmptyTaskHandle() TaskHandle {
	return TaskHandle{
		Id: uuid.New().String(),
	}
}

type TextureRef struct {
	Id  string
	Uri string
}
