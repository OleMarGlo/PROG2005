package main

import (
	exampleLib "example/testLib"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello World")
	s := exampleLib.Student{}
	s.SetName("Kebab")
	fmt.Println(s.String())

	s1 := exampleLib.Student{Age: 10}

	sArr := make([]exampleLib.Student, 2)

	sArr = append(sArr, s1, s1, s1)
	fmt.Println(sArr)
	fmt.Println(s1)

	fmt.Println("For loop")

	for i := 0; i < len(sArr); i++ {
		fmt.Println(sArr[i])
	}

	fmt.Println("While loop")
	ct := 0
	for ct < len(sArr) {
		fmt.Println(sArr[ct])
		ct++
	}
	fmt.Println("Range")
	//k is the range av v is the value
	for k, v := range sArr {
		fmt.Println(k, v)
	}

	//maps
	sMap := make(map[int]exampleLib.Student, 5)
	sMap[0] = s1

	//you can iterate over it the exactly same
	for k, v := range sMap {
		fmt.Println("Key " + strconv.Itoa(k) + ", Value: " + v.String())
	}

}
