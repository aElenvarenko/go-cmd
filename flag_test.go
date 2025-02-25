package cmd_test

import (
	"testing"

	"github.com/aElenvarenko/go-cmd"
)

func TestFlag(t *testing.T) {
	flagName := "test"
	flagAlias := "t"
	flagDesc := "test flag"
	flagReq := false
	flagDefault := "test"
	flag := cmd.NewFlag(flagName, flagAlias, flagDesc, flagReq, flagDefault)

	if flag.GetName() != flagName {
		t.Errorf("Flag name not equal")
	}

	if flag.GetAlias() != flagAlias {
		t.Errorf("Flag alias not equal")
	}

	if flag.GetDesc() != flagDesc {
		t.Errorf("Flag description not equal")
	}

	if flag.GetRequired() != flagReq {
		t.Errorf("Flag required not equal")
	}

	if flag.GetDefaultValue() != flagDefault {
		t.Errorf("Flag defaultValue not equal")
	}
}
