package nosql

import (
	"context"

	"github.com/tapiaw38/reservation-service-be/internal/platform/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client interface {
	Disconnect(context.Context) error
	Ping(context.Context) error
	GetCollection() *mongo.Collection
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
