package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/antonyho/go-congkit/internal/data"
	"github.com/antonyho/go-congkit/internal/db"
)

var (
	source string
	target string
)

const (
	DefaultTarget = "congkit.db"

	SourceFileUsage = "Congkit source table file path"
	TargetFileUsage = "Target database file path"
	HelpUsage       = "Print usages"
)

func init() {
	flag.StringVar(&source, "source", "", SourceFileUsage)
	flag.StringVar(&source, "s", "", SourceFileUsage)

	flag.StringVar(&target, "target", "", TargetFileUsage)
	flag.StringVar(&target, "t", "", TargetFileUsage)

	flag.BoolFunc("help", HelpUsage, helpFunc)
	flag.BoolFunc("h", HelpUsage, helpFunc)
}

func main() {
	flag.Parse()

	var err error

	var sourceTable [][]string
	if source == "" {
		fmt.Println("Using built-in Congkit radicals table")

		sourceTable, err = data.ReadBuiltinTable()
		if err != nil {
			log.Fatalf("Failed reading data from built-in Congkit table.\n%v\n", err)
		}
	} else {
		fmt.Println("Using provided Congkit radicals table ", source)

		sourceTableFile, err := os.Open(source)
		if err != nil {
			log.Fatalf("Failed opening Congkit table file %s.\n%v\n", source, err)
		}
		sourceTable, err = data.ReadTable(sourceTableFile)
		if err != nil {
			log.Fatalf("Failed reading data from Congkit table.\n%v\n", err)
		}
	}

	if target == "" {
		fmt.Println("Using default target database file path")

		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalln("Cannot get current working directory")
		}
		target = path.Join(cwd, DefaultTarget)
	}
	fmt.Printf("Target database file path: %s\n", target)

	if err = db.Generate(sourceTable, target); err != nil {
		log.Fatalf("Failed generating Congkit database file.\n%v\n", err)
	}
}

func helpFunc(_ string) error {
	flag.Usage()
	os.Exit(0)

	return nil
}
