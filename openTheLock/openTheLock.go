package openTheLock

import (
	"strings"
)

func openLock(deadends []string, target string) int {
	//set up visited set
	visited := make(map[string]bool)
	for _, d := range deadends {
		visited[d] = true
	}

	if visited["0000"] {
		return -1
	} else {
		visited["0000"] = true
	}

	rotate := 0
	queue := []string{"0000"}

	//BFS
	for len(queue) > 0 {
		newQueue := []string{}
		for _, str := range queue {
			if str == target {
				return rotate
			}
			for i, size := 0, len(str); i < size; i++ {
				//anticlockwise
				var anticlockwiseBuilder strings.Builder
				if str[i] == '0' {
					anticlockwiseBuilder.WriteString("9")
				} else {
					anticlockwiseBuilder.WriteByte(str[i] - 1)
				}
				anticlockwiseStr := str[:i] + anticlockwiseBuilder.String() + str[i+1:]
				if !visited[anticlockwiseStr] {
					visited[anticlockwiseStr] = true
					newQueue = append(newQueue, anticlockwiseStr)
				}
				//clockwise
				var clockwiseBuilder strings.Builder
				if str[i] == '9' {
					clockwiseBuilder.WriteString("0")
				} else {
					clockwiseBuilder.WriteByte(str[i] + 1)
				}
				clockwiseStr := str[:i] + clockwiseBuilder.String() + str[i+1:]
				if !visited[clockwiseStr] {
					visited[clockwiseStr] = true
					newQueue = append(newQueue, clockwiseStr)
				}
			}
		}
		queue = newQueue
		rotate++
	}
	return -1
}
