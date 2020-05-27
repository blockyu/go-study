package mystruct

import "fmt"

type myStruct struct {
	name string
	age  int
}

func TestStruct() {
	//aStruct := make(myStruct, 0)
	aStruct := new(myStruct) // return memory
	aStruct.age = 1
	aStruct.name = "hello"
	fmt.Println(*aStruct)

}
