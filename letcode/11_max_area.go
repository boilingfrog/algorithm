package letcode

func MaxArea(height []int) int {

	max := 0
	_ = max
	i := 0
	j := len(height)
	nowData := 0
	_ = nowData
	for {
		if i == j {
			break
		}
		// 双指针进行判断
		if height[i] > height[j-1] {
			j--
			nowData = height[j-1] * (j - i)
		} else {
			i++
			nowData = height[j-1] * (j - i)

		}
		if max < nowData {
			max = nowData
		}
	}
	return max
}
