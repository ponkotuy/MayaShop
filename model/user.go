package model

import ara "github.com/diegogub/aranGO"

type User struct {
	Name string
	Coin int64
	Wares []WaresRecord
	ara.Document
}

func NewUser(name string) User {
	var u User
	u.Name = name
	u.Wares = make([]WaresRecord, 0)
	return u
}

type WaresRecord struct {
	WaresId int64
	SellCount int
	HaveCount int
	ara.Document
}
