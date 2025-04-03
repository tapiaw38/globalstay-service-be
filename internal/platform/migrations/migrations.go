package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Migration struct {
	Version int
	Up      func(db *mongo.Database) error
}

func ExecuteHotelMigrations(collectionName string) []Migration {
	return []Migration{
		{
			Version: 1,
			Up: func(db *mongo.Database) error {
				_, err := db.Collection(collectionName).UpdateMany(
					context.Background(),
					bson.M{},
					bson.M{"$set": bson.M{"location_name": "tinogasta"}},
				)
				return err
			},
		},
		// Add more migrations here
	}
}
