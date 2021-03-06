// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 99.

// Graph shows how to use a map of maps to represent a directed graph.
package main

import "fmt"

//!+
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	fmt.Print(from, to, ":")
	return graph[from][to]
}

//!-

func main() {
	addEdge("a", "b")
	addEdge("c", "d")
	addEdge("a", "d")
	addEdge("d", "a")
	fmt.Println(graph)
	fmt.Println(hasEdge("a", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("a", "d"))
	fmt.Println(hasEdge("d", "a"))
	fmt.Println(hasEdge("x", "b"))
	fmt.Println(hasEdge("c", "d"))
	fmt.Println(hasEdge("x", "d"))
	fmt.Println(hasEdge("d", "x"))
	fmt.Println(hasEdge("c", "b"))
/*
map[a:map[b:true d:true] c:map[d:true] d:map[a:true]]
ab:true
cd:true
ad:true
da:true
xb:false
cd:true
xd:false
dx:false
cb:false
*/
}
