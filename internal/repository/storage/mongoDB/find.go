package mongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"modEventFeedback/internal/repository/storage"
	"slices"
)

// Еhe "Find" tool for searching by filter, *option and automatically writing to the structure type
// [storage.Course, storage.Lecture, storage.Question storage.AnswerQuestion]
// TODO: Edit log answer on err
// TODO: ПОдумать на счет уменьшения кода *findOne и *findAll
// TODO: Добавить фильтрацию по option

// findAll tool for findAll element in collection by filter and *option
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

	// Write result in bson.M
	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		db.log.Error("Error with write cursor in result", err)
	}

	return result
}

// FindAllCourse The Tool find all course by filter, *option
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

// FindAllLecture The Tool find all lecture by filter, *option
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

// FindAllQuestion The Tool find all course by question, *option
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

// FindAllAnswerQuestion The Tool find all answer on question by filter, *option
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

/* Дописать. Посмотреть как передовать и возвращять разные структуры
func (db *MongoDB) findOne(filter bson.D, tableName, collectionName string, struc interface{}) interface{}{
	var answerQuestion storage.AnswerQuestion
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

	collection := client.Database(tableName).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), filter).Decode(&answerQuestion); err != nil {
		db.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return answerQuestion
}
*/

// FindOneAnswerQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneAnswerQuestion(filter bson.D) storage.AnswerQuestion {
	tableName, collectionName := tableDB, "answerQuestion"
	var answerQuestion storage.AnswerQuestion

	// Check correct collection and table name
	if !slices.Contains(collectionDB, collectionName) {
		return answerQuestion
	}
	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	collection := client.Database(tableName).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), filter).Decode(&answerQuestion); err != nil {
		db.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return answerQuestion
}

// FindOneLecture The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneLecture(filter bson.D) storage.Lecture {
	tableName, collectionName := tableDB, "lecture"
	var lecture storage.Lecture

	// Check correct collection and table name
	if !slices.Contains(collectionDB, collectionName) {
		return lecture
	}
	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	collection := client.Database(tableName).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), filter).Decode(&lecture); err != nil {
		db.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return lecture
}

// FindOneQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneQuestion(filter bson.D) storage.Question {
	tableName, collectionName := tableDB, "question"
	var question storage.Question

	// Check correct collection and table name
	if !slices.Contains(collectionDB, collectionName) {
		return question
	}
	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	collection := client.Database(tableName).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), filter).Decode(&question); err != nil {
		db.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return question
}

// FindOneCourse The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneCourse(filter bson.D) storage.Course {
	tableName, collectionName := tableDB, "course"
	var course storage.Course

	// Check correct collection and table name
	if !slices.Contains(collectionDB, collectionName) {
		return course
	}
	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	collection := client.Database(tableName).Collection(collectionName)
	if err := collection.FindOne(context.TODO(), filter).Decode(&course); err != nil {
		db.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return course
}
