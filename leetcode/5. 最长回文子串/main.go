package main

import (
	"fmt"
)

func main() {
	// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

	// 示例 1：

	// 输入: "babad"
	// 输出: "bab"
	// 注意: "aba" 也是一个有效答案。
	// 示例 2：

	// 输入: "cbbd"
	// 输出: "bb"

	// 来源：力扣（LeetCode）
	// 链接：https://leetcode-cn.com/problems/longest-palindromic-substring
	// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

	// table := make(map[string]string)
	// table["babadb"] = "badb"
	// table["cbbd"] = "bb"
	// table[""] = ""
	// table["a"] = "a"
	// table["ccc"] = "ccc"
	// table["aacabdkacaa"] = "aca"
	// for k, v := range table {
	// 	excepted := v
	// 	got := longestPalindrome(k)
	// 	if got != excepted {
	// 		log.Fatal(fmt.Sprintf("%s excepted: %v, got: %v", k, excepted, got))
	// 	}
	// }

	fmt.Println(longestPalindrome("aacabdkacaa"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	table := make(map[uint8]int)
	var (
		startIndex int
		endIndex   int
	)
	for i, c := range s {
		index, ok := table[uint8(c)]
		matched := false
		if ok && i-index > (endIndex-startIndex) {
			matched = true
			startIndex = index
			endIndex = i
		}

		if !matched || i == len(s)-1 || s[i] != s[i+1] {
			table[uint8(c)] = i
		}
	}
	return s[startIndex : endIndex+1]
}

func longestPalindrome2(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	table := make(map[uint8]int)
	var (
		startIndex int
		endIndex   int
	)
	for i, c := range s {
		index, ok := table[uint8(c)]
		if ok && i-index > (endIndex-startIndex) {
			startIndex = index
			endIndex = i
		}
		table[uint8(c)] = i
	}
	return s[startIndex : endIndex+1]
}
