package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var input string

	fmt.Fscan(r, &input)
	fmt.Fprintf(w, "Hello, %s!", input)
}
