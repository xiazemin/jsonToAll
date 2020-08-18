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
	"github.com/spf13/cobra"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/file"
	"github.com/xiazemin/json-parser/antlr/jsonToAll/tool"
)

// unionCmd represents the union command
var unionCmd = &cobra.Command{
	Use:   "union",
	Short: "union two json file ",
	Long: `union the keys of two json
./jsonToIdl union -l ../t.json -r ../t.json
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("union called")
		if jsonName1=="" || jsonName2==""{
			fmt.Println("two json file is required",jsonName1,jsonName2)
			return
		}
		j1:=file.GetJson(jsonName1)
		j2:=file.GetJson(jsonName2)
		file.Save("union_"+ getName(jsonName1)+"_"+ getName(jsonName2),tool.UnionJSON(string(j1),string(j2)))
	},
}

func init() {
	rootCmd.AddCommand(unionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	unionCmd.Flags().StringVarP(&jsonName1, "json1", "l", "", "json name (required)")
	unionCmd.MarkFlagRequired("json1")
	unionCmd.Flags().StringVarP(&jsonName2, "json2", "r", "", "json name (required)")
	unionCmd.MarkFlagRequired("json2")
}
