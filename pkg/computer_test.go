package pkg

import (
	"fmt"
	"testing"
)

type ResultSet struct {
	Program   []int
	Input int
	Output int
	Assertion func(c *Computer, t *testing.T)
}

func AssertOutput(value int) func(c *Computer, t *testing.T) {
    return func(c *Computer, t *testing.T) {
        if c.OutputRegister != value {
            t.Errorf("output register should contain %v, found %v", value, c.OutputRegister)
        }
    }
}

func AssertMemoryValue(index, value int) func(c *Computer, t *testing.T) {
	return func(c *Computer, t *testing.T) {
		if c.Memory[index] != value {
			t.Errorf("memory at %v should be %v, found %v", index, value, c.Memory[index])
		}
	}
}

func Aggregate(assertions ...func(c *Computer, t *testing.T)) func(c *Computer, t *testing.T) {
	return func(c *Computer, t *testing.T) {
		for _, a := range assertions {
			a(c, t)
		}
	}
}

func TestComputer(t *testing.T) {
	var expectedResults = []ResultSet{
		{
			Program:   []int{1, 0, 0, 0, 99},
			Assertion: AssertMemoryValue(0, 2),
		},
		{
			Program:   []int{2, 3, 0, 3, 99},
			Assertion: AssertMemoryValue(3, 6),
		},
		{
			Program:   []int{2, 4, 4, 5, 99, 0},
			Assertion: AssertMemoryValue(5, 9801),
		},
		{
			Program: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			Assertion: Aggregate(
				AssertMemoryValue(0, 30),
				AssertMemoryValue(5, 5),
			),
		},
        {
            Program: []int {3, 0, 4, 0, 99},
            Input: 123,
            Output: 444,
            Assertion: AssertOutput(123),
        },
        {
            Program: []int {1003, 0, 1004, 0, 99},
            Input: 123,
            Output: 444,
            Assertion: AssertOutput(123),
        },
        {
            Program: []int {1002,4,3,4,33},
            Assertion: AssertMemoryValue(4, 99),
        },
	}

	for i, v := range expectedResults {
		t.Run(fmt.Sprintf("test-%v", i), func(t *testing.T) {
			computer := NewComputer(v.Program)
			computer.InputRegister = v.Input
			computer.OutputRegister = v.Output
			computer.Run()
			v.Assertion(computer, t)
		})
	}
	{
		computer, err := ComputerFromFile("../input_files/day_2.txt")
		if err != nil {
			t.Error(err)
		} else {
			computer.Memory[1] = 12
			computer.Memory[2] = 2
			computer.Run()
			t.Logf("Address #0 contains %v", computer.Memory[0])
		}
	}
	{
		computer, err := ComputerFromFile("../input_files/day_5.txt")
		if err != nil {
			t.Error(err)
		} else {
			computer.InputRegister = 1
			computer.Run()
			t.Logf("Diagnostics ran successfully! Output: %v", computer.OutputRegister)
		}
	}

}

func TestNounAndVerb(t *testing.T) {
	expectedValue := 19690720
	for k := 0; k < 100; k++ {
		for v := 0; v < 100; v++ {
			computer, err := ComputerFromFile("../input_files/day_2.txt")
			if err != nil {
				t.Error(err)
			} else {
				computer.Memory[1] = k
				computer.Memory[2] = v
				computer.Run()
				if computer.Memory[0] == expectedValue {
					t.Logf("Noun: %v, Verb: %v, result: %v", computer.Memory[1], computer.Memory[2], 100*computer.Memory[1]+computer.Memory[2])
				}
			}
		}
	}
}
