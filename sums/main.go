package main

import (
	"fmt"
	"sort"
)

type testCase struct {
	N       int
	Nums    []int
	Request []int
}

var cache = make(map[int64]bool)

func sumsOfSumsRec(nums []int, from, to int) []int {
	if _, ok := cache[int64(from)<<31+int64(to)]; ok {
		return []int{}
	}
	if from == to {
		return nums[from:from]
	}
	var res []int
	cache[int64(from)<<31+int64(to)] = true
	res = append(res, sumsOfSumsRec(nums, from+1, to)...)
	res = append(res, sumsOfSumsRec(nums, from, to-1)...)
	fmt.Printf("from %d to %d\n", from, to)
	if len(res) > 0 {
		res = append(res, res[len(res)-1] + nums[to-1])
	} else {
		res = append(res, sums(nums, from, to))
	}
	return res
}

func sumsOfSums(nums []int) []int {
	var res []int
	for s := 0; s < len(nums); s++ {
		for i := 0; i+s < len(nums); i++ {
			res = append(res, sums(nums, i, i+s+1))
		}
	}
	sort.Ints(res)
	return res
}

func sums(nums []int, from, to int) int {
	var acc int
	for i := from; i < to; i++ {
		acc += nums[i]
	}
	return acc
}

func main() {
	var testCases []testCase
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var tCase testCase
		var nbN, nbR int
		tCase.N = i + 1
		fmt.Scan(&nbN, &nbR)
		for j := 0; j < nbN; j++ {
			var n int
			fmt.Scan(&n)
			tCase.Nums = append(tCase.Nums, n)
		}
		for j := 0; j < nbR; j++ {
			var from, to int
			fmt.Scan(&from, &to)
			tCase.Request = append(tCase.Request, from, to)
		}
		testCases = append(testCases, tCase)
	}
	var a []int
	for i := 0; i < 20000; i++ {
		a = append(a, i)
	}
	res := sumsOfSumsRec(a, 0, len(a))
	fmt.Println(sums(res, 0, 10))

	for _, c := range testCases {
		cache = make(map[int64]bool)
		fmt.Printf("Case #%d:\n", c.N)
		//res := sumsOfSums(c.Nums)
		resRec := sumsOfSumsRec(c.Nums, 0, len(c.Nums))
		sort.Ints(resRec)
		//fmt.Println("res: ", res)
		fmt.Println("resRec: ", resRec)
		for i := 0; i < len(c.Request); i += 2 {
			fmt.Printf("%d\n", sums(resRec, c.Request[i]-1, c.Request[i+1]))
			//fmt.Printf("(%d)\n", sums(resRec, c.Request[i]-1, c.Request[i+1]))
		}
	}

}
