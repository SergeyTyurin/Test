package main

import (
	"io"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func setTestDefaultFlags() {
	isNotDouble = false
	isReverse = false
	clmnSort = 0
	isNumeric = false
}

func TestSortStrings(t *testing.T) {
	data := struct {
		input  Tab
		output []Row
	}{
		input: Tab{rows: []Row{
			[]string{"a", "2"},
			[]string{"5", "1"},
			[]string{"ф", "0"},
			[]string{"1", "a"},
			[]string{"1", "a"},
			[]string{"1", "d"},
			[]string{"2", "sfsfsf"},
		}},
		output: []Row{
			[]string{"1", "a"},
			[]string{"1", "a"},
			[]string{"1", "d"},
			[]string{"2", "sfsfsf"},
			[]string{"5", "1"},
			[]string{"a", "2"},
			[]string{"ф", "0"},
		},
	}

	setTestDefaultFlags()
	SortTab(data.input)

	require.Equal(t, data.output, data.input.rows)
}

func TestSortUniqueStrings(t *testing.T) {
	data := struct {
		input  io.Reader
		output []Row
	}{
		input: strings.NewReader(`a 2
5 1
ф 0
1 a
1 a
1 d
2 sfsfsf`),
		output: []Row{
			[]string{"1", "a"},
			[]string{"1", "d"},
			[]string{"2", "sfsfsf"},
			[]string{"5", "1"},
			[]string{"a", "2"},
			[]string{"ф", "0"},
		},
	}
	setTestDefaultFlags()

	isNotDouble = true
	tab := NewTab(data.input)
	SortTab(tab)

	require.Equal(t, data.output, tab.rows)
}

func TestSortReverseStrings(t *testing.T) {
	data := struct {
		input  io.Reader
		output []Row
	}{
		input: strings.NewReader(`a 2
5 1
ф 0
1 a
1 a
1 d
2 sfsfsf`),
		output: []Row{
			[]string{"ф", "0"},
			[]string{"a", "2"},
			[]string{"5", "1"},
			[]string{"2", "sfsfsf"},
			[]string{"1", "d"},
			[]string{"1", "a"},
			[]string{"1", "a"},
		},
	}
	setTestDefaultFlags()

	isNotDouble = false
	isReverse = true
	tab := NewTab(data.input)
	SortTab(tab)

	require.Equal(t, data.output, tab.rows)
}

func TestSortColumnStrings(t *testing.T) {
	data := struct {
		input  io.Reader
		output []Row
	}{
		input: strings.NewReader(`a 2
5 1
ф 0
1 a
1 a
1 d
2 sfsfsf`),
		output: []Row{
			[]string{"ф", "0"},
			[]string{"5", "1"},
			[]string{"a", "2"},
			[]string{"1", "a"},
			[]string{"1", "a"},
			[]string{"1", "d"},
			[]string{"2", "sfsfsf"},
		},
	}

	setTestDefaultFlags()
	clmnSort = 1
	tab := NewTab(data.input)
	sort.Sort(tab)

	require.Equal(t, data.output, tab.rows)
}

func TestSortColumnNumericalStrings(t *testing.T) {
	data := struct {
		input  io.Reader
		output []Row
	}{
		input: strings.NewReader(`a 2
5 1
ф 0
1 a
1 a
1 d
2 sfsfsf`),
		output: []Row{
			[]string{"1", "a"},
			[]string{"1", "a"},
			[]string{"1", "d"},
			[]string{"2", "sfsfsf"},
			[]string{"ф", "0"},
			[]string{"5", "1"},
			[]string{"a", "2"},
		},
	}
	setTestDefaultFlags()
	clmnSort = 1
	isNumeric = true
	tab := NewTab(data.input)
	sort.Sort(tab)

	require.Equal(t, data.output, tab.rows)
}

func TestSortDiffLenStrings(t *testing.T) {

	input := `fnjn 4\ jhiowjn dbgoi oswagh
wjfo34 43oi4o4 3o34 i2n4
mje u2343 i2n4
fie wru 2io3 
uh`

	testData := []struct {
		columnNumber int
		output       []Row
	}{
		{
			columnNumber: 0,
			output: []Row{
				[]string{"fie", "wru", "2io3", ""},
				[]string{"fnjn", "4\\", "jhiowjn", "dbgoi", "oswagh"},
				[]string{"mje", "u2343", "i2n4"},
				[]string{"uh"},
				[]string{"wjfo34", "43oi4o4", "3o34", "i2n4"},
			},
		},
		{
			columnNumber: 1,
			output: []Row{
				[]string{"wjfo34", "43oi4o4", "3o34", "i2n4"},
				[]string{"fnjn", "4\\", "jhiowjn", "dbgoi", "oswagh"},
				[]string{"mje", "u2343", "i2n4"},
				[]string{"fie", "wru", "2io3", ""},
				[]string{"uh"},
			},
		},
		{
			columnNumber: 2,
			output: []Row{
				[]string{"fie", "wru", "2io3", ""},
				[]string{"wjfo34", "43oi4o4", "3o34", "i2n4"},
				[]string{"mje", "u2343", "i2n4"},
				[]string{"fnjn", "4\\", "jhiowjn", "dbgoi", "oswagh"},
				[]string{"uh"},
			},
		},
		{
			columnNumber: 3,
			output: []Row{
				[]string{"fie", "wru", "2io3", ""},
				[]string{"fnjn", "4\\", "jhiowjn", "dbgoi", "oswagh"},
				[]string{"mje", "u2343", "i2n4"},
				[]string{"uh"},
				[]string{"wjfo34", "43oi4o4", "3o34", "i2n4"},
			},
		},
		{
			columnNumber: 4,
			output: []Row{
				[]string{"fie", "wru", "2io3", ""},
				[]string{"fnjn", "4\\", "jhiowjn", "dbgoi", "oswagh"},
				[]string{"mje", "u2343", "i2n4"},
				[]string{"uh"},
				[]string{"wjfo34", "43oi4o4", "3o34", "i2n4"},
			},
		},
		{
			columnNumber: 5,
			output: []Row{
				[]string{"fie", "wru", "2io3", ""},
				[]string{"fnjn", "4\\", "jhiowjn", "dbgoi", "oswagh"},
				[]string{"mje", "u2343", "i2n4"},
				[]string{"uh"},
				[]string{"wjfo34", "43oi4o4", "3o34", "i2n4"},
			},
		},
	}

	for _, data := range testData {
		setTestDefaultFlags()
		clmnSort = data.columnNumber
		tab := NewTab(strings.NewReader(input))
		SortTab(tab)
		require.Equal(t, data.output, tab.rows)
	}
}
