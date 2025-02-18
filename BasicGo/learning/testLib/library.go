package exampleLib

type Student struct {
	name string
	Age  int
}

func (stu *Student) String() string {
	return stu.name
}

func (stu *Student) SetName(name string) {
	stu.name = name
}
