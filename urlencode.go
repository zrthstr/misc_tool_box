// moderatly stric url encode
// encodes optional char such as: +
// does not encode _
//
// -a for full encode, of everything

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {
	encodeAll := flag.Bool("a", false, "Encode all characters, including those normally not required")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		output := encodeInput(input, *encodeAll)
		fmt.Println(output)
	}
}

func encodeInput(input string, encodeAll bool) string {
	if encodeAll {
		return encodeAllChars(input)
	}
	// Normal URL encoding
	return url.QueryEscape(input)
}

func encodeAllChars(input string) string {
	var encoded string
	for i := 0; i < len(input); i++ {
		encoded += "%" + fmt.Sprintf("%02X", input[i])
	}
	return encoded
}
