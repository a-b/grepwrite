package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

func init() {
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}
}

func main() {
	var dryRun bool

	app := cli.NewApp()
	app.Name = "gw"
	app.Usage = "read error format and write it back into corresponding files"
	// app.Action = func(c *cli.Context) error {
	// 	fmt.Println("boom! I say!")
	// 	return nil
	// }

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:   "dry-run, d",
			Usage:  "run script without writing to the files",
			EnvVar: "DRYRUN",

			Destination: &dryRun,
		},

		cli.StringFlag{
			Name:     "file, f",
			Usage:    "input file",
			FilePath: "error_file_input",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if !ctx.Bool("dry-run") {
			fmt.Println("Hola")
			return cli.NewExitError("it is not in the soup", 86)
		}
		return nil
	}
	// app.Action = func(c *cli.Context) error {
	// 	name := "someone"
	// 	if c.NArg() > 0 {
	// 		name = c.Args()[0]
	// 	}
	// 	if dryRun {
	// 		fmt.Println("Hola", name)
	// 	} else {
	// 		fmt.Println("Hello", name)
	// 	}
	// 	return nil
	// }

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
	// gw := gw.GrepWrite{In: os.Stdin, Out: os.Stdout, Err: os.Stderr}
	// exitCode := gw.Run(os.Args[1:])
	// os.Exit(exitCode)
}
