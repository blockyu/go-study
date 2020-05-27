package myslice

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func printSlice(sliceName string, slice []aStructure) {
	fmt.Println("name : ", sliceName)
	for i, value := range slice {
		//log.Printf("idx :%d value:%s\n",i, value)
		fmt.Printf("idx :%d value:%v\n", i, value)
	}
}

func TestSlice() {

	/**
	myslice
	*/

	//mySlice := []aStructure{{"init",0,0}}
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"hello2", 2, 2})
	mySlice = append(mySlice, aStructure{"hello1", 1, 1})
	mySlice = append(mySlice, aStructure{"hello3", 3, 3})
	mySlice = append(mySlice, aStructure{"hello5", 5, 5})
	mySlice = append(mySlice, aStructure{"hello4", 4, 4})

	printSlice("orgin mySlice", mySlice)

	cpSlice := make([]aStructure, len(mySlice))
	copy(cpSlice[0:], mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})

	printSlice("sorted mySlice", mySlice)
	printSlice("origin cpSlice", cpSlice[0:])

}
