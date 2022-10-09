package main

import (
	"fmt"
	"math/bits"
)

func main() {
	min := minBitFlips(10, 7)
	fmt.Println(min)
	//https://leetcode.cn/problems/minimum-bit-flips-to-convert-number
	//转换数字的最小反转次数
}
func minBitFlips(start int, goal int) int {
	return bits.OnesCount(uint(start ^ goal))
}
