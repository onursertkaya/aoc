package main

import (
	"fmt"
)

func arrays() {
	// array:
	// a contagious sequence of typed elements.
	// size is part of the type, i.e. type([4]int) != type([6]int)
	int_array := [4]int

	// bytes: an array of 8-bit `bytes`.
	byte_array := [8]byte

	type [32]byte ThirtyTwoBytes
	large_byte_array := ThirtyTwoBytes{}

	// len
	fmt.Println(len(large_byte_array) == 32)

	// indexing
	large_byte_array[3] = 14
	large_byte_array[7] = 20

	// iteration
	for idx, b := range byte_array {
		fmt.Println(idx)
	}
}

func slices()
{
	some_array := [10]byte

	// slice: a view to an array object.
	// - its type is `[]type`` (no size, because it does not manage memory)
	// - only holds a pointer to `type`, a length variable and a capacity varible.

	// idiomatic definition for a full slice of an array
	slice := some_array[:]

	// verbose definition for a subregion of an array
	var another_slice []byte = some_array[3:8]

	// a slice can be sliced, it will still point to original array.
	slice_of_slice := slice[0:2]

	// a slice can also be defined to point at an anonymous array.
	slice_to_anon_array := []int{1,2,3,4,5}

	// capacity retains the actual total size of the underlying array. trying to grow the slice
	// beyond the capacity will crash the program, as it literally means out-of-bounds access.
	length := len(another_slice)    // 5 -> 3:8
	capacity := cap(another_slice)  // 7 -> len(some_array) - 3

	// As it stores a pointer, making a copy (return from/pass-by-value to functions) still ties to
	// the underlying array.
	// > sliceCopyingFunction(slice)

	// however, if modifying the *slice itself* is required, it must be passed by pointer, or the
	// returned value from function must be assigned to the slice.
	// > sliceModifyingFunction(&slice)
	// > slice = sliceCopyingFunction(slice)
}

func makeFunction()
{
	// when dynamic memory allocation is needed, make() comes into play. it is a shorthand funtion
	// for allocating a new anonymous array and tying a slice to it.
	// type, initial length, total capacity
	made_slice_with_room := make(int[], 10, 15)
	made_slice_to_larger_array := make(int[], 10, 2*(cap(made_slice)))

	// capacity argument is optional, so following is equivalent to cap=size
	default_capacity_slice := make(int[], 10)
}

func copyFunction()
{
	// to copy underlying arrays use copy through slices. it uses the min(len(source), len(target))
	// to identify the range
	source := make(int[], 10)
	target := make(int[], 20)
	num_copied_elems := copy(target, source)
}

func insertAnElement()
{
	// as the arguments to copy are slices, it might be that they are pointing to the same
	// underlying array with different offsets.
	index_to_insert_element := 4
	element_to_insert := 10
	insertion_slice := make([]int, 10)

	temp := insertion_slice       // copy the slice
    temp = temp[0 : len(temp)+1]  // increase slice size
    copy(temp[index_to_insert_element+1:], temp[index_to_insert_element:])  // make a hole
    temp[index_to_insert_element] = element_to_insert  // put the value in the hole

	insertion_slice = temp  // update the slice
}

func main() {





	// string: a read-only slice of bytes. does not have to be valid characters
}
