package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(s), cap(s), s, s == nil)
	var t = []int{}
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(t), cap(t), t, t == nil)
	var u = make([]int, 5)
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(u), cap(u), u, u == nil)
	var v = make([]int, 0, 5)
	fmt.Printf("%d %d %T %5t %#[3]v\n", len(v), cap(v), v, v == nil)
}
