package base

// 常考排序：快速排序、归并排序、堆排序
// 此外有十大经典排序算法（分为比较类排序和非比较类排序）
// 比较类排序：冒泡、快速（交换）；简单插入、希尔排序（插入）；堆排序、简单选择（选择）；二路归并排序、多路归并排序（归并）；
// 非比较类排序：计数排序、桶排序、基数排序

// 快速排序
func QuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}
func quickSort(nums []int, start, end int) {
	if start < end {
		// 分治法
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}
// 定一个基准点，先将数组分为小于该基准点的和大于等于该基准点的两部分
func partition(nums []int, start, end int) int {
	// 数组最后一位当作基准点
	p := nums[end]
	i := start
	// 最后一位是基准点，所以就不需要比较最后一位
	for j:=start;j<end;j++{
		// i将一直和j保持同步，或者i停留在大于等于p的地方
		if nums[j] < p {
			swap(nums, i, j)
			i++
		}
	}
	// 最后将基准值交换到它的位置
	swap(nums, i, end)
	// 返回基准值的位置，i左边都是小于基准值的值，i右边都是大于等于基准值的值
	return i
}
func swap(nums []int, i,j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}