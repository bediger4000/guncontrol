package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	aProbability := flag.Int("A", 75, "Proability in percent that Marksman A hits the target")
	bProbability := flag.Int("B", 25, "Proability in percent that Marksman B hits the target")
	reps := flag.Int("N", 10000, "Number of rounds of shooting to simulate")
	flag.Parse()

	var aCount, bCount int
	var hit0, hit1, hit2 int

	for i := 0; i < *reps; i++ {
		var aHit, bHit int
		r1 := rand.Intn(100)
		if r1 < *aProbability {
			aHit = 1
		}
		r2 := rand.Intn(100)
		if r2 < *bProbability {
			bHit = 1
		}
		switch aHit + bHit {
		case 0:
			hit0++
		case 1:
			hit1++
			aCount += aHit
			bCount += bHit
		case 2:
			hit2++
		}
	}

	fmt.Printf("A %.03f\n", float64(aCount)/float64(hit1))
	fmt.Printf("B %.03f\n", float64(bCount)/float64(hit1))
	fmt.Printf("0: %d (%.03f)\t1: %d (%.03f)\t2: %d (%.03f)\n",
		hit0,
		float64(hit0)/float64(*reps),
		hit1,
		float64(hit1)/float64(*reps),
		hit2,
		float64(hit2)/float64(*reps),
	)
}
