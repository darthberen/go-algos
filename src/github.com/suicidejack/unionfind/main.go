package main

/************************************************************************
 * Author: Matt Price
 * Last Modified: Thu Jun 19, 2014 18:41:53 EDT
 * Purpose: Display the running of various quick-find algorithms as taken
 *          from the coursera class algorithms part 1 starting June 2014.
 ************************************************************************/

import (
	"fmt"
	eager "github.com/suicidejack/unionfind/eager"
)

func main() {
	eager.Init(10)
	fmt.Println(eager.String())
}
