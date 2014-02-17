package main

import (
	"flag"
	"fmt"
	"github.com/3xian/elephantlist"
	"github.com/ryszard/goskiplist/skiplist"
	"math/rand"
	"time"
)

var (
	n = flag.Int("n", 100000, "run data in [0,n)")
)

func Sequential() {
	fmt.Println("\n[benchmark 1]")

	e := elephantlist.NewIntList()
	s := skiplist.NewIntMap()

	eSeqStart := time.Now()
	for i := 0; i < *n; i++ {
		e.Set(i, i)
	}
	eSeqDuration := time.Since(eSeqStart)

	sSeqStart := time.Now()
	for i := 0; i < *n; i++ {
		s.Set(i, i)
	}
	sSeqDuration := time.Since(sSeqStart)

	fmt.Printf("sequential set %d elements\n", *n)
	fmt.Println("\telephantlist:\t", eSeqDuration)
	fmt.Println("\tskiplist:\t", sSeqDuration)

	query := rand.Perm(*n)
	eSeqStart = time.Now()
	for _, x := range query {
		e.Get(x)
	}
	eSeqDuration = time.Since(eSeqStart)

	sSeqStart = time.Now()
	for _, x := range query {
		s.Get(x)
	}
	sSeqDuration = time.Since(sSeqStart)
	fmt.Printf("random get %d elements\n", *n)
	fmt.Println("\telephantlist:\t", eSeqDuration)
	fmt.Println("\tskiplist:\t", sSeqDuration)
}

func Random() {
	fmt.Println("\n[benchmark 2]")

	e := elephantlist.NewIntList()
	s := skiplist.NewIntMap()

	query := rand.Perm(*n)
	eRandStart := time.Now()
	for _, x := range query {
		e.Set(x, x)
	}
	eRandDuration := time.Since(eRandStart)

	sRandStart := time.Now()
	for _, x := range query {
		s.Set(x, x)
	}
	sRandDuration := time.Since(sRandStart)

	fmt.Printf("random set %d elements\n", *n)
	fmt.Println("\telephantlist:\t", eRandDuration)
	fmt.Println("\tskiplist:\t", sRandDuration)

	query = rand.Perm(*n)
	eRandStart = time.Now()
	for _, x := range query {
		e.Get(x)
	}
	eRandDuration = time.Since(eRandStart)

	sRandStart = time.Now()
	for _, x := range query {
		s.Get(x)
	}
	sRandDuration = time.Since(sRandStart)
	fmt.Printf("random get %d elements\n", *n)
	fmt.Println("\telephantlist:\t", eRandDuration)
	fmt.Println("\tskiplist:\t", sRandDuration)
}

func main() {
	flag.Parse()

	Sequential()
	Random()
}
