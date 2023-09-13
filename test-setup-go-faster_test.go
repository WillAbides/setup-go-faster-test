package testsetupgofaster

import (
	"fmt"
	"testing"
)

func TestNothing(t *testing.T) {
	fmt.Println("starting test")
	fmt.Println("ending test")
}

func TestFailure(t *testing.T) {
	t.Errorf("This test should fail")
}
