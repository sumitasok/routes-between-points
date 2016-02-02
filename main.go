package main

import (
	"bufio"
	"fmt"
	"os"
)

type Routes struct {
	Path map[string]*Path
}

type Path map[string]int

func (r *Routes) Add(a, b string, pathCount int) {
	a, b = alphabeticSort(a, b)
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
