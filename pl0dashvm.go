package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ryo33/pl0dashvm/vm"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "pl0dashrun"
	app.Usage = ""
	app.Action = action
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "trace, t",
			Usage: "print detailed reporting",
		},
	}
	app.Run(os.Args)
}

func action(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		log.Fatal("filename required")
	}
	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	option := vm.NewOption()
	if c.Bool("trace") {
		option.Trace()
	}
	result, err := vm.Run(lines, option)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
