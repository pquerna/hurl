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
	"github.com/pquerna/hurl/workers"
)

func RunWorkers(conf *Config) error {
	err := workers.Run(&conf.BasicConfig)
	if err != nil {
		return err
	}
	// TODO: results holder
	return nil
}

func init() {
	workers.Register("http-v1", NewWorker)
}

type Worker struct {
	workers.LocalWorker
	conf *Config
}

func NewWorker() *Worker {
	return &Worker{}
}
