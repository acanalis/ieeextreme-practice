// https://csacademy.com/ieeextreme-practice/task/bear-sums/statement/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f := os.Stdin
	if fin, err := os.Open("ins.txt"); err == nil {
		f = fin
	}
	nextint := NextInt(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	T := nextint()
	L := make([]int, 20000)
	for t := 0; t < T; t++ {
		S := nextint()
		E := nextint()
		for e := 0; e < E; e++ {
			L[e] = nextint()
		}
		fmt.Fprintln(out, solve(S, E, L[0:E]))
	}
}

func solve(S, E int, L []int) string {
	if E == 0 {
		return "!OK"
	}
	seen := make(map[int]string)
	for _, l := range L {
		side := "l"
		if l > S/2 {
			side = "r"
			l = S - l
		}
		if seen[l] != "" && (seen[l] != side || 2*l == S) {
			return fmt.Sprint(l, S-l)
		}
		seen[l] = side
	}
	return "!OK"
}

func NextInt(r io.Reader) func() int {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	return func() int {
		if !s.Scan() {
			return 0 //EOF
		}
		n, e := strconv.Atoi(s.Text())
		if e != nil {
			panic(e)
		}
		return n
	}
}
