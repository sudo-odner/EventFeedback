package storage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Create struct for "course", "lecture", "question", "answerQuestion"
// Полезная ссылка https://stackoverflow.com/questions/39153419/golang-mongo-insert-with-self-generated-id-using-bson-newobjectid-resulting-i

type Course struct {
	ID    *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string              `bson:"title,omitempty" json:"title,omitempty"`
}

type Lecture struct {
	ID          *primitive.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string               `bson:"title,omitempty" json:"title,omitempty"`
	Description string               `bson:"description,omitempty" json:"description,omitempty"`
	Lector      string               `bson:"lector,omitempty" json:"lector,omitempty"`
	QuestionsID []primitive.ObjectID `bson:"questionID,omitempty" json:"questionID,omitempty""`
	FeedbacksID []primitive.ObjectID `bson:"feedbackID,omitempty" json:"feedbackID,omitempty"`
}

type Question struct {
	ID    *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string              `bson:"title,omitempty" json:"title,omitempty"`
}

type AnswerQuestion struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Question   string              `bson:"question,omitempty" json:"question,omitempty"`
	Answer     string              `bson:"answer,omitempty" json:"answer,omitempty"`
	IsRelevant bool                `bson:"isRelevant,omitempty" json:"isRelevant,omitempty"`
	IsPositive bool                `bson:"isPositive,omitempty" json:"isPositive,omitempty"`
	Object     string              `bson:"object,omitempty" json:"object,omitempty"`
}
