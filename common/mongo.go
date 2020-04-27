package common

import (
	"gopkg.in/mgo.v2"
)

type mongo struct {
	Tasks *mgo.Collection
}

var DB *mongo

// Start a MongoDB session
func connectDB() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	DB = &mongo{session.DB("tasksdb").C("tasks")}
}
