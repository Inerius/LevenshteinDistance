package levenshteindistance

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// Pair - struct for storing <str, distance>
type Pair struct {
	str      string
	distance int
}

// QuickSort - sorts pairs by distance
func QuickSort(pairs []Pair) []Pair {
	if len(pairs) < 2 {
		return pairs
	}
	left, right := 0, len(pairs)-1
	pivot := rand.Int() % len(pairs)
	pairs[pivot], pairs[right] = pairs[right], pairs[pivot]
	for i := range pairs {
		if pairs[i].distance < pairs[right].distance {
			pairs[left], pairs[i] = pairs[i], pairs[left]
			left++
		}
	}
	pairs[left], pairs[right] = pairs[right], pairs[left]
	QuickSort(pairs[:left])
	QuickSort(pairs[left+1:])
	return pairs
}

// OpenFile - returns Scanner for input file
func OpenFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}

// GetStrings - get string from input file
func GetStrings(fname string) (string, []string) {
	var strlst []string
	scanner := OpenFile(fname)
	scanner.Scan()
	keyname := scanner.Text()
	for scanner.Scan() {
		strlst = append(strlst, scanner.Text())
	}
	return keyname, strlst
}

// MinOfThree - returns min(min(a, b), c) but in other way
func MinOfThree(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// DpInit - initializates dynamic programming slice
func DpInit(s1, s2 string) [][]int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 0; i <= len(s1); i++ {
		dp[i][0] = i
	}
	for i := 0; i <= len(s2); i++ {
		dp[0][i] = i
	}
	return dp
}

// LevenschteinDistance - calculates Levenschtein distance between s1 and s2 using slice
func LevenschteinDistance(s1, s2 string, dp [][]int) int {
	var c int
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				c = 0
			} else {
				c = 1
			}
			dp[i][j] = MinOfThree(dp[i-1][j-1]+c, dp[i-1][j]+1, dp[i][j-1]+1)
		}
	}
	return dp[len(s1)][len(s2)]
}

// MakePairs - makes []Pair
func MakePairs(keyword string, strlst []string) []Pair {
	pairs := make([]Pair, len(strlst))
	for i := range strlst {
		dp := DpInit(keyword, strlst[i])
		pairs[i] = Pair{strlst[i], LevenschteinDistance(keyword, strlst[i], dp)}
	}
	return QuickSort(pairs)
}

// PrintPairs - prints every pair
func PrintPairs(pairs []Pair) {
	for i := range pairs {
		fmt.Printf("%20s : %d\n", pairs[i].str, pairs[i].distance)
	}
}
