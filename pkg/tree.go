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