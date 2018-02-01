package main

/*
	1,3,5 coin  num = min{ f(x) }
	DP

	d[0] = 0
	d[1] = d[0] + 1
	d[2] = d[1] + 1
	d[3] = min{d[3 - 1] + 1, d[3 - 3] + 1}
	d[i] = min{d[i-1]+1, d[i-3]+1, d[i-5]+1}

	array cache result
*/

//const UINT_MIN uint = 0
//const UINT_MAX = ^uint(0)
//const INT_MAX = int(^uint(0) >> 1)
//const INT_MIN = ^INT_MAX

import (
	"fmt"
	"os"
	"strconv"
)

type Data struct {
	count int
	array [100]int
}

var coin [3]int = [3]int{1, 3, 5}

func Dynamic2(num int) {
	dp := make([]Data, num+1)

	for i := 0; i < num+1; i++ {
		dp[i].count = i

		for m := 0; m < i; m++ {
			dp[i].array[m] = 1
		}

		for j := 0; j < 3; j++ {
			if coin[j] <= i && dp[i-coin[j]].count+1 < dp[i].count {
				dp[i].count = dp[i-coin[j]].count + 1

				for m := 0; m < dp[i-coin[j]].count; m++ {
					dp[i].array[m] = dp[i-coin[j]].array[m]
				}

				dp[i].array[dp[i].count-1] = coin[j]
			}
		}
	}

	fmt.Println("result: ", dp[num].count)
	for i := 0; i < dp[num].count; i++ {
		fmt.Println(dp[num].array[i])
	}
}

func Dynamic1(num int) {
	dp := make([]int, num+1)

	for i := 0; i < num+1; i++ {

		dp[i] = i

		for j := 0; j < 3; j++ {
			if coin[j] <= i && dp[i-coin[j]]+1 < dp[i] {
				dp[i] = dp[i-coin[j]] + 1
			}
		}
	}

	fmt.Println("result: ", dp[num])
}

func main() {
	alen := len(os.Args)
	if alen != 2 {
		fmt.Println("Usage: ./coin 11")
		return
	}

	num, _ := strconv.Atoi(os.Args[1])

	Dynamic1(num)

	Dynamic2(num)
}
