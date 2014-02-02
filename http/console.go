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

package http

import (
	"fmt"
	"github.com/pquerna/hurl/common"
	"github.com/pquerna/hurl/workers"
	"github.com/spf13/cobra"
)

var g_config = common.HttpConfig{}

func ConsoleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Send basic HTTP or HTTPS traffic to a server.",
		Long:  ``,
		Run:   ConsoleRun,
	}

	g_config.AddFlags(cmd.Flags())

	return cmd
}

func ConsoleRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		common.ConsoleErr(cmd, fmt.Sprintf("Error: Expected 1 URL, got: %s", args))
		return
	}

	g_config.Url = args[0]

	err := g_config.Validate()
	if err != nil {
		common.ConsoleErr(cmd, fmt.Sprintf("Error: %s", err))
		return
	}

	err = workers.Run("http", &g_config)
	if err != nil {
		common.ConsoleErr(cmd, fmt.Sprintf("Error: %s", err))
		return
	}

	// Get Reports.
	// Render Reports.
}
