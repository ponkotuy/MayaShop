package main

import (
	model "./model"
)

func Hello() string {
	return "Hello world!"
}

func CreateUser(user UserForm) (int, string) {
	err := MayaDB.Col("user").Save(model.NewUser(user.Name))
	return errorToStr(err)
}

func errorToStr(err error) (int, string) {
	if err == nil {
		return 200, "Success"
	} else {
		return 500, "Failure"
	}
}
