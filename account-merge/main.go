package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	accounts := [][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}

	res := accountsMerge(accounts)

	fmt.Println(res)

}

func accountsMerge(accounts [][]string) [][]string {
	return merge2(accounts)
}

type ListNode struct {
	email string
	next  *ListNode
}

var visited map[string]bool
var adjacencyMap map[string]map[string]bool

func merge2(accounts [][]string) [][]string {
	visited = make(map[string]bool)
	adjacencyMap = make(map[string]map[string]bool)
	for _, accounts := range accounts {
		for _, account := range accounts[1:] {
			if adjacencyMap[account] == nil {
				adjacencyMap[account] = make(map[string]bool)
			}
			for _, account2 := range accounts {
				if account == account2 {
					continue
				}
				adjacencyMap[account][account2] = true
			}
		}
	}

	result := [][]string{}

	for key := range adjacencyMap {
		if !visited[key] {
			name, emails := bfs(key)
			result = append(result, append([]string{name}, emails...))
		}
	}
	return result
}

func bfs(email string) (string, []string) {
	ans := make(map[string]bool)
	ans[email] = true
	var name string
	for key := range adjacencyMap[email] {
		if !strings.Contains(key, "@") {
			name = key
			break
		}
	}
	head := &ListNode{email: email, next: nil}
	tail := head
	visited[email] = true
	for head != nil {
		for key := range adjacencyMap[head.email] {
			if !ans[key] && key != name {
				tail.next = &ListNode{email: key, next: nil}
				tail = tail.next
				ans[key] = true
				visited[key] = true
			}
		}
		head = (*head).next
	}
	result := []string{}
	for key := range ans {
		result = append(result, key)
	}
	sort.Strings(result)

	return name, result
}
