package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	r := NewReader("ins.txt")
	t := r.Int()
	for i := 0; i < t; i++ {
		n, m := r.Int(), r.Int()
		nodemap := r.EdgeList(n, m)
		s := solve(nodemap)
		fmt.Println(s)
	}
}

func solve(nodemap map[int]*node) int {
	for _, n := range nodemap {
		if n.visit {
			continue
		}
		n.visit = true
		queue := n.child
		for ; len(queue) > 0; queue = queue[1:] {
			e := queue[0]
			if e.n1 == nil && e.n2 == nil {
				continue
			}
			var n *node
			switch {
			case e.n1.visit && !e.n2.visit:
				n = e.n2
			case !e.n1.visit && e.n2.visit:
				n = e.n1
			case e.n1.visit && e.n2.visit:
				return 1
			}
			queue = append(queue, n.child...)
			*e = edge{}
			n.visit = true
		}
	}
	return 0
}

type Reader struct {
	io.Reader
}

func NewReader(path string) *Reader {
	f, err := os.Open(path)
	if err != nil {
		f = os.Stdin
	}
	return &Reader{f}
}

func (r *Reader) Int() (n int) {
	_, e := fmt.Fscan(r, &n)
	if e != nil && e != io.EOF {
		panic(e)
	}
	return n
}

type node struct {
	id    int
	visit bool
	child []*edge
}

func (e *edge) String() string {
	if e.n1 == nil && e.n2 == nil {
		return "nil"
	}
	return fmt.Sprint(e.n2.id, e.n1.id, ",")
}

type edge struct {
	n1 *node
	n2 *node
}

func (r *Reader) EdgeList(numnodes int, numedges int) (nodes map[int]*node) {
	var from, to int
	nodes = make(map[int]*node, numnodes)
	for i := 0; i < numedges; i++ {
		_, e := fmt.Fscan(r, &from, &to)
		if e != nil && e != io.EOF {
			panic(e)
		}
		N1, ok := nodes[from]
		if !ok {
			N1 = new(node)
			N1.id = from
			nodes[from] = N1
		}
		N2, ok := nodes[to]
		if !ok {
			N2 = new(node)
			N2.id = to
			nodes[to] = N2
		}
		edge := &edge{N1, N2}
		N1.child = append(N1.child, edge)
		N2.child = append(N2.child, edge)
	}
	return nodes
}
