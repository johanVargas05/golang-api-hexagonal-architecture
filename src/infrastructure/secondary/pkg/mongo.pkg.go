package pkg

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	mongoOnce      sync.Once
)

const (
	connectionString       = "mongodb://%s:%s/test?retryWrites=false"
	connectTimeOut         = 10
	errorConnectionMessage = "error connecting to MongoDB: %v"
)

type MongoConnectionPort interface {
	Connection() *mongo.Database
	Close() error
}

func newConnection() *mongo.Client {
	mongoOnce.Do(func() {
		dbHost := os.Getenv("MONGO_DB_HOST")
		dbPort := os.Getenv("MONGO_DB_PORT")
		connectionURI := fmt.Sprintf(connectionString, dbHost, dbPort)
		clientInstance = connect(connectionURI)
	})
	return clientInstance
}

func connect(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeOut*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Errorf(errorConnectionMessage, err))
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Errorf(errorConnectionMessage, err))
	}

	return client
}



type MongoConnection struct{}

func NewMongoConnection() *MongoConnection {
	return &MongoConnection{}
}

func (*MongoConnection) Connection() *mongo.Database {
	dbName := os.Getenv("MONGO_DB_DATABASE")
	if clientInstance == nil {
		connection := newConnection()
		return connection.Database(dbName)
	}

	return clientInstance.Database(dbName)
}

func (*MongoConnection) Close() error {
	if clientInstance != nil {
		return clientInstance.Disconnect(context.Background())
	}
	return nil
}
