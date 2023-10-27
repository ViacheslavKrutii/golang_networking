package main

type teacher struct {
	ID       string
	Name     string
	ClassNum classNum
}

var teachersDB = map[string]teacher{
	"1": {ID: "1", Name: "Franc", ClassNum: 5},
}
