package namaste

import (
	"fmt"
	"testing"
)

var testDir = "./testdata/"

func TestInitialization(t *testing.T) {
	g, err := Initialize(testDir)
	if err != nil {
		t.Errorf("Could not initialize: %v", err)
	}

	if g.RootDir != testDir {
		t.Errorf("Could not Initialize Root Directory")
	}

	if len(g.blueprints) == 0 {
		t.Errorf("Testdata could not be loaded")
	} else {
		fmt.Printf("%v Blueprints could be found, first ones Name is %v ", len(g.blueprints), g.blueprints[0].Name)
	}

}

func TestInitializationWithCreateDir(t *testing.T) {
	_, err := Initialize("~/.namaste")
	if err != nil {
		t.Errorf("Could not create new Directory: %v", err)
	}

}

func TestReadDirectory(t *testing.T) {
	fi := readDirectory(testDir)
	if len(fi) < 1 {
		t.Error("Could not Read Testdirectory")
	}
}

func TestDirContainsBlueprint(t *testing.T) {
	b := dirContainsBlueprintConfig(testDir + "nodejs/fes/")
	if !b {
		t.Errorf("Cant find the Masterplan")
	}
}
