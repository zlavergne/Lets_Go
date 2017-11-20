package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mgoSession   *mgo.Session
	databaseName = "exampleDB"
	location     = "localhost"
	collection   = "examples"
)

type exampleStruct struct {
	ID    string `bson:"id" json:"id"`
	Fname string `bson:"fname" json:"name"`
	Lname string `bson:"lname" json:"lname"`
}

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(location)
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

func usingCollection(collection string) *mgo.Collection {
	session := getSession()
	return session.DB(databaseName).C(collection)
}

func getExampleByID(id string) (example exampleStruct, err error) {
	err = usingCollection(collection).Find(bson.M{"id": id}).One(&example)
	return
}

func insertExample(example exampleStruct) (err error) {
	err = usingCollection(collection).Insert(example)
	return
}

func main() {

}
