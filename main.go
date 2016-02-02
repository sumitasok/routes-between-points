package main

import (
	"bufio"
	"fmt"
	"os"
)

type Routes struct {
	Path map[string]map[string]int
}

func (r *Routes) Increment(a, b string, pathCount int) {
	a, b = alphabeticSort(a, b)

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
