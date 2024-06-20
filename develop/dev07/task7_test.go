package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCloseSingleChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-joinChannels(
		sig(3*time.Second),
		sig(1*time.Second),
	)

	require.Equal(t, 1, int(time.Since(start).Seconds()))
}

func TestCloseSingleChannelEmpty(t *testing.T) {
	start := time.Now()
	<-joinChannels()

	require.Equal(t, 0, int(time.Since(start).Seconds()))
}
