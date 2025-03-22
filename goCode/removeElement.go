package main

func removeElement(nums []int, val int) int {
	j := 0
	for i := range nums {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return len(nums)
}

func main() {
	numList := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeElement(numList, 2)
}
