// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	//"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

const (
	undiscovered = iota
	discovered
	finished
)

func main() {
	topo, err := topoSort(prereqs)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for i, course := range topo {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]int)
	var visitAll func(from string, items []string,
		visited *map[string]int) error

	// Returns an error if a cycle is encountered,
	//  which happens when we encounter a back edge
	// Since this is a directed graph, we don't have
	//  to check if it's where we came from
	visitAll = func(from string, items []string, visited *map[string]int) error {
		for _, item := range items {
			if (*visited)[item] == discovered {
				return fmt.Errorf("Cycle in input graph from %s to %s",
					from, item)
			}
			if (*visited)[item] == undiscovered {
				(*visited)[item] = discovered
				errval := visitAll(item, m[item], visited)
				if errval != nil {
					return errval
				}
				order = append(order, item)
				(*visited)[item] = finished
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	hasCycleError := visitAll("", keys, &seen)
	if hasCycleError != nil {
		return nil, hasCycleError
	}
	return order, nil
}
