package models

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"hello-beego-api/db"
	"log"
)

const (
	host = ""
)

var (
	Movies map[string]*Movie
)

type Movie struct {
	ObjectId bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Director string        `bson:"director"`
}

func init() {
	Movies = make(map[string]*Movie)
}

func GetMovies() ([]Movie, error) {
	session := db.GetSession()
	defer session.Close()
	c := session.DB("hellodb").C("movies")

	result := make([]Movie, 0, 10)
	err := c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}

func AddMovie(movie Movie) (Movie, error) {
	session := db.GetSession()
	defer session.Close()
	c := session.DB("hellodb").C("movies")

	id := bson.NewObjectId()
	movie.ObjectId = id
	log.Println(movie)
	err := c.Insert(&movie)
	if err != nil {
		log.Println(err)
	}
	return movie, err
}

func DeleteMovie(id string) error {
	session := db.GetSession()
	defer session.Close()
	c := session.DB("hellodb").C("movies")

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func UpdateMovie(id string, movie Movie) (err error) {
	log.Println(id, movie)
	if bson.ObjectIdHex(id) != movie.ObjectId {
		return errors.New("id not match")
	}

	session := db.GetSession()
	defer session.Close()
	c := session.DB("hellodb").C("movies")

	log.Println(movie, "inmodels")
	err = c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, movie)

	return err
}
