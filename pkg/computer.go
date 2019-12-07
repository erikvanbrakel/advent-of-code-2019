package pkg

import (
    "bufio"
    "bytes"
    "errors"
    "fmt"
    "os"
    "strconv"
)

type Computer struct {
    Memory             []int
    InstructionPointer int
    InputRegister      int
    OutputRegister     int
    Halted             bool
    Debugger  []string
}

var opCodes = map[int]func(c *Computer, direct []bool){
    // add
    1: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "ADD")
        result := c.GetMemoryValue(c.InstructionPointer+1, direct[0]) + c.GetMemoryValue(c.InstructionPointer+2, direct[1])
        c.SetMemoryValue(c.InstructionPointer+3, result)
        c.InstructionPointer += 4
    },
    // multiply
    2: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "MUL")
        result := c.GetMemoryValue(c.InstructionPointer+1, direct[0]) * c.GetMemoryValue(c.InstructionPointer+2, direct[1])
        c.SetMemoryValue(c.InstructionPointer+3, result)
        c.InstructionPointer += 4
    },
    // read input
    3: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "IN")
        c.SetMemoryValue(c.InstructionPointer+1, c.InputRegister)
        c.InstructionPointer += 2
    },
    // write output
    4: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "OUT")
        c.OutputRegister = c.GetMemoryValue(c.InstructionPointer+1, direct[0])
        c.InstructionPointer += 2
    },
    // jump if true
    5: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "JMPTRUE")
        if c.GetMemoryValue(c.InstructionPointer+1, direct[0]) != 0 {
            c.InstructionPointer = c.GetMemoryValue(c.InstructionPointer+2, direct[1])
        } else {
            c.InstructionPointer += 3
        }
    },
    // jump if false
    6: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "JMPFALSE")
        if c.GetMemoryValue(c.InstructionPointer+1, direct[0]) == 0 {
            c.InstructionPointer = c.GetMemoryValue(c.InstructionPointer+2, direct[1])
        } else {
            c.InstructionPointer += 3
        }
    },
    // less than
    7: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "LESS")
        if c.GetMemoryValue(c.InstructionPointer+1, direct[0]) < c.GetMemoryValue(c.InstructionPointer+2, direct[1]) {
            c.SetMemoryValue(c.InstructionPointer+3, 1)
        } else {
            c.SetMemoryValue(c.InstructionPointer+3, 0)
        }
        c.InstructionPointer += 4
    },
    // equal to
    8: func(c *Computer, direct []bool) {
        c.Debugger = append(c.Debugger, "GREATER")
        if c.GetMemoryValue(c.InstructionPointer+1, direct[0]) == c.GetMemoryValue(c.InstructionPointer+2, direct[1]) {
            c.SetMemoryValue(c.InstructionPointer+3, 1)
        } else {
            c.SetMemoryValue(c.InstructionPointer+3, 0)
        }
        c.InstructionPointer += 4
    },
    // halt
    99: func(c *Computer, direct []bool) {
        c.Halted = true
    },
}

func NewComputer(initialMemory []int) *Computer {
    return &Computer{
        Memory:             initialMemory,
        InstructionPointer: 0,
        Halted:             false,
    }
}

func (c *Computer) GetMemoryValue(index int, direct bool) int {
    if direct {
        return c.Memory[index]
    } else {
        return c.GetMemoryValue(c.Memory[index], true)
    }
}

func (c *Computer) SetMemoryValue(index, value int) {
	c.Memory[c.Memory[index]] = value
}

type Instruction struct {
    OpCode       int
    addressModes []bool
}

func (c *Computer) NextInstruction() Instruction {
    opCode := c.GetMemoryValue(c.InstructionPointer, true)

    return Instruction{
        opCode % 100,
        []bool{
            (opCode/100)%10 == 1,
            (opCode/1000)%10 == 1,
            (opCode/10000)%10 == 1,
        },
    }
}

func (c *Computer) Run() error {
    for {
        if c.Halted {
            return nil
        }
        instruction := c.NextInstruction()
        handler := opCodes[instruction.OpCode]
        if handler == nil {
            return errors.New(fmt.Sprintf("Invalid opcode: %v", instruction.OpCode))
        }
        handler(c, instruction.addressModes)
    }
}

func ComputerFromFile(path string) (*Computer, error) {
    file, err := os.Open(path)
    defer file.Close()

    if err != nil {
        return nil, err
    }

    initialMemory := []int{}
    scanner := bufio.NewScanner(file)
    scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {

        search := []byte(",")
        if i := bytes.Index(data, search); i >= 0 {
            return i + len(search), data[0:i], nil
        }

        if atEOF && len(data) != 0 {
            return len(data), data, nil
        }

        return 0, nil, nil
    })
    for scanner.Scan() {
        memory, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return nil, err
        }
        initialMemory = append(initialMemory, memory)
    }

    return &Computer{
        Memory:             initialMemory,
        InstructionPointer: 0,
        Halted:             false,
    }, nil
}
