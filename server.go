package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

func main() {
	m := martini.Classic()
	m.Group("/api/v1", func(r martini.Router) {
		r.Post("/user/create", binding.Bind(UserForm{}), CreateUser)
		r.Get("/gear/all", GearAll)
	})
	m.Get("**", GetPublic)
	m.Run()
}

type UserForm struct {
	Name string `form:"name" binding:"required"`
}
