/*
 * Copyright © 2021 Andrew Roth <roth.andy@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/rothandrew/convert-to-helm-raw-values/pkg/yaml"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "convert-to-helm-raw-values",
		Short: "Converts K8s YAML to a version that helm incubator/raw can use",
		Long:  "Converts K8s YAML to a version that helm incubator/raw can use",
		Run:   runRoot,
	}
	infile  string
	outfile string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(
		&infile,
		"infile",
		"i",
		"",
		"Input file - Must be compliant with K8s YAML",
	)
	_ = rootCmd.MarkFlagRequired("infile")
	rootCmd.Flags().StringVarP(
		&outfile,
		"outfile",
		"o",
		"",
		`Output file - Will be formatted such that it can be used as values.yaml for helm chart incubator/raw.
Will always overwrite if the file already exists.`,
	)
	_ = rootCmd.MarkFlagRequired("outfile")
}

func runRoot(_ *cobra.Command, _ []string) {
	// Read the infile
	infileContent, err := ioutil.ReadFile(infile)
	if err != nil {
		log.Fatalf("ioutil.ReadFile() failed with '%s'\n", err)
	}

	// Convert to new format
	outfileContent, err := yaml.ConvertToHelmRawValues(string(infileContent))
	if err != nil {
		log.Fatalf("yaml.ConvertToHelmRawValues() failed with '%s'\n", err)
	}

	// Write to outfile
	err = ioutil.WriteFile(outfile, []byte(outfileContent), 0644)
	if err != nil {
		log.Fatalf("ioutil.WriteFile() failed with '%s'\n", err)
	}
}
