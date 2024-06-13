package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"modEventFeedback/internal/config"
	"modEventFeedback/internal/repository/storage"
	"slices"
)

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

// Work with Course

func (db *MongoDB) findAll(tableName, collectionName string, filter bson.D) []bson.M {
	// Check correct collection and table name
	if tableName != tableDB {
		return []bson.M{bson.M{}}
	}
	if !slices.Contains(collectionDB, collectionName) {
		return []bson.M{bson.M{}}
	}
	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	// Create cursor in collection
	collection := client.Database(tableName).Collection(collectionName)
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		db.log.Error("[mongoDB][FindAll] Cursor not created")
	}
	defer db.closeCursor(cursor)

	// write result in bson.M

	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		db.log.Error("Error with write cursor in result", err)
	}

	return result
}

// FindAllCourse for find all element with filter bson.D
func (db *MongoDB) FindAllCourse(filter bson.D) []storage.Course {
	resultBson := db.findAll(tableDB, "course", filter)

	dataCourse := make([]storage.Course, 0, len(resultBson))
	for _, result := range resultBson {
		var course storage.Course
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &course); err != nil {
			db.log.Error("Can't read result")
		}
		dataCourse = append(dataCourse, course)
		fmt.Println(course)
	}
	return dataCourse
}

// FindAllLecture for find all element with filter bson.D
func (db *MongoDB) FindAllLecture(filter bson.D) []storage.Lecture {
	resultBson := db.findAll(tableDB, "lecture", filter)

	dataLecture := make([]storage.Lecture, 0, len(resultBson))
	for _, result := range resultBson {
		var lecture storage.Lecture
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &lecture); err != nil {
			db.log.Error("Can't read result")
		}
		dataLecture = append(dataLecture, lecture)
		fmt.Println(lecture)
	}
	return dataLecture
}

// FindAllQuestion for find all element with filter bson.D
func (db *MongoDB) FindAllQuestion(filter bson.D) []storage.Question {
	resultBson := db.findAll(tableDB, "lecture", filter)

	dataQuestion := make([]storage.Question, 0, len(resultBson))
	for _, result := range resultBson {
		var question storage.Question
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &question); err != nil {
			db.log.Error("Can't read result")
		}
		dataQuestion = append(dataQuestion, question)
		fmt.Println(question)
	}
	return dataQuestion
}

// FindAllAnswerQuestion for find all element with filter bson.D
func (db *MongoDB) FindAllAnswerQuestion(filter bson.D) []storage.AnswerQuestion {
	resultBson := db.findAll(tableDB, "lecture", filter)

	dataAnswerQuestion := make([]storage.AnswerQuestion, 0, len(resultBson))
	for _, result := range resultBson {
		var answerQuestion storage.AnswerQuestion
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &answerQuestion); err != nil {
			db.log.Error("Can't read result")
		}
		dataAnswerQuestion = append(dataAnswerQuestion, answerQuestion)
		fmt.Println(answerQuestion)
	}
	return dataAnswerQuestion
}

// TODO: Method drop database
// TODO: Расписать нужные методы

// Скорее всего каждый раз нужно их в методах использовать
// TODO: Create session to DB (Connect)
// TODO: Close session to DB

// TODO: Создать реализацию
