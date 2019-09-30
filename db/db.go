package db

import "github.com/globalsign/mgo"

var (
	url = "localhost:27017"
)

func GetSession() *mgo.Session{
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	return session

}