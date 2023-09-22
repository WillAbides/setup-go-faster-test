package testsetupgofaster

import (
	"fmt"
	"testing"
)

func TestNothing(t *testing.T) {
	fmt.Println("starting test")
	fmt.Println("ending test")
}

func TestLogging(t *testing.T) {
	t.Log("hello from TestLogging")
}
