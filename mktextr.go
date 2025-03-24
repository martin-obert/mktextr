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
	repository     domain.ITextureRefRepository
	taskManager    domain.ITaskManager
	textureStorage domain.ITextureStorage
}

func (m mktextrsrvc) GetTextureByID(ctx context.Context, payload *mktextr.GetTextureByIDPayload) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m mktextrsrvc) GetTextureByCoordinates(ctx context.Context, payload *mktextr.GetTextureByCoordinatesPayload) (res *mktextr.GetTextureByCoordinatesResult, err error) {
	texRef, err := m.repository.GetTextureRefByCoordinates(ctx, payload.WorldID, payload.X, payload.Y)
	if err != nil {
		if errors.Is(err, domain.TextureRefNotFound) {
			t, cErr := m.taskManager.CreateTextureRenderingTask(ctx, payload.WorldID, payload.X, payload.Y)
			if cErr != nil {
				return nil, cErr
			}

			return &mktextr.GetTextureByCoordinatesResult{
				XmktextrTaskID: &t,
				Location:       nil,
			}, nil
		}

		return nil, err
	}

	return &mktextr.GetTextureByCoordinatesResult{
		Location:       &texRef.Uri,
		XmktextrTaskID: nil,
	}, nil
}

func (m mktextrsrvc) CompleteTask(ctx context.Context, payload *mktextr.CompleteTaskPayload) (err error) {
	//TODO implement me
	panic("implement me")
}

//
//func (s *mktextrsrvc) CompleteTask(ctx context.Context, payload *mktextr.CompleteTaskPayload) (res *mktextr.TextureReferencePayload, err error) {
//	tr, err := s.taskManager.CompleteTask(ctx, payload.TaskID)
//	if err != nil {
//		return nil, err
//	}
//
//	m, err := s.textureStorage.StoreTexture(payload.File)
//	if err != nil {
//		return nil, err
//	}
//	r, err := s.repository.InsertTextureRef(ctx, m, tr.WorldId, tr.X, tr.Y)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &mktextr.TextureReferencePayload{
//		ID: &r.Id,
//	}, err
//}
//
//func (s *mktextrsrvc) GetTextureByCoordinates(ctx context.Context, payload *mktextr.GetTextureByCoordinatesPayload) (res *mktextr.TextureReferencePayload, err error) {
//	texRef, err := s.repository.GetTextureRefByCoordinates(ctx, payload.WorldID, payload.X, payload.Y)
//	if err != nil {
//		if errors.Is(err, domain.TextureRefNotFound) {
//			t, cErr := s.taskManager.CreateTextureRenderingTask(ctx, payload.WorldID, payload.X, payload.Y)
//			if cErr != nil {
//				return nil, cErr
//			}
//
//			return &mktextr.TextureReferencePayload{}, nil
//		}
//
//		return nil, err
//	}
//
//	return &mktextr.TextureReferencePayload{
//		ID: &texRef.Id,
//	}, nil
//}

// NewMktextr returns the mktextr service implementation.
func NewMktextr() mktextr.Service {
	dbContext, err := data_access.NewMongoDbContext(MONGODB_URL)
	if err != nil {
		panic(err)
	}

	taskRepository := data_access.NewRedisTaskRepository("localhost:6379", "", 0)

	return &mktextrsrvc{
		repository:     data_access.NewMongoTextureRefRepository(dbContext),
		taskManager:    domain.NewTaskManager(taskRepository),
		textureStorage: data_access.NewLocalTextureStorage("./upload"),
	}
}
