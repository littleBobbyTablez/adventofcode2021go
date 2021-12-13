package main

import (
	"fmt"
	"strings"
)

type path struct {
	doubleSmall bool
	path        string
}

func findPaths2(input map[string][]string) {

	visit := []string{"start"}

	paths := make(map[path]bool)
	paths[path{false, "start"}] = true
	step2(input, visit, paths)
}

func step2(input map[string][]string, toVisit []string, paths map[path]bool) {
	if len(toVisit) == 0 {
		fmt.Println(len(paths))
		return
	}

	node := toVisit[0]
	nextNodes := input[node]

	paths = appendIfEndsWithNode2(nextNodes, paths, node)

	if isLower(node) && !needed2(node, paths) {
		for k, v := range input {
			input[k] = remove(v, node)
		}
		delete(input, node)
	}

	step2(input, append(toVisit[1:], nextNodes...), paths)
}

func needed2(node string, paths map[path]bool) bool {
	needed := false
	if node == "start" {
		return false
	}
	for p, _ := range paths {
		isNotFinished := p.path[len(p.path)-3:] != "end"
		if (isNotFinished && !strings.Contains(p.path, node)) || (isNotFinished && !p.doubleSmall) {
			needed = true
		}
	}
	return needed
}

func appendIfEndsWithNode2(nextNodes []string, paths map[path]bool, node string) map[path]bool {
	if node == "end" || nextNodes == nil {
		return paths
	}
	for p, _ := range paths {
		if p.path[len(p.path)-len(node):] == node {
			for _, nextNode := range nextNodes {
				if !isLower(nextNode) || !strings.Contains(p.path, nextNode) {
					paths[path{p.doubleSmall, p.path + "," + nextNode}] = true
				} else if strings.Contains(p.path, nextNode) && !p.doubleSmall {
					paths[path{true, p.path + "," + nextNode}] = true
				}
			}
			delete(paths, p)
		}
	}

	return paths
}
