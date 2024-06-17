package mongoDB

import (
	"go.mongodb.org/mongo-driver/bson"
	"modEventFeedback/internal/repository/storage"
)

// TODO
// FindAllCourse() +
// FindOneCourse() +
// CreateCourse()
// SetCourse()
// DeleteCourse()

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

// FindOneCourse The Tool find one answer on question by filter, *option
func (db *MongoDB) FindOneCourse(filter bson.D) storage.Course {
	table, collection := tableDB, "course"
	var course storage.Course

	resultBson := db.findOne(table, collection, filter)
	bsonBytes, _ := bson.Marshal(resultBson)
	if err := bson.Unmarshal(bsonBytes, &course); err != nil {
		db.log.Error("Can't read result")
	}

	return course
}
