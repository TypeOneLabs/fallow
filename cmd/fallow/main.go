// Command fallow finds and prices cloud waste.
package main

import (
	"fmt"
	"os"

	"github.com/TypeOneLabs/fallow/internal/cli"
)

func main() {
	code, err := cli.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "fallow: "+err.Error())
	}
	os.Exit(code)
}
