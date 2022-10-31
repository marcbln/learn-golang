package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, '.', tabwriter.Debug)

	fmt.Fprintln(w, "some\trow\tA")
	fmt.Fprintln(w, "some\trow\tB")
	fmt.Fprintln(w, "some\trow\tC")
	fmt.Fprintln(w, "some\trow\tD")

	w.Flush()
}
