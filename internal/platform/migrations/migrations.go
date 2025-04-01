package migrations

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Migration struct {
	Version int
	Up      func(db *mongo.Database) error
}

func ExecuteBusinessMigrations(collectionName string) []Migration {
	return []Migration{
		// {
		// 	Version: 1,
		// 	Up: func(db *mongo.Database) error {
		// 		_, err := db.Collection(collectionName).UpdateMany(
		// 			context.Background(),
		// 			bson.M{},
		// 			bson.M{"$set": bson.M{"is_promoted": false}},
		// 		)
		// 		return err
		// 	},
		// },
		// Add more migrations here
	}
}
