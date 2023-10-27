package main

type class struct {
	Teachers []teacher
	Students []student
}

var class5 = class{Teachers: []teacher{teachersDB["1"]}, Students: []student{studentsDB["1"], studentsDB["2"]}}
var clsses = map[classNum]class{5: class5}
