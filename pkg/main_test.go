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
	1969: 966,
	100756: 50346,
}

func TestDay1(t *testing.T) {
	for mass,expected := range expectedResults {
		result := CalculateTotalFuel(mass)

		if result != expected {
			t.Errorf("expected CalculateTotalFuel(%v) to yield %v, got %v", mass, expected, result)
		}
	}
}

func TestDay1Sum(t *testing.T) {
	file, err := os.Open("../input_files/day_1.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()


	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += CalculateTotalFuel(mass)
	}

	t.Log(total)
}
