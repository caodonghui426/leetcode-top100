package main

import "fmt"

func moveZeroes(nums []int) {
	// 定义两个指针，a，b同时指向0
	a := 0
	b := 0
	// b往前走，如果遇到不是0的，就交换a，b的位置，同时a往前移动一位
	// [***a,00b]，a可以理解为分开0与非0元素的界限
	for ; b < len(nums); b++ {
		if nums[b] != 0 {
			nums[a], nums[b] = nums[b], nums[a]
			a += 1
		}
	}
}

func main() {
	nums := []int{1, 0, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
