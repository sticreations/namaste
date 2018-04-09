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
