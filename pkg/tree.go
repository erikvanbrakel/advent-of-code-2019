package pkg

import (
    "bufio"
    "os"
    "strings"
)

type Node struct {
    Data string
    Parent *Node
    Children []*Node
}

func BuildTreeFromFile(path string) ([]*Node, error) {
    file, err := os.Open(path)
    defer file.Close()

    if err != nil {
        return nil, err
    }

    data := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        data = append(data, scanner.Text())
    }

    return BuildTree(data), nil
}

func BuildTree(data []string) []*Node {
    nodes := map[string]*Node {}
    for _, r := range data {
        parts := strings.Split(r, ")")
        if nodes[parts[0]] == nil {
            nodes[parts[0]] = &Node{ Data: parts[0] }
        }
        parent := nodes[parts[0]]

        if nodes[parts[1]] == nil {
            nodes[parts[1]] = &Node{ Data: parts[1] }
        }

        nodes[parts[1]].Parent = parent
        parent.AddChild(nodes[parts[1]])
    }

    roots := []*Node {}

    for _,v := range nodes {
        if v.Parent == nil {
            roots = append(roots, v)
        }
    }

    return roots
}

func (n *Node) AddChild(child *Node) {
    n.Children = append(n.Children, child)
}

func (n *Node) FindAncestors() []*Node {
    if n.Parent == nil {
        return []*Node{}
    }

    return append(n.Parent.FindAncestors(), n.Parent)
}

func (n *Node) FindDescendants() []*Node {
    if len(n.Children) == 0 {
        return []*Node {}
    }

    descendants := n.Children
    for _,c := range n.Children {
        descendants = append(descendants, c.FindDescendants()...)
    }
    return descendants
}

func (n *Node) FindDescendant(data string) *Node {
    if n.Data == data { return n }
    for _,c := range n.Children {
        descendant := c.FindDescendant(data)
        if descendant != nil {
            return descendant
        }
    }
    return nil
}

func (n *Node) DistanceTo(other *Node) int {
    a1 := n.FindAncestors()
    a2 := other.FindAncestors()
    hashmap := map[string]*Node {}
    for _, v := range a1 {
        hashmap[v.Data] = v
    }
    for j:=len(a2)-1;j>=0;j-- {
        v := a2[j]
        if hashmap[v.Data] != nil {
            distance := 0
            for i:=0;i<len(a1);i++ {
                v2 := a1[len(a1)-i-1]
                if v2.Data == v.Data {
                    distance += i
                    break
                }
            }
            for i:=0;i<len(a2);i++ {
                v2 := a2[len(a2)-i-1]
                if v2.Data == v.Data {
                    distance += i
                    break
                }
            }
            return distance
        }
    }
    return 0
}