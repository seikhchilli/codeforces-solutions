package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var in, out = bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

func main() {
	defer out.Flush()
	var t int
	var n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		readSlice(arr)
		sortedArr := slices.Clone(arr)
		slices.Sort(sortedArr)
		mn, mx := sortedArr[0], sortedArr[n-1]
		ans := mx - mn
		isAlreadySorted := true
		for i := 0; i < n; i++ {
			if arr[i] != sortedArr[i] {
				isAlreadySorted = false
				ans = min(ans, max(mx-arr[i], arr[i]-mn))
			}
		}
		if isAlreadySorted {
			fmt.Fprintln(out, -1)
			continue
		}
		fmt.Fprintln(out, ans)
	}
}

func readSlice(arr []int) {
	for i := range arr {
		fmt.Fscan(in, &arr[i])
	}
}

func printSlice(arr []int) {
	for i, v := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

func minOfSlice(arr []int) int {
	if len(arr) == 0 {
		panic("Trying to get length of slice size 0.")
	}
	m := arr[0]
	for _, v := range arr {
		m = min(m, v)
	}
	return m
}

func maxOfSlice(arr []int) int {
	if len(arr) == 0 {
		panic("Trying to get length of slice size 0.")
	}
	m := arr[0]
	for _, v := range arr {
		m = max(m, v)
	}
	return m
}
