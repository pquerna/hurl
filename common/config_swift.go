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

type SwiftConfig struct {
	BasicConfig
	DatastoreConfig
	Username string
	ApiKey   string
	AuthUrl  string
	Region   string
}

func (conf *SwiftConfig) AddFlags(flags *flag.FlagSet) {
	conf.BasicConfig.AddFlags(flags)
	conf.DatastoreConfig.AddFlags(flags)
	flags.StringVarP(&conf.Username, "username", "u", "", "Username")
	flags.StringVarP(&conf.ApiKey, "apikey", "k", "", "API Key")
	flags.StringVarP(&conf.AuthUrl, "authurl", "a", "https://identity.api.rackspacecloud.com/v2.0", "Authenication URL")
	flags.StringVarP(&conf.Region, "region", "r", "IAD", "Target Region")
}

func (conf *SwiftConfig) Validate() error {

	err := conf.BasicConfig.Validate()

	if err != nil {
		return err
	}

	err = conf.DatastoreConfig.Validate()

	if err != nil {
		return err
	}

	if conf.Username == "" {
		return errors.New("Username missing.")
	}

	if conf.ApiKey == "" {
		return errors.New("API Key missing.")
	}

	if conf.Region == "" {
		return errors.New("Region missing.")
	}

	return nil
}

func (conf *SwiftConfig) GetSwiftConfig() *SwiftConfig {
	return conf
}

func (conf *SwiftConfig) GetDatastoreConfig() *DatastoreConfig {
	return &conf.DatastoreConfig
}
