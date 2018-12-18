package main

import (
	"fmt"
	"time"

	"../levenshteindistance"
)

func main() {
	start := time.Now()
	keyname, strlist := levenshteindistance.GetStrings("./input.txt")
	levenshteindistance.PrintPairs(levenshteindistance.MakePairs(keyname, strlist))
	fmt.Println("finished in : ", time.Now().Sub(start))
}
