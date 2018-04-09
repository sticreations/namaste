package namaste

import (
	"testing"
)

func TestReadMasterplan(t *testing.T) {
	m := readMasterplan("./testdata/nodejs/fes/masterplan.yaml")
	if m.Vars == nil || m.VarsNeeded == nil {
		t.Error("Could not read vars")
	}
}
