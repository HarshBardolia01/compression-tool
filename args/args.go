package args

import "flag"

type Args struct {
	InputFilePath  string
	OutputFilePath string
	ToCompress     bool
}

func ReadFlags() *Args {
	var args Args
	flag.StringVar(&args.InputFilePath, "i", "", "Input File")
	flag.StringVar(&args.OutputFilePath, "o", "", "Output File")
	flag.BoolVar(&args.ToCompress, "c", true, "Compress")
	return &args
}
