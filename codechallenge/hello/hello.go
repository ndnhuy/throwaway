package main

import "fmt"

func main() {
	// fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
    arr := []int{1,2,3,4,5}
    fmt.Println(binarySearch(arr, 0, 0, len(arr)-1))
}

func threeSum(nums []int) [][]int {
	rs := [][]int{}
	for i, ni := range nums {
		for j, nj := range nums[i+1:] {
			for _, nk := range nums[i+j+1:] {
				sum := ni + nj + nk
				if sum == 0 {
					rs = append(rs, []int{ni, nj, nk})
				}
			}
		}
	}

	return rs
}

// nums: sorted array
func binarySearch(nums []int, val int, low int, high int) bool {
    if low > high {
        return false
    }
    mid := (low+high) / 2
    if nums[mid] == val {
        return true
    } else {
        if val > nums[mid] {
            return binarySearch(nums, val, mid+1, high)
        } else {
            return binarySearch(nums, val, low, mid-1)
        }
    }
}