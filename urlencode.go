// moderatly stric url encode
// encodes optional char such as: +
// does not encode _

package main

import (
	"bufio"
	"net/url"
	"os"
)

func main() {
	err := urlencode(os.Stdin, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
}

func urlencode(input *os.File, output *os.File) error {
	scanner := bufio.NewScanner(input)
	firstLine := true
	for scanner.Scan() {
		encoded := url.QueryEscape(scanner.Text())
		if firstLine {
			firstLine = false
		} else {
			if _, err := output.WriteString("%0A"); err != nil {
				return err
			}
		}
		if _, err := output.WriteString(encoded); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
