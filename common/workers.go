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
)

type ResultSaver interface {
	SaveRecord(interface{}) error
}

type Worker interface {
	Start(interface{}) (string, error)
	Halt(string) error
	Status(string) (string, error)
	Results(string, ResultSaver) error
}

func AssembleWorkers(conf *BasicConfig) ([]Worker, error) {
	if conf.Cluster != "" {
		// TODO: add ClusterWorker
		return nil, fmt.Errorf("TODO: Cluster support")
	}

	return []Worker{&LocalWorker{}}, nil
}

type LocalWorker struct{}

func (lw *LocalWorker) Start(conf interface{}) (string, error) {
	return "", nil
}

func (lw *LocalWorker) Halt(id string) error {
	return nil
}

func (lw *LocalWorker) Status(id string) (string, error) {
	return "TODO", nil
}

func (lw *LocalWorker) Results(id string, rs ResultSaver) error {
	return fmt.Errorf("TODO: get results")
}
