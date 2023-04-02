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

	encodedInput := url.QueryEscape(input)

	fmt.Println(encodedInput)
}
