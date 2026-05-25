package greedy

import (
	"sort"
	"strconv"
)

// 给定一个数n，比如23121。再给定一个数组{2，4，9}，求由数组中元素组成的小于n的最大数，如小于23121的最大数为22999
func maxLessThanN(n int, nums []int) int {
	// 先给元素排序
	sort.Ints(nums)

	// 先把n转换为字符串，方便从高位到低位遍历
	s := strconv.Itoa(n)
	ans := []int{}

	// 从高位到低位遍历数字n，查找nums中等于或小于数字n该位值的索引
	for i := 0; i < len(s); i++ {
		d := int(s[i] - '0')

		// 找 >=d 的位置
		idx := sort.SearchInts(nums, d)

		// 当前位有相等数字
		if idx < len(nums) && nums[idx] == d {
			ans = append(ans, d)
			continue
		}

		// 找小于 d 的最大数字
		smaller := idx - 1

		if smaller >= 0 {
			ans = append(ans, nums[smaller])

			// 后续填最大值
			for j := i + 1; j < len(s); j++ {
				ans = append(ans, nums[len(nums)-1])
			}

			return digitsToInt(ans)
		}

		// 没有找到等于或者小于d的数字，当前位就不能填，需要向前回退
		for len(ans) > 0 {
			prev := ans[len(ans)-1]
			ans = ans[:len(ans)-1]

			pos := sort.SearchInts(nums, prev)

			if pos > 0 {
				ans = append(ans, nums[pos-1])

				// 后续补最大值
				for len(ans) < len(s) {
					ans = append(ans, nums[len(nums)-1])
				}

				return digitsToInt(ans)
			}
		}

		// 回退到最高位也没有找到较小的元素，那么就不能构造同等长度的数字
		res := 0
		for i := 0; i < len(s)-1; i++ {
			res = res*10 + nums[len(nums)-1]
		}
		return res
	}

	// 如果完全匹配到n，那么则需要向前回退
	for len(ans) > 0 {
		prev := ans[len(ans)-1]
		ans = ans[:len(ans)-1]

		pos := sort.SearchInts(nums, prev)

		if pos > 0 {
			ans = append(ans, nums[pos-1])

			// 后续补最大值
			for len(ans) < len(s) {
				ans = append(ans, nums[len(nums)-1])
			}

			return digitsToInt(ans)
		}
	}
	// 回退到最高位也没有找到较小的元素，那么就不能构造同等长度的数字
	res := 0
	for i := 0; i < len(s)-1; i++ {
		res = res*10 + nums[len(nums)-1]
	}
	return res
}

func digitsToInt(arr []int) int {
	res := 0
	for _, v := range arr {
		res = res*10 + v
	}
	return res
}
