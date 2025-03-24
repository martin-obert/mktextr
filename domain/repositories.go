package domain

import (
	"context"
	"fmt"
)

var TextureRefNotFound = fmt.Errorf("texture not found")
var RenderImageTaskNotFound = fmt.Errorf("rendering task not found")

type ITextureRefRepository interface {
	InsertTextureRef(ctx context.Context, meta StoredTextureMeta, worldId string, x int, y int) (TextureRef, error)
	GetTextureRefById(ctx context.Context, texRefId string) (TextureRef, error)
	GetTextureRefByCoordinates(ctx context.Context, worldId string, x int, y int) (TextureRef, error)
	QueryTextureRefs(ctx context.Context) ([]TextureRef, error)
}
