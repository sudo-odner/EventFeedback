package mongoDB

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"modEventFeedback/internal/config"
	"slices"
)

func accessConnectToTableAndCollection(table, collection string) error {
	if table != tableDB {
		return errors.New(fmt.Sprintf("table %s is not exist", table))
	}
	if !slices.Contains(collectionDB, collection) {
		return errors.New(fmt.Sprintf("collection %s is not exist", collection))
	}
	return nil
}

var (
	tableDB      = "feedback"
	collectionDB = []string{"course", "lecture", "question", "answerQuestion"}
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

func (db *MongoDB) closeCursor(client *mongo.Cursor) {
	if err := client.Close(context.TODO()); err != nil {
		db.log.Error("Error with close cursor", err)
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

func (db *MongoDB) CreateDatabaseFeedback() {
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	if result, err := client.ListDatabaseNames(context.TODO(), bson.D{{"name", tableDB}}); err != nil {
		fmt.Println(result)
		if len(result) == 0 {
			table := client.Database(tableDB)
			for _, item := range collectionDB {
				if err := table.CreateCollection(context.TODO(), item); err != nil {
					msg := fmt.Sprintf("Creating collection %s has error:", item)
					db.log.Error(msg, err)
				}
			}
			db.log.Info("[New DataBase] DataBase feedback and collection [course,lecture,question,answerQuestion] created")
		} else {
			db.log.Info("DataBase is already created. Check collection")
			table := client.Database(tableDB)
			resul, _ := table.ListCollectionNames(context.TODO(), bson.D{})
			for _, item := range resul {
				if slices.Contains(collectionDB, item) {
					msg := fmt.Sprintf("Collection %s is alrady created", item)
					db.log.Info(msg)
				} else {
					if err := table.CreateCollection(context.TODO(), item); err != nil {
						msg := fmt.Sprintf("Creating collection %s has error:", item)
						db.log.Error(msg, err)
					}
					msg := fmt.Sprintf("Collection %s is created", item)
					db.log.Info(msg)
				}
				db.log.Info("[Collection] DataBase feedback and collection [course,lecture,question,answerQuestion] created")
			}
		}
	}
}

// TODO: Method drop database
// TODO: Расписать нужные методы

// Скорее всего каждый раз нужно их в методах использовать
// TODO: Create session to DB (Connect)
// TODO: Close session to DB

// TODO: Создать реализацию
