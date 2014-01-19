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
	"github.com/spf13/cobra"
)

var g_config = Config{}

func ConsoleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "smash",
		Short: "Send basic HTTP or HTTPS traffic to a server.",
		Long:  ``,
		Run:   ConsoleRun,
	}

	flags := cmd.Flags()
	flags.StringVarP(&g_config.method, "method", "m", "GET", "HTTP method to use.")
	flags.IntVarP(&g_config.numRequests, "numrequests", "n", 1000, "Number of requests")
	flags.BoolVarP(&g_config.keepalive, "keepalive", "k", true, "Enable HTTP Keepalive")

	return cmd
}

func ConsoleRun(cmd *cobra.Command, args []string) {
	// Assemble Smashers.
	// Distribute load.
	// RUN IT
	// Get Reports.
	// Render Reports.
}
