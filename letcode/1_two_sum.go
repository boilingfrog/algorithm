package letcode

func twoSum(nums []int, target int) []int {

	var res []int
	_ = res
	mapp := make(map[int]int)
	for key, item := range nums {

		if _, ok := mapp[target-item]; ok {
			//存在
			//res[1]=1
			//res[mapp[target-item]]=1
			res = append(res, key)
			res = append(res, mapp[target-item])

		}
		mapp[item] = key
	}
	return res
}
