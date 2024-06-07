package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDB struct {
	clientOpts *options.ClientOptions
}

func New(uri string) *MongoDB {
	clientOptions := options.Client().ApplyURI(uri)
	// Добавить проверку что он существует и происходить подключение
	return &MongoDB{
		clientOpts: clientOptions,
	}
}

func disconnect(client *mongo.Client, err error) {
	if err = client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err) // Описать
	}
}

func Ping(db *MongoDB) (err error) {
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

	if client, err := mongo.Connect(context.TODO(), db.clientOpts); err == nil {
		if err = client.Ping(context.TODO(), nil); err == nil {
			return nil
		}
	}
	return err
}

// Скорее всего каждый раз нужно их в методах использовать
// TODO: Create session to DB (Connect)
// TODO: Close session to DB

// TODO: Создать реализацию
