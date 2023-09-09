package main

func CalculateSeriesSum(n int) float64 {
	if n==1 {
		return 1.0
	}
	var ans float64
	ans = 1/float64(n)
	return ans+(CalculateSeriesSum(n-1))
}