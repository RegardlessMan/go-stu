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

// 轮转数组
func rotate(nums []int, k int) {
	tmp := make([]int, k)
	n := len(nums)
	j := 0
	k %= n
	for i := n - k; i < n; i++ {
		tmp[j] = nums[i]
		j++
	}
	for i := n - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}
func rotateImplementTwo(nums []int, k int) {
	n := len(nums)
	k %= n
	reverseNums(nums, 0, n-1)
	reverseNums(nums, 0, k-1)
	reverseNums(nums, k, n-1)
}

func reverseNums(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

/*
*
买卖股票的最佳时机
*/
//func maxProfit(prices []int) int {
//	maxValue := 0
//	minValue := math.MaxInt64
//	for i := 0; i < len(prices); i++ {
//		if prices[i] < minValue {
//			minValue = prices[i]
//		}
//		if prices[i]-minValue > maxValue {
//			maxValue = prices[i] - minValue
//		}
//	}
//
//	return maxValue
//}

// maxProfit 买卖股票的最佳时机 II  动态规划解决
func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

// canJump 跳跃游戏
func canJump(nums []int) bool {
	n := len(nums)
	high := 0
	for i := 0; i < n; i++ {
		if i <= high {
			high = max(high, i+nums[i])
			if high >= n-1 {
				return true
			}
		}
	}

	return false
}

// jump 跳跃游戏 II
func jump(nums []int) int {
	maxPosition := 0
	end := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}
