package mktextrapi

import (
	"context"
	"errors"
	"mktextr/data_access"
	"mktextr/domain"
	mktextr "mktextr/gen/mktextr"
)

const MONGODB_URL = "mongodb://localhost:27017"

// mktextr service example implementation.
// The example methods log the requests and return zero values.
type mktextrsrvc struct {
	repository  domain.ITextureRefRepository
	taskManager domain.ITaskManager
}

func (s *mktextrsrvc) GetTextureByCoordinates(ctx context.Context, payload *mktextr.GetTextureByCoordinatesPayload) (res *mktextr.TextureReferencePayload, err error) {
	texRef, err := s.repository.GetTextureRefByCoordinates(ctx, payload.X, payload.Y, payload.WorldID)
	if err != nil {
		if errors.Is(err, domain.TextureRefNotFound) {
			s.taskManager.Enqueue()
		}

		return nil, err
	}

	return &mktextr.TextureReferencePayload{
		ID: &texRef.Id,
	}, nil
}

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
	dbContext, err := data_access.NewMongoDbContext(MONGODB_URL)
	if err != nil {
		panic(err)
	}

	localStorage := domain.NewLocalTextureStorage("./")

	return &mktextrsrvc{
		repository:  data_access.NewMongoTextureRefRepository(dbContext),
		taskManager: domain.NewTaskManager(localStorage),
	}
}
