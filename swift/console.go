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

package swift

import (
	"fmt"
	"github.com/pquerna/hurl/common"
	"github.com/pquerna/hurl/workers"
	"github.com/spf13/cobra"
)

var g_config = common.SwiftConfig{}

func ConsoleCommand(ui common.UI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swift",
		Short: "swift bencharming client.",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			ConsoleRun(ui, cmd, args)
		},
	}

	g_config.AddFlags(cmd.Flags())

	return cmd
}

func ConsoleRun(ui common.UI, cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		common.ConsoleErr(cmd, fmt.Sprintf("Error: Missing mode: %s", args))
		return
	}

	// https://auth.api.rackspacecloud.com/v1.0

	g_config.Mode = args[0]

	ui.ConfigSet(&g_config)

	common.ConsoleRun(workers.Run, "swift", ui, cmd, args)
}
