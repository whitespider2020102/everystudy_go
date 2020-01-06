/*
!/usr/bin/env GO114
@time    :  2020/01/06
@Author  :  WhiteSpider
@File    :  day2_Palindrome.py
@Version :  1.14
@Software:  GoLand
@Blog    :  https://github.com/whitespider2020102/everystudy_go
说明：具体参考leetcode，回文数
题目：判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
具体思路:回文数是正序读和倒序读都一样的整数，那我们只要根据输入的target，计算出其倒序的整数，如果x==palindrome，那么x就是回文数。
*/

package main

import "fmt"

func plindrome(target int) bool {
	if target < 0 {
		return false
	}
	palindrome := 0 // palindrome最后得到的值
	x := target     // x最初的target
	for target > 0 {
		palindrome = palindrome*10 + target%10
		target /= 10
	}
	if x == palindrome {
		return true
	} else {
		return false
	}
}

func main() {
	target := 1221
	fmt.Println(plindrome(target))
}
