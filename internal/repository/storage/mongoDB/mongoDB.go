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

// Проверка на дурака, что бы не обращаться к колекциям и бд вне зоны работы
func accessConnectToTableAndCollection(table, collection string) error {
	if table != tableDB {
		return errors.New(fmt.Sprintf("table %s is not exist", table))
	}
	if !slices.Contains(collectionDB, collection) {
		return errors.New(fmt.Sprintf("collection %s is not exist", collection))
	}
	return nil
}

// подтянуть из конфига
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

// Еhe "Find" tool for searching by filter, *option and automatically writing to the structure type
// [storage.Course, storage.Lecture, storage.Question storage.AnswerQuestion]
// TODO: Edit log answer on err
// TODO: Добавить фильтрацию по option
//
//// findAll tool for find all element in collection by filter and *option
//func (db *MongoDB) findAll(table, collection string, filter bson.D) []bson.M {
//	// Check correct name collection and table
//	if err := accessConnectToTableAndCollection(table, collection); err != nil {
//		db.log.Error("access is close, err:", err)
//		return []bson.M{}
//	}
//
//	// Connect to mongoDB
//	client, err := mongo.Connect(context.TODO(), db.clientOpts)
//	defer db.closeConnection(client)
//	if err != nil {
//		db.log.Error("Connection with MongoDB is not created", err)
//	}
//
//	// Create connect to collection
//	c := client.Database(table).Collection(collection)
//
//	// Create cursor in collection
//	cursor, err := c.Find(context.TODO(), filter)
//	if err != nil {
//		db.log.Error("[mongoDB][FindAll] Cursor not created")
//	}
//	defer db.closeCursor(cursor)
//
//	// Write result in []bson.M
//	var result []bson.M
//
//	if err = cursor.All(context.TODO(), &result); err != nil {
//		db.log.Error("Error with write cursor in result", err)
//	}
//
//	return result
//}
//
//// findOne tool for find one element in collection by filter and *option
//func (db *MongoDB) findOne(table, collection string, filter bson.D) bson.M {
//	// Check correct name collection and table
//	if err := accessConnectToTableAndCollection(table, collection); err != nil {
//		db.log.Error("access is close, err:", err)
//		return bson.M{}
//	}
//
//	// Connect to mongoDB
//	client, err := mongo.Connect(context.TODO(), db.clientOpts)
//	defer db.closeConnection(client)
//	if err != nil {
//		db.log.Error("Connection with MongoDB is not created", err)
//		return bson.M{}
//	}
//
//	// Create connect to collection
//	c := client.Database(table).Collection(collection)
//
//	// Write result in bson.M
//	var result bson.M
//
//	if err := c.FindOne(context.TODO(), filter).Decode(&result); err != nil {
//		db.log.Error("Error with find one element in answerQuestion collection:", err)
//	}
//	return result
//}
//
//// connectToCollection to collection
//func (db *MongoDB) connectToCollection(table, collection string) (*mongo.Collection, error) {
//	// Check correct name collection and table
//	if err := accessConnectToTableAndCollection(table, collection); err != nil {
//		db.log.Error("Access is close, err:", err)
//		return nil, err
//	}
//	// Connect to mongoDB
//	client, err := mongo.Connect(context.TODO(), db.clientOpts)
//	if err != nil {
//		db.log.Error("Connection with MongoDB is not created", err)
//		return nil, err
//	}
//	coll := client.Database(table).Collection(collection)
//	return coll, nil
//}
//
//// CreateItem
//func (db *MongoDB) createItem(table, collection string, item interface{}) primitive.ObjectID {
//	client, err := db.connectToCollection(table, collection)
//	if err != nil {
//		// Написать выход из функции
//	}
//	defer db.closeConnection(client)
//	// Create connect to collection
//	c := client.Database(table).Collection(collection)
//	// Insert one element in collection
//	insertResult, err := c.InsertOne(context.TODO(), item)
//	if err != nil {
//		db.log.Error("New item not created:", err)
//	}
//	insertedID := insertResult.InsertedID.(primitive.ObjectID)
//
//	return insertedID
//}
//
//// deleteItem
//func (db *MongoDB) deleteItem(table, collection string, filter bson.D) {
//	// Check correct name collection and table
//	if err := accessConnectToTableAndCollection(table, collection); err != nil {
//		db.log.Error("Access is close, err:", err)
//	}
//	// Connect to mongoDB
//	client, err := mongo.Connect(context.TODO(), db.clientOpts)
//	defer db.closeConnection(client)
//	if err != nil {
//		db.log.Error("Connection with MongoDB is not created", err)
//	}
//
//	// Create connect to collection
//	c := client.Database(table).Collection(collection)
//}
//
//// https://www.mongodb.com/docs/drivers/go/current/quick-start/ - Для вставки и
//// TODO: Method drop database
//// TODO: Расписать нужные методы
