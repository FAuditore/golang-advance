package main

import "fmt"

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"A", "C", "D"},
		"C": {"A", "B", "D", "E"},
		"D": {"B", "C", "E", "F"},
		"E": {"C", "D"},
		"F": {"D"},
	}
	//var result = make(chan string)
	fmt.Println("->BFS:")
	BFS(graph, "A")
	fmt.Println("->DFS:")
	DFS(graph, "A")
}

func BFS(graph map[string][]string, start string) {
	visit := make(map[string]struct{})
	var queue []string
	queue = append(queue, start)
	visit[start] = struct{}{}

	for len(queue) > 0 {
		//访问下一个节点 并出队
		nextNode := queue[0]
		fmt.Println(nextNode)
		queue = queue[1:]

		//该节点所有相邻且未访问的节点入队
		for _, node := range graph[nextNode] {
			if _, ok := visit[node]; !ok {
				queue = append(queue, node)
				visit[node] = struct{}{}
			}
		}
	}
}
func DFS(graph map[string][]string, start string) {
	visit := make(map[string]struct{})
	var queue []string
	queue = append(queue, start)
	visit[start] = struct{}{}

	for len(queue) > 0 {
		//访问下一个节点 并出队
		nextNode := queue[len(queue)-1]
		fmt.Println(nextNode)
		queue = queue[:len(queue)-1]

		//该节点所有相邻且未访问的节点入队
		for _, node := range graph[nextNode] {
			if _, ok := visit[node]; !ok {
				queue = append(queue, node)
				visit[node] = struct{}{}
			}
		}
	}
}
