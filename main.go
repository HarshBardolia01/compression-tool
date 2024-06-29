package main

import (
	"github.com/HarshBardolia01/compression-tool/args"
	"github.com/HarshBardolia01/compression-tool/compressor"
	"github.com/HarshBardolia01/compression-tool/decompressor"
)

func main() {
	arg := args.ReadFlags()

	if arg.ToCompress {
		compressor.Compress(arg.InputFilePath, arg.OutputFilePath)
	} else {
		decompressor.Decompress(arg.InputFilePath, arg.OutputFilePath)
	}
}
