package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIncorrectNtpServer(t *testing.T) {
	ntpServer = "incorrect"
	require.Error(t, Time())
}
