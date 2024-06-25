# 1. 哈希
## Q1. 两数之和

### 题目描述
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：
> 输入：nums = [2,7,11,15], target = 9 
> 
> 输出：[0,1]
> 
> 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
> 输入：nums = [3,2,4], target = 6
>
> 输出：[1,2]
> 
示例 3：
> 输入：nums = [3,3], target = 6
>
> 输出：[0,1]

### 题解：
1. 两层for循环进行遍历，复杂度为 $O(N^2)$
```go
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
            }           
        }   
    }
	return nil
}
```
2. 使用哈希表，寻找是否存在target - x，复杂度为 $O(N)$ 
```go
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target - x]; ok {
			return []int{p, i}
        }
		hashTable[x] = i
    }
	return nil
}
```
## Q49. 字母异位词分组
### 题目描述
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

> 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
> 
> 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

示例 2:
> 输入: strs = [""]
> 
> 输出: [[""]]

示例 3:
> 输入: strs = ["a"]
> 
> 输出: [["a"]]
### 题解
首先对每一个单词内部字母进行排序，然后将相同的分为一组。分组的过程是性能瓶颈，与两数之和一样，如果暴力的话还是两层for循环，然后时间复杂度还是$O(N^2)$ ，所以还是使用哈希表进行存储。
```go
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		// 对单词内部进行ASIIC码排序
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		//fmt.Println(s)
		sortedStr := string(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0)
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
```

## Q128. 最长连续序列
### 题目描述
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

> 输入：nums = [100,4,200,1,3,2]
> 
> 输出：4
> 
> 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
> 输入：nums = [0,3,7,2,5,8,4,6,0,1]
> 
> 输出：9
### 题解
1. 首先判断是 nums 长度否小于等于1，如果小于等于1，直接返回对应长度即可，然后设置起始长度 ans 和临时长度 temp 为1，对 nums 进行生序排序，遍历nums，如果 `nums[i] == nums[i-1]+1` 就让 temp 加1，否则当`nums[i] != nums[i-1]` 求`ans = max(ans, temp)` ，然后重置temp为1，重新开始计数。在for循环外也需要`ans = max(ans, temp)`。寻找的时间复杂度为$O(N)$，但是快排的时间复杂度为$O(NlogN)$。
```go
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
```
2. 首先转换为哈希集合，`set := map[int]bool{}`，然后遍历map集合，判断set[num - 1]是否存在，如果存在就说明不是连续子序列的第一个，跳过。找到连续子序列的第一个，开始向上计算连续子序列的长度。
```go
func longestConsecutive(nums []int) int {
    // 转化成哈希集合，方便快速查找是否存在某个元素
    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    res := 0

    for num := range set {
        if set[num-1] {
            // num 不是连续子序列的第一个，跳过
            continue
        }
        // num 是连续子序列的第一个，开始向上计算连续子序列的长度
        curNum := num
        curLen := 1

        for set[curNum+1] {
            curNum += 1
            curLen += 1
        }
        // 更新最长连续序列的长度
        res = max(res, curLen)
    }

    return res
}
```