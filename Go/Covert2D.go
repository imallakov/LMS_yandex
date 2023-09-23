package main

func Convert2D(nums []int, m, n int) [][]int {
	var ans [][]int
	
	for i:=0;i<m;i++ {
		var temp []int
		for j:=0;j<n;j++{
			temp = append(temp, nums[i*n+j])
		}
		ans = append(ans, temp)
	}

	return ans
}