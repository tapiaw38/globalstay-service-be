package business

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Delete(ctx context.Context, id string) (string, error) {
	result, err := r.client.GetCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return "", err
	}

	if result.DeletedCount == 0 {
		return "", nil
	}

	return id, nil
}
