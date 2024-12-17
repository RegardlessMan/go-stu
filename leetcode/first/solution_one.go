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

/*
*
移除元素
*/
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

/*
*
删除有序数组中的重复项
*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l := 1
	for i := 1; i < n; i++ {
		if nums[i-1] != nums[i] {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}

/*
*
多数元素
*/
func majorityElement(nums []int) int {
	n := len(nums)
	tmp := nums[0]
	count := 0
	for i := 0; i < n; i++ {
		if nums[i] == tmp {
			count++
		} else if count <= 0 {
			tmp = nums[i]
			count++
		} else {
			count--
		}
	}

	return tmp
}
