package main

func MoveZeroes(nums []int) []int {
	var ans []int

	for i:=0;i<len(nums);i++ {
		if nums[i]!=0 {
			ans = append(ans, nums[i])
		}
	}

	for i:=0;i<len(nums);i++ {
		if nums[i]==0 {
			ans = append(ans, nums[i])
		}
	}

	return ans
}