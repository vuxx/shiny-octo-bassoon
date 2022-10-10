package main

import (
	"fmt"
	"math/bits"
)

func main() {
	x := hammingWeight(10)
	fmt.Println(x)
}

// https://leetcode.cn/problems/number-of-1-bits/submissions/
func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}
