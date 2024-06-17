package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
	"modEventFeedback/internal/config"
)

type ICollection interface {
	FindAll(filter bson.D) []bson.M
	FindOne(filter bson.D) bson.M
	Create(item interface{}) primitive.ObjectID
	Set(updateBson bson.D, filter bson.D)
	Delete(filter bson.D)
	Disconnect()
}

// TODO: Получше расписать ошибки + выход с err описать
// TODO: Изменить работу с методами

type Collection struct {
	client     *mongo.Client
	collection *mongo.Collection
	cfg        *config.MongoDB
	log        *slog.Logger
}

func Connect(table, collection string, clientOpts *options.ClientOptions, cfg *config.MongoDB, log *slog.Logger) (ICollection, error) {
	// Check correct name collection and table
	if err := accessConnectToTableAndCollection(table, collection); err != nil {
		log.Error("access is close, err:", err)
		return nil, err
	}
	// Connect to mongoDB without disconnect
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Error("Connection with MongoDB is not created", err)
		return nil, err
	}
	// Create connect to collection
	coll := client.Database(table).Collection(collection)
	return &Collection{
		client:     client,
		collection: coll,
		cfg:        cfg,
		log:        log,
	}, nil
}

func (c *Collection) closeCursor(cursor *mongo.Cursor) {
	if err := cursor.Close(context.TODO()); err != nil {
		c.log.Error("Error with close cursor", err)
	}
}

func (c *Collection) Disconnect() {
	if err := c.client.Disconnect(context.TODO()); err != nil {
		c.log.Error("Error with disconnect client", err)
	}
}

func (c *Collection) FindAll(filter bson.D) []bson.M {
	// Create cursor in collection
	cursor, err := c.collection.Find(context.TODO(), filter)
	if err != nil {
		c.log.Error("[mongoDB][FindAll] Cursor not created")
	}
	defer c.closeCursor(cursor)

	// Write result in []bson.M
	var result []bson.M

	if err = cursor.All(context.TODO(), &result); err != nil {
		c.log.Error("Error with write cursor in result", err)
	}

	return result
}

func (c *Collection) FindOne(filter bson.D) bson.M {
	// Write result in bson.M
	var result bson.M

	if err := c.collection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		c.log.Error("Error with find one element in answerQuestion collection:", err)
	}
	return result
}

func (c *Collection) Create(item interface{}) primitive.ObjectID {
	insertResult, err := c.collection.InsertOne(context.TODO(), item)
	if err != nil {
		c.log.Error("New item not created:", err)
	}
	insertedID := insertResult.InsertedID.(primitive.ObjectID)

	return insertedID
}

func (c *Collection) Set(updateBson bson.D, filter bson.D) {
	//filter := bson.D{{"_id", id}}
	// Creates instructions to add the "avg_rating" field to documents
	update := bson.D{{"$set", updateBson}}
	// Updates the first document that has the specified "_id" value
	_, err := c.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.log.Error("Error with set object: ", err)
	}
}

func (c *Collection) Delete(filter bson.D) {
	// Deletes the first document that has a "title" value of "Twilight"
	_, err := c.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		c.log.Error("Error with delete object: ", err)
	}
}
