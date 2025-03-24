// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mktextr service
//
// Command:
// $ goa gen mktextr/design

package mktextr

import (
	"context"
)

// Texture store
type Service interface {
	// GetTextureByID implements getTextureById.
	GetTextureByID(context.Context, *GetTextureByIDPayload) (res *TextureReferencePayload, err error)
	// GetTextureByCoordinates implements getTextureByCoordinates.
	GetTextureByCoordinates(context.Context, *GetTextureByCoordinatesPayload) (res *TextureReferencePayload, err error)
	// CompleteTask implements completeTask.
	CompleteTask(context.Context, *TaskCompletionPayload) (err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "mktextr"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "mktextr"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"getTextureById", "getTextureByCoordinates", "completeTask"}

// GetTextureByCoordinatesPayload is the payload type of the mktextr service
// getTextureByCoordinates method.
type GetTextureByCoordinatesPayload struct {
	// Texture X
	X int
	// Texture y
	Y int
	// WorldId
	WorldID string
}

// GetTextureByIDPayload is the payload type of the mktextr service
// getTextureById method.
type GetTextureByIDPayload struct {
	// Texture ID
	ID string
}

// Complete task
type TaskCompletionPayload struct {
	// Unique identifier
	TaskID string
	// The texture
	Texture []byte
}

// Texture reference
type TextureReferencePayload struct {
	// Unique identifier
	ID *string
}
