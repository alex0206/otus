package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func Top10(text string) []string {
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err.Error())
	}

	return nil
}
