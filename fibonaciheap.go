package main

import (
	"fmt"
)

type FibHeap struct {
	roots       *Fibtree
	nodes       int
	degree      int
	numoftress  int
	minheap     *Fibtree
	numofmarked int
}

type Fibtree struct {
	Parent, Child, next, prev *Fibtree
	Key                       int
	Degree                    int
	Mark                      bool
	key                       int
	value                     string
}

func CreateNewHeap() *FibHeap {
	heap := new(FibHeap)
	heap.roots = nil
	heap.nodes = 0
	heap.degree = 0
	heap.minheap = nil
	heap.numoftress = 0
	heap.numofmarked = 0
	return heap
}

func CreateSingeltonTree(ket int, val string) *Fibtree {
	tree1 := new(Fibtree)
	tree1.Degree = 0
	tree1.Parent = nil
	tree1.Child = nil
	tree1.next = tree1
	tree1.prev = tree1
	tree1.Mark = false
	tree1.key = ket
	tree1.value = val
	return tree1
}
func (fb *FibHeap) Insert(ket int, val string) {
	tree := CreateSingeltonTree(ket, val)
	fb.minheap = mergeintolist(fb.minheap, tree)
	fb.nodes += 1
}

func mergeintolist(one, two *Fibtree) *Fibtree {
	if one == nil && two == nil {
		return nil
	} else if one == nil && two != nil {
		return two
	} else if one != nil && two == nil {
		return one
	} else {
		y := one.prev
		one.prev = two
		two.next = one
		two.prev = y
		y.next = two
		if two.Key < one.Key {
			return two
		} else {
			return one
		}
	}
}
func printSpaces(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Print(" ")
	}
}
func PrintChild(val int) {
	printSpaces(val)
}
func (fb *FibHeap) PrintRoot() {
	k := fb.minheap
	y := k.next
	fmt.Println("root is ", k)
	for y != k {
		fmt.Println("root is ", y)
		y = y.next
	}
}

func (fb *FibHeap) DeleteMin() {
	k := fb.minheap
	//heap is empty
	if k != nil {
		addchildToHeap(fb, k.Child)
		k.Parent = nil
		if k.next == k {
			fb.minheap = nil
		} else {
			fb.minheap = k.next
			Consolidate(fb)
		}
	}
	fb.nodes -= 1
	return k
}

func Consolidate(fb *FibHeap) {

}
func addchildToHeap(fb *FibHeap, chil *Fibtree) {
	if chil == nil {
		return
	} else {
		track := fb.minheap.prev
		track.next = chil
		chil.prev = track
	}
}
func main() {
	heap1 := CreateNewHeap()
	fmt.Println("heap is ", heap1)
	heap1.Insert(3, "priya")
	heap1.Insert(7, "jdfnc")
	heap1.PrintRoot()
}
