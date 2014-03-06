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
	Mode        string
	Records     int64
	AtRecord    int64
	FieldCount  int
	FieldSize   int
	ReadRatio   float64
	UpdateRatio float64
	// TODO(pquerna): m
	//	InsertRatio float32
	//	ScanRatio   float32
}

func (conf *DatastoreConfig) AddFlags(flags *flag.FlagSet) {
	flags.StringVarP(&conf.Mode, "mode", "m", "load", "[load|run]")
	flags.Int64VarP(&conf.Records, "records", "", 1000000, "Number of Records.")
	flags.IntVar(&conf.FieldCount, "fieldcount", 10, "Number of Fields.")
	flags.IntVar(&conf.FieldSize, "fieldsize", 100, "Size of Fields.")
	flags.Float64Var(&conf.ReadRatio, "readratio", 0.95, "Ratio of Reads.")
	flags.Float64Var(&conf.UpdateRatio, "updateratio", 0.05, "Ratio of Updates.")
	//	flags.Float64Var(&conf.InsertRatio, "readratio", 0.0, "Ratio of Inserts.")
	//	flags.Float64Var(&conf.ScanRatio, "readratio", 0.0, "Ratio of Scan operations.")
}

func (conf *DatastoreConfig) Validate() error {

	if conf.Mode != "load" && conf.Mode != "run" {
		return errors.New("Method must be run or load.")
	}

	return nil
}

func (conf *DatastoreConfig) GetDatastoreConfig() *DatastoreConfig {
	return conf
}
