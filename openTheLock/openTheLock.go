package openTheLock

import (
	"strconv"
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
			strArr := strings.Split(str, "")
			for i, d := range strArr {
				//anti-clockwise
				antiClockwiseDigit, _ := strconv.Atoi(d)
				antiClockwiseString := strings.Join(strArr[:i], "") + strconv.Itoa((antiClockwiseDigit+1)%10) + strings.Join(strArr[i+1:], "")
				if !visited[antiClockwiseString] {
					visited[antiClockwiseString] = true
					newQueue = append(newQueue, antiClockwiseString)
				}
				//clockwise
				clockwiseDigit, _ := strconv.Atoi(d)
				clockwiseString := strings.Join(strArr[:i], "") + strconv.Itoa((clockwiseDigit-1+10)%10) + strings.Join(strArr[i+1:], "")
				if !visited[clockwiseString] {
					visited[clockwiseString] = true
					newQueue = append(newQueue, clockwiseString)
				}
			}
		}
		queue = newQueue
		rotate++
	}
	return -1
}
