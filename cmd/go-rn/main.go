package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/urfave/cli.v1"
	rn "github.com/random-names/go"
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
		
		opt := &rn.Options{
			Max:    c.Float64("max"),
			Number: c.Int("number"),
			Real:   c.Bool("real"),
		}

		names := []string{}
		var err error
		if source != "" {
			names, err = rn.GetRandomNames(source, opt)
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
