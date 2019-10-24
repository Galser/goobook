// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import "fmt"

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root, 0)
	PrintTree(root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree, r int64) []int {
	r++
	if t != nil {
		//fmt.Println("Entering recusion level : ", r)
		//fmt.Println("enter ", values, " tree : ", t)
		values = appendValues(values, t.left, r)
		//fmt.Println("1-st append  ", values, " tree : ", t)
		values = append(values, t.value)
		//fmt.Println("slice append  ", values, " tree : ", t)
		values = appendValues(values, t.right, r)
		//fmt.Println("last tree append  ", values, " tree : ", t)

	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func preOrder(t *tree, s []string, d int, dir int) []string {
	//fmt.Println("preOrder", s, d, dir)
	if len(s) <= d {
		s = append(s, " ")
	}
	if dir < 0 {
		// adding left
		s[d] = s[d] + "  " + fmt.Sprint(t.value) + "  "
	} else {
		// adding right
		s[d] = "  " + fmt.Sprint(t.value) + "  " + s[d]
	}
	if t.left != nil {
		s = preOrder(t.left, s, d+1, -1)
	}
	if t.right != nil {
		s = preOrder(t.right, s, d+1, +1)
	}
	return s
}

// PrintTree - Prints tree
func PrintTree(t *tree) {
	var s []string
	d := 0
	s = preOrder(t.left, s, d, -1)
	s = preOrder(t.right, s, d, +1)
	//fmt.Println(s)
	for _, level := range s {
		fmt.Println(level)
	}
}

//!-
