package pkg

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

var expectedResults = map[int]int {
	12: 2,
	14: 2,
	1969: 654,
	100756: 33583,
}

func TestDay1(t *testing.T) {
	for mass,expected := range expectedResults {
		result := CalculateFuel(mass)

		if result != expected {
			t.Errorf("expected CalculateFuel(%v) to yield %v, got %v", mass, expected, result)
		}
	}
}

func TestDay1Sum(t *testing.T) {
	file, err := os.Open("../input_files/day_1.txt")
	defer file.Close()
	if err != nil {
		t.Error(err)
	}

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += CalculateFuel(mass)
	}

	t.Log(total)
}
