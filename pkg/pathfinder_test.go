package pkg

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

type PathfinderResultSet struct {
	Input [][]string
	Output int
}
func TestPathfinder(t *testing.T) {
	expectedResults := []PathfinderResultSet{
		{
			Input: [][]string {
				{"R8","U5","L5","D3"},
				{"U7","R6","D4","L4"},
			},
			Output: 6,
		},
		{
			Input: [][]string {
				{"R75","D30","R83","U83","L12","D49","R71","U7","L72"},
				{"U62","R66","U55","R34","D71","R55","D58","R83"},
			},
			Output: 159,
		},
		{
			Input: [][]string {
				{"R98","U47","R26","D63","R33","U87","L62","D20","R33","U53","R51"},
				{"U98","R91","D20","R16","D67","R40","U7","R15","U6","R7"},
			},
			Output: 135,
		},
	}

	for _, expected := range expectedResults {
		p1 := Pathfinder{}
		p2 := Pathfinder{}
		for _, v := range expected.Input[0] {
			err := p1.Move(v)
			if err != nil {
				t.Error(err)
			}
		}
		for _, v := range expected.Input[1] {
			err := p2.Move(v)
			if err != nil {
				t.Error(err)
			}
		}
		intersection := p1.Path.Intersect(&p2.Path)

		shortestDistance := 1000000
		for _,i := range intersection {
			distance := i.ManhattanDistance(&Coordinate {0,0})
			if distance < shortestDistance { shortestDistance = distance }
		}

		if shortestDistance != expected.Output {
			t.Errorf("Expected %v, got %v", expected.Output, shortestDistance)
		}
	}
}

func TestPathfinderFromFile(t *testing.T) {
	file, err  := os.Open("../input_files/day_3.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	pf1 := Pathfinder{}
	instructions := strings.Split(scanner.Text(), ",")
	if err := pf1.Move(instructions...); err != nil {
		t.Error(err)
	}

	scanner.Scan()
	pf2 := Pathfinder{}
	instructions = strings.Split(scanner.Text(), ",")
	if err := pf2.Move(instructions...); err != nil {
		t.Error(err)
	}

	intersection := pf1.Path.Intersect(&pf2.Path)

	shortestDistance := 1000000
	for _,i := range intersection {
		distance := i.ManhattanDistance(&Coordinate {0,0})
		if distance < shortestDistance { shortestDistance = distance }
	}

	t.Logf("Shortest distance is %v", shortestDistance)
}