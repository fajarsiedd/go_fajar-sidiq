package main

import (
	"fmt"
)

// IntSet mempresentasikan kumpulan bilangan bulat.
type IntSet struct {
	Data map[int]struct{} // Menggunakan struct{} untuk menghemat memori karena yang diguanakan hanyalah keys
}

// NewIntSet untuk membuat dan mengembalikan instance baru dari IntSet
func NewIntSet() *IntSet {
	return &IntSet{
		Data: make(map[int]struct{}),
	}
}

// Menggunakan pointer agar data bisa langsung diubah
func (set *IntSet) Add(value int) {
	set.Data[value] = struct{}{}
}

func (set *IntSet) Get() []int {
	result := []int{}
	for key := range set.Data {
		result = append(result, key)
	}
	return result
}

func (set *IntSet) Remove(key int) {
	delete(set.Data, key)
}

func main() {
	set := NewIntSet()

	set.Add(1)
	set.Add(2)
	set.Add(1)
	set.Add(3)
	set.Add(4)
	set.Add(5)

	set.Remove(100)

	fmt.Println(set.Get())
}
