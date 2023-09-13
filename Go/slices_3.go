package main

func SumOfSlice(slice []int) int {
	var sum int
	for i:=0;i<len(slice);i++ {
		sum+=slice[i]
	}

	return sum
}