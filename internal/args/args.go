package args

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	InputFilePath  string
	OutputFilePath string
	ToCompress     bool
}

func ReadFlags() *Args {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	var args Args
	flag.StringVar(&args.InputFilePath, "i", "", "Path to the Input File")
	flag.StringVar(&args.OutputFilePath, "o", "", "Path to the Output File")
	flag.BoolVar(&args.ToCompress, "c", true, "true -> Compress | false -> Decompress")

	flag.Parse()
	validateInput(args)

	return &args
}

func validateInput(args Args) {
	_, err := os.Open(args.InputFilePath)

	if err != nil {
		fmt.Printf("Error opening the Input File: %s\n", err.Error())
		os.Exit(1)
	}

	_, err = os.Open(args.OutputFilePath)

	if err != nil {
		fmt.Printf("Error opening the Output File: %s\n", err.Error())
		os.Exit(1)
	}
}
