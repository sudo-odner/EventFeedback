package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"modEventFeedback/internal/repository/storage"
)

// TODO
// FindAllAnswerQuestion() +
// FindOneAnswerQuestion() +
// CreateAnswerQuestion() +
// SetAnswerQuestion()
// DeleteAnswerQuestion() +

// FindOneAnswerQuestion The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneAnswerQuestion(filter bson.D) storage.AnswerQuestion {
	table, collection := tableDB, "answerQuestion"
	var answerQuestion storage.AnswerQuestion

	resultBson := db.findOne(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &answerQuestion); err != nil {
		db.log.Error("Can't read result")
	}

	return answerQuestion
}

func (db *MongoDB) CreateAnswerQuestion(item storage.AnswerQuestion) primitive.ObjectID {
	table, collection := tableDB, "answerQuestion"

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	// Create connect to collection
	c := client.Database(table).Collection(collection)

	insertResult, err := c.InsertOne(context.TODO(), item)
	if err != nil {
		db.log.Error("New item not created:", err)
	}
	insertedID := insertResult.InsertedID.(primitive.ObjectID)

	return insertedID
}

func (db *MongoDB) SetAnswerQuestion(filter bson.D, set bson.D) {
	table, collection := tableDB, "answerQuestion"

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	// Create connect to collection
	c := client.Database(table).Collection(collection)

	// Creates instructions to add the "avg_rating" field to documents
	update := bson.D{{"$set", set}}

	// Updates the first document that has the specified "_id" value
	_, err = c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
}

func (db *MongoDB) DeleteAnswerQuestion(filter bson.D) {
	table, collection := tableDB, "answerQuestion"

	// Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), db.clientOpts)
	defer db.closeConnection(client)
	if err != nil {
		db.log.Error("Connection with MongoDB is not created", err)
	}

	// Create connect to collection
	c := client.Database(table).Collection(collection)

	// Deletes the first document that has a "title" value of "Twilight"
	_, err = c.DeleteOne(context.TODO(), filter)
	if err != nil {
		db.log.Error("Error with delete object: ", err)
	}
}
