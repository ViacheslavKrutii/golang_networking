package main

type class struct {
	Teachers []teacher
	Students []student
}

var class1 = class{Teachers: []teacher{teachersDB["1"]}, Students: []student{studentsDB["1"], studentsDB["2"]}}
var class2 class
var class3 class
var classes = map[classNum]class{
	1: class1,
	2: class2,
	3: class3,
}

func addTeacherToClass(t teacher) {
	curentClass := classes[t.ClassNum]
	curentClass.Teachers = append(curentClass.Teachers, t)
	classes[t.ClassNum] = curentClass
}

func addStudentToClass(s student) {
	curentClass := classes[s.ClassNum]
	curentClass.Students = append(curentClass.Students, s)
	classes[s.ClassNum] = curentClass
}
