package domain

import (
	"context"
	"errors"
)

type ITextureStorage interface {
	StoreTexture(rawData []byte) (TextureRef, error)
	GetTextureById(textureId string) ([]byte, error)
	DeleteTextureById(textureId string) error
}

type ITextureSetStore interface {
	GetOrCreateTextureRefByAddress(ctx context.Context, address TextureAddress) (TextureSet, error)
	InsertTextureSet(ctx context.Context, m TextureSet) error
	SetTextureAt(ctx context.Context, address TextureAddress, textureType TextureType, p TextureRef) error
}

type TextureSetStore struct {
	repo  ITextureSetRepository
	cache ITextureSetCache
}

func (t TextureSetStore) SetTextureAt(ctx context.Context, address TextureAddress, textureType TextureType, p TextureRef) error {
	set, err := t.repo.GetTextureSetByAddress(ctx, address)
	if err != nil {
		return err
	}

	// TODO: what if already ready?
	set.TextureSet[textureType] = TextureRef{
		Uri:    p.Uri,
		FileId: p.FileId,
		State:  TextureReady,
	}

	err = t.repo.UpdateTextureSet(ctx, set)
	if err != nil {
		return err
	}
	_, err = t.cache.SetTextureSet(ctx, set)
	if err != nil {
		// TODO: log!
	}

	return nil
}

func (t TextureSetStore) GetOrCreateTextureRefByAddress(ctx context.Context, address TextureAddress) (TextureSet, error) {
	r, err := t.GetTextureSetByAddress(ctx, address)
	if err == nil {
		return r, nil
	}

	if errors.Is(err, TextureRefNotFoundErr) {
		r = NewTextureSet(address)
		insertErr := t.InsertTextureSet(ctx, r)
		if insertErr != nil {
			return TextureSet{}, insertErr
		}

		return r, nil
	}

	return TextureSet{}, err
}

func (t TextureSetStore) GetTextureSetByAddress(ctx context.Context, address TextureAddress) (TextureSet, error) {
	r, err := t.cache.GetTextureSetByAddress(ctx, address)

	// return from cache
	if err == nil {
		return r, nil
	}

	// just not found in cache
	if errors.Is(err, ErrTextureSetNotFoundInCache) {

		// reach to a database
		r, err = t.repo.GetTextureSetByAddress(ctx, address)

		// we found record in a database
		if err == nil {

			// cache it since it was not cached before
			_, err = t.cache.SetTextureSet(ctx, r)
			if err != nil {
				// TODO: log!
			}

			// return result
			return r, nil
		}

		// not in cache nor in a database
		if errors.Is(err, TextureRefNotFoundErr) {
			return TextureSet{}, TextureRefNotFoundErr
		}
	}

	// general error
	return TextureSet{}, err
}

func (t TextureSetStore) InsertTextureSet(ctx context.Context, m TextureSet) error {

	// cache it first
	_, err := t.cache.SetTextureSet(ctx, m)
	if err != nil {
		return err
	}

	id, err := t.repo.InsertTextureSet(ctx, m)
	if err != nil {
		return err
	}

	// sanity id check
	if id != m.Id {
		return ErrInsertIdsMismatch
	}

	return nil
}

func NewTextureSetStore(repo ITextureSetRepository, cache ITextureSetCache) ITextureSetStore {
	return &TextureSetStore{
		repo:  repo,
		cache: cache,
	}
}
