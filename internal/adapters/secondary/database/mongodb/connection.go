package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection represents a MongoDB connection
type Connection struct {
	client *mongo.Client
}

// NewConnection creates a new MongoDB connection
func NewConnection(uri string) (*Connection, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB")
	return &Connection{client: client}, nil
}

// GetDatabase returns a database instance
func (c *Connection) GetDatabase(name string) *mongo.Database {
	return c.client.Database(name)
}

// Close closes the connection
func (c *Connection) Close(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}

// IsConnected checks if the connection is alive
func (c *Connection) IsConnected() bool {
	err := c.client.Ping(context.TODO(), nil)
	return err == nil
}
