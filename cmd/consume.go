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
	"encoding/xml"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xoviat/jcad/lib"
)

// consumeCmd represents the importLibrary command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume an eagle library.",
	Long:  `Consume the symbols and footprints in an eagle library.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		src := args[0]

		library, _ := lib.NewDefaultLibrary()

		fpsrc, err := os.Open(src)
		if err != nil {
			fmt.Printf("failed to open file: %s\n", err)
			return
		}

		elibrary := lib.EagleLibrary{}
		dec := xml.NewDecoder(fpsrc)
		err = dec.Decode(&elibrary)
		if err != nil {
			fmt.Printf("failed to decode library: %s\n", err)
			return
		}

		for _, pkg := range elibrary.Packages {
			fmt.Println("importing package: " + pkg.Name)
		}

		for _, symbol := range elibrary.Symbols {
			fmt.Println("importing symbol: " + symbol.Name)
		}

		library.AddPackages(elibrary.Packages)
		library.AddSymbols(elibrary.Symbols)
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
