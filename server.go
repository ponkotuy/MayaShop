package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/auth"
)

func main() {
	m := martini.Classic()
	m.Group("/api/v1", func(r martini.Router) {
		r.Get("/user/info", UserInfo)
	}, auth.BasicFunc(MayaAuth))
	m.Group("/open/v1", func(r martini.Router) {
		r.Post("/user/create", binding.Bind(UserForm{}), CreateUser)
		r.Get("/gear/all", GearAll)
	})
	m.Get("**", GetPublic)
	m.Run()
}

type UserForm struct {
	Name string `form:"name" binding:"required"`
	Passwd string `form:"passwd" binding:"required"`
}
