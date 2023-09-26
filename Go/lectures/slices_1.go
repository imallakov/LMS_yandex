package main

func ReverseSlice(slice []int) []int {
	var new_slice []int = nil
	for i:=len(slice)-1;i>=0;i-- {
		new_slice=append(new_slice, slice[i])
	}

	return new_slice
}