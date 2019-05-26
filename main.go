package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "random names"
	app.Usage = "generate random human names"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "number, n",
			Value: 1,
			Usage: "how many names to generate",
		},
		cli.Float64Flag{
			Name:  "max, m",
			Usage: "the maximum of the random number",
		},
		cli.BoolFlag{
			Name:  "real, r",
			Usage: "use the real percentage",
		},
	}

	app.Action = func(c *cli.Context) error {
		var source string
		if c.NArg() > 0 {
			source = c.Args().Get(0)
		}

		max := c.Float64("max")
		if max > 100 {
			max = 100
		} else if max < 0 {
			max = 0
		}

		number := c.Int("number")
		if number <= 0 {
			number = 1
		}

		opt := &options{
			max:    max,
			number: number,
			real:   c.Bool("real"),
		}

		names := []string{}
		var err error
		if source != "" {
			names, err = GetRandomNames(source, opt)
		}

		if err != nil {
			log.Fatal(err)
			return err
		}

		if len(names) > 0 {
			fmt.Println(strings.Join(names, ", "))
		} else {
			fmt.Println("No database seleted")
		}
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
