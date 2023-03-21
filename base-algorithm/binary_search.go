package base

// 二分搜索模板
// 给一个有序数组和目标值，找第一次/最后一次/任何一次出现的索引，如果没有出现就返回-1
// 模板四点要素：
// 1. 初始化：start=0, end=len-1
// 2. 循环条件：start+1<end
// 3. 比较中点和目标值：A[mid]==、<、> traget
// 4. 判断最后两个元素是否符合：A[start]、A[end] ==? target
// 时间复杂度O(logn)，使用场景一般是有序数组的查找

// 典型示例
// 给定一个n个元素有序的（升序）整型数组nums和一个目标值target，写一个函数搜索nums中的target，如果目标值存在返回下标，否则返回-1

// 二分搜索最常用模板
func search(nums []int, target int) int {
	// 1. 初始化start、end
	start := 0
	end := len(nums) - 1
	// 2. 处理for循环
	for start+1 < end {
		mid := start + (end-start)/2
		// 3. 比较a[mid]和target值
		if nums[mid] == target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	// 4. 最后剩下两个元素，手动判断
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

// 大部分的二分查找都可以用这个模板，然后做一些特殊逻辑即可

// 另外二分查找还有其他一些模板，比如下面这个模板：
// 模板1
// left := 0; right := length-1;
// while(left<=right){
// 	mid=left+(right-left)/2;
// 	if nums[mid]==target {
// 		return mid;
// 	}else if nums[mid]<target{
// 		left = mid+1;
// 	}else {
// 		right = mid-1;
// 	}
// }

// 模板2很少用到

// 模板3
// left := 0; right := length - 1;
// while (left + 1 < right){
// 	mid = left + (right - left)/2;
// 	if(num[mid] < target){
// 		left = mid;
// 	}else {
// 		right = mid;
// 	}
// }

// 以下是常见题目

// 给定一个包含n个整数的排序数组，找出给定目标值target的起始和结束位置，如果如果目标值不在数组中，则返回[-1,-1]
// 思路：核心要点就是找第一个target的索引，和最后一个target的索引，所以用两次二分搜索分别查找第一个和最后一个的位置
func searchRange(A []int, target int) []int {
	if len(A) == 0 {
		return []int{-1, -1}
	}
	res := make([]int, 2)
	start := 0
	end := len(A) - 1
	// 先找第一个位置
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] < target {
			start = mid
		} else if A[mid] > target {
			end = mid
		} else { // 往左边找，找第一个位置
			end = mid
		}
	}
	if A[start] == target { //先找左边的索引
		res[0] = start
	} else if A[end] == target {
		res[0] = end
	} else {
		res[0] = -1
		res[1] = -1
		return res
	}
	// 再找最后一个位置
	start = 0
	end = len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] < target {
			start = mid
		} else if A[mid] > target {
			end = mid
		} else { // 往右边找，找最后一个位置
			start = mid
		}
	}
	if A[end] == target { // 先找右边的索引
		res[1] = end
	} else if A[start] == target {
		res[1] = start
	}
	return res
}

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置
// 思路：找到第一个大于等于目标值的位置
func searchInsert(A []int, target int) int {
	start := 0
	end := len(A) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		if A[mid] > target {
			end = mid
		} else if A[mid] < target {
			start = mid
		} else if A[mid] == target {
			end = target
		}
	}
	if A[start] >= target {
		return target
	} else if A[end] >= target {
		return end
	} else if A[end] < target {
		return end + 1
	}
	return 0
}

// 编写一个高效的算法来判断m x n矩阵中，是否存在一个目标值，该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列
// 每行的第一个整数大于前一行的最后一个整数
func searchMatrix(matrix [][]int, target int) bool {
	// 思路：将二维数组转换为一维数组，进行二分搜索
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start := 0
	end := row*col - 1
	for start+1 < end {
		mid := start+(end-start)/2
		// 获取二维数组对应值
		val := matrix[mid/col][mid%col]
		if val < target {
			start = mid
		} else if val > target {
			end = mid
		}else {
			return true
		}
	}
	if matrix[start/col][start%col]==target || matrix[end/col][end%col]==target {
		return true
	}
	return false
}

// 假设你有n个版本[1,2,3,,,n]，你想找出导致之后所有版本出错的第一个错误的版本。你可以通过调用bool isBadVersion(version)接口
// 来判断版本号version是否在单元测试中出错。实现一个函数来查找第一个错误的版本，你应该尽量减少调用API的次数。
func firstBadVersion(n int) int {
	// 思路：二分搜索
	start := 0
	end := n
	for start+1 < end {
		mid := start+(end-start)/2
		if isBadVersion(mid) {
			end = mid
		}else if isBadVersion(mid)==false {
			start=mid
		}
	}
	if isBadVersion(start){
		return start
	}
	return end
}
func isBadVersion(n int) bool {
	return true
}

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转（例如，数组[0,1,2,4,5,6,7]
// 可能变成[4,5,6,7,0,1,2]。请找出其中最小的元素。
func findMin(nums []int) int {
	// 思路：
}
