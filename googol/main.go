package main

import (
	"fmt"
	"math"
)

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func switchStr(s string) string {
	res := []rune(s)
	for i, c := range s {
		if c == '0' {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

func stringify(l uint64, r []uint64) string {
	res := ""
	var i uint64
	var j int
	for ; i < l; i++ {
		if i == r[j] {
			res += "1"
			j++
		} else {
			res += "0"
		}
	}
	return res
}

func isOne(l, k uint64) bool {
	if k == (l / 2) {
		return false
	} else if k < l/2 {
		return isOne(l/2, k)
	} else {
		return !isOne(l/2, l-k-1)
	}
}

func findKthCharacter(k uint64) string {
	var idx uint64
	for idx <= k {
		idx = idx*2 + 1
	}
	if isOne(idx, k) {
		return "1"
	}
	return "0"
}

func smarterGenerateGoogleString(limit int) (uint64, []uint64) {
	res := make([]uint64, 1<<30-1)
	var idx uint64 = 5
	var l uint64 = 3
	var ci = 1
	res[0] = 2
	for i := 0; i < (limit-2) && float64(l) < math.Pow10(9); i++ {
		tmp := ci
		//fmt.Println(res[:ci])
		//fmt.Println(stringify(l, res[:ci]))
		for j := 0; j < tmp; j++ {
			if idx < res[j]+l+uint64(1) {
				res[ci] = idx
				idx = idx*2 + 1
				ci++
			}
			res[ci] = res[j] + l + uint64(1)
			ci++
		}
		l = (l * 2) + 1
		//os.Stderr.WriteString(fmt.Sprintf("%v", res))
		//os.Stderr.WriteString(fmt.Sprintf("[%d] Current length: %d\n", i, l))
	}
	return l, res
}

func smartGenerateGoogleString(limit int) string {
	res := []rune{'0'}
	var idx int
	for i := 0; i < (limit-1) && float64(len(res)) < math.Pow10(18); i++ {
		fmt.Println(string(res))
		tail := make([]rune, len(res))
		copy(tail, res)
		tail[idx] = '1'
		res = append(res, '0')
		res = append(res, tail...)
		idx = idx*2 + 1
	}
	return string(res)
}

func generateGoogolString(limit int) string {
	res := ""
	for i := 0; i < limit; i++ {
		if float64(len(res)) > math.Pow10(18) {
			return res
		}
		res = res + "0" + switchStr(reverseStr(res))
		fmt.Println(res)
	}
	return res
}

func main() {
	var max int
	fmt.Scan(&max)
	for i := 0; i < max; i++ {
		var k uint64
		fmt.Scan(&k)
		fmt.Printf("Case #%d: %s\n", i+1, findKthCharacter(k-1))
	}
}
