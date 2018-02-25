package main

import (
	"fmt"
)

type GBus struct {
	From int
	To   int
}

type TestCase struct {
	Nb           int
	GBuses       []GBus
	Cities       []int
	WantedAnswer []int
}

func readInput() []TestCase {
	var l int
	var testCases []TestCase
	fmt.Scan(&l)
	for i := 0; i < l; i++ {
		var tc TestCase
		tc.Cities = make([]int, 6000)
		var nbGbus int
		var nbAnswers int
		fmt.Scan(&nbGbus)
		for j := 0; j < nbGbus; j++ {
			var gbus GBus
			fmt.Scan(&gbus.From, &gbus.To)
			tc.GBuses = append(tc.GBuses, gbus)
		}
		fmt.Scan(&nbAnswers)
		for j := 0; j < nbAnswers; j++ {
			var a int
			fmt.Scan(&a)
			tc.WantedAnswer = append(tc.WantedAnswer, a)
		}
		testCases = append(testCases, tc)
	}
	return testCases
}

func resolve(testCase TestCase) []int {
	var answers []int
	for _, bus := range testCase.GBuses {
		for i := bus.From; i <= bus.To; i++ {
			testCase.Cities[i] += 1
		}
	}
	for _, idx := range testCase.WantedAnswer {
		answers = append(answers, testCase.Cities[idx])
	}
	return answers
}

func main() {
	testCases := readInput()
	for i, testCase := range testCases {
		res := resolve(testCase)
		fmt.Printf("Case #%d:", i+1)
		for _, a := range res {
			fmt.Printf(" %d", a)
		}
		fmt.Println("")
	}
}
