package mktextr

import (
	"encoding/json"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Task struct {
	id         string
	parameters json.RawMessage
}

type TaskHandle struct {
	id uuid.UUID
}

func EmptyTaskHandle() TaskHandle {
	return TaskHandle{
		id: uuid.New(),
	}
}

type TaskProcessor struct {
	id uuid.UUID
}

type TextureRefDataModel struct {
	id  bson.ObjectID
	uri string
}

type TextureRef struct {
	id  string
	uri string
}

func EmptyTextureRef() TextureRef {
	return TextureRef{}
}

type PagedResult[T any] struct {
	Items      []T `json:"items"`
	TotalCount int `json:"totalCount"`
}
