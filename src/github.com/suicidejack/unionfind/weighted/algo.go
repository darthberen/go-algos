package weighted

/************************************************************************
 * Author: Matt Price
 * Last Modified: Thu Jun 19, 2014 21:05:36 EDT
 * Purpose: Algorithm for the quick-find union find
 ************************************************************************/

import (
	"fmt"
	"strconv"
)

var id []int
var sz []int

func Init(numElements int) {
	id = make([]int, numElements)
	sz = make([]int, numElements)
	for i := 0; i < numElements; i++ {
		id[i] = i
		sz[i] = 1
	}
}

func String() (str string) {
	str = ""
	for _, val := range id {
		str += strconv.Itoa(val) + " "
	}
	return fmt.Sprintf("%s", str)
}

func root(i int) int {
	root := -1
	for root != id[i] {
		root = id[i]
	}
	return root
}

func Connected(p, q int) bool {
	return root(p) == root(q)
}

func Union(p, q int) {
	i := root(p)
	j := root(q)
	if i == j {
		return
	}
	if sz[i] < sz[j] {
		id[i] = j
		sz[j] += sz[i]
	} else {
		id[j] = i
		sz[i] += sz[j]
	}
}
