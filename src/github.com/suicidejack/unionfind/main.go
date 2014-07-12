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
	weighted.Union(2, 8)
	weighted.Union(3, 6)
	weighted.Union(1, 3)
	weighted.Union(6, 4)
	weighted.Union(0, 5)
	weighted.Union(1, 7)
	weighted.Union(0, 2)
	weighted.Union(0, 3)
	weighted.Union(0, 9)
	fmt.Println(weighted.String())
}

func runEager() {
	eager.Init(10)
	eager.Union(0, 3)
	eager.Union(0, 6)
	eager.Union(0, 7)
	eager.Union(7, 4)
	eager.Union(0, 1)
	eager.Union(0, 9)
	fmt.Println(eager.String())
}
