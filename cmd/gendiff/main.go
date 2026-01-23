package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var firstFile, secondFile, format string
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Usage:       "output format",
				DefaultText: "stylish",
				Value: "stylish",
				Destination: &format,
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "first",
				Destination: &firstFile,
			},
			&cli.StringArg{
				Name:        "second",
				Destination: &secondFile,
			},
		},
		ArgsUsage: "<first_file> <second_file>",
		Action: func(ctx context.Context, c *cli.Command) error {

			if firstFile == "" || secondFile == "" {
				firstFile, secondFile = "../../testdata/nested_json1.json", "../../testdata/nested_json2.json"

				//				return errors.New("must be two files")
			}
			res, err := code.GenDiff(firstFile, secondFile, format)
			if err != nil {
				return err
			}
			fmt.Println(res)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
