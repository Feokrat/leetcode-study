package TwoSum

func twoSum(nums []int, target int) []int {
	var res []int
	k := 1
	for i, v := range nums {
		tmp := target - v
		for j, el := range nums[i+1:] {
			if el == tmp {
				res = append(res, i)
				res = append(res, j+k)
				return res
			}
		}
		k++
	}
	return res
}
