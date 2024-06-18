package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"modEventFeedback/internal/repository/storage"
)

// TODO
// Ошибки + лог
// Переделать тип выхода
// FindAllAnswerQuestion() +
// FindOneAnswerQuestion() +
// CreateAnswerQuestion() +
// SetAnswerQuestion()
// DeleteAnswerQuestion() +

// Какой то конвертор в эту ебень bson
// Find(count: int, filter: узнать как делать) -> struct: [struct, ...], err: error
// Create(item: struct) -> id: primitive.ObjectID, err: error
// Delete(filter: узнать как делать) -> err: error
// Set(item:struct, filter: узнать как делать)) -> err: error

// FindOneAnswerQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneAnswerQuestion(filter bson.D) storage.AnswerQuestion {
	table, collection := tableDB, "answerQuestion"
	var answerQuestion storage.AnswerQuestion

	newCollection, err := Connect(table, collection, db.clientOpts, db.cfg, db.log)
	if err != nil {
		// Ошибка
		// Выход из метода
	}
	defer newCollection.Disconnect()

	resultBson := newCollection.FindOne(filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &answerQuestion); err != nil {
		db.log.Error("Can't read result")
	}

	return answerQuestion
}

func (db *MongoDB) CreateAnswerQuestion(item storage.AnswerQuestion) primitive.ObjectID {
	table, collection := tableDB, "answerQuestion"

	newCollection, err := Connect(table, collection, db.clientOpts, db.cfg, db.log)
	if err != nil {
		// Ошибка
		// Выход из метода
	}
	defer newCollection.Disconnect()

	insertedID := newCollection.Create(item)

	return insertedID
}

func (db *MongoDB) SetAnswerQuestion(filter bson.D, set bson.D) {
	table, collection := tableDB, "answerQuestion"

	newCollection, err := Connect(table, collection, db.clientOpts, db.cfg, db.log)
	if err != nil {
		// Ошибка
		// Выход из метода
	}
	defer newCollection.Disconnect()

	// Creates instructions to add the "avg_rating" field to documents
	update := bson.D{{"$set", set}}

	newCollection.Set(update, filter)
}

func (db *MongoDB) DeleteAnswerQuestion(filter bson.D) {
	table, collection := tableDB, "answerQuestion"

	newCollection, err := Connect(table, collection, db.clientOpts, db.cfg, db.log)
	if err != nil {
		// Ошибка
		// Выход из метода
	}
	defer newCollection.Disconnect()

	newCollection.Delete(filter)
	newCollection.Disconnect()
}
