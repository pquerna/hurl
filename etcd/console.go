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

package etcd

import (
	"fmt"
	"github.com/pquerna/hurl/common"
	"github.com/pquerna/hurl/workers"
	"github.com/spf13/cobra"
)

var g_config = common.EtcdConfig{}

func ConsoleCommand(ui common.UI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etcd",
		Short: "etcd bencharming client.",
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
		common.ConsoleErr(cmd, fmt.Sprintf("Error: Expected 1 URL, got: %s", args))
		return
	}

	g_config.Url = args[0]

	ui.ConfigSet(&g_config)

	common.ConsoleRun(workers.Run, "etcd", ui, cmd, args)
}
