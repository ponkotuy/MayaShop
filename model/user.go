package model

import ara "github.com/diegogub/aranGO"

type User struct {
	Name string
	Coin int64
	Wares []WaresRecord
	Workers []Worker
	ara.Document
}

func NewUser(name string) User {
	var u User
	u.Name = name
	u.Wares = make([]WaresRecord, 0)
	u.Workers = make([]Worker, 0)
	return u
}

type WaresRecord struct {
	WaresId int64
	SellCount int
	HaveCount int
	ara.Document
}

func NewWaresRecord(waresId int64, have int) WaresRecord {
	var w WaresRecord
	w.WaresId = waresId
	w.HaveCount = have
	return w
}

type Worker struct {
	GearId int					// Gearに置いてないのは-1
	SetDate int64
	ara.Document
}

func NewWorker() Worker {
	var w Worker
	w.GearId = -1
	return w
}
