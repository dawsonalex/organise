package main

import (
	"fmt"

	"github.com/dawsonalex/organise/cmd/flags"
)

func main() {
	options := flags.ParseFlags()
	fmt.Printf("running in dir: %s\n", *options.Dir)
}
