package mktextrapi

import (
	"context"
	mktextr "mktextr/gen/mktextr"
)

// mktextr service example implementation.
// The example methods log the requests and return zero values.
type mktextrsrvc struct{}

func (s *mktextrsrvc) GetTextureByID(ctx context.Context, payload *mktextr.GetTextureByIDPayload) (res *mktextr.TextureReferencePayload, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *mktextrsrvc) CompleteTask(ctx context.Context, payload *mktextr.TaskCompletionPayload) (err error) {
	//TODO implement me
	panic("implement me")
}

// NewMktextr returns the mktextr service implementation.
func NewMktextr() mktextr.Service {
	return &mktextrsrvc{}
}
