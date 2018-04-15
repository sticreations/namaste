package namaste

import (
	"fmt"
	"os/user"
	"testing"
)

var testDir = "./testdata/"

func TestGetBlueprints(t *testing.T) {
	var n Namaste
	n, err := Initialize(testDir)
	if err != nil {
		t.Errorf("Could not initialize Project: %v", err)
	}
	bp := n.GetBlueprints()

	if !(len(bp) > 0) {
		t.Errorf("There are no Blueprints initialized")
	}

}

func TestCreateNewProject(t *testing.T) {
	n, err := Initialize(testDir)
	if err != nil {
		t.Error("Could not initialize")
	}
	bp := n.GetBlueprints()

	n.CreateNewProject(bp[0], "./testdata/generated")

}

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
	}
}

func TestInitializationWithCreateDir(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		fmt.Errorf("Could not get User: %v", err)
	}
	_, err = Initialize(usr.HomeDir + "/.namaste/")
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
	b := dirContainsBlueprintConfig(testDir + "nodejs/testservice/")
	if !b {
		t.Errorf("Cant find the Masterplan")
	}
}
