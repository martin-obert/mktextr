package data_access

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	DATABASE_NAME               = "mktextr"
	TEXTURE_REF_COLLECTION_NAME = "texture_refs"
)

type MongoDbContext struct {
	client                *mongo.Client
	database              *mongo.Database
	textureRefsCollection *mongo.Collection
}

func NewMongoDbContext(connectionUri string) (*MongoDbContext, error) {
	client, err := mongo.Connect(options.Client().
		ApplyURI(connectionUri))
	if err != nil {
		return nil, err
	}
	database := client.Database(DATABASE_NAME)
	return &MongoDbContext{client: client,
		textureRefsCollection: database.Collection(TEXTURE_REF_COLLECTION_NAME),
		database:              database,
	}, nil
}

type TextureRefDataModel struct {
	Id      bson.ObjectID `bson:"_id,omitempty"`
	Uri     string        `bson:"uri"`
	FileId  string        `bson:"file_id"`
	WorldId string        `bson:"world_id"`
	X       int           `bson:"x"`
	Y       int           `bson:"y"`
}
