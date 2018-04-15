package namaste

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type Namaste interface {
	GetBlueprints() []*Blueprint
	CreateNewProject(bp *Blueprint, dir string) error
}

type namaste struct {
	RootDir    string
	blueprints []*Blueprint
}

type Blueprint struct {
	Name string

	Masterplan *Masterplan

	dir string

	files []string

	folders []string
}

func (n *namaste) CreateNewProject(bp *Blueprint, dir string) error {
	template.New("File").ParseFiles(bp.files)
	return nil
}

func (n *namaste) GetBlueprints() []*Blueprint {
	return n.blueprints
}

func (n *namaste) newBlueprint(path string, file os.FileInfo) {
	bp := &Blueprint{dir: path, Name: file.Name}
	bp.Masterplan = readMasterplan(path + "/masterplan.yaml")
	n.blueprints = append(n.blueprints, bp)
}

func Initialize(root string) (*namaste, error) {
	n := &namaste{RootDir: root}
	if !checkIfDirectoryExists(root) {
		err := os.Mkdir(root, 0755)
		if err != nil {
			return n, fmt.Errorf("Could not Create Directory: %v", err)
		}
	}
	n.scanForBlueprints()
	return n, nil
}

func checkIfDirectoryExists(dir string) bool {
	fn, err := ioutil.ReadDir(dir)
	if err != nil || len(fn) == 0 {
		return false
	}
	return true
}

func (n *namaste) scanForBlueprints() {
	e := filepath.Walk(n.RootDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Could not Walk Path: %v", err)
			return err
		}
		if f.IsDir() && dirContainsBlueprintConfig(path) {
			n.newBlueprint(path, f.Name())
		}
		return err
	})
	if e != nil {
		fmt.Errorf("Could not Walk FilePath. There is something wrong : %v", e)
	}
}

func dirContainsBlueprintConfig(dir string) bool {
	fi := readDirectory(dir)
	for _, f := range fi {
		if f.Name() == "masterplan.yaml" {
			return true
		}
	}
	return false
}

func readDirectory(dir string) []os.FileInfo {
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Errorf("Could not Read Blueprint Directory. Make sure Blueprint Directory exist")
	}
	return fi
}
