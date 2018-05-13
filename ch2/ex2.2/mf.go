package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/takatori/gopl/ch2/ex2.2/distconv"
)

func main() {

	m := flag.Float64("m", 0, "meter")
	ft := flag.Float64("ft", 0, "feet")
	flag.Parse()

	if *m != 0 {
		meter := distconv.Meter(*m)
		fmt.Printf("%s = %s\n", meter, distconv.MToF(meter))
        return
	}

	if *ft != 0 {
		feet := distconv.Feet(*ft)
		fmt.Printf("%s = %s\n", feet, distconv.FToM(feet))
        return
	}

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		meter := distconv.Meter(t)
		feet := distconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s\n", meter, distconv.MToF(meter), feet, distconv.FToM(feet))
	}
}
