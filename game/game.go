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
	"strconv"
	"strings"
	"time"
)

//word structure
type Word struct {
	kata        string
	comment     string
	translation string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var correctStreak = 0

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

	for {
		rand.Seed(time.Now().UTC().UnixNano())
		var random = rand.Intn(rowCount - 1)
		fmt.Print(words[random].kata)
		fmt.Print("\nPlease provide meaning\n")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")

		if strings.EqualFold(words[random].translation, input) || strings.Contains(words[random].translation, input) {
			correctStreak = correctStreak + 1
			fmt.Print("\ncorrect!\n")
			fmt.Print("\nStreak🔥:\n" + strconv.Itoa(correctStreak) + "\n")

		} else {
			correctStreak = 0
			fmt.Print("\nnope 💩!\n")
			fmt.Print("\n" + words[random].translation + "\n")
		}
	}
}
