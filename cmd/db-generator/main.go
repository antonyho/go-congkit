package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/antonyho/go-cangjie/internal/data"
	"github.com/antonyho/go-cangjie/internal/db"
)

var (
	source string
	target string
)

const (
	DefaultTarget = "cangjie.db"

	SourceFileUsage = "Cangjie source table file path"
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
		fmt.Println("Using built-in Cangjie radicals table")

		sourceTable, err = data.ReadBuiltinTable()
		if err != nil {
			log.Fatalf("Failed reading data from built-in Cangjie table.\n%v\n", err)
		}
	} else {
		fmt.Println("Using provided Cangjie radicals table ", source)

		sourceTableFile, err := os.Open(source)
		if err != nil {
			log.Fatalf("Failed opening Cangjie table file %s.\n%v\n", source, err)
		}
		sourceTable, err = data.ReadTable(sourceTableFile)
		if err != nil {
			log.Fatalf("Failed reading data from Cangjie table.\n%v\n", err)
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
		log.Fatalf("Failed generating Cangjie database file.\n%v\n", err)
	}
}

func helpFunc(_ string) error {
	flag.Usage()
	os.Exit(0)

	return nil
}
