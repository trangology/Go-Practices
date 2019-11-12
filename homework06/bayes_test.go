package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"testing"

	"github.com/jbrukh/bayesian"
)

const (
	ham  bayesian.Class = "Ham"
	spam bayesian.Class = "Spam"

	dataPath = "data/SMSSpamCollection"
	testPath = "data/SMSTest"
)

var hamMessages []string
var spamMessages []string
var allMessages []string
var allLabels []string

// TODO: check error when parsing inputFile
func parseData() {
	inputFile, _ := os.Open(dataPath)
	reader := csv.NewReader(bufio.NewReader(inputFile))
	reader.Comma = '\t'
	reader.LazyQuotes = true

	messages, err := reader.ReadAll()
	check(err)

	for _, message := range messages {
		if message[0] == "ham" {
			hamMessages = append(hamMessages, message[1])
		} else {
			spamMessages = append(spamMessages, message[1])
		}
		allMessages = append(allMessages, message[1])
		allLabels = append(allLabels, message[0])
	}
}

func readFile(path string) []string {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func classifyMessages(sms []string) []string {
	classifier := bayesian.NewClassifier(ham, spam)
	classifier.Learn(hamMessages, ham)
	classifier.Learn(spamMessages, spam)

	var result []string

	for _, m := range sms {
		proba, _, _ := classifier.ProbScores([]string{m})
		if proba[0] > proba[1] {
			result = append(result, "ham")
		} else {
			result = append(result, "spam")
		}
	}

	return result
}

func Test_Bayes(T *testing.T) {
	parseData()
	sms := readFile(testPath)
	result1 := classifyMessages(sms)

	var result2 []string
	fitData(allMessages, allLabels)
	for _, m := range sms {
		label := predict(m)
		result2 = append(result2, label)
	}

	// if reflect.DeepEqual(result1, result2) == false {
	// 	T.Fatalf("Your algorithm is wrong")
	// }

	for i := 0; i < len(result2); i++ {
		if result2[i] != result1[i] {
			T.Errorf("Your prediction was incorrect with the message %d", i+1)
		}
	}
}
