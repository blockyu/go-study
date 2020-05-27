package mymap

import "fmt"

func printMap(mapName string, mmap map[string]int) {
	fmt.Println("name : ", mapName)
	for key, value := range mmap {
		//log.Printf("idx :%d value:%s\n",i, value)
		fmt.Printf("idx :%v value:%v\n", key, value)
	}

}

func TestMap() {
	/**
	mymap
	*/

	//mmap := mymap[string]int {
	//	"test": 1,
	//}

	mmap := make(map[string]int)
	mmap["a"] = 1
	mmap["b"] = 2
	mmap["c"] = 3

	printMap("mmap", mmap)

	delete(mmap, "a")
	delete(mmap, "a")
	printMap("mmap deleted a", mmap)

	_, ok := mmap["a"] // map의 원소 존재 여부를 확인하는 방법
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}
}
