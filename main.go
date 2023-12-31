package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/antonyho/go-congkit/engine"
)

var (
	version    int
	simplified bool
	easy       bool
	prediction bool
	db         string
)

const (
	DefaultDB             = "congkit.db"
	DefaultCongkitVersion = engine.CongkitV5
)

const (
	HelpUsage        = "Print usages"
	VersionUsage     = "Congkit version(3/5)"
	SimplifiedUsage  = "Output simplified Chinese word"
	EasyIMUsage      = "Use 'Easy' input method"
	PredicationUsage = "Predict the possible typing word"
	DBUsage          = "Custom database file path"
)

func init() {
	flag.IntVar(&version, "version", int(DefaultCongkitVersion), VersionUsage)
	flag.IntVar(&version, "v", int(DefaultCongkitVersion), VersionUsage)

	flag.BoolVar(&simplified, "simplified", false, SimplifiedUsage)
	flag.BoolVar(&simplified, "s", false, SimplifiedUsage)

	flag.BoolVar(&easy, "easy", false, EasyIMUsage)
	flag.BoolVar(&easy, "e", false, EasyIMUsage)

	flag.BoolVar(&prediction, "prediction", false, PredicationUsage)
	flag.BoolVar(&prediction, "p", false, PredicationUsage)

	flag.StringVar(&db, "database", DefaultDB, DBUsage)
	flag.StringVar(&db, "d", DefaultDB, DBUsage)

	flag.BoolFunc("help", HelpUsage, helpFunc)
	flag.BoolFunc("h", HelpUsage, helpFunc)
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Printf("%s [congkit_radicals]\n\n", os.Args[0])
		flag.Usage()
		os.Exit(0)
	}

	options := []engine.Option{engine.WithDatabase(db)}

	switch engine.CongkitVersion(version) {
	case engine.CongkitV3:
		options = append(options, engine.WithCongkitV3())
	case engine.CongkitV5:
		options = append(options, engine.WithCongkitV5())

	default:
		options = append(options, engine.WithCongkitV5())
	}

	if simplified {
		options = append(options, engine.WithSimplified())
	}

	if easy {
		options = append(options, engine.WithEasy())
	}

	if prediction {
		options = append(options, engine.WithPrediction())
	}

	eng := engine.New(options...)
	defer eng.Close()
	result, err := eng.Encode(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resultStrings := make([]string, len(result))
	for i, r := range result {
		resultStrings[i] = string(r)
	}

	fmt.Println(resultStrings)
}

func helpFunc(_ string) error {
	flag.Usage()
	os.Exit(0)

	return nil
}
