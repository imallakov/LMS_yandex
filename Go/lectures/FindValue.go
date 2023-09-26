package main

func FindValue(nums []int) int {
	nums = sort(nums)
	if nums[0]!=nums[1] {
		return nums[0]
	}
	if nums[len(nums)-1]!=nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	for i:=1;i<len(nums)-1;i++ {
		if nums[i]!=nums[i-1] && nums[i]!=nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

func sort(nums []int) []int {
	var n int = len(nums)
	for i:=1;i<=n;i++ {
		for j:=1;j<n;j++ {
			if nums[j]<nums[j-1] {
				nums[j],nums[j-1]=SwapNumbers(nums[j],nums[j-1])
			}
		}
	}
	return nums
}

func SwapNumbers (a int ,b int) (int, int) {
	return b,a
}