package errkit

import (
	"fmt"
	"os"
)

func FinalErr(msg string) {
	fmt.Fprintln(os.Stderr, "Error: "+msg)
	os.Exit(1)
}
