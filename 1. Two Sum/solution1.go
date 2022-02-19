package TwoSum

func twoSum(nums []int, target int) []int {
    diffs := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        diffs[target-nums[i]] = i
        for j := i+1; j < len(nums); j++ {
            _, isPresent := diffs[nums[j]]
            if isPresent {
                return []int{i, j}
            }
        }
    }
    return nil
}