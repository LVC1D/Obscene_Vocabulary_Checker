package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// write your code here
	var filename, phraseToCheck string
	fmt.Scan(&filename)

	taboo := make(map[string]struct{})

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		taboo[scanner.Text()] = struct{}{}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	reader := bufio.NewScanner(os.Stdin)

	for reader.Scan() {
		phraseToCheck = reader.Text()
		if phraseToCheck == "exit" {
			fmt.Println("Bye!")
			break
		}

		phraseSlice := strings.Split(phraseToCheck, " ")

		for i, wordToCheck := range phraseSlice {
			_, ok := taboo[strings.ToLower(wordToCheck)]
			if ok {
				phraseSlice[i] = strings.Repeat("*", len(wordToCheck))
			} else {
				phraseSlice[i] = wordToCheck
			}
		}

		fmt.Println(strings.Join(phraseSlice, " "))
	}
}
