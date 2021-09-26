package leetcode

/*
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/

func Divide(dividend int, divisor int) int {
	flag := 0
	if dividend < 0 {
		dividend = 0 - dividend
		flag++
	}
	if divisor < 0 {
		divisor = 0 - divisor
		flag++
	}
	ret := 0
	ret = dd(dividend, divisor)
	if flag == 1 {
		ret = 0 - ret
	}
	if ret > (1<<31)-1 {
		ret = (1 << 31) - 1
	}
	return ret
}

func dd(dividend int, divisor int) int {
	if dividend < divisor {
		return 0
	}
	if dividend == divisor {
		return 1
	}
	ret := 0
	count := 1
	divisor1 := divisor
	for dividend > divisor1 {
		ret += count
		dividend -= divisor1
		divisor1 += divisor1
		count += count
	}
	return ret + dd(dividend, divisor)
}
