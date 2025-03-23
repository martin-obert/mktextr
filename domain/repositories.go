package domain

import (
	"context"
)

type ITextureRefRepository interface {
	InsertTextureRef(ctx context.Context, uri string) (TextureRef, error)
	GetTextureRefById(ctx context.Context, texRefId string) (TextureRef, error)
	QueryTextureRefs(ctx context.Context) ([]TextureRef, error)
}
