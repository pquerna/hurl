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
	"errors"
	flag "github.com/spf13/pflag"
)

type DatastoreConfig struct {
	Method string
}

func (conf *DatastoreConfig) AddFlags(flags *flag.FlagSet) {
	flags.StringVarP(&conf.Method, "mode", "m", "load", "[load|run]")
}

func (conf *DatastoreConfig) Validate() error {

	if conf.Method != "load" && conf.Method != "run" {
		return errors.New("Method must be run or load.")
	}

	return nil
}
