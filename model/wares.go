package model

import ara "github.com/diegogub/aranGO"

type Wares struct {
	Name string
	Coin int64
	ara.Document
}

func NewWares(name string, coin int64) Wares {
	var w Wares
	w.Name = name
	w.Coin = coin
	return w
}
