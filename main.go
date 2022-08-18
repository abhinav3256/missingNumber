package main

import "fmt"

func main() {
	arr := []int{3, 0, 1}

	result := missingNumber(arr)
	fmt.Println(result)

}

func missingNumber(nums []int) int {

	i := 0
	for i < len(nums) {

		if nums[i] == i {
			i++
		} else {
			if nums[i] < len(nums) {
				swap(nums, i, nums[i])
			} else {
				i++
			}

		}

	}
	fmt.Println(nums)
	result := findMissing(nums)

	return result

}
func findMissing(nums []int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	return -1
}

func swap(nums []int, i, j int) []int {

	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp

	return nums
}
