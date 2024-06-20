package task2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	data := []struct {
		Input  string
		Output string
		Error  error
	}{
		{"abcd", "abcd", nil},
		{"a4bc2d5e", "aaaabccddddde", nil},
		{"45", "", fmt.Errorf("invalid string: 45")},
		{"abc11d", "abcccccccccccd", nil},
		{"", "", nil},
		{".8", "........", nil},
		{"ab0", "a", nil},
		{"f3", "fff", nil},
		{"f", "f", nil},
		{"-10", "----------", nil},
		{"абв10", "абвввввввввв", nil},
		{"а5бпак6в10", "ааааабпакккккквввввввввв", nil},
		{`\a0bc`, `bc`, nil},
		{`\\`, `\`, nil},
		{`\`, "", fmt.Errorf(`invalid string: \`)},
		{`\45`, `44444`, nil},
		{`\\6`, `\\\\\\`, nil},
		{`\d3`, `ddd`, nil},
		{`\3`, `3`, nil},
		{`qwe\3\45`, `qwe344444`, nil},
		{`\\\`, "", fmt.Errorf(`invalid string: \\\`)},
	}

	for _, item := range data {
		res, err := Unpack(item.Input)
		require.Equal(t, item.Output, res, item.Input)
		require.Equal(t, item.Error, err, item.Input)
	}

}
