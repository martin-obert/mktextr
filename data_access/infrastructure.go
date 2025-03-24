package data_access

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"io"
	"mktextr/domain"
	"os"
	"path/filepath"
)

type MongoTextureRefRepository struct {
	dbContext *MongoDbContext
}

func (m MongoTextureRefRepository) GetTextureRefByCoordinates(ctx context.Context, worldId string, x int, y int) (domain.TextureRef, error) {
	c := m.dbContext.textureRefsCollection.FindOne(ctx, bson.M{"world_id": worldId, "x": x, "y": y}, options.FindOne())

	var result TextureRefDataModel

	err := c.Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.TextureRef{}, domain.TextureRefNotFound
		}
		return domain.TextureRef{}, err
	}

	return toDomain(result), nil
}

func (m MongoTextureRefRepository) InsertTextureRef(ctx context.Context, meta domain.StoredTextureMeta, worldId string, x int, y int) (domain.TextureRef, error) {

	var data = TextureRefDataModel{
		Uri:     meta.Uri,
		FileId:  meta.FileId,
		WorldId: worldId,
		X:       x,
		Y:       y,
	}
	one, err := m.dbContext.textureRefsCollection.InsertOne(ctx, data)
	if err != nil {
		return domain.TextureRef{}, err
	}

	return toDomain(TextureRefDataModel{
		Id:     one.InsertedID.(bson.ObjectID),
		Uri:    data.Uri,
		FileId: data.FileId,
	}), nil

}

func (m MongoTextureRefRepository) GetTextureRefById(ctx context.Context, texRefId string) (domain.TextureRef, error) {
	var data TextureRefDataModel
	err := m.dbContext.textureRefsCollection.FindOne(ctx, bson.M{"_id": texRefId}).Decode(&data)
	if err != nil {
		return domain.TextureRef{}, err
	}

	return toDomain(data), nil
}

func (m MongoTextureRefRepository) QueryTextureRefs(ctx context.Context) ([]domain.TextureRef, error) {
	cursor, err := m.dbContext.textureRefsCollection.Find(ctx, bson.D{})
	if err != nil {
		return []domain.TextureRef{}, err
	}

	var results []TextureRefDataModel
	err = cursor.All(ctx, &results)

	if err != nil {
		return []domain.TextureRef{}, err
	}
	return toDomainArr(results), nil
}

func NewMongoTextureRefRepository(ctx *MongoDbContext) domain.ITextureRefRepository {
	return &MongoTextureRefRepository{dbContext: ctx}
}

type RedisTaskRepository struct {
	redisClient *redis.Client
}

func (r RedisTaskRepository) DeleteTask(ctx context.Context, taskId string) error {
	c := r.redisClient.Del(ctx, taskId)
	return c.Err()
}

func (r RedisTaskRepository) GetTextureRenderingTask(ctx context.Context, taskId string) (domain.RenderImageTask, error) {

	// Get the JSON string back from Redis
	jsonResult, err := r.redisClient.Get(ctx, taskId).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.RenderImageTask{}, domain.RenderImageTaskNotFound
		}
		return domain.RenderImageTask{}, err
	}

	// Unmarshal back to struct
	var task domain.RenderImageTask
	err = json.Unmarshal([]byte(jsonResult), &task)

	if err != nil {
		return domain.RenderImageTask{}, err
	}

	return task, nil
}

func (r RedisTaskRepository) InsertTextureRenderingTask(ctx context.Context, task domain.RenderImageTask) error {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = r.redisClient.Set(ctx, domain.TextureTaskId(task.WorldId, task.X, task.Y), jsonData, 0).Err()
	if err != nil {
		return err
	}

	return err
}

func NewRedisTaskRepository(address string, password string, db int) domain.ITaskRepository {
	return &RedisTaskRepository{redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})}
}

type LocalTextureStorage struct {
	rootFolder string
}

func (t LocalTextureStorage) DeleteTextureById(textureId string) error {
	var path = t.getFilePath(textureId)
	return os.Remove(path)
}

func (t LocalTextureStorage) getFilePath(textureId string) string {
	p, err := filepath.Abs(t.rootFolder)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(p, os.ModeDir|os.ModePerm)
	if err != nil {
		panic(err)
	}

	return filepath.Join(p, fmt.Sprintf("%s.png", textureId))
}

func (t LocalTextureStorage) GetTextureById(textureId string) ([]byte, error) {
	var path = t.getFilePath(textureId)

	// Read binary data from file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Open error:", err)
		return nil, err
	}
	defer func(file *os.File) {
		closeErr := file.Close()
		if closeErr != nil {

		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Read error:", err)
		return nil, err
	}

	return data, nil
}

func (t LocalTextureStorage) StoreTexture(rawData []byte) (domain.StoredTextureMeta, error) {
	fileId := uuid.New().String()
	// Create (or overwrite) the file
	var path = t.getFilePath(fileId)
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return domain.StoredTextureMeta{}, err
	}

	defer func(file *os.File) {
		closeError := file.Close()
		if closeError != nil {
			fmt.Println("Error closing file:", closeError)
		}
	}(file)

	// Write binary data to the file
	_, err = file.Write(rawData)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return domain.StoredTextureMeta{}, err
	}

	return domain.StoredTextureMeta{
		Uri:    path,
		FileId: fileId,
	}, nil
}

func NewLocalTextureStorage(rootFolder string) domain.ITextureStorage {
	return &LocalTextureStorage{rootFolder: rootFolder}
}
