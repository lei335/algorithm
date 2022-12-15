package search

// 给定一个  haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回  -1。
// 思路：核心点遍历给定字符串字符，判断以当前字符开头字符串是否等于目标字符串
func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
// 需要注意点
// 循环时，i 不需要到 len-1
// 如果找到目标字符串，len(needle)==j


// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 思路：这是一个典型的应用回溯法的题目，简单来说就是穷尽所有可能性，算法模板如下