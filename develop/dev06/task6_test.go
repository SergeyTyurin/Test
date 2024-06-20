package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCut(t *testing.T) {
	fields = []int{2, 4}

	testData := []struct {
		Input  []Row
		Output []Row
	}{
		{
			Input: []Row{
				{
					Separator: true,
					Column:    []string{"1", "2", "3", "4", "5"},
				},
			},
			Output: []Row{
				{
					Column: []string{"2", "4"},
				},
			},
		},

		{
			Input: []Row{
				{
					Separator: false,
					Column:    []string{"1gjg5"},
				},
			},
			Output: []Row{
				{
					Column: []string{"1gjg5"},
				},
			},
		},

		{
			Input: []Row{
				{
					Separator: true,
					Column:    []string{"1gjg5"},
				},
			},
			Output: []Row{
				{
					Column: nil,
				},
			},
		},

		{
			Input: []Row{
				{
					Separator: false,
					Column:    []string{"2", "4"},
				},
			},
			Output: []Row{
				{
					Column: []string{"2"},
				},
			},
		},
	}

	for _, data := range testData {
		var tab Tab = data.Input

		var tabRes Tab = data.Output
		res := cut(tab)
		require.Equal(t, tabRes, res)
	}
}
