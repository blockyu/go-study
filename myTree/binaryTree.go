package myTree

import (
	"fmt"
	"math/rand"
	"time"
)

type BTree struct {
	Left *BTree
	Value int
	Right *BTree
}

func traverse(t *BTree) {
	if t==nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *BTree {
	var t *BTree
	rand.Seed(time.Now().Unix())
	for i:=0; i<2*n; i++ {
		tmp := rand.Intn(n*2)
		t = insert(t, tmp)
	}
	return t
}

func insert(t *BTree, v int) *BTree {
	if t==nil {
		return &BTree{nil, v, nil}
	}
	if v==t.Value {
		return t
	}
	if v<t.Value {
		t.Left = insert(t.Left, v)
	}else {
		t.Right = insert(t.Right,v)
	}
	return t
}
func TestTree() {
	tree := create(10)

	fmt.Println("Create Tree :")
	traverse(tree)
	fmt.Println("Complete\n\n")

	insert(tree, 100)
	fmt.Println("Insert 100 in Tree :")
	traverse(tree)
	fmt.Println("Complete\n\n")



}
