package main

/************************************************************************
 * Author: Matt Price
 * Last Modified: Thu Jun 19, 2014 20:52:42 EDT
 * Purpose: Display the running of various quick-find algorithms as taken
 *          from the coursera class algorithms part 1 starting June 2014.
 ************************************************************************/

import (
	"fmt"
	eager "github.com/suicidejack/unionfind/eager"
	weighted "github.com/suicidejack/unionfind/weighted"
)

func main() {
	runEager()
	runWeighted()
}

func runWeighted() {
	weighted.Init(10)
	weighted.Union(8, 4)
	weighted.Union(9, 7)
	weighted.Union(6, 1)
	weighted.Union(5, 0)
	weighted.Union(7, 2)
	weighted.Union(2, 3)
	weighted.Union(5, 3)
	weighted.Union(8, 1)
	weighted.Union(5, 8)
	fmt.Println(weighted.String())
}

func runEager() {
	eager.Init(10)
	eager.Union(9, 1)
	eager.Union(3, 6)
	eager.Union(0, 5)
	eager.Union(0, 7)
	eager.Union(2, 0)
	eager.Union(8, 9)
	fmt.Println(eager.String())
}
