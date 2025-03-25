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

func (m MongoTextureRefRepository) GetTextureSetById(ctx context.Context, id string) (domain.TextureSet, error) {
	c := m.dbContext.textureRefsCollection.FindOne(ctx,
		bson.M{
			"_id": id,
		}, options.FindOne())

	var result TextureSetDataModel

	err := c.Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.TextureSet{}, domain.TextureRefNotFoundErr
		}
		return domain.TextureSet{}, err
	}

	return result.toDomain(), nil
}

func (m MongoTextureRefRepository) GetTextureSetByAddress(ctx context.Context, address domain.TextureAddress) (domain.TextureSet, error) {
	c := m.dbContext.textureRefsCollection.FindOne(ctx,
		bson.M{
			"address": bson.M{
				"world_id": address.WorldId,
				"x":        address.X,
				"y":        address.Y,
			},
		}, options.FindOne())

	var result TextureSetDataModel

	err := c.Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.TextureSet{}, domain.TextureRefNotFoundErr
		}
		return domain.TextureSet{}, err
	}

	return result.toDomain(), nil
}

func (m MongoTextureRefRepository) InsertTextureSet(ctx context.Context, model domain.TextureSet) (string, error) {
	data := NewTextureSetDataModel(model)
	o, err := m.dbContext.textureRefsCollection.InsertOne(ctx, data)
	return o.InsertedID.(bson.ObjectID).Hex(), err
}

func (m MongoTextureRefRepository) UpdateTextureSet(ctx context.Context, model domain.TextureSet) error {
	data := NewTextureSetDataModel(model)
	filter := bson.M{"_id": data.Id}
	c, err := m.dbContext.textureRefsCollection.ReplaceOne(ctx, filter, data)
	if err != nil {
		return err
	}
	if c.MatchedCount == 0 {
		return fmt.Errorf("texture set not updated")
	}
	return err
}

func NewMongoTextureRefRepository(ctx *MongoDbContext) domain.ITextureSetRepository {
	return &MongoTextureRefRepository{dbContext: ctx}
}

type RedisTaskRepository struct {
	redisClient *redis.Client
}

func (r RedisTaskRepository) DeleteTaskFromQueue(ctx context.Context, taskId string) error {
	return r.redisClient.SRem(ctx, taskQueueKey, taskId).Err()
}

func (r RedisTaskRepository) GetTaskQueue(ctx context.Context) ([]string, error) {
	return r.redisClient.SMembers(ctx, taskQueueKey).Result()
}

func (r RedisTaskRepository) DeleteTask(ctx context.Context, taskId string) error {
	key := taskId

	p := r.redisClient.Pipeline()

	err := p.SRem(ctx, taskQueueKey, key).Err()
	if err != nil {
		return err
	}

	err = p.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	_, err = p.Exec(ctx)

	return err
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

const taskQueueKey = "task_queue"

func (r RedisTaskRepository) InsertTextureRenderingTasks(ctx context.Context, tasks []domain.RenderImageTask) error {
	p := r.redisClient.Pipeline()

	for _, task := range tasks {
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(task)
		if err != nil {
			return err
		}
		key := task.Id
		err = p.Set(ctx, key, jsonData, 0).Err()
		if err != nil {
			return err
		}

		err = p.SAdd(ctx, taskQueueKey, task.Id).Err()
		if err != nil {
			return err
		}
	}

	_, err := p.Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func NewRedisTaskRepository(client *redis.Client) domain.ITaskRepository {
	return &RedisTaskRepository{
		redisClient: client,
	}
}

type LocalTextureStorage struct {
	rootFolder string
	uriPrefix  string
}

func (t LocalTextureStorage) DeleteTextureById(textureId string) error {
	var path = t.getFileAbsPath(textureId)
	return os.Remove(path)
}

func (t LocalTextureStorage) getFileUri(textureId string) string {
	return fmt.Sprintf("%s%s", t.uriPrefix, textureId)

}

func (t LocalTextureStorage) getFileAbsPath(textureId string) string {
	p, err := filepath.Abs(t.rootFolder)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(p, os.ModeDir|os.ModePerm)
	if err != nil {
		panic(err)
	}

	return filepath.Join(p, textureId)
}

func (t LocalTextureStorage) GetTextureById(textureId string) ([]byte, error) {
	var path = t.getFileAbsPath(textureId)

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

func (t LocalTextureStorage) StoreTexture(rawData []byte, extension string) (domain.TextureRef, error) {
	fileId := fmt.Sprintf("%s%s", uuid.New().String(), extension)
	// Create (or overwrite) the file
	var absPath = t.getFileAbsPath(fileId)
	var uri = t.getFileUri(fileId)
	file, err := os.Create(absPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return domain.TextureRef{}, err
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
		return domain.TextureRef{}, err
	}

	return domain.TextureRef{
		Uri:    uri,
		FileId: fileId,
	}, nil
}

func NewLocalTextureStorage(rootFolder string, uriPrefix string) domain.ITextureStorage {
	return &LocalTextureStorage{rootFolder: rootFolder, uriPrefix: uriPrefix}
}

type RedisTextureSetCache struct {
	client *redis.Client
}

func NewRedisTextureCache(client *redis.Client) domain.ITextureSetCache {
	return &RedisTextureSetCache{
		client: client,
	}
}

func getKeyFromId(id string) string {
	return fmt.Sprintf("texture_set_cache:%s", id)
}

func (r RedisTextureSetCache) GetTextureSetByAddress(ctx context.Context, address domain.TextureAddress) (domain.TextureSet, error) {
	key := getKeyFromId(address.String())
	jsonResult, err := r.client.Get(ctx, key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return domain.TextureSet{}, domain.ErrTextureSetNotFoundInCache
		}
		return domain.TextureSet{}, err
	}

	// Unmarshal back to struct
	var textureSet domain.TextureSet
	err = json.Unmarshal([]byte(jsonResult), &textureSet)

	if err != nil {
		return domain.TextureSet{}, err
	}

	return textureSet, nil
}

func (r RedisTextureSetCache) SetTextureSet(ctx context.Context, textureSet domain.TextureSet) (string, error) {
	jsonData, err := json.Marshal(textureSet)
	if err != nil {
		return "", err
	}

	key := getKeyFromId(textureSet.Address.String())

	err = r.client.Set(ctx, key, jsonData, 0).Err()
	if err != nil {
		return "", err
	}
	return key, nil
}

func (r RedisTextureSetCache) DelTextureSet(ctx context.Context, textureSet domain.TextureSet) error {
	key := getKeyFromId(textureSet.Address.String())
	return r.client.Del(ctx, key).Err()
}
