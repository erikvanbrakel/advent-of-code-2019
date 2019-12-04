package pkg

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

type Computer struct {
	Memory []int
	Cursor int
	Halted bool
}

var opCodes = map[int]func(c *Computer) {
	1: func (c *Computer) {
		addr1 := c.Memory[c.Cursor + 1]
		addr2 := c.Memory[c.Cursor + 2]
		addr3 := c.Memory[c.Cursor + 3]

		c.Memory[addr3] = c.Memory[addr1] + c.Memory[addr2]
		c.Cursor += 4
	},
	2: func(c *Computer) {
		addr1 := c.Memory[c.Cursor + 1]
		addr2 := c.Memory[c.Cursor + 2]
		addr3 := c.Memory[c.Cursor + 3]

		c.Memory[addr3] = c.Memory[addr1] * c.Memory[addr2]
		c.Cursor += 4
	},
	99: func(c *Computer) {
		c.Halted = true
	},
}


func NewComputer(initialMemory []int) *Computer {
	return &Computer {
		Memory: initialMemory,
		Cursor: 0,
		Halted: false,
	}
}

func (c *Computer) NextInstruction() int {
	opCode := c.Memory[c.Cursor]
	return opCode
}

func (c *Computer) Run() {
	for {
		if c.Halted {
			return
		}
		opCode := c.NextInstruction()
		opCodes[opCode](c)
	}
}

func ComputerFromFile(path string) (*Computer, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	initialMemory := []int {}
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
		Memory: initialMemory,
		Cursor: 0,
		Halted: false,
	}, nil
}