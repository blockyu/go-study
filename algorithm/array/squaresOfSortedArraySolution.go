package main

import (
	"fmt"
	"math"
)

/*

Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.



Example 1:

Input: nums = [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Explanation: After squaring, the array becomes [16,1,0,9,100].
After sorting, it becomes [0,1,9,16,100].
Example 2:

Input: nums = [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums is sorted in non-decreasing order.


Follow up: Squaring each element and sorting the new array is very trivial, could you find an O(n) solution using a different approach?

1. 0 혹은 음수 양의 정수가 나뉘어 지는 지점을 찾는다
2. 0 이 있으면 가장 앞에 추가한다
3. left, right 의 square 값을 비교후 작은 값을 앞쪽에 추가한 후 이동한다.

 */

func sortedSquares(nums []int) []int {
	lptr, rptr := 0, len(nums)-1
	arr := make([]int, len(nums))
	for lptr < rptr {
		if math.Abs(float64(lptr)) > math.Abs(float64(rptr)) {
			arr = append(arr, int(math.Sqrt(float64(nums[lptr]))))
			lptr++
		}else {
			arr = append(arr, int(math.Sqrt(float64(nums[rptr]))))
			rptr--
		}
	}
	fmt.Print(arr)
	return arr
}

func main() {
	sortedSquares([]int{-7,-3,2,3,11})
}