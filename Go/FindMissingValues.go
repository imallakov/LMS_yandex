package main

func FindMissingValues(nums []int) []int {
	var n int = len(nums)
	var ans []int
	for i:=1;i<=n;i++ {
		var found bool = false
		for j:=0;j<n && !found;j++ {
			if nums[j]==i {
				found=true
			}
		}
		if !found {
			ans = append(ans, i)
		}
	}
	return ans
}