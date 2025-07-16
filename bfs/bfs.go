package bfs

import "slices"

// BFS performs a breadth-first search starting from the given start node.
// It returns the order in which nodes are visited.
// Note: this implementation does not account for cyclical graphs
func BFS(edges [][]int, start int) []int {
	graph := make(map[int][]int)

	// If not provided the graph itself, need to create a graph from the edges
	// This assumes there are no duplicate edges
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	
	// A queue to track any remaining nodes and the order they should be visited
	queue := []int{start}

	// A visited slice to indicate if a given node has been visited
	visited := []int{}

	for len(queue) > 0 {
			if !slices.Contains(visited, queue[0]) {
				visited = append(visited, queue[0])
				
				for _, neighbor := range graph[queue[0]] {

					if !slices.Contains(visited, neighbor) {
						queue = append(queue, neighbor)
					}
				}

				queue = queue[1:]
			}
	}

	return visited
}