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
	for {
		if p1 >= m || p2 >= n {
			break
		}
		if nums1[p1] < nums2[p2] {
			sortNums[p1+p2] = nums1[p1]
			p1++
		} else {
			sortNums[p1+p2] = nums2[p2]
			p2++
		}
	}
	for ; p1 < m; p1++ {
		sortNums[p1+p2] = nums1[p1]
	}
	for ; p2 < n; p2++ {
		sortNums[p1+p2] = nums2[p2]
	}
	copy(nums1, sortNums)
}
