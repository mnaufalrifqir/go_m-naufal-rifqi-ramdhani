package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	var jumpOneTime, jumpTwoTimes int
	cost := make([]int, len(jumps))
	cost[0] = 0
	for i := 1; i < len(jumps); i++ {
		jumpOneTime = cost[i-1] + int(math.Abs(float64(jumps[i]-jumps[i-1])))
		if i > 1 {
			jumpTwoTimes = cost[i-2] + int(math.Abs(float64(jumps[i]-jumps[i-2])))
		}
		if i == 1 {
			cost[i] = jumpOneTime
		} else {
			cost[i] = int(math.Min(float64(jumpOneTime), float64(jumpTwoTimes)))
		}
	}
	return cost[len(jumps)-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
