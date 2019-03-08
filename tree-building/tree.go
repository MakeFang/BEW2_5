package tree

import (
    // "fmt"
    "sort"
    "errors"
)


type Node struct {
    ID int
    Children []*Node
}

type Record struct {
    ID int
    Parent int
}

func Build(records []Record) (*Node, error) {
    if len(records) < 1 {
        return nil, nil
    }
    sort.Slice(records[:], func(i, j int) bool {
        return records[i].ID < records[j].ID
    })
    // fmt.Println(records)
    valInd := 0
    var resNodes []*Node
    for _, value := range records {
    // fmt.Println(index, value)
        if value.ID != valInd || value.ID < value.Parent {
            return nil, errors.New("construction error")
        }
        if value.ID == value.Parent {
            if value.ID > 0 {
                return nil, errors.New("construction error")
            }
            newNode := Node{value.ID, nil}
            resNodes = append(resNodes, &newNode)
        } else {
            newNode := Node{value.ID, nil}
            resNodes = append(resNodes, &newNode)
            resNodes[value.Parent].Children = append(resNodes[value.Parent].Children, &newNode)
        }
    // fmt.Println(resNodes)
        valInd += 1
    }
    return resNodes[0], nil
}
