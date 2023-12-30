package preference

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func NewPreferenceRepo() PreferenceRepo {
	return &PreferenceRepoImpl{}
}

func (r *PreferenceRepoImpl) Create(ctx context.Context, userId string) (UserPreference, error) {
	data := UserPreference{
		UserId:    userId,
		Tags:      []string{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if result, err := r.BaseRepo.Create(ctx, data); err != nil {
		return data, err
	} else {
		data.Id = result
	}
	return data, nil
}

func (r *PreferenceRepoImpl) FindByUserId(ctx context.Context, userId string) (UserPreference, error) {
	var data UserPreference
	err := r.FindOneByQuery(ctx, bson.M{"userId": userId}, &data)
	return data, err
}
