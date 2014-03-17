package List

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
