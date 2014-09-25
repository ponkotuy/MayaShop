package main

import (
	"math"
	"./model"
	"encoding/json"
)

func Hello() string {
	return "Hello world!"
}

func CreateUser(user UserForm) (int, string) {
	err := MayaDB.Col("user").Save(model.NewUser(user.Name))
	return errorToStr(err)
}

func GearAll() (int, string) {
	gears, err := MayaDB.Col("gear").All(0, math.MaxInt32)
	if err != nil { return 500, "DB Error" }
	js, err := json.Marshal(gears.Result)
	if err != nil { return 500, "JSON Error" }
	return 200, string(js)
}

func errorToStr(err error) (int, string) {
	if err == nil {
		return 200, "Success"
	} else {
		return 500, "Failure"
	}
}
