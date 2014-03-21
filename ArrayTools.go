package List

import (
	"math/rand"
)

// And array of int that auto increase its capacity when needed
func AppendInt(slice []int, elements ...int) []int {
	ori_size := len(slice)
	final_size := len(slice) + len(elements)
	total := cap(slice)
	// We double the size as long as the size is not big enough
	for total < final_size {
		total = 2*total + 1
	}
	if cap(slice) < final_size {
		newSlice := make([]int, len(slice)+len(elements), total)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:final_size]
	copy(slice[ori_size:], elements)
	return slice
}

// And array of int that auto increase its capacity when needed
func AppendString(slice []string, elements ...string) []string {
	ori_size := len(slice)
	final_size := len(slice) + len(elements)
	total := cap(slice)
	// We double the size as long as the size is not big enough
	for total < final_size {
		total = 2*total + 1
	}
	if cap(slice) < final_size {
		newSlice := make([]string, len(slice)+len(elements), total)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:final_size]
	copy(slice[ori_size:], elements)
	return slice
}

// Shuffle an array of int. Same values move at random position
func ShuffleInt(slice []int) {
	N := len(slice)
	for i := 0; i < len(slice); i++ {
		rand := i + rand.Intn(N-i)
		slice[i], slice[rand] = slice[rand], slice[i]
	}
}

// Shuffle an array of string. Same values move at random position
func ShuffleString(slice []string) {
	N := len(slice)
	for i := 0; i < len(slice); i++ {
		rand := i + rand.Intn(N-i)
		slice[i], slice[rand] = slice[rand], slice[i]
	}
}
