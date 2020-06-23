package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"reflect"
	"testing"
)

func TestLoadQuiz(t *testing.T) {
	t.Run("Should load the quiz", func(t *testing.T) {
		problems := loadQuiz(false)
		if len(problems) == 0 {
			t.Errorf("the quiz slice returned is empty")
		}

		if !isOrderSameAsProblemsCSV(problems) {
			t.Errorf("the quiz did not return the questions in the expected order")
		}
	})

	t.Run("Should load the quiz in a random order", func(t *testing.T) {
		problems := loadQuiz(true)
		if len(problems) == 0 {
			t.Errorf("the quiz slice returned is empty")
		}

		if isOrderSameAsProblemsCSV(problems) {
			t.Errorf("the quiz did not randomise the questions")
		}
	})
}

func TestRandomiseOrder(t *testing.T) {
	sliceThatWillChange := []QA{
		QA{"q", "answer1"},
		QA{"q", "answer2"},
		QA{"q", "answer3"},
		QA{"q", "answer4"},
		QA{"q", "answer5"},
		QA{"q", "answer6"},
	}

	originalSlice := make([]QA, len(sliceThatWillChange))
	copy(originalSlice, sliceThatWillChange)

	randomiseOrder(&sliceThatWillChange)
	if len(originalSlice) != len(sliceThatWillChange) {
		t.Errorf("the quiz slice returned is not the same length")
	}

	if reflect.DeepEqual(originalSlice, sliceThatWillChange) {
		t.Errorf("the quiz slice returned is the same as the original")
	}
}

// Helper functions

func isOrderSameAsProblemsCSV(quiz []QA) bool {
	file, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(file))
	allRecords, _ := reader.ReadAll()

	for i, record := range allRecords {
		qa := QA{record[0], record[1]}
		if quiz[i] != qa {
			return false
		}
	}
	return true
}
