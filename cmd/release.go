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
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "release",
	Short: "Add a release to your repo",
	Long: `This command allows you to add a new release to your repo.
It will first analyze the name of the chart.  If the chart is passed in
"repo/chart" format, it will assume it needs to rely on one of the repos
in the repo manifest.  Else, it will look in your charts directory to
generate the chart with.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	addCmd.AddCommand(repoCmd)
	repoCmd.DisableFlagsInUseLine = true

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
