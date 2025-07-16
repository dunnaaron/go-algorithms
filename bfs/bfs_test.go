package bfs

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		start    int
		expected []int
	}{
		{
			name: "linear graph",
			edges: [][]int{
				{1, 2},
				{2, 3},
			},
			start:    1,
			expected: []int{1, 2, 3},
		},
		{
			name: "branching graph",
			edges: [][]int{
				{1, 2},
				{1, 3},
				{2, 4},
				{3, 5},
			},
			start:    1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "disconnected graph",
			edges: [][]int{
				{1, 2},
				{3, 4},
			},
			start:    3,
			expected: []int{3, 4},
		},
		{
			name: "single node",
			edges: [][]int{},
			start:    1,
			expected: []int{1},
		},
		{
			name: "empty graph",
			edges: [][]int{},
			start:    0,
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BFS(tt.edges, tt.start)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BFS() = %v, want %v", result, tt.expected)
			}
		})
	}
}
