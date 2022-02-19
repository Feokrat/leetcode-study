// best

package TwoSum

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int)
	for i, v := range nums {
		if ind, isPresent := diffs[target-v]; isPresent {
			return []int{ind, i}
		}
		diffs[v] = i
	}
	return nil
}
