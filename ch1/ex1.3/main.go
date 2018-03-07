package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func slowEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func fastEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	slowEcho()
	secs1 := time.Since(start).Seconds()

	start = time.Now()
	fastEcho()
	secs2 := time.Since(start).Seconds()

	fmt.Printf("slow %f\n", secs1)
	fmt.Printf("fast %f\n", secs2)

}
