package pkg

import (
    "errors"
    "fmt"
    "strconv"
)

type Coordinate struct {
    X, Y int
}

type Path struct {
    Coordinates []Coordinate
}

type Pathfinder struct {
    Path    Path
    Current Coordinate
}

func (c *Coordinate) ManhattanDistance(o *Coordinate) int {
    return absInt(c.X-o.X) + absInt(c.Y-o.Y)
}

func (p *Path) Intersect(p2 *Path) []Coordinate {
    intersections := []Coordinate{}

    hashMap := map[string]int{}
    for _, c := range p.Coordinates {
        hashMap[fmt.Sprintf("%v-%v", c.X, c.Y)] = 1
    }
    for _, c := range p2.Coordinates {
        if hashMap[fmt.Sprintf("%v-%v", c.X, c.Y)] == 1 {
            intersections = append(intersections, c)
        }
    }

    return intersections
}

func (p *Path) AddCoordinate(c Coordinate) {
    p.Coordinates = append(p.Coordinates, c)
}

func absInt(i int) int {
    if i < 0 {
        return -i
    } else {
        return i
    }
}

func (p *Pathfinder) Move(instructions ...string) error {
    for _, instruction := range instructions {
        direction := instruction[0]
        amount, err := strconv.Atoi(instruction[1:])
        if err != nil {
            return err
        }
        switch direction {
        case 'U':
            for y := 0; y < amount; y++ {
                p.Current.Y += 1
                p.Path.AddCoordinate(p.Current)
            }
            break
        case 'D':
            for y := 0; y < amount; y++ {
                p.Current.Y -= 1
                p.Path.AddCoordinate(p.Current)
            }
            break
        case 'L':
            for x := 0; x < amount; x++ {
                p.Current.X -= 1
                p.Path.AddCoordinate(p.Current)
            }
            break
        case 'R':
            for x := 0; x < amount; x++ {
                p.Current.X += 1
                p.Path.AddCoordinate(p.Current)
            }
            break
        default:
            return errors.New("oops!")
        }
    }
    return nil
}
