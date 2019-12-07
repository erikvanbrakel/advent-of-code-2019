package pkg

import (
    "bufio"
    "os"
    "testing"
)

func TestOrbits(t *testing.T) {
    testData := []string {
        "COM)B",
        "B)C",
        "C)D",
        "D)E",
        "E)F",
        "B)G",
        "G)H",
        "D)I",
        "E)J",
        "J)K",
        "K)L",
    }

    countOrbits([]string { "COM)B"}, 1, t)
    countOrbits([]string { "COM)B", "COM)C"}, 2, t)
    countOrbits([]string { "COM)B", "COM)C", "C)D"}, 4, t)
    countOrbits([]string { "COM)B", "COM)C", "C)D", "COM1)B1"}, 5, t)
    countOrbits([]string { "COM)A", "A)B", "B)C" }, 6, t)
    countOrbits([]string { "A)B", "COM)A", "B)C", }, 6, t)
    countOrbits(testData, 42, t)
}

func countOrbits(testData []string, expectedValue int, t *testing.T) {
    roots := BuildTree(testData)

    totalOrbits := 0

    for _, root := range roots {
        descendants := root.FindDescendants()
        for _, n := range descendants {
            totalOrbits += len(n.FindAncestors())
        }
    }
    if totalOrbits != expectedValue {
        t.Errorf("Expected %v orbits, got %v", expectedValue, totalOrbits)
    }
}


func TestOrbitsFromFile(t *testing.T) {
    root, err := BuildTreeFromFile("../input_files/day_6.txt")

    if err != nil {
        t.Error(err)
    }

    totalOrbits := 0
    for _, r := range root {
        descendants := r.FindDescendants()
        for _, n := range descendants {
            totalOrbits += len(n.FindAncestors())
        }
    }
    file, err := os.Open("../input_files/day_6.txt")
    if err != nil {
        t.Error(err)
    }
    defer file.Close()
    lines := []string {}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    t.Log(totalOrbits)
}