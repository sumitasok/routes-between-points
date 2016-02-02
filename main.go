package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	path := strings.Split(route, "")

	for i := 0; i < len(path)-1; i++ {
		a, b := alphabeticSort(path[i], path[i+1])

		if _, oka := r.Path[a]; oka {
			if _, okb := r.Path[a][b]; okb {
				if r.Path[a][b] > 0 {
					r.Path[a][b] = r.Path[a][b] - 1
				} else {
					r.Reset()
					return false, 0
				}
			} else {
				r.Reset()
				return false, 0
			}
		} else {
			r.Reset()
			return false, 0
		}
	}
	// start := path[0]
	// end := path[len(path)-1]
	// fmt.Println(start, end)
	r.Reset()
	return true, 1
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
	// fmt.Println(route)

	r := Routes{}

	r.Increment("a", "b", 1)
	r.Increment("b", "c", 3)
	r.Increment("c", "d", 3)
	r.Increment("b", "d", 2)
	r.Increment("d", "e", 1)

	route = strings.Replace(route, "\n", "", -1)
	// fmt.Println(route)

	ok, count := r.Traverse(route)
	fmt.Println(ok)
	fmt.Println(count)
}
