package main

import (
	"os"

	gw "github.com/a-b/grepwrite"
)

// func init() {
// 	if cpu := runtime.NumCPU(); cpu == 1 {
// 		runtime.GOMAXPROCS(2)
// 	} else {
// 		runtime.GOMAXPROCS(cpu)
// 	}
// }

func main() {
	gw := gw.GrepWrite{In: os.Stdin, Out: os.Stdout, Err: os.Stderr}

	exitCode := gw.Run{os.Args[1:]}
	os.Exit(exitCode)
}
