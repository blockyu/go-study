package main

import (
	"fmt"
	"sort"
)

//bfs
type node struct {
	e []int
	vis bool
}

func dfs(nodes []node, v int) {
	c := &nodes[v]
	sort.Ints(c.e)
	fmt.Printf("%d ",v)
	c.vis = true
	for _, v := range c.e {
		if !nodes[v].vis {
			dfs(nodes, v)
		}
	}
}

func bfs(nodes []node, v int) {
	qu := []int{v}
	for len(qu) > 0 {
		v := qu[0] // front
		qu = qu[1:] // pop
		c := &nodes[v]
		if c.vis {
			continue
		}
		c.vis=true
		sort.Ints(c.e)
		fmt.Printf("%d ",v)
		for _, t := range c.e {
			if !nodes[t].vis {
				qu = append(qu, t)
			}
		}
	}
}
func main() {
	var D, E, S int // vertex, edge, size
	_, err := fmt.Scan(&D, &E, &S)
	nodes := make([]node, D+1)
	if err != nil {
		fmt.Println(err)
	}
	for i:=0; i<E; i++ {
		var s, e int
		_, err := fmt.Scan(&s, &e)
		if err != nil {
			fmt.Println(err)
		}
		nodes[s].e = append(nodes[s].e, e)
		nodes[e].e = append(nodes[e].e, s)
	}
	cpNode := make([]node, len(nodes))
	copy(cpNode, nodes)
	dfs(nodes, S)
	fmt.Println()
	copy(nodes, cpNode)
	bfs(nodes, S)
}
