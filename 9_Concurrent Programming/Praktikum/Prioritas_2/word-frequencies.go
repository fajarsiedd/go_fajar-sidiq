package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	frequenciesCh := make(chan map[string]int)
	wordCh := make(chan string)
	done := make(chan bool)

	words := extractTextFile("document.txt")

	go func() {
		for _, word := range words {
			wordCh <- word
		}
		close(wordCh)
	}()

	go calculateFrequency(wordCh, frequenciesCh)

	go generateCSV(frequenciesCh, done)

	<-done
}

func generateCSV(frequenciesCh chan map[string]int, done chan bool) {
	wordsMap := <-frequenciesCh
	records := [][]string{{"word", "frequencies"}}

	for k, v := range wordsMap {
		records = append(records, []string{k, strconv.Itoa(v)})
	}

	file, err := os.Create("word-frequencies.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	done <- true
}

func calculateFrequency(wordCh chan string, frequenciesCh chan map[string]int) {
	wordsMap := make(map[string]int)

	for word := range wordCh {
		wordsMap[word]++
	}

	frequenciesCh <- wordsMap
}

func extractTextFile(filename string) (result []string) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splittedText := strings.FieldsFunc(line, func(r rune) bool {
			return r == ' ' || r == ',' || r == '.'
		})
		result = append(result, splittedText...)
	}

	return result
}
