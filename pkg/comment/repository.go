package comment

import (
	"context"

	b "github.com/forumGamers/Octo-Cat/pkg/base"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCommentRepo() CommentRepo {
	return &CommentRepoImpl{b.NewBaseRepo(b.GetCollection(b.Comment))}
}

func (r *CommentRepoImpl) CreateComment(ctx context.Context, data *Comment) error {
	result, err := r.Create(ctx, &data)
	if err != nil {
		return err
	}
	data.Id = result
	return nil
}

func (r *CommentRepoImpl) FindById(ctx context.Context, id primitive.ObjectID, data *Comment) error {
	return r.FindOneById(ctx, id, data)
}

func (r *CommentRepoImpl) DeleteOne(ctx context.Context, id primitive.ObjectID) error {
	return r.DeleteOneById(ctx, id)
}

func (r *CommentRepoImpl) CreateMany(ctx context.Context, datas []any) (*mongo.InsertManyResult, error) {
	return r.InsertMany(ctx, datas)
}

func (r *CommentRepoImpl) CreateReply(ctx context.Context, id primitive.ObjectID, data *ReplyComment) error {
	if _, err := r.UpdateOneByQuery(ctx, id, bson.M{
		"$push": bson.M{
			"reply": data,
		},
	}); err != nil {
		return err
	}
	return nil
}

func (r *CommentRepoImpl) DeleteReplyByPostId(ctx context.Context, postId primitive.ObjectID) error {
	cursor, err := r.FindByQuery(ctx, bson.M{"postId": postId})
	if err != nil {
		return err
	}

	defer cursor.Close(ctx)
	var commentIds []primitive.ObjectID
	for cursor.Next(ctx) {
		var comment struct {
			CommentId primitive.ObjectID `bson:"_id"`
		}
		if err := cursor.Decode(&comment); err != nil {
			return err
		}
		commentIds = append(commentIds, comment.CommentId)
	}

	if len(commentIds) > 0 {
		if err := r.DeleteManyByQuery(ctx, bson.M{
			"commentId": bson.M{
				"$in": commentIds,
			},
		}); err != nil {
			return err
		}
	}
	return nil
}

func (r *CommentRepoImpl) FindReplyById(ctx context.Context, id, replyId primitive.ObjectID, data *ReplyComment) error {
	if err := r.FindOneByQuery(ctx, bson.M{
		"_id": id,
		"reply": bson.M{
			"$elemMatch": bson.M{
				"_id": replyId,
			},
		},
	}, &data); err != nil {
		return err
	}
	return nil
}

func (r *CommentRepoImpl) DeleteOneReply(ctx context.Context, id, replyId primitive.ObjectID) error {
	if _, err := r.UpdateOneByQuery(ctx, id, bson.M{
		"reply": bson.M{
			"$pull": bson.M{
				"$elemMatch": bson.M{
					"_id": replyId,
				},
			},
		},
	}); err != nil {
		return err
	}
	return nil
}

func (r *CommentRepoImpl) DeleteMany(ctx context.Context, postId primitive.ObjectID) error {
	return r.DeleteManyByQuery(ctx, bson.M{"postId": postId})
}
