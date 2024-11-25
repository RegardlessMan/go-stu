/**
 * @Author QG
 * @Date  2024/11/25 22:07
 * @description
**/

package first

/*
*
两数之和
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	sortNums := make([]int, m+n)

	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		if p1 < m {
			sortNums[p1+p2] = nums1[p1]
			p1++
		}
		if p2 < n {
			sortNums[p1+p2] = nums2[p2]
			p2++
		}
	}

	if p1 < m {
		copy(sortNums[p1+p2:], nums1[p1:])
	}
	if p2 < n {
		copy(sortNums[p1+p2:], nums2[p2:])
	}
}
