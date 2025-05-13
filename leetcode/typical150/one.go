package typical150

// 合并两个数据
// 解题思路
// 1、从尾部开始填充
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	end := len(nums1) - 1

	for p2 >= 0 {
		// 如果 nums1 已经比较完，则直接拷贝 nums2 剩余部分
		if p1 < 0 || nums2[p2] > nums1[p1] {
			nums1[end] = nums2[p2]
			p2--
		} else {
			nums1[end] = nums1[p1]
			p1--
		}
		end--
	}
}

// 移除数组元素
func removeElement(nums []int, val int) int {
	n := len(nums)
	l := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
