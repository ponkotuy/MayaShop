package main

import (
	"math"
	"log"
	"os"
	"bytes"
	"encoding/json"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"./model"
	"./lib"
)

func MayaAuth(uName, passwd string) bool {
	var user model.User
	err := MayaDB.Col("user").First(map[string]string{"Name": uName}, &user)
	if err != nil { return false }
	hashed := lib.CreatePasswordHash(passwd, user.Salt)
	return bytes.Equal(hashed, user.Passwd)
}

func UserInfo(authUser auth.User) (int, string) {
	var user model.User
	err := MayaDB.Col("user").First(map[string]string{"Name": string(authUser)}, &user)
	if err != nil { return 500, "DB Error" }
	js, err := json.Marshal(user)
	if err != nil { return 500, "JSON Error" }
	return 200, string(js)
}

func CreateUser(form UserForm) (int, string) {
	user := model.NewUser(form.Name, form.Passwd)
	user.Wares = append(user.Wares,
		model.NewWaresRecord(1, 20),
		model.NewWaresRecord(2, 15),
		model.NewWaresRecord(3, 5))
	for i := 0; i < 3; i++ { user.Workers = append(user.Workers, model.NewWorker()) }
	err := MayaDB.Col("user").Save(user)
	return errorToStr(err)
}

func GearAll() (int, string) {
	gears, err := MayaDB.Col("gear").All(0, math.MaxInt32)
	if err != nil { return 500, "DB Error" }
	js, err := json.Marshal(gears.Result)
	if err != nil { return 500, "JSON Error" }
	return 200, string(js)
}

func GetPublic(params martini.Params) (int, string) {
	base := params["_1"]
	if base[len(base)-1] == '/' { base = base + "index.html" }
	file, errOpen := os.Open("public" + base)
	if errOpen != nil { return 404, "Not Found" }
	bytes := make([]byte, 1000000) // Max File Size is 1MB
	_, errRead := file.Read(bytes)
	if errRead != nil { return 500, "File Read Error" }
	return 200, string(bytes)
}

func errorToStr(err error) (int, string) {
	if err == nil {
		return 200, "Success"
	} else {
		log.Println(err)
		return 500, "Failure"
	}
}
