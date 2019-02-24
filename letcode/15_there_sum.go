package letcode

func ThreeSum(nums []int) [][]int {

	var res = [][]int{}

	max := 0
	_ = max
	target := -nums[0]
	i := 1
	j := len(nums) - 1
	nowData := 0
	_ = nowData
	for {
		if i >= j {
			break
		}
		// 双指针进行判断
		if (nums[i] + nums[j]) > target {
			//nowData = nums[j-1] * (j - i - 1)
			j--
		} else if (nums[i] + nums[j]) < target {
			//nowData = nums[i] * (j - i - 1)
			i++
		} else if (nums[i] + nums[j]) == target {
			//fmt.Println("henghengheng", nums[i], nums[j], nums[0])
			var resNow = []int{nums[i], nums[j], nums[0]}
			res = append(res, resNow)
			j--
			i++
		}
	}
	return res
}
