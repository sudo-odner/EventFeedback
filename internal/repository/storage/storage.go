package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create struct for "course", "lecture", "question", "answerQuestion"

type Course struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title"`
}

type Lecture struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	Title       string               `bson:"title,omitempty"`
	Description string               `bson:"description,omitempty"`
	Lector      string               `bson:"lector,omitempty"`
	QuestionsID []primitive.ObjectID `bson:"questionID,omitempty"`
	FeedbacksID []primitive.ObjectID `bson:"feedbackID,omitempty"`
}

type Question struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Title string             `bson:"title,omitempty"`
}

type AnswerQuestion struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Question   string             `bson:"question,omitempty"`
	Answer     string             `bson:"answer,omitempty"`
	IsRelevant bool               `bson:"isRelevant,omitempty"`
	IsPositive bool               `bson:"isPositive,omitempty"`
	Object     string             `bson:"object,omitempty"`
}
