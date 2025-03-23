package main

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}

func main() {
	numList := []int{1, 2, 1}
	getConcatenation(numList)
}
