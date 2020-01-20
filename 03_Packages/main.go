package main

// To import multiple packages, wrap them in (-grouping-)
// and do not parse with commas
// packages disappear if you don't use them (on VSCode)

import (
	"fmt"
	"math"

	"github.com/CZAR-1/goCzar/03_Packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7), ", ", math.Ceil(2.7), ", ", math.Sqrt(16), ", ", strutil.Reverse("hallo"))
}
