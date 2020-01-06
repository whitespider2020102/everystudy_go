/*
日期：2020-01-06
项目组成员：吴金翰
说明：具体参考leetcode，整数反转
题目：给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转
具体思路:1. Go 的取模运算很方便，不用单独判断正负。
        2.int 取值范围为 [-2^{31}, 2^{31} - 1]，如果翻转数字溢出，则立即返回 0。
git渠道：GitHub个人库
*/

package main

import "fmt"

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}

func main() {
	x := 123
	fmt.Println(reverse(x))
}
