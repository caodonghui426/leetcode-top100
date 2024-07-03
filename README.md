# 🚲 LeetCode-Hot100

## 1. 哈希

### Q1. 两数之和

#### 题目描述

给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：

> 输入：nums = \[2,7,11,15], target = 9
>
> 输出：\[0,1]
>
> 解释：因为 nums\[0] + nums\[1] == 9 ，返回 \[0, 1] 。

示例 2：

> 输入：nums = \[3,2,4], target = 6
>
> 输出：\[1,2]

示例 3：

> 输入：nums = \[3,3], target = 6
>
> 输出：\[0,1]

#### 题解：

1. 两层for循环进行遍历，复杂度为 $$O(N^2)$$

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

2. 使用哈希表，寻找是否存在target - x，复杂度为 $$O(N)$$

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

### Q49. 字母异位词分组

#### 题目描述

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

> 输入: strs = \["eat", "tea", "tan", "ate", "nat", "bat"]
>
> 输出: \[\["bat"],\["nat","tan"],\["ate","eat","tea"]]

示例 2:

> 输入: strs = \[""]
>
> 输出: \[\[""]]

示例 3:

> 输入: strs = \["a"]
>
> 输出: \[\["a"]]

#### 题解

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

### Q128. 最长连续序列

#### 题目描述

给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：

> 输入：nums = \[100,4,200,1,3,2]
>
> 输出：4
>
> 解释：最长数字连续序列是 \[1, 2, 3, 4]。它的长度为 4。

示例 2：

> 输入：nums = \[0,3,7,2,5,8,4,6,0,1]
>
> 输出：9

#### 题解

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

2. 首先转换为哈希集合，`set := map[int]bool{}`，然后遍历map集合，判断set\[num - 1]是否存在，如果存在就说明不是连续子序列的第一个，跳过。找到连续子序列的第一个，开始向上计算连续子序列的长度。

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

## 2. 双指针

### Q283. 移动零

#### 题目描述

给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

**请注意** ，必须在不复制数组的情况下原地对数组进行操作。&#x20;

**示例 1:**

<pre><code><strong>输入: nums = [0,1,0,3,12]
</strong><strong>输出: [1,3,12,0,0]
</strong></code></pre>

**示例 2:**

<pre><code><strong>输入: nums = [0]
</strong><strong>输出: [0]
</strong></code></pre>

**提示**:

* `1 <= nums.length <= 104`
* `-231 <= nums[i] <= 231 - 1`

进阶：你能尽量减少完成的操作次数吗？

#### 题解

使用双指针，指针a和指针b都位于起点0，然后b往前走，如果遇到不是0的，就交换a，b的位置，同时a往前移动一位。\[\*\*\*a,00b]，a可以理解为分开0与非0元素的界限。

```go
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
```

### Q11. 盛最多水的容器

#### 题目描述

给定一个长度为 `n` 的整数数组 `height` 。有 `n` 条垂线，第 `i` 条线的两个端点是 `(i, 0)` 和 `(i, height[i])` 。

找出其中的两条线，使得它们与 `x` 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

**说明：**你不能倾斜容器。

**示例 1：**

![](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question\_11.jpg)

<pre><code><strong>输入：[1,8,6,2,5,4,8,3,7]
</strong><strong>输出：49 
</strong><strong>解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
</strong></code></pre>

**示例 2：**

<pre><code><strong>输入：height = [1,1]
</strong><strong>输出：1
</strong></code></pre>

**提示：**

* `n == height.length`
* `2 <= n <= 105`
* `0 <= height[i] <= 104`

#### 题解

要想盛最多的水，那么就要尽可能的保证height足够高，宽足够大。所以可以先从两边开始，使用双指针指向两端，然后计算当前面积，然后如果左边height比较小，就移动左边的，如果右边比较小就移动右边的，直至相遇。

这里面的逻辑就是移动最小的，保留更大的有可能在后面会增大面积，这是一种贪心思想。

```go
func maxArea(height []int) int {
    left, right := 0, len(height)-1
    maxArea := 0
    for left < right {
       currentArea := min(height[left], height[right]) * (right - left)
       maxArea = max(maxArea, currentArea)
       if height[left] < height[right] {
          left++
       } else {
          right--
       }
    }
    return maxArea
}

func max(a, b int) int {
    if a > b {
       return a
    } else {
       return b
    }
}

func min(a, b int) int {
    if a < b {
       return a
    } else {
       return b
    }
}
```

### Q15. 三叔之和

#### 题目描述：

给你一个整数数组 `nums` ，判断是否存在三元组 `[nums[i], nums[j], nums[k]]` 满足 `i != j`、`i != k` 且 `j != k` ，同时还满足 `nums[i] + nums[j] + nums[k] == 0` 。请

你返回所有和为 `0` 且不重复的三元组。

**注意：**答案中不可以包含重复的三元组。

**示例 1：**

<pre><code><strong>输入：nums = [-1,0,1,2,-1,-4]
</strong><strong>输出：[[-1,-1,2],[-1,0,1]]
</strong><strong>解释：
</strong>nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
</code></pre>

**示例 2：**

<pre><code><strong>输入：nums = [0,1,1]
</strong><strong>输出：[]
</strong><strong>解释：唯一可能的三元组和不为 0 。
</strong></code></pre>

**示例 3：**

<pre><code><strong>输入：nums = [0,0,0]
</strong><strong>输出：[[0,0,0]]
</strong><strong>解释：唯一可能的三元组和为 0 。
</strong></code></pre>

#### 题解：

1. 暴力：

这个思路没问题，但是会超时。思想就是先进行排序，时间复杂度是 $O(nlogn)$，然后进行全部的遍历，时间复杂度是$O(N^3)$。

这里需要注意的是排除可能重复的组合，即如果当前一个元素和上一个元素值一样，跳过。

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}
```

2. 双指针

在最外层循环中与第一种解法保持一致，然后对于后两个数的选择，使用左右指针指向两边，不断向中间靠拢。

如果找到了，就缩小范围，left++，right--，同时减少可能存在的重复组合。

如果sum < 0， left++

如果sum > 0，right--

所以整体时间复杂度为$O(N^2)$&#x20;

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
```
