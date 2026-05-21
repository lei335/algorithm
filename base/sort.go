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

// 归并排序
func MergeSort(nums []int) []int {
	return mergeSort(nums)
}
func mergeSort(nums []int) []int {
	// 分治法
	// 先确定返回条件
	if len(nums)<=1 {
		return nums
	}
	// 再找出中点
	mid := len(nums)/2
	// 递归处理
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	// nums = merge(left, right) // 这里不应该直接赋给nums
	result := merge(left, right)
	return result
}
func merge(left, right []int) []int {
	var result []int
	l,r := 0,0
	// 合并两个有序数组
	for l<len(left) && r<len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		}else {
			result = append(result, right[r])
			r++
		}
	}
	// 合并剩余部分
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// 堆排序
// 完全二叉树：除最后一层之外，其他所有层都是满的且最后一层的叶子节点都在最左边
// 堆包含大顶堆和小顶堆。大顶堆：所有根节点都大于等于其左右孩子节点的完全二叉树；小顶堆：所有根节点都小于等于其左右孩子节点的完全二叉树
func HeapSort(a []int) []int {
	// 先构造成一个大顶堆（升序用大顶堆，降序用小顶堆）
	// 1. 构造大顶堆：从最后一个非叶子节点开始构造（找出其和其左右孩子中最大的节点当根节点）
	for i:=len(a)/2-1;i>=0;i--{
		sink(a, i, len(a))
	}

	// 2. 交换a[0]和a[len-1], 即将堆顶节点和最后一个节点互换，
	// 3. 然后将前n-1个节点（总共n个节点）再次下沉保持成一个大顶堆
	// 4. 再次将剩余的n-1个节点构造的大顶堆中，堆顶节点和最后一个节点（即a[len-2]）互换，依此类推
	for i:=len(a)-1;i>=1;i--{
		heapswap(a, i, 0)
		sink(a, 0, i)
	}
	return a
}
func sink(a []int, i, length int) {
	// 最后一个非叶子节点是len(a)/2-1（索引从0开始）
	for {
		idx := i // 记录根、左、右三个节点中最大值的索引
		l := 2*i + 1 // 左节点索引为2*i+1
		r := 2*i + 2 // 右节点索引为2*i+2
		// 存在左节点，且左节点值较大，则取左节点
		if l < length && a[l] > a[idx] {
			idx = l
		}
		// 存在右节点，且右节点值较大，则取右节点
		if r < length && a[r] > a[idx] {
			idx = r 
		}
		// 根节点较大，则不用交换值
		if idx == i {
			break
		}
		// 如果根节点较小，则交换值，并且继续下沉
		heapswap(a, idx, i)
		// 继续下沉idx节点
		i = idx
	}
}
func heapswap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}