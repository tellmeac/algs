package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rv(v int) *int {
	return &v
}

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice    []*int
		expected *Node
	}{
		{
			name:  "No nils",
			slice: []*int{rv(1), rv(2), rv(3), rv(4)},
			expected: &Node{
				V: 1,
				L: &Node{
					V: 2,
					L: &Node{
						V: 4,
					},
				},
				R: &Node{
					V: 3,
				},
			},
		},
		{
			name:  "With nils",
			slice: []*int{rv(1), rv(2), nil, rv(4)},
			expected: &Node{
				V: 1,
				L: &Node{
					V: 2,
					L: &Node{
						V: 4,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, isEqual(FromSlice(tt.slice), tt.expected))
		})
	}
}

func isEqual(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.V == b.V && isEqual(a.L, b.L) && isEqual(a.R, b.R)
}

func TestInOrderDFS(t *testing.T) {
	tests := []struct {
		name          string
		root          *Node
		expectedOrder []int
	}{
		{
			name:          "Common case",
			expectedOrder: []int{4, 2, 1, 3},
			root: &Node{
				V: 1,
				L: &Node{
					V: 2,
					L: &Node{
						V: 4,
					},
				},
				R: &Node{
					V: 3,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedOrder, InOrderDFS(tt.root))
		})
	}
}
