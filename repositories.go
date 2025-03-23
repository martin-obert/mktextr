package mktextr

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ITextureRefRepository interface {
	InsertTextureRef(uri string) (TextureRef, error)
	GetTextureRefById(texRefId uuid.UUID) (TextureRef, error)
	QueryTextureRefs() ([]TextureRef, error)
}

type MongoTextureRefRepository struct {
	client *mongo.Client
}

func (r MongoTextureRefRepository) getTextureCollection() *mongo.Collection {
	return r.client.Database("test").Collection("test")
}

func (r MongoTextureRefRepository) InsertTextureRef(ctx context.Context, uri string) (TextureRef, error) {
	var ref = TextureRefDataModel{
		uri: uri,
	}

	one, err := r.getTextureCollection().InsertOne(ctx, ref, options.InsertOne())

	if err != nil {
		return TextureRef{}, err
	}

	return TextureRef{
		id:  one.InsertedID.(bson.ObjectID).String(),
		uri: uri,
	}, nil
}

func (r MongoTextureRefRepository) GetTextureRefById(texRefId uuid.UUID) (TextureRef, error) {
	//TODO implement me
	panic("implement me")
}

func (r MongoTextureRefRepository) QueryTextureRefs() ([]TextureRef, error) {
	//TODO implement me
	panic("implement me")
}

func NewMongoTextureRefRepository(connectionUri string) (ITextureRefRepository, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI(connectionUri))
	if err != nil {
		return nil, err
	}
	return &MongoTextureRefRepository{client: client}, nil
}
