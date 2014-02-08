/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package main

import (
	"github.com/pquerna/hurl/ui"
	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Use:   "hurl",
		Short: "hurl is a tool to hurl traffic at URLs",
		Long: `hurl is a flexiable benchmarking tool with the ability to scale out.

Complete documentation is available online at:
	https://github.com/pquerna/hurl/`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.UsageFunc()(cmd)
		},
	}

	subcmds := ui.ConsoleCommands()

	for _, c := range subcmds {
		cmd.AddCommand(c)
	}

	cmd.Execute()
}
