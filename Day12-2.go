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
	paths := []path{{false, "start"}}
	step2(input, visit, paths)
}

func step2(input map[string][]string, toVisit []string, paths []path) {
	if len(toVisit) == 0 {
		output := make(map[path]bool)
		for i, path := range paths {
			split := strings.Split(path.path, ",")
			if split[len(split)-1] == "end" {
				output[paths[i]] = true
			}
		}
		fmt.Println(len(output))
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

func needed2(node string, paths []path) bool {
	needed := false
	if node == "start" {
		return false
	}
	for _, p := range paths {
		split := strings.Split(p.path, ",")
		if (split[len(split)-1] != "end" && !strings.Contains(p.path, node)) || (split[len(split)-1] != "end" && !p.doubleSmall) {
			needed = true
		}
	}
	return needed
}

func appendIfEndsWithNode2(nextNodes []string, paths []path, node string) []path {
	if node == "end" || nextNodes == nil {
		return paths
	}
	var newPaths []path
	var pathsToRemove []path
	for i, p := range paths {
		split := strings.Split(p.path, ",")
		if split[len(split)-1] == node {
			for _, nextNode := range nextNodes {
				if !isLower(nextNode) {
					newPaths = append(newPaths, path{p.doubleSmall, strings.Join(append(split, nextNode), ",")})
				} else if !strings.Contains(p.path, nextNode) {
					newPaths = append(newPaths, path{p.doubleSmall, strings.Join(append(split, nextNode), ",")})
				} else if strings.Contains(p.path, nextNode) && !p.doubleSmall {
					newPaths = append(newPaths, path{true, strings.Join(append(split, nextNode), ",")})
				}
			}
			pathsToRemove = append(pathsToRemove, paths[i])
		}
	}

	for _, p := range pathsToRemove {
		paths = remove2(paths, p)
	}
	return append(paths, newPaths...)
}

func remove2(s []path, n path) []path {
	length := len(s)
	found := false
	for i := 0; i < length; i++ {
		subject := s[i]
		if subject == n {
			s[i] = s[length-1]
			found = true
		}
	}

	if found {
		return s[:length-1]
	}
	return s
}
