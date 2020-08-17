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
	"github.com/xiazemin/json-parser/antlr/jsonToAll/generator"
	"strings"
)

// toIdlCmd represents the toIdl command
var jsonName string
var dstDir string

var toIdlCmd = &cobra.Command{
	Use:   "toIdl",
	Short: " src.json " ,
	Long: `generate idl from json with the same prefix
when define idl from complex json, it's a hard work,
   so we can use this tool to improve your efficiency`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("toIdl called")
        if jsonName==""{
        	fmt.Println("src json file is required")
			return
		}
		na:=strings.Split(jsonName,"/")
		if len(na)<1{
			fmt.Println("src json file name is not valid")
			return
		}
		name:= na[len(na)-1]

		_,strIdlGen,subStructs:=generator.Gen(jsonName,"idl")
		//fmt.Println(strIdlOri)
		//log.Println(strIdlGen)
		//fmt.Println(subStructs,len(subStructs))
		if dstDir==""{
			file.PutIdl("./"+name+".idl",subStructs,strIdlGen)
		}else{
			file.PutIdl( dstDir+"/"+name+".idl",subStructs,strIdlGen)
		}

	},
}


func init() {
	rootCmd.AddCommand(toIdlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toIdlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toIdlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	toIdlCmd.Flags().StringVarP(&jsonName, "json", "j", "", "json name (required)")
	rootCmd.MarkFlagRequired("json")

	toIdlCmd.Flags().StringVarP(&dstDir, "dst", "d", ".", "dest dir (optional)")
}
