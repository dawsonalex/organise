package flags

import (
	"flag"
	"os"
)

// FlagOptions contains options for controlling
// the application based on command line arguments.
type FlagOptions struct {
	Dir *string
}

// ParseFlags parses command line options and returns a flagOptions struct pointer
// to read options from.
func ParseFlags() *FlagOptions {
	dir := flag.String("dir", defaultDir(), "The directory to organise files in.")

	flag.Parse()

	return &FlagOptions{
		dir,
	}
}

func defaultDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		return os.TempDir()
	}
	return dir
}
