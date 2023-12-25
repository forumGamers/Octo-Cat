package post

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewPostService(repo PostRepo) PostService {
	return &PostServiceImpl{repo}
}

func (s *PostServiceImpl) InsertManyAndBindIds(ctx context.Context, datas []Post) error {
	var payload []any

	for _, data := range datas {
		payload = append(payload, data)
	}

	ids, err := s.Repo.CreateMany(ctx, payload)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids.InsertedIDs); i++ {
		id := ids.InsertedIDs[i].(primitive.ObjectID)
		datas[i].Id = id
	}
	return nil
}
