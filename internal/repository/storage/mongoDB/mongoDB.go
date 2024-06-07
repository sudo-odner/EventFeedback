package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	Client mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

func New() *MongoDB {

}
