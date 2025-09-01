package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetScore(t *testing.T) {

	type testCase struct {
		// value inputs
		name   string
		stamp  []ScoreStamp
		offset int

		//value outputs
		expected Score
	}

	t.Run("Testing major cases", func(t *testing.T) {
		tests := []testCase{
			{name: "Empty slice", stamp: []ScoreStamp{}, offset: 10, expected: Score{0, 0}},
			{name: "Offset lower than first", stamp: []ScoreStamp{
				{Offset: 5, Score: Score{1, 0}},
			}, offset: 3, expected: Score{1, 0}},
			{name: "Offset higher than last", stamp: []ScoreStamp{
				{Offset: 5, Score: Score{1, 0}},
				{Offset: 10, Score: Score{2, 0}},
			}, offset: 15, expected: Score{2, 0}},
			{name: "Offset exact", stamp: []ScoreStamp{
				{Offset: 5, Score: Score{1, 0}},
			}, offset: 5, expected: Score{1, 0}},
			{name: "Offset between", stamp: []ScoreStamp{
				{Offset: 5, Score: Score{1, 0}},
				{Offset: 10, Score: Score{2, 0}},
			}, offset: 7, expected: Score{1, 0}},
		}

		for _, test := range tests {
			actual, err := getScore(test.stamp, test.offset)
			assert
		}
	})

}
