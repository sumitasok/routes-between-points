package main

import (
	"bufio"
	"fmt"
	"os"
)

type Routes struct {
	Path  map[string]map[string]int
	Build []Path
}

type Path struct {
	A     string
	B     string
	Count int
}

func (r *Routes) Increment(a, b string, pathCount int) {
	a, b = alphabeticSort(a, b)

	p := Path{a, b, pathCount}
	r.Build = append(r.Build, p)

	if r.Path == nil {
		r.Path = map[string]map[string]int{}
	}

	if _, ok := r.Path[a]; ok {
		r.Path[a][b] = r.Path[a][b] + pathCount
	} else {
		r.Path[a] = map[string]int{
			b: pathCount,
		}
	}
}

func (r *Routes) Reset() {
	r.Path = map[string]map[string]int{}

	for _, path := range r.Build {
		r.Increment(path.A, path.B, path.Count)
	}
}

func (r *Routes) Traverse(route string) (bool, int) {
	r.Reset()
	return true, 0
}

func alphabeticSort(a, b string) (string, string) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReader(os.Stdin)
	route, _ := reader.ReadString('\n')
	fmt.Println(route)
}
