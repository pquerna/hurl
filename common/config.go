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
	flag "github.com/spf13/pflag"
)

type ConfigGetter interface {
	// TODO(pquerna): This interface sucks. IT REALLY SUCKS. FIX ME.
	GetBasicConfig() *BasicConfig
	GetDatastoreConfig() *DatastoreConfig
	GetHttpConfig() *HttpConfig
	GetEtcdConfig() *EtcdConfig
	GetSwiftConfig() *SwiftConfig
	Validate() error
}

type BasicConfig struct {
	Url         string
	Cluster     string
	NumRequests int64
	Concurrency int
}

func (conf *BasicConfig) AddFlags(flags *flag.FlagSet) {
	flags.Int64VarP(&conf.NumRequests, "numrequests", "n", 10000, "Number of requests.")
	flags.IntVarP(&conf.Concurrency, "concurrency", "c", 100, "Number of concurrent workers.")
	flags.StringVarP(&conf.Cluster, "cluster", "", "", "Peer nodes to use, if none, use this process.")
}

func (conf *BasicConfig) Validate() error {

	if conf.Concurrency > 250000 {
		return fmt.Errorf("concurrency of %d is unlikely to work well.", conf.Concurrency)
	}

	if conf.NumRequests < 0 {
		return fmt.Errorf("numrequests of %d is less than zero.", conf.NumRequests)
	}

	return nil
}

func (conf *BasicConfig) GetBasicConfig() *BasicConfig {
	return conf
}

func (conf *BasicConfig) GetHttpConfig() *HttpConfig {
	return nil
}

func (conf *BasicConfig) GetEtcdConfig() *EtcdConfig {
	return nil
}

func (conf *BasicConfig) GetSwiftConfig() *SwiftConfig {
	return nil
}

func (conf *BasicConfig) GetDatastoreConfig() *DatastoreConfig {
	return nil
}
