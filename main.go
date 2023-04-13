package main

import "fmt"

func main() {
	nums := []int{3,1,3,4,18,10}
	fmt.Println(twoSum(nums, 6))

}

func twoSum (nums []int, target int) []int {
	sliceMap := make(map[int]int)
	result := []int{0,0}
	for index,value := range nums {
		if count, ok := sliceMap[target - value]; ok && count != index {
			result[0] = index
			result[1] = count
			break
		}
		sliceMap[value] = index
	}
	return result
}
