package binary

// 基本操作
// a=0^a=a^0    a与0异或等于a
// 0=a^a
// 由上面两个推导出：a=a^b^b

// 交换两个数
// a=a^b
// b=a^b
// a=a^b   从而达到交换两个数的目的

// 移除最后一个1
// a=n&(n-1)

// 获取最后一个1
// diff=(n&(n-1))^n

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次，找出那个只出现了一次的元素
func singleNumber(nums []int) int {
	// 10 ^ 10 = 0
	// 两个数异或就变成0
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次，找出那个只出现了一次的元素
func singleNumber2(nums []int) int {
	// 统计每位上1的个数
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计1的个数
			sum += (nums[j] >> i) & 1
		}
		// 还原位 00^10=10 或者用 | 也可以
		result ^= (sum % 3) << i
	}
	return result
}

// 给定一个整数数组nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。找出只出现一次的那两个元素
// 思路：可根据a=a^b^b得到两个元素异或的结果，两个元素设为a、b，则得到x=a^b
// 怎么把两个元素分开？可以根据x&(x-1)去掉最后一位1，再根据x&(x-1)^x得到最后一位1的位置
// 最后一位1代表a在这位上为1，b在这位上为0（或者a在这位上为0，b为1），从而可以分开两个元素
func singleNumber3(nums []int) []int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	result := []int{diff, diff}
	diff = (diff & (diff - 1)) ^ diff // 得到最后一位1的位置
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			result[0] ^= nums[i]
		} else {
			result[1] ^= nums[i]
		}
	}
	return result
}

// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为‘1’的个数（也被称为汉明重量）
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

// 给定一个非负整数num，对于0 <= i <= num范围中的每个数字i，计算其二进制数中的1的数目并将它们作为数组返回
func countBits1(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[i] = count1(i)
	}
	return res
}
func count1(i int) (res int) {
	for i != 0 {
		i = i & (i - 1)
		res++
	}
	return
}
// 另一种解法：利用动态规划法
func countBits2(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1 // 上一个缺1的元素+1即可
	}
	return res
}

// 颠倒给定的32位无符号整数的二进制位
// 思路：依次颠倒即可
func reverseBits(num uint32) uint32 {
	res := uint32(0)
	bits := uint32(31)
	for num != 0 {
		res += (num & 1) << bits
		num = num >> 1
		bits--
	}
	return res
}

// 给定范围[m,n]，其中0<=m<=n<=2147483647，返回此范围内所有数字的按位与（包含m,n两端点）
// 思路：如果m的位数小于n的位数的话，那么所有数字的按位与肯定是0；
// 如果m的位数等于n的位数的话，那么两个数最高那几位相同的部分按位与是1，剩下的低位按位与肯定是0；总之结果取决于m
// 比如m=1100 0010, n=1100 1100, 那么结果就是1100 0000；即得出两个数最左边相同的部分，剩下的都为0
func rangeBitwiseAnd(m,n int) int {
	var count int // 找出要补0的位数
	for m!=n {
		m >>= 1
		n >>= 1
		count++
	}
	m <<= count
	return m
}