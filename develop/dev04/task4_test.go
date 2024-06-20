package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnagranm(t *testing.T) {
	input := []string{"мааал", "пятак", "НоГа", "слиток", "пятка", "тяпка", "листок", "ламаа", "столик", "рука", "нога", "столик"}
	output := map[string][]string{
		"мааал":  {"ламаа", "мааал"},
		"пятак":  {"пятак", "пятка", "тяпка"},
		"слиток": {"листок", "слиток", "столик"},
	}

	require.Equal(t, output, Anagram(input))
}
