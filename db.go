package main

import ara "github.com/diegogub/aranGO"

var dbUrl = "http://localhost:8529"
var dbName = "maya"
var MayaDB = NewDB(dbUrl, dbName)

type DB struct {
	Url string
	Name string
	session *ara.Session
}

func NewDB(url string, name string) DB {
	var db DB
	db.Url = url
	db.Name = name
	var err error
	db.session, err = ara.Connect(url, "", "", false)
	if err != nil { panic(err) }

	db.session.CreateDB(name, nil)
	return db
}

func (db *DB) DB() *ara.Database {
	return db.session.DB(db.Name)
}

func (db *DB) Col(colName string) *ara.Collection {
	createCol(db.DB(), colName)
	return db.DB().Col(colName)
}

func createCol(db *ara.Database, name string) {
	if !db.ColExist(name) {
		docs1 := ara.NewCollectionOptions(name, true)
		db.CreateCollection(docs1)
	}
}
