package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	records := readCsvFile("survey.csv")

	var avgMentorSatisf, avgLearningSatisf float64
	courseGroup := map[string]map[string]int{}
	for i := 1; i < len(records); i++ {
		mentorSatifCell, _ := strconv.ParseFloat(records[i][3], 64)
		avgMentorSatisf += mentorSatifCell / float64(len(records)-1)

		learningSatisfCell, _ := strconv.ParseFloat(records[i][4], 64)
		avgLearningSatisf += learningSatisfCell / float64(len(records)-1)

		_, ok := courseGroup[records[i][2]]
		if ok {
			courseGroup[records[i][2]]["sum"] += int(mentorSatifCell)
			courseGroup[records[i][2]]["count"]++
		} else {
			courseGroup[records[i][2]] = make(map[string]int)
			courseGroup[records[i][2]]["count"] = 1
		}
	}

	fmt.Printf("Overall average mentor satisfaction = %.2f\n", avgMentorSatisf)
	fmt.Printf("Overall average learning satisfaction = %.2f\n", avgLearningSatisf)
	fmt.Printf("Highest average mentor satisfcation = %s", findHighestAvg(courseGroup))
}

func findHighestAvg(data map[string]map[string]int) (result string) {
	max := 0.0
	for k, v := range data {
		avg := float64(v["sum"]) / float64(v["count"])
		if avg > max {
			max = avg
			result = k
		}
	}

	return result
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
