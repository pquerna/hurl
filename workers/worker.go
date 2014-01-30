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

package workers

import (
	"fmt"
	"github.com/pquerna/hurl/common"
)

type ResultSaver interface {
	SaveRecord(string) error
}

type Worker interface {
	Start() (string, error)
	Halt(string) error
	Status(string) (string, error)
	Results(string, ResultSaver) error
}

type newWorker func(common.ConfigGetter) (Worker)

func Run(conf *common.BasicConfig) (error) {
//	if clusterConf != "" {
//		// TODO: add ClusterWorker
//		return nil, fmt.Errorf("TODO: Cluster support")
//	}
//	return []Worker{&LocalWorker{WorkerType: "blah"}}, nil
	return fmt.Errorf("NOT IMPLEMENTED")
}

var g_workers_types map[string]newWorker

func init() {
	g_workers_types = make(map[string]newWorker)
}

func Register(wt string,  nw newWorker) {
	g_workers_types[wt] = nw
}

type LocalWorker struct {
	WorkerType string
}

func (lw *LocalWorker) Start() (string, error) {
	return "", fmt.Errorf("unknown worker type: %s", lw.WorkerType)
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
