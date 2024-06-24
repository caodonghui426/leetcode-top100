package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	output := groupAnagrams(strs)
	fmt.Println(output)
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		//fmt.Println(s)
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
