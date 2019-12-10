package pkg

import (
    "errors"
    "fmt"
	"strings"
	"testing"
)

type ResultSet struct {
    Name      string
    Program   []int
    Input     int
    Output    int
    Assertion func(c *Computer, t *testing.T)
}

func AssertOutput(value int) func(c *Computer, t *testing.T) {
    return func(c *Computer, t *testing.T) {
        if c.OutputRegister != value {
            t.Errorf("output register should contain %v, found %v", value, c.OutputRegister)
            t.Log(strings.Join(c.Debugger, "\n"))
        }
    }
}

func AssertMemoryValue(index, value int) func(c *Computer, t *testing.T) {
    return func(c *Computer, t *testing.T) {
        if c.Memory[index] != value {
            t.Errorf("memory at %v should be %v, found %v", index, value, c.Memory[index])
            t.Log(strings.Join(c.Debugger, "\n"))
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
            Name: "add-position-mode",
            Program:   []int{1, 0, 0, 0, 99},
            Assertion: AssertMemoryValue(0, 2),
        },
        {
            Name: "add-direct-mode",
            Program:   []int{1101, 1, 2, 0, 99},
            Assertion: AssertMemoryValue(0, 3),
        },
        {
            Name : "multiply-position-mode",
            Program:   []int{2, 3, 0, 3, 99},
            Assertion: AssertMemoryValue(3, 6),
        },
        {
            Name : "multiply-direct-mode",
            Program:   []int{1102, 3, 3, 3, 99},
            Assertion: AssertMemoryValue(3, 9),
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
            Program:   []int{3, 0, 4, 0, 99},
            Input:     123,
            Output:    444,
            Assertion: AssertOutput(123),
        },
        {
            Program:   []int{1003, 0, 1004, 0, 99},
            Input:     123,
            Output:    444,
            Assertion: AssertOutput(123),
        },
        {
            Program:   []int{1002, 4, 3, 4, 33},
            Assertion: AssertMemoryValue(4, 99),
        },
        {
            Name:      "position-mode-equal-to-8",
            Program:   []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
            Input:     8,
            Assertion: AssertOutput(1),
        },
        {
            Name:      "position-mode-not-equal-to-8",
            Program:   []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
            Input:     7,
            Assertion: AssertOutput(0),
        },
        {
            Name:      "immediate-mode-not-less-than-8",
            Program:   []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
            Input:     8,
            Assertion: AssertOutput(0),
        },
        {
            Name:      "immediate-mode-less-than-8",
            Program:   []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
            Input:     5,
            Assertion: AssertOutput(1),
        },
        {
            Program:   []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
            Input:     8,
            Assertion: AssertOutput(1),
        },
        {
            Program:   []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
            Input:     7,
            Assertion: AssertOutput(0),
        },
        {
            Program:   []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
            Input:     8,
            Assertion: AssertOutput(0),
        },
        {
            Program:   []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
            Input:     5,
            Assertion: AssertOutput(1),
        },
    }

    for i, v := range expectedResults {
        if v.Name == "" {
            v.Name = fmt.Sprintf("test-%v", i)
        }
        t.Run(v.Name, func(t *testing.T) {
            computer := NewComputer(v.Program)
            computer.Inputs = []int { v.Input }
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
}

func TestDiagnostics(t *testing.T) {
	for _, systemId := range []int{1, 5} {
		computer, err := ComputerFromFile("../input_files/day_5.txt")
		if err != nil {
			t.Error(err)
		} else {
			computer.Inputs = []int { systemId }
			err := computer.Run()
			if err != nil {
				t.Log(strings.Join(computer.Debugger, "\n"))
				t.Error(err)
			} else {
				t.Logf("Diagnostics for system %v ran successfully! Output: %v", systemId, computer.OutputRegister)
			}
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

// Heap's algorithm
func generatePermutations(a []int, size int) (result [][]int) {
    if size == 1 {
        b := make([]int, len(a))
        copy(b,a)
        result = append(result, b)
    }

    for i := 0; i < size; i++ {
        result = append(result, generatePermutations(a, size-1)...)

        if size%2 == 1 {
            a[0], a[size-1] = a[size-1], a[0]
        } else {
            a[i], a[size-1] = a[size-1], a[i]
        }
    }
    return result
}
func TestThrusterInputs(t *testing.T) {

    combinations := generatePermutations([]int { 0,1,2,3,4 }, 5)

    findMaxThrusterSetting := func(program []int, expectedOutput int) (int, error) {
        maxOutput := 0
        for _, c := range combinations {
                computer := NewComputer(program)

            output := 0
            for _, p := range c {
                computer.Reset()
                computer.Inputs = []int{p, output}
                computer.Run()
                output = computer.OutputRegister
            }
            if maxOutput < output {
                maxOutput = output
            }
        }
        if maxOutput != expectedOutput {
            return maxOutput, errors.New(fmt.Sprintf("expected %v, found %v", expectedOutput, maxOutput))
        }
        return maxOutput, nil
    }

    if _, err := findMaxThrusterSetting([]int { 3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0 }, 43210); err != nil { t.Error(err) }
    if _, err := findMaxThrusterSetting([]int { 3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0 }, 54321); err != nil { t.Error(err) }
    if _, err := findMaxThrusterSetting([]int { 3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33, 1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0 }, 65210); err != nil { t.Error(err) }

    program, err := LoadProgram("../input_files/day_7.txt")
    if err != nil {
        t.Error(err)
    }

    maxOutput, _ := findMaxThrusterSetting(program, 0)
    t.Log(maxOutput)

}
