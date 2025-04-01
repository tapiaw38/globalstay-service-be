package nosql

import (
	"context"
	"log"

	"github.com/tapiaw38/reservation-service-be/internal/platform/config"
	"github.com/tapiaw38/reservation-service-be/internal/platform/migrations"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	Disconnect(context.Context) error
	Ping(context.Context) error
	GetCollection() *mongo.Collection
	RunMigrations(*mongo.Collection, []migrations.Migration) error
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type client struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewClient(cfg *config.NoSQLCollectionConfig) (Client, error) {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		cfg.DatabaseURI,
	))
	if err != nil {
		return nil, err
	}
	if err := mongoClient.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return &client{
		client:     mongoClient,
		database:   cfg.Database,
		collection: cfg.Collection,
	}, nil
}

func (m *client) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}

func (m *client) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, nil)
}

func (m *client) GetCollection() *mongo.Collection {
	return m.client.Database(m.database).Collection(m.collection)
}

func (m *client) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return m.GetCollection().InsertOne(ctx, document)
}

func (m *client) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.GetCollection().Find(ctx, filter, opts...)
}

func (m *client) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.GetCollection().FindOne(ctx, filter, opts...)
}

func (m *client) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return m.GetCollection().UpdateOne(ctx, filter, update, opts...)
}

func (m *client) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return m.GetCollection().DeleteOne(ctx, filter, opts...)
}

func (m *client) RunMigrations(migrationCollection *mongo.Collection, migrations []migrations.Migration) error {
	for _, migration := range migrations {
		var existing bson.M
		err := migrationCollection.FindOne(context.Background(), bson.M{
			"version":    migration.Version,
			"collection": m.GetCollection().Name(),
		}).Decode(&existing)
		if err == nil {
			log.Printf("Migration version %d already exists, skipping", migration.Version)
			continue
		}

		if err != mongo.ErrNoDocuments {
			return err
		}

		log.Printf("Running migration version %d", migration.Version)
		if err := migration.Up(m.client.Database(m.database)); err != nil {
			return err
		}

		_, err = migrationCollection.InsertOne(context.Background(), bson.M{
			"version":    migration.Version,
			"collection": m.GetCollection().Name(),
		})
		if err != nil {
			return err
		}
	}

	log.Println("All migrations are up to date")
	return nil
}
