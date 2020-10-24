package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed  = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "left default 1")
	end   = flag.Int("end", 6, "right default 6")
)

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(l, r int) int {
	return l + rand.Intn(r+1)
}

func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	// Dice roll 1..6
	if *start > *end {
		fmt.Println("start > end")
		return
	}
	fmt.Println(randInterval(*start, *end))
}
