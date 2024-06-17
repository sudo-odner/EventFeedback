package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"modEventFeedback/internal/repository/storage"
)

// TODO
// FindAllLecture() +
// FindOneLecture() +
// CreateLecture()
// SetLecture()
// DeleteLecture()

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

// FindOneLecture The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneLecture(filter bson.D) storage.Lecture {
	table, collection := tableDB, "lecture"
	var lecture storage.Lecture

	resultBson := db.findOne(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &lecture); err != nil {
		db.log.Error("Can't read result")
	}

	return lecture
}
