package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"modEventFeedback/internal/config"
)

type MongoDB struct {
	log        *slog.Logger
	cfg        *config.MongoDB
	clientOpts *options.ClientOptions
}

// New Создание нового объекта подключения к MongoDB
func New(cfg *config.MongoDB, log *slog.Logger) *MongoDB {
	clientOptions := options.Client().ApplyURI(cfg.Uri)
	// Добавить проверку что он существует и происходить подключение
	return &MongoDB{
		log:        log,
		cfg:        cfg,
		clientOpts: clientOptions,
	}
}

func (db *MongoDB) closeConnection(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		db.log.Error("Error with disconnect client", err)
	}
}

func (db *MongoDB) Ping() {
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
		return
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		db.log.Error("Ping is not work", err)
		return
	}
	return
}

// Скорее всего каждый раз нужно их в методах использовать
// TODO: Create session to DB (Connect)
// TODO: Close session to DB

// TODO: Создать реализацию
