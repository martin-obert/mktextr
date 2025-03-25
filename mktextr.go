package mktextrapi

import (
	"context"
	"github.com/redis/go-redis/v9"
	"mktextr/data_access"
	"mktextr/domain"
	mktextr "mktextr/gen/mktextr"
)

const MONGODB_URL = "mongodb://localhost:27017"

// mktextr service example implementation.
// The example methods log the requests and return zero values.
type mktextrsrvc struct {
	store          domain.ITextureSetStore
	taskManager    domain.ITaskManager
	textureStorage domain.ITextureStorage
}

func (m mktextrsrvc) GetTaskQueue(ctx context.Context) (res *mktextr.GetTaskQueueResult, err error) {
	q, err := m.taskManager.GetTaskQueue(ctx)
	if err != nil {
		return nil, err
	}
	queue := make([]string, len(q))
	for i, s := range q {
		queue[i] = domain.EncodeTaskId(s)
	}

	return &mktextr.GetTaskQueueResult{Tasks: queue}, nil
}

func (m mktextrsrvc) GetTextureByCoordinates(ctx context.Context, payload *mktextr.GetTextureByCoordinatesPayload) (res *mktextr.GetTextureByCoordinatesResponse, err error) {

	address := domain.TextureAddress{
		WorldId: payload.WorldID,
		X:       payload.X,
		Y:       payload.Y,
	}

	texRef, err := m.store.GetOrCreateTextureRefByAddress(ctx, address)

	if err != nil {
		return &mktextr.GetTextureByCoordinatesResponse{}, err
	}

	if texRef.State() != domain.TextureSetStateComplete {

		cErr := m.taskManager.CreateTextureSetTasks(ctx, texRef)
		if cErr != nil {
			return nil, cErr
		}
	}
	state := texRef.State().ToDesign()

	bMap := texRef.TextureSet[domain.TextureTypeBaseMap].Uri

	cMap := texRef.TextureSet[domain.TextureTypeContourMap].Uri

	return &mktextr.GetTextureByCoordinatesResponse{
		TextureSetState: &state,
		BaseMapURL:      &bMap,
		ContourMapURL:   &cMap,
	}, nil

}

//func (m mktextrsrvc) GetTextureByCoordinates(ctx context.Context, payload *mktextr.GetTextureByCoordinatesPayload) (res *mktextr.GetTextureByCoordinatesResult, err error) {

//texRef, err := m.repository.GetTextureRefByCoordinates(ctx, payload.WorldID, payload.X, payload.Y)
//if err != nil {
//	if errors.Is(err, domain.TextureRefNotFoundErr) {
//		t, cErr := m.taskManager.CreateTextureSetTasks(ctx, payload.WorldID, payload.X, payload.Y)
//		if cErr != nil {
//			return nil, cErr
//		}
//
//		return &mktextr.GetTextureByCoordinatesResult{
//			XmktextrTaskID: &t,
//			Location:       nil,
//		}, nil
//	}
//
//	return nil, err
//}
//
//return &mktextr.GetTextureByCoordinatesResult{
//	Location:       &texRef.BaseMapUri,
//	XmktextrTaskID: nil,
//}, nil
//}

func (m mktextrsrvc) CompleteTask(ctx context.Context, payload *mktextr.CompleteTaskPayload) (err error) {
	tId, err := domain.DecodeTaskId(payload.TaskID)
	if err != nil {
		return err
	}

	t, err := m.taskManager.CompleteTask(ctx, tId)
	if err != nil {
		return err
	}

	p, err := m.textureStorage.StoreTexture(payload.File)

	if err != nil {
		return err
	}

	return m.store.SetTextureAt(ctx, t.Address, t.TextureType, p)
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
//	r, err := s.repository.InsertTextureSet(ctx, m, tr.WorldId, tr.X, tr.Y)
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
//		if errors.Is(err, domain.TextureRefNotFoundErr) {
//			t, cErr := s.taskManager.CreateTextureSetTasks(ctx, payload.WorldID, payload.X, payload.Y)
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

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	taskRepository := data_access.NewRedisTaskRepository(redisClient)
	repo := data_access.NewMongoTextureRefRepository(dbContext)
	cache := data_access.NewRedisTextureCache(redisClient)

	return &mktextrsrvc{
		store:          domain.NewTextureSetStore(repo, cache),
		taskManager:    domain.NewTaskManager(taskRepository),
		textureStorage: data_access.NewLocalTextureStorage("./upload"),
	}
}
