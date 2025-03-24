package data_access

import "mktextr/domain"

func toDomain(model TextureRefDataModel) domain.TextureRef {
	return domain.TextureRef{
		Id:  model.id.String(),
		Uri: model.uri,
	}
}

func toDomainArr(models []TextureRefDataModel) []domain.TextureRef {
	result := make([]domain.TextureRef, 0, len(models)) // preallocate capacity for performance

	for _, model := range models {
		result = append(result, toDomain(model))
	}

	return result
}
