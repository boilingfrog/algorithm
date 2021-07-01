package letcode

/*
编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。


示例 1：

输入：matrix = [
		[1,4,7,11,15],
		[2,5,8,12,19],
		[3,6,9,16,22],
		[10,13,14,17,24],
		[18,21,23,26,30],
	],
target = 5
输出：true


示例 2：

输入：matrix = [
		[1,4,7,11,15],
		[2,5,8,12,19],
		[3,6,9,16,22],
		[10,13,14,17,24],
		[18,21,23,26,30],
		],
target = 20
输出：false


提示：

m == matrix.length
n == matrix[i].length
1 <= n, m <= 300
-109 <= matix[i][j] <= 109
每行的所有元素从左到右升序排列
每列的所有元素从上到下升序排列
-109 <= target <= 109


解题思路

左下角开始查询，

如果当前数值小于目标值，则数组的二维节点++

如果当前值大于目标值，则数组的第一维度--


错误的想法

刚开始我的思路是从左上角开始

但是发现，这种方法比较麻烦，因为目标值可能存在在横向和纵向的某个值

如果开始一直向横向找，当值大于目标值的时候开始纵向找，这时候要考虑额二维下标的--操作，

应该也行不行但是要考虑的情况比较多

*/

func SearchMatrix(matrix [][]int, target int) bool {
	var i, j = len(matrix) - 1, 0

	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		}
	}
	return false
}
