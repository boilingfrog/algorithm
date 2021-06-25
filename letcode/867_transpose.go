package letcode

/*
给定一个矩阵 A， 返回 A 的转置矩阵。

矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。



示例 1：

输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
示例 2：

输入：[[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]


提示：

1 <= A.length <= 1000
1 <= A[0].length <= 1000
*/

func transpose(A [][]int) [][]int {
	len1, len2 := len(A), len(A[0])
	b := make([][]int, len2)
	for i := 0; i < len2; i++ {
		var c []int
		for j := 0; j < len1; j++ {
			c = append(c, A[j][i])
		}
		b[i] = c
	}
	return b
}
