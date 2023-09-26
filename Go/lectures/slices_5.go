package main

func SortSlice(slice []int) []int {
	for i:=0;i<len(slice);i++ {
		for j:=1;j<len(slice);j++ {
			if slice[j]<slice[j-1] {
				temp := slice[j-1]
				slice[j-1]=slice[j]
				slice[j]=temp
			}
		}
	}

	return slice
}