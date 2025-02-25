package cmd_test

import (
	"testing"

	"github.com/aElenvarenko/go-cmd"
)

func TestArgument(t *testing.T) {
	argumentName := "name"
	argumentValue := "value"
	argument := cmd.NewArgument(argumentName, argumentValue)

	if argument.GetName() != argumentName {
		t.Errorf("Argument name not equal")
	}

	if argument.GetValue() != argumentValue {
		t.Errorf("Argument value not equal")
	}
}
