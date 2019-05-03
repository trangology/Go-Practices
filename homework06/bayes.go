package main

import (
	"math"
	"strings"
)

const alpha  = float64(1)

type Classifier struct {
	labels string
	frequencies map[string]int
}


var data []Classifier

// amountWord is list of amount different words according labels
var	amountWord []int

// allWord is total different words trained
var allWords = 0


func NaiveBayesClassifier(X []string, y []string)  {
	// finding unique labels from given list labels
	for _, label := range y{
		for _, value := range data{
			if value.labels != label{
				var c Classifier
				c.labels = label
				data = append(data, c)
			}
		}
	}


	for _, value := range data{
		// frequencies is map which included word with frequency
		frequency := make(map[string]int)
		for i, x := range X{
			if y[i] == value.labels{
				wordList := strings.Fields(x)
				for _, word := range wordList{
					_, ok := frequency[word]
					if ok {
						frequency[word] += 1
					} else {
						frequency[word] = 1
					}
				}
			}
		}
		amountWord = append(amountWord, len(frequency))
		allWords += len(frequency)
		value.frequencies = frequency
	}
	// end fit data
}

func findLabel(arr []float64) (result string) {
	maxValue := float64(0)
	for i, value := range arr{
		if value > maxValue{
			result = data[i].labels
		}
	}
	return
}


func predict(x string) (result string) {

	// probaTitle is array of probabilities with each label
	var probaTitle []float64

	words := strings.Fields(x)

	for i := 0; i < len(data); i++{
		proba := float64(0)
		for _, word := range words{
			frequency := data[i].frequencies[word]
			if frequency != 0{
				proba += math.Log((float64(frequency) + alpha) / float64(amountWord[i] + allWords))
			} else {
				proba += math.Log(alpha / float64(amountWord[i] + allWords))
			}
		}
		proba += math.Log(float64(amountWord[i] / allWords))
		probaTitle = append(probaTitle, proba)
	}
	result = findLabel(probaTitle)

	return
}