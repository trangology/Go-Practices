package main

import (
	"math"
	"strings"
)

const alpha = float64(1)

type Classifier struct {
	labels      string
	frequencies map[string]int
}

var data []Classifier

// labelWord is list of amount different words according labels
var labelWord []int
var uniqueLabels []string

// allWord is total different words trained
var allWords = 0

func fitData(X []string, y []string) {
	// finding unique labels from given list labels
	for _, label := range y {
		exist := false
		for _, curLabel := range uniqueLabels {
			if curLabel == label {
				exist = true
			}
		}
		if !exist {
			uniqueLabels = append(uniqueLabels, label)
		}
	}

	// for _, value := range data {
	// frequencies is map which included word with frequency

	for _, label := range uniqueLabels {
		frequency := make(map[string]int)
		for i, message := range X {
			if y[i] == label {
				wordList := strings.Fields(message)
				for _, word := range wordList {
					_, exist := frequency[word]
					if exist {
						frequency[word]++
					} else {
						frequency[word] = 1
					}
				}
			}
		}
		labelWord = append(labelWord, len(frequency))
		allWords += len(frequency)
		var c Classifier
		c.labels = label
		c.frequencies = frequency
		data = append(data, c)
	}
}

func findLabel(arr []float64) (result string) {
	maxValue := arr[0]
	result = data[0].labels
	for i, value := range arr {
		if value > maxValue {
			result = data[i].labels
		}
	}
	return
}

func predict(x string) (result string) {

	// probaTitle is array of probabilities with each label
	var probaTitle []float64

	words := strings.Fields(x)

	for i := 0; i < len(data); i++ {
		proba := float64(0)
		for _, word := range words {
			frequency, ok := data[i].frequencies[word]
			if frequency != 0 {
				proba += math.Log((float64(frequency) + alpha) / float64(labelWord[i]+allWords))
			} else if !ok {
				proba += math.Log(alpha / float64(labelWord[i]+allWords))
			}
		}
		proba += math.Log(float64(labelWord[i]) / float64(allWords))
		probaTitle = append(probaTitle, proba)
	}
	result = findLabel(probaTitle)

	return
}
