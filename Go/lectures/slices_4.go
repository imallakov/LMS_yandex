package main

func IntersectionOfSlices(slice1, slice2 []int) []int {
	if len(slice1)==0 || len(slice2)==0 {
		return nil
	}
	
	var new_slice []int = nil

	for i:=0;i<len(slice1);i++ {
		for j:=0;j<len(slice2);j++ {
			if slice1[i]==slice2[j] {
				new_slice=append(new_slice, slice1[i])
			}
		}
	}

	return new_slice

}