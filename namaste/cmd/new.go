// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os/user"

	"github.com/manifoldco/promptui"
	"github.com/sticreations/goject"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new Project based on your Blueprint",
	Long:  `I already said what it did. creates projects!`,
	Run: func(cmd *cobra.Command, args []string) {
		e := generatePrompt()
		if e != nil {
			fmt.Printf("something went terribly wrong")

		}
	},
}

func generatePrompt() error {
	usr, err := user.Current()
	if err != nil {
		fmt.Errorf("Could not get User: %v", err)
	}
	nam, err := namaste.Initialize(usr.HomeDir + "/.namaste/")
	if err != nil {
		return fmt.Errorf("Im very sorry. But you have no blueprints in your Path")
	}

	bp := nam.GetBlueprints()
	bpMap := map[string]*namaste.Blueprint{}
	var names []string
	for _, b := range bp {
		bpMap[b.Name] = b
		names = append(names, b.Name)
	}
	prompt := promptui.Select{
		Label: "Select Project you want to generate",
		Items: names,
	}
	_, res, err := prompt.Run()
	if err != nil {
		fmt.Errorf("Could not Read User Input: %v", err)
	}
	fmt.Printf(res)
	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
