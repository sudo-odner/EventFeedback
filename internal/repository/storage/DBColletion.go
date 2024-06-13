package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create struct for "course", "lecture", "question", "answerQuestion"

type Course struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title"`
}

type Lecture struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Title       string               `bson:"title"`
	Description string               `bson:"description"`
	Lector      string               `bson:"lector"`
	QuestionsID []primitive.ObjectID `bson:"questionID"`
	FeedbacksID []primitive.ObjectID `bson:"feedbackID"`
}

type Question struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title"`
}

type AnswerQuestion struct {
	ID         primitive.ObjectID `bson:"_id"`
	Question   string             `bson:"question"`
	Answer     string             `bson:"answer"`
	IsRelevant bool               `bson:"isRelevant"`
	IsPositive bool               `bson:"isPositive"`
	Object     string             `bson:"object"`
}
