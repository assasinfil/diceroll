package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed     = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start    = flag.Int("start", 1, "left default 1")
	end      = flag.Int("end", 6, "right default 6")
	n        = flag.Int("n", 1, "round count default 1")
	norepeat = flag.Bool("norepeat", false, "no repeat numbers")
)

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(l, r int) int {
	return l + rand.Intn(r+1)
}

func main() {
	flag.Parse()

	if *start > *end {
		fmt.Println("start > end")
		return
	}

	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}

	if *norepeat {
		length := *end - *start
		if *n > length {
			fmt.Println("numbers count more than length")
			return
		} else {
			var result []int
			step := 0
			for step < *n {
				tmp := randInterval(*start, *end)
				found := false

				for _, value := range result {
					if tmp == value {
						found = true
						break
					}
				}

				if !found {
					result = append(result, tmp)
					step++
				}
			}
			for _, value := range result {
				fmt.Print(value)
				fmt.Print(" ")
			}
		}
	} else {
		for i := 0; i < *n; i++ {
			fmt.Println(randInterval(*start, *end))
		}
	}

}
