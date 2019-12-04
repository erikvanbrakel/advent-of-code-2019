package pkg

import "testing"

type ResultSet struct {
	Input  []int
	Output []int
}

func TestComputer(t *testing.T) {
	var expectedResults = []ResultSet{
		{
			Input:  []int{1, 0, 0, 0, 99},
			Output: []int{2, 0, 0, 0, 99},
		},
		{
			Input:  []int{2, 3, 0, 3, 99},
			Output: []int{2, 3, 0, 6, 99},
		},
		{
			Input:  []int{2, 4, 4, 5, 99, 0},
			Output: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			Input:  []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			Output: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, v := range expectedResults {
		computer := NewComputer(v.Input)
		computer.Run()
		for i, a := range computer.Memory {
			if v.Output[i] != a {
				t.Error("memory doesn't match")
			}
		}
	}

	computer, err := ComputerFromFile("input_files/day_2.txt")
	if err != nil {
		t.Error(err)
	} else {
		computer.Memory[1] = 12
		computer.Memory[2] = 2
		computer.Run()
		t.Logf("Address #0 contains %v", computer.Memory[0])
	}
}

func TestNounAndVerb(t *testing.T) {
	expectedValue := 19690720
	for k := 0; k < 100; k++ {
		for v := 0; v < 100; v++ {
			computer, err := ComputerFromFile("input_files/day_2.txt")
			if err != nil {
				t.Error(err)
			} else {
				computer.Memory[1] = k
				computer.Memory[2] = v
				computer.Run()
				if computer.Memory[0] == expectedValue {
					t.Logf("Noun: %v, Verb: %v, result: %v", computer.Memory[1], computer.Memory[2], 100 * computer.Memory[1] + computer.Memory[2])
				}
			}
		}
	}
}
