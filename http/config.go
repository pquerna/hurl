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
	"github.com/pquerna/hurl/common"
	flag "github.com/spf13/pflag"
)

type Config struct {
	common.BasicConfig
	method    string
	url       string
	keepalive bool
}

func (config *Config) AddFlags(flags *flag.FlagSet) {
	config.BasicConfig.AddFlags(flags)
	flags.StringVarP(&config.method, "method", "m", "GET", "HTTP method to use.")
	flags.BoolVarP(&config.keepalive, "keepalive", "k", true, "Enable HTTP Keepalive")
}
