package main

func FindMinMaxInSlice (slice []int) (int, int) {
	if len(slice)==0 {
		return 0,0
	}
	
	var max,min int = slice[0],slice[0]

	for i:=1;i<len(slice);i++ {
		if max<slice[i] {
			max = slice[i]
		}
		if min>slice[i] {
			min = slice[i]
		}
	}

	return min,max
}