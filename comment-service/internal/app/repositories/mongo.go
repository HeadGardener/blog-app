package repositories

import (
	"context"
	"github.com/HeadGardener/blog-app/comment-service/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

const (
	commentsCollection = "comments"
)

func NewMongoDBCollection(ctx context.Context, config configs.DBConfig) (*mongo.Collection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port))
	if err != nil {
		return nil, err
	}

	ctxConn, _ := context.WithTimeout(ctx, 10*time.Second)
	err = client.Connect(ctxConn)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctxConn, readpref.Primary())
	if err != nil {
		return nil, err
	}

	db := client.Database(config.DBName)

	return db.Collection(commentsCollection), nil
}
