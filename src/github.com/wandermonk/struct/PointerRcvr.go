package main

import (
	"fmt"
)

type Student struct {
	name string
	id   string
	Marks
}

type Marks struct {
	english int64
	maths   int64
	science int64
	social  int64
}

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) setId(id string) {
	s.id = id
}

func (s *Student) setMarks(marks Marks) {
	s.Marks = marks
}

func (s *Student) printInfo(){
	fmt.Printf("The name is :: %s\n",s.name)
	fmt.Printf("The id is :: %s\n",s.id)
	fmt.Printf("Marks in english :: %d\n",s.Marks.english)
        fmt.Printf("Marks in maths :: %d\n",s.Marks.maths)
	fmt.Printf("Marks in science :: %d\n",s.Marks.science)
	fmt.Printf("Marks in social :: %d\n",s.Marks.social)
}

func (m *Marks) getAllMarks() {
	fmt.Printf("English :: %d\n", m.english)
	fmt.Printf("Maths :: %d\n", m.maths)
	fmt.Printf("Science :: %d\n", m.science)
	fmt.Printf("Social :: %d\n", m.social)
}

func main() {
	m1 := &Marks{
		english: 30,
		maths:   60,
		science: 70,
		social:  20,
	}

	s1 := &Student{
		name: "phani",
		id:   "#1242",
		Marks: *m1,
	}

	m1.getAllMarks()
	fmt.Println("<-------------------->")
	s1.printInfo()
}
