/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/file"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/tool"

	"github.com/spf13/cobra"
)

// uniondiffCmd represents the uniondiff command
var uniondiffCmd = &cobra.Command{
	Use:   "uniondiff",
	Short: "merge the diff keys of json1 and json2",
	Long: `the union of keys in json1 not in json2 ,and keys in json2 not in json1
./jsonToIdl uniondiff -l ../t.json -r ../t.json
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("uniondiff called")
		if jsonName1=="" || jsonName2==""{
			fmt.Println("two json file is required",jsonName1,jsonName2)
			return
		}
		j1:=file.GetJson(jsonName1)
		j2:=file.GetJson(jsonName2)
		file.Save("uniondiff_"+ getName(jsonName1)+"_"+ getName(jsonName2),tool.UnionDiffJson(string(j1),string(j2)))

	},
}

func init() {
	rootCmd.AddCommand(uniondiffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uniondiffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uniondiffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	uniondiffCmd.Flags().StringVarP(&jsonName1, "json1", "l", "", "json name (required)")
	uniondiffCmd.MarkFlagRequired("json1")
	uniondiffCmd.Flags().StringVarP(&jsonName2, "json2", "r", "", "json name (required)")
	uniondiffCmd.MarkFlagRequired("json2")
}
