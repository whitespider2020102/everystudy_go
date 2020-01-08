/*
!/usr/bin/env GO114
@time    :  2020/01/08
@Author  :  WhiteSpider
@File    :  day2_longestCommonPrefix.py
@Version :  1.14
@Software:  GoLand
@Blog    :  https://github.com/whitespider2020102/everystudy_go
说明：具体参考leetcode，最长公共前缀
题目：编写一个函数来查找字符串数组中的最长公共前缀,如果不存在公共前缀，返回空字符串 ""。
具体思路：
根据题意，将第一个元素作为基准元素（假设基准元素为x，后面的元素依次为x1,x2,x3,x4.....），可以得出以下结论：
1、最长公共前缀一定产生于基准元素（或没有） --- 专门列出这个是有些同学容易忽视掉该条件
2、最长公共前缀的长度小于或等于基准元素长度。
3、如果strings.Index(x1,x) == 0,则直接跳过x1继续与x2进行比较
4、如果strings.Index(x1,x) != 0 则依次截取掉x的末端元素，继续与x1进行比较。
4/1 、如果满足3条件则跳过，继续与x2元素进行比较
4/2 、如果x截取到空，则无需进行比较，x为空串
*/

package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, code := range strs {
		for strings.Index(string(code), prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[0 : len(prefix)-1]
		}
	}
	return string(prefix)
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
