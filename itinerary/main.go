package main

import "fmt"

func findItinerary(flight map[string]string) []string {
	var itenerary []string

	for len(flight) > 0 {
		from := ""
		for k, v := range flight {
			if _, ok := flight[v]; !ok {
				from = k
				break
			}
		}
		to := flight[from]
		itenerary = append([]string{from, to}, itenerary...)
		delete(flight, from)
		from = to
	}
	return itenerary
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var f int
		fmt.Scan(&f)
		flights := make(map[string]string)
		for j := 0; j < f; j++ {
			from, to := "", ""
			fmt.Scan(&from, &to)
			flights[from] = to
		}
		itinerary := findItinerary(flights)
		fmt.Printf("Case #%d:", i+1)
		for k := 0; k < len(itinerary); k += 2 {
			fmt.Printf(" %s-%s", itinerary[k], itinerary[k+1])
		}
		fmt.Println("")
	}
}
