package model

import ara "github.com/diegogub/aranGO"

type Wares struct {
	Id int64
	Name string
	Coin int64
	ara.Document
}

func NewWares(id int64, name string, coin int64) Wares {
	var w Wares
	w.Id = id
	w.Name = name
	w.Coin = coin
	return w
}
