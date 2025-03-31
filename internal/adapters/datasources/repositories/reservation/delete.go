package reservation

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) Delete(ctx context.Context, id string) error {
	_, err := r.client.GetCollection().DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
