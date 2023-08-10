package testsetupgofaster

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNothing(t *testing.T) {
	fmt.Println("starting test")
	require.True(t, true)
	fmt.Println("ending test")
}
