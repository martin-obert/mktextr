package mktextrapi

import (
	"context"
	"github.com/caarlos0/env/v11"
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

func (m mktextrsrvc) CompleteTask(ctx context.Context, payload *mktextr.CompleteTaskPayload) (err error) {
	tId, err := domain.DecodeTaskId(payload.TaskID)
	if err != nil {
		return err
	}

	t, err := m.taskManager.CompleteTask(ctx, tId)
	if err != nil {
		return err
	}

	p, err := m.textureStorage.StoreTexture(payload.File, payload.Extension)

	if err != nil {
		return err
	}

	return m.store.SetTextureAt(ctx, t.Address, t.TextureType, p)
}

type config struct {
	MongoUrl            string `env:"MONGO_CONNECTION_STRING" envDefault:"mongodb://localhost:27017"`
	RedisAddr           string `env:"REDIS_ADDRESS" envDefault:"localhost:6379"`
	RedisPwd            string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDb             int    `env:"REDIS_DB" envDefault:"0"`
	LocalStoreRoot      string `env:"LOCAL_STORE_ROOT" envDefault:"./uploads"`
	LocalStoreUriPrefix string `env:"LOCAL_STORE_URI_PREFIX" envDefault:"file:///C:/Repositories/obert/mktextr/uploads/"`
}

// NewMktextr returns the mktextr service implementation.
func NewMktextr() mktextr.Service {
	cfg, err := env.ParseAs[config]()
	if err != nil {
		panic(err)
	}
	dbContext, err := data_access.NewMongoDbContext(cfg.MongoUrl)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPwd,
		DB:       cfg.RedisDb,
	})

	taskRepository := data_access.NewRedisTaskRepository(redisClient)
	repo := data_access.NewMongoTextureRefRepository(dbContext)
	cache := data_access.NewRedisTextureCache(redisClient)

	return &mktextrsrvc{
		store:          domain.NewTextureSetStore(repo, cache),
		taskManager:    domain.NewTaskManager(taskRepository),
		textureStorage: data_access.NewLocalTextureStorage(cfg.LocalStoreRoot, cfg.LocalStoreUriPrefix),
	}
}
