package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Word struct {
	kanji       string
	kata        string
	translation string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var url = "https://docs.google.com/spreadsheets/d/e/2PACX-1vRW71MEJqbKIsCL1EulKsV90D1CfVAk1f4xK6DH8occ6ZjQrbpVoV4ZVTfH91fiy4rk4SzRWbQrOjJb/pub?output=csv"
	var filename = "translations.csv"
	var timeout = int64(100)
	client := http.Client{
		Timeout: time.Duration(timeout * int64(time.Second)),
	}
	resp, _ := client.Get(url)
	b, _ := ioutil.ReadAll(resp.Body)
	ioutil.WriteFile(filename, b, 0644)

	csvFile, _ := os.Open(filename)
	csvFileContent := csv.NewReader(csvFile)
	var rowCount = 0

	var words = []Word{}

	for {
		record, err := csvFileContent.Read()
		if err == io.EOF {
			break
		}

		word := Word{record[0], record[1], record[2]}
		words = append(words, word)
		rowCount = rowCount + 1
	}

	rand.Seed(time.Now().UTC().UnixNano())
	var random = rand.Intn(rowCount - 1)

	fmt.Print(words[random].kanji)
	fmt.Print("\nPlease provide meaning 🔥 🔥 🔥\n")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	if words[random].translation == text {
		fmt.Print("\ncorrect!\n")
	} else {
		fmt.Print("\nnope!\n")
	}
}
