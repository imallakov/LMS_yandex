package main

func CountSubslices(slice []int) int {
	var count, temp int
	for i:=0;i<len(slice);i++{
		temp+=slice[i]
	}
	temp/=len(slice)

	for i:=0;i<len(slice);i++ {
		var sum int
		for j:=i;j<len(slice);j++ {
			sum+=slice[j]
			if sum>temp {
				count++
			}
		}
	}

	return count
}