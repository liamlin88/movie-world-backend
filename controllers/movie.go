package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hello-beego-api/models"
	"log"
)

// Operations about object
type MovieController struct {
	beego.Controller
}

func (u *MovieController) GetMovies() {
	movies, _ := models.GetMovies()
	u.Data["json"] = movies
	u.ServeJSON()
}

func (u *MovieController) AddMovie() {
	var movie models.Movie
	json.Unmarshal(u.Ctx.Input.RequestBody, &movie)
	fmt.Println(movie)
	movie, _ = models.AddMovie(movie)
	u.Data["json"] = movie
	u.ServeJSON()
}

func (u *MovieController) DeleteMovie() {
	id := u.Ctx.Input.Param(":id")
	err := models.DeleteMovie(id)
	if err != nil {
		log.Println("delete", id, err)
	}
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *MovieController) UpdateMovie() {
	id := u.Ctx.Input.Param(":id")
	var movie models.Movie
	json.Unmarshal(u.Ctx.Input.RequestBody, &movie)

	log.Println(movie)
	err := models.UpdateMovie(id, movie)
	if err != nil {
		u.Data["json"] = err.Error() + " " + id
	} else {
		u.Data["json"] = "update success!"
	}
	u.ServeJSON()
}