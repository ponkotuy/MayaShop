package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func main() {
	m := martini.Classic()
	m.Get("/", Hello)
	m.Post("/user/create", binding.Bind(UserForm{}), CreateUser)
	m.Run()
}

type UserForm struct {
	Name string `form:"name" binding:"required"`
}
