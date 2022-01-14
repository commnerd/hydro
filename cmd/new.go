/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var PathSeparatorString = string(os.PathSeparator)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [repo]",
	Short: "Generate a new hydro repository",
	Long:  `Generate a new hydro repository`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 1:
			return initialize(cmd, args)
		}

		return cmd.Help()
	},
}

var initialDirs = [...]string{
	"charts",
	"releases",
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.DisableFlagsInUseLine = true

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initDirs(path string) error {
	err := makeDir(path)
	if err != nil {
		return err
	}

	for _, dir := range initialDirs {
		curPath := path + PathSeparatorString + dir
		err = makeDir(curPath)
		if err != nil {
			return err
		}

		makeEmptyFile(curPath + PathSeparatorString + ".gitkeep")
	}

	return nil
}

func initialize(cmd *cobra.Command, args []string) error {
	err := initDirs(args[0])
	if err != nil {
		return err
	}

	return nil
}

func makeDir(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return os.Mkdir(path, os.ModePerm)
	}

	return nil
}

func makeEmptyFile(path string) error {
	emptyFile, err := os.Create(path)
	if err != nil {
		return err
	}
	return emptyFile.Close()
}
