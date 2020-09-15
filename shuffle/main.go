package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix [][]int

func main() {
	file, err := os.Open("ins.txt")
	if err != nil {
		file = os.Stdin
	}
	reader := bufio.NewReader(file)
	var N int
	fmt.Fscanf(reader, "%d\n", &N)
	K := parse(N, reader)
	M := process(K)
	fmt.Println(M.Range())
}

func process(m matrix) matrix {
	var n int = len(m)
	var M matrix
	for i := 0; i < 2*n; i++ {
		M = append(M, make([]int, n*n))
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if m[i][j] == 1 {
				M[i][count] = 1
				M[n-1+j][count] = 1
				count++
			}
		}
	}
	for i, m := range M {
		M[i] = m[:2*n]
	}
	return M
}

func parse(N int, reader *bufio.Reader) matrix {
	var K matrix
	for i := 0; i < N; i++ {
		K = append(K, make([]int, N))
	}
	K.AddInt(1)

	for i := 0; i < N; i++ {
		str, _ := reader.ReadString('\n')
		for _, js := range strings.Split(str, " ") {
			js = strings.TrimSpace(js)
			j, _ := strconv.Atoi(js)
			K[i][j] = 0
		}
		K[i][i] = 0
	}
	return K
}

func (m matrix) String() string {
	var str []string
	for _, k := range m {
		str = append(str, fmt.Sprint(k))
	}
	str = append(str, "")
	return strings.Join(str, "\n")
}

func (m matrix) AddInt(x int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			m[i][j] += x
		}
	}
}

func (m matrix) SwapRow(a, b int) {
	m[a], m[b] = m[b], m[a]
}

func (m matrix) AddRow(a, b, x int) {
	for i := range m[0] {
		m[b][i] += x * m[a][i]
	}
}

func (m matrix) MultRow(a, x int) {
	for i := range m[0] {
		m[a][i] *= x
	}
}

func (m matrix) Range() (res int) {
	res = len(m)
	for c := 0; c < len(m); c++ {
		pivotrow := -1
		for r := c; r < len(m); r++ {
			if m[r][c] != 0 {
				pivotrow = r
				break
			}
		}
		if pivotrow == -1 {
			res--
			continue
		}
		m.SwapRow(c, pivotrow)
		for r := c; r < len(m); r++ {
			if r == c {
				continue
			}
			if m[r][c] == 0 {
				continue
			}
			m.MultRow(c, m[r][c])
			m.MultRow(r, m[c][c])
			sign := 1
			if 0 < m[c][c]*m[r][c] {
				sign = -1
			}
			m.AddRow(c, r, sign)
		}
	}
	return res
}
