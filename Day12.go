package main

import (
	"strings"
	"unicode"
)

func findPaths(input map[string][]string) int {
	visit := []string{"start"}
	return step(input, visit, visit)
}

func step(input map[string][]string, toVisit []string, paths []string) int {
	if len(toVisit) == 0 {
		output := make(map[string]bool)
		for i, path := range paths {
			split := strings.Split(path, ",")
			if split[len(split)-1] == "end" {
				output[paths[i]] = true
			}
		}
		return len(output)
	}

	node := toVisit[0]
	nextNodes := input[node]

	paths = appendIfEndsWithNode(nextNodes, paths, node)

	if isLower(node) && !needed(node, paths) {
		for k, v := range input {
			input[k] = remove(v, node)
		}
		delete(input, node)
	}

	return step(input, append(toVisit[1:], nextNodes...), paths)
}

func needed(node string, paths []string) bool {
	needed := false
	for _, path := range paths {
		if path[len(path)-3:] != "end" && !strings.Contains(path, node) {
			needed = true
		}
	}
	return needed
}

func appendIfEndsWithNode(nextNodes []string, paths []string, node string) []string {
	if node == "end" || nextNodes == nil {
		return paths
	}
	var newPaths []string
	var pathsToRemove []string
	for i, path := range paths {
		if path[len(path)-len(node):] == node {
			for _, nextNode := range nextNodes {
				if !strings.Contains(path, nextNode) || !isLower(nextNode) {
					newPaths = append(newPaths, path+","+nextNode)
				}
			}
			pathsToRemove = append(pathsToRemove, paths[i])
		}
	}

	for _, p := range pathsToRemove {
		paths = remove(paths, p)
	}
	return append(paths, newPaths...)
}

func remove(s []string, n string) []string {
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

func isLower(s string) bool {
	return unicode.IsLower(rune(s[0]))
}
