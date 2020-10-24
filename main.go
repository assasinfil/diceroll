package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
)

// Фукнция должна вернуть число из интервала [l,r]
func randInterval(l, r int) int {
	return l + rand.Intn(r)
}

func main() {
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().Unix())
	} else {
		rand.Seed(int64(*seed))
	}
	// Dice roll 1..6
	fmt.Println(randInterval(1, 6))
}
