package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 13
	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if j, ok := hashTable[target-x]; ok {
			return []int{i, j}
		}
		hashTable[x] = i // 增加这一行，将当前元素及其索引存入哈希表
	}
	return nil
}
