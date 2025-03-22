package main

func removeDuplicates(nums []int) int {
	j := 0
	for i := range nums {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}

func main() {
	numList := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(numList)
}
