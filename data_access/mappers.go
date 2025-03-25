package data_access

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"mktextr/domain"
)

func (m *TextureSetDataModel) toDomain() domain.TextureSet {
	t := map[domain.TextureType]domain.TextureRef{}
	for k, v := range m.Textures {
		t[k] = v.toDomain()
	}
	return domain.TextureSet{
		Id:         m.Id.Hex(),
		TextureSet: t,
		Address:    m.Address.toDomain(),
	}
}

func (m *TextureRefDataModel) toDomain() domain.TextureRef {
	return domain.TextureRef{
		Uri:    m.Uri,
		FileId: m.FileId,
		State:  m.State,
	}
}

type TextureAddressDataModel struct {
	WorldId string `bson:"world_id"`
	X       int    `bson:"x"`
	Y       int    `bson:"y"`
}

func (m *TextureAddressDataModel) toDomain() domain.TextureAddress {
	return domain.TextureAddress{
		WorldId: m.WorldId,
		X:       m.X,
		Y:       m.Y,
	}
}

func NewTextureAddressDataModel(d domain.TextureAddress) TextureAddressDataModel {
	return TextureAddressDataModel{
		WorldId: d.WorldId,
		X:       d.X,
		Y:       d.Y,
	}
}

func NewTextureSetDataModel(d domain.TextureSet) TextureSetDataModel {
	t := map[domain.TextureType]TextureRefDataModel{}
	for textureType, ref := range d.TextureSet {
		t[textureType] = NewTextureRefDataModel(ref)
	}
	id, err := bson.ObjectIDFromHex(d.Id)
	if err != nil {
		panic(err)
	}
	return TextureSetDataModel{
		Id:       id,
		Textures: t,
		Address:  NewTextureAddressDataModel(d.Address),
	}
}
func NewTextureRefDataModel(d domain.TextureRef) TextureRefDataModel {
	return TextureRefDataModel{
		Uri:    d.Uri,
		FileId: d.FileId,
		State:  d.State,
	}
}
