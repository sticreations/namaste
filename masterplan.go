package namaste

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Masterplan struct {
	VarsNeeded map[string]string `yaml:"vars_needed"`
	Vars       map[string]string
}

func (m *Masterplan) GetVarsNeeded() map[string]string {
	return m.VarsNeeded
}

func (m *Masterplan) SetVarNeeded(k, v string) error {
	m.VarsNeeded[k] = v
	return nil
}

func (m *Masterplan) ValidateVarsNeeded() error {
	for k, v := range m.GetVarsNeeded() {
		if v == "" {
			return fmt.Errorf("var %v is still empty. Please fill it", k)
		}
	}
	return nil
}

func readMasterplan(file string) *Masterplan {
	m := &Masterplan{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("Could not read Masterplan File %v : %v", file, err)
	}
	err = yaml.Unmarshal(b, m)
	if err != nil {
		fmt.Errorf("Could not Unmarshal YAML")
	}
	return m
}
