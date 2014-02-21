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
	"net"
	"net/url"
)

type EtcdConfig struct {
	BasicConfig
}

func (conf *EtcdConfig) Validate() error {

	err := conf.BasicConfig.Validate()

	if err != nil {
		return err
	}

	u, err := url.Parse(conf.Url)

	if err != nil {
		return err
	}

	if u.IsAbs() != true {
		return errors.New("URL must include scheme. (hint, http:// missing?)")
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return errors.New("URL Scheme must be http or https.")
	}

	var host = u.Host

	if HasPort(u.Host) {
		host, _, err = net.SplitHostPort(u.Host)

		if err != nil {
			return err
		}
	}

	_, err = net.LookupIP(host)
	if err != nil {
		return err
	}

	return nil
}

func (conf *EtcdConfig) GetEtcdConfig() *EtcdConfig {
	return conf
}
