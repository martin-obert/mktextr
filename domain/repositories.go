package domain

import (
	"context"
	"fmt"
)

var TextureRefNotFound = fmt.Errorf("texture not found")

type ITextureRefRepository interface {
	InsertTextureRef(ctx context.Context, uri string) (TextureRef, error)
	GetTextureRefById(ctx context.Context, texRefId string) (TextureRef, error)
	GetTextureRefByCoordinates(ctx context.Context, x int, y int, worldId string) (TextureRef, error)
	QueryTextureRefs(ctx context.Context) ([]TextureRef, error)
}
