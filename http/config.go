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
	"errors"
	"fmt"
	"github.com/pquerna/hurl/common"
	flag "github.com/spf13/pflag"
	"net/url"
)

type Config struct {
	common.BasicConfig
	Method    string
	Keepalive bool
}

func (conf *Config) AddFlags(flags *flag.FlagSet) {
	conf.BasicConfig.AddFlags(flags)
	flags.StringVarP(&conf.Method, "method", "m", "GET", "HTTP method to use.")
	flags.BoolVarP(&conf.Keepalive, "keepalive", "k", true, "Enable HTTP Keepalive")
}

func (conf *Config) Validate() error {
	u, err := url.Parse(conf.Url)

	if err != nil {
		return err
	}

	if u.IsAbs() != true {
		return errors.New("URL must include scheme. (hint, http:// missing?)")
	}

	if conf.Concurrency > 250000 {
		return fmt.Errorf("Concurrency of %d is unlikely to work well. Consider scaling out configuration?", conf.Concurrency)
	}

	return nil
}
