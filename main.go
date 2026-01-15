package main

import (
	"context"
	fileparser "file-parser"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var firstFile, secondFile string
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"f"},
				Usage:       "output format",
				DefaultText: "stylish",
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
			return fileparser.ParseFiles(firstFile, secondFile)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
