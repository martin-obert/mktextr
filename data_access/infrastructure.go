package data_access

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"mktextr/m/domain"
)

type MongoTextureRefRepository struct {
	dbContext *MongoDbContext
}

func (m MongoTextureRefRepository) InsertTextureRef(ctx context.Context, uri string) (domain.TextureRef, error) {
	var data = TextureRefDataModel{
		uri: uri,
	}
	one, err := m.dbContext.textureRefsCollection.InsertOne(ctx, data)
	if err != nil {
		return domain.TextureRef{}, err
	}

	return domain.TextureRef{
		Id:  one.InsertedID.(bson.ObjectID).String(),
		Uri: uri,
	}, nil

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

func NewMongoTextureRefRepository(ctx *MongoDbContext) *MongoTextureRefRepository {
	return &MongoTextureRefRepository{dbContext: ctx}
}
