package main

import (
	"encoding/json"
	"fmt"
)

type TestJsonType struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func main() {
	//data := map[string]string{["name":"test"]}
	data := TestJsonType{"yujin", 10}
	mData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("1. %v\n", data)
	fmt.Printf("2. %v\n", mData)
	fmt.Printf("3. %v\n", string(mData))

	var umData TestJsonType
	err = json.Unmarshal(mData, &umData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("4. %v\n", umData)
}
