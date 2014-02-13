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

package common

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type RunCmd func(ui UI, taskType string) error

func ConsoleErr(cmd *cobra.Command, str string) {
	cmd.Printf(str)
	cmd.Println("")
	cmd.UsageFunc()(cmd)
	os.Exit(1)
}

func ConsoleRun(runner RunCmd, workerType string, ui UI, cmd *cobra.Command, args []string) {
	conf := ui.ConfigGet()

	err := conf.Validate()
	if err != nil {
		ConsoleErr(cmd, fmt.Sprintf("Error: %s", err))
		return
	}

	err = runner(ui, workerType)
	if err != nil {
		ConsoleErr(cmd, fmt.Sprintf("Error: %s", err))
		return
	}
}
