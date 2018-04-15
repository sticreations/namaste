package namaste

import (
	"testing"
)

func TestReadMasterplan(t *testing.T) {
	m := readMasterplan("./testdata/nodejs/testservice/masterplan.yaml")
	if m.Vars == nil || m.VarsNeeded == nil {
		t.Error("Could not read vars")
	}
}

func TestGetVarsNeeded(t *testing.T) {
	m := readMasterplan("./testdata/nodejs/testservice/masterplan.yaml")
	vars := m.GetVarsNeeded()
	if _, ok := vars["name"]; !ok {
		t.Error("Variable 'name' is not in VarsNeeded")
	}
}

func TestSetVarNeeded(t *testing.T) {
	const testName = "this is a test!"
	m := readMasterplan("./testdata/nodejs/testservice/masterplan.yaml")
	m.SetVarNeeded("name", testName)
	vars := m.GetVarsNeeded()
	f, ok := vars["name"]
	if !ok || f != testName {
		t.Error("Variable 'name' is not in VarsNeeded")
	}
}

func TestValidateVarsNeeded(t *testing.T) {
	const testName = "this is a test!"
	m := readMasterplan("./testdata/nodejs/testservice/masterplan.yaml")
	err := m.ValidateVarsNeeded()
	if err == nil {
		t.Error("There are missing Vars.")
	}
	m.SetVarNeeded("name", testName)
	err = m.ValidateVarsNeeded()
	if err != nil {
		t.Errorf("Validate Vars did not work: %v", err)
	}
}
