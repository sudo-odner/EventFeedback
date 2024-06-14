package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"modEventFeedback/internal/repository/storage"
)

// TODO
// FindAllAnswerQuestion() +
// FindOneAnswerQuestion() +
// CreateAnswerQuestion()
// SetAnswerQuestion()
// DeleteAnswerQuestion()

// FindAllAnswerQuestion The Tool find all answer on question by filter, *option
func (db *MongoDB) FindAllAnswerQuestion(filter bson.D) []storage.AnswerQuestion {
	tableName, collectionName := tableDB, "answerQuestion"
	resultBson := db.findAll(tableName, collectionName, filter)

	dataAnswerQuestion := make([]storage.AnswerQuestion, 0, len(resultBson))
	for _, result := range resultBson {
		var answerQuestion storage.AnswerQuestion
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &answerQuestion); err != nil {
			db.log.Error("Can't read result")
		}
		dataAnswerQuestion = append(dataAnswerQuestion, answerQuestion)
	}
	return dataAnswerQuestion
}

// FindOneAnswerQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneAnswerQuestion(filter bson.D) storage.AnswerQuestion {
	table, collection := tableDB, "answerQuestion"
	var answerQuestion storage.AnswerQuestion

	resultBson := db.findAll(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &answerQuestion); err != nil {
		db.log.Error("Can't read result")
	}

	return answerQuestion
}
