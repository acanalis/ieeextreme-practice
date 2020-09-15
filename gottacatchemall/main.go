// https://csacademy.com/ieeextreme-practice/task/e610aba28810ebcf2d3998692269b5a0/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Stdin
	if fin, err := os.Open("ins.txt"); err == nil {
		f = fin
	}
	r := bufio.NewReader(f)

	N := NextInt(r)
	M := NextInt(r)
	R := IntMatrix(r, N, M)
	G := IntMatrix(r, N, M)
	L := IntMatrix(r, N, M)
	for i, row := range R {
		copy(G[i], row)
		copy(L[i], row)
	}
	P := IntMatrix(r, N, M)

	// Be Greedy, Maximize Life at all squares
	for i := 1; i < N; i++ {
		G[i][0] += G[i-1][0]
	}
	for j := 1; j < M; j++ {
		G[0][j] += G[0][j-1]
	}
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if G[i][j-1] > G[i-1][j] {
				G[i][j] += G[i][j-1]
			} else {
				G[i][j] += G[i-1][j]
			}
		}
	}

	// Look for a path with smaller pinchpoint
	for i := 1; i < N; i++ {
		P[i][0] = Min(G[i-1][0], G[i][0])
		L[i][0] = G[i][0]
	}
	for j := 1; j < M; j++ {
		P[0][j] = Min(G[0][j-1], G[0][j])
		L[0][j] = G[0][j]
	}
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if P[i][j-1] > P[i-1][j] {
				L[i][j] += L[i][j-1]
			} else {
				L[i][j] += L[i-1][j]
			}

			if L[i][j] < P[i][j-1] && L[i][j] < P[i-1][j] {

			}

			P[i][j] = Max(
				Min(P[i][j-1], G[i][j-1]+R[i][j]),
				Min(P[i-1][j], G[i-1][j]+R[i][j]),
			)
		}
	}

	res := 1 - P[N-1][M-1]
	if res < 1 {
		res = 1
	}
	fmt.Println(res)
}

func NextInt(r *bufio.Reader) (n int) {
	_, err := fmt.Fscan(r, &n)
	if err == nil || err == io.EOF {
		return n
	}
	panic(err)
}

func Min(N ...int) int {
	min := N[0]
	for _, n := range N {
		if n < min {
			min = n
		}
	}
	return min
}

func Max(N ...int) int {
	max := N[0]
	for _, n := range N {
		if n > max {
			max = n
		}
	}
	return max
}

func DiscardLine(r *bufio.Reader) {
	_, err := r.ReadBytes('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
}

func IntMatrix(r *bufio.Reader, N, M int) [][]int {
	var out [][]int
	for n := 0; n < N; n++ {
		row := make([]int, N)
		for m := 0; m < M; m++ {
			row[m] = NextInt(r)
		}
		out = append(out, row)
	}
	DiscardLine(r)
	return out
}

func String(m [][]int) {
	for _, r := range m {
		fmt.Println(r)
	}
	fmt.Println()
}
