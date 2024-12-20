package main

import "fmt"

func main() {
	fmt.Println(minPartitions("27346209830709182346"))
}

func minPartitions(n string) int {
	maxN := 0
	for _, char := range n {
		num := int(char - '0')
		if num > maxN {
			maxN = num
		}
		if maxN == 9 {
			break
		}
	}
	return maxN
}
