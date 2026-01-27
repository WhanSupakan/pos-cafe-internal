package mongo

import (
	"cafe-pos/internal/domain/logger"
	"context"
	"time"
)

type MongoCollection interface {
	InsertOne(ctx context.Context, document interface{}) (interface{}, error)
}

type MongoLogger struct {
	collection MongoCollection
}

func NewMongoLogger(collection MongoCollection) logger.Logger {
	return &MongoLogger{
		collection: collection,
	}
}

func (l *MongoLogger) LogError(code, backendMessage string, details map[string]string, err error) {
	doc := map[string]interface{}{
		"code":           code,
		"backendMessage": backendMessage,
		"timestamp":      time.Now(),
		"details":        details,
	}

	if err != nil {
		doc["error"] = err.Error()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, _ = l.collection.InsertOne(ctx, doc)
}
