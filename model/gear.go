package model

import ara "github.com/diegogub/aranGO"

type Gear struct {
	Name string
	Id int
	Elems []GearElem
	ara.Document
}

type GearElem struct {
	Number int					// 1-indexed
	Minutes int
	Explanation string
	ara.Document
}
