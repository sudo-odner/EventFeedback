package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"modEventFeedback/internal/repository/storage"
)

// Еhe "Find" tool for searching by filter, *option and automatically writing to the structure type
// [storage.Course, storage.Lecture, storage.Question storage.AnswerQuestion]
// TODO: Edit log answer on err
// TODO: Добавить фильтрацию по option

// findAll tool for find all element in collection by filter and *option
func (db *MongoDB) findAll(table, collection string, filter bson.D) []bson.M {
	// Check correct name collection and table
	if err := accessConnectToTableAndCollection(table, collection); err != nil {
		db.log.Error("access is close, err:", err)
		return []bson.M{}
	}

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	// Create connect to collection
	c := client.Database(table).Collection(collection)

	// Create cursor in collection
	cursor, err := c.Find(context.TODO(), filter)
	if err != nil {
		db.log.Error("[mongoDB][FindAll] Cursor not created")
	}
	defer db.closeCursor(cursor)

	// Write result in []bson.M
	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		db.log.Error("Error with write cursor in result", err)
	}

	return result
}

// FindAllCourse The Tool find all course by filter, *option
func (db *MongoDB) FindAllCourse(filter bson.D) []storage.Course {
	tableName, collectionName := tableDB, "course"
	resultBson := db.findAll(tableName, collectionName, filter)

	dataCourse := make([]storage.Course, 0, len(resultBson))
	for _, result := range resultBson {
		var course storage.Course
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &course); err != nil {
			db.log.Error("Can't read result")
		}
		dataCourse = append(dataCourse, course)
	}
	return dataCourse
}

// FindAllLecture The Tool find all lecture by filter, *option
func (db *MongoDB) FindAllLecture(filter bson.D) []storage.Lecture {
	tableName, collectionName := tableDB, "lecture"
	resultBson := db.findAll(tableName, collectionName, filter)

	dataLecture := make([]storage.Lecture, 0, len(resultBson))
	for _, result := range resultBson {
		var lecture storage.Lecture
		bsonBytes, _ := bson.Marshal(result)
		if err := bson.Unmarshal(bsonBytes, &lecture); err != nil {
			db.log.Error("Can't read result")
		}
		dataLecture = append(dataLecture, lecture)
	}
	return dataLecture
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

// findOne tool for find one element in collection by filter and *option
func (db *MongoDB) findOne(table, collection string, filter bson.D) bson.M {
	// Check correct name collection and table
	if err := accessConnectToTableAndCollection(table, collection); err != nil {
		db.log.Error("access is close, err:", err)
		return bson.M{}
	}

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
		return bson.M{}
	}

	// Create connect to collection
	c := client.Database(table).Collection(collection)

	// Write result in bson.M
	var result bson.M

	if err := c.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		db.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return result
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

// FindOneLecture The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneLecture(filter bson.D) storage.Lecture {
	table, collection := tableDB, "lecture"
	var lecture storage.Lecture

	resultBson := db.findAll(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &lecture); err != nil {
		db.log.Error("Can't read result")
	}

	return lecture
}

// FindOneQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneQuestion(filter bson.D) storage.Question {
	table, collection := tableDB, "question"
	var question storage.Question

	resultBson := db.findAll(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &question); err != nil {
		db.log.Error("Can't read result")
	}

	return question
}

// FindOneCourse The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneCourse(filter bson.D) storage.Course {
	table, collection := tableDB, "course"
	var course storage.Course

	resultBson := db.findAll(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &course); err != nil {
		db.log.Error("Can't read result")
	}

	return course
}
