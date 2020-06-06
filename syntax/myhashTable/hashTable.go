package myhashTable

import "fmt"

const SIZE=15

type Node struct {
	Value int
	Next *Node
}

type HashTable struct {
	Table map[int]*Node
	Size int
}

func hashFunction(i, size int) int {
	return i%size
}

func insert(hash *HashTable, value int) int {
	index := hashFunction(value, hash.Size)
	element := Node{Value:value, Next:hash.Table[index]}
	hash.Table[index] = &element
	return index
}

func traverse(hash *HashTable) {
	for k := range hash.Table {
		fmt.Println(k)
		if hash.Table[k] != nil {
			t := hash.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}
func TestHashTable() {
	table := make(map[int]*Node, SIZE)
	hash := &HashTable{table, SIZE}
	fmt.Println("Number of spaces: ", hash.Size)
	for i:=0; i<30; i++ {
		insert(hash, i)
	}
	traverse(hash)
}