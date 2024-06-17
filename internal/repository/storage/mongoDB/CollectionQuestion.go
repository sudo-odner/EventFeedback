package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"modEventFeedback/internal/repository/storage"
)

// TODO
// FindAllQuestion() +
// FindOneQuestion() +
// CreateQuestion()
// SetQuestion()
// DeleteQuestion()

func (db *MongoDB) TestFindAllQuestion(filter bson.D) []storage.Question {

	tableName, collectionName := tableDB, "question"
	resultBson := db.findAll(tableName, collectionName, filter)

	dataQuestion := make([]storage.Question, 0, len(resultBson))
	for _, result := range resultBson {
		var question storage.Question
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &question); err != nil {
			db.log.Error("Can't read result")
		}
		dataQuestion = append(dataQuestion, question)
	}
	return dataQuestion
}

// FindAllQuestion The Tool find all course by question, *option
func (db *MongoDB) FindAllQuestion(filter bson.D) []storage.Question {
	tableName, collectionName := tableDB, "question"
	resultBson := db.findAll(tableName, collectionName, filter)

	dataQuestion := make([]storage.Question, 0, len(resultBson))
	for _, result := range resultBson {
		var question storage.Question
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &question); err != nil {
			db.log.Error("Can't read result")
		}
		dataQuestion = append(dataQuestion, question)
	}
	return dataQuestion
}

// FindOneQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneQuestion(filter bson.D) storage.Question {
	table, collection := tableDB, "question"
	var question storage.Question

	resultBson := db.findOne(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &question); err != nil {
		db.log.Error("Can't read result")
	}

	return question
}
