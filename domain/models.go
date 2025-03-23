package domain

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Task struct {
	Id         string
	Parameters json.RawMessage
}

type TaskHandle struct {
	Id uuid.UUID
}

func EmptyTaskHandle() TaskHandle {
	return TaskHandle{
		Id: uuid.New(),
	}
}

type TaskProcessor struct {
	Id uuid.UUID
}

type TextureRef struct {
	Id  string
	Uri string
}

func EmptyTextureRef() TextureRef {
	return TextureRef{}
}
