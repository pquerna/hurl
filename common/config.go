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
	"github.com/spf13/cobra"
)

type Config struct {
	numRequests int
	concurrency int
}

func (config *Config) AddCommonFlags(cmd *cobra.Command) {
	flags := cmd.Flags()
	flags.IntVarP(&config.numRequests, "numrequests", "n", 10000, "Number of requests.")
	flags.IntVarP(&config.concurrency, "concurrency", "c", 100, "Number of concurrent requests.")
}
