package main

func FindMaxKey(m map[int]int) int {
	var max, max_key int
	for key,value := range m {
		if value>max {
			max=value
			max_key=key
		}
	}
	return max_key
}