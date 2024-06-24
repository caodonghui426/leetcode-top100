package main

import (
	"fmt"
	"sort"
)

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	ans := 1
	temp := 1
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		//fmt.Println(nums[i])
		if nums[i] == nums[i-1]+1 {
			temp += 1
			//fmt.Println(temp)
		} else if nums[i] != nums[i-1] {
			ans = max(ans, temp)
			//fmt.Println(ans)
			temp = 1
		}
	}
	ans = max(ans, temp)
	return ans
}

func main() {
	nums := []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
}
