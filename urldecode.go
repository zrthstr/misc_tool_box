package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	decodedInput, err := url.QueryUnescape(input)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(decodedInput)
}
