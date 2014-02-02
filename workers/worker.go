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
	"github.com/dchest/uniuri"
	"github.com/pquerna/hurl/common"
	"sync"
	"time"
)

type ResultSaver interface {
	SaveRecord(string) error
}

type WorkerTask interface {
	Work() (*common.Result, error)
}

type Worker interface {
	Start(wg *sync.WaitGroup, NumRequests int64) error
	Halt() error
}

type newTask func(common.ConfigGetter) WorkerTask

var g_workers_tasks map[string]newTask

func init() {
	g_workers_tasks = make(map[string]newTask)
}

func Register(wt string, nt newTask) {
	g_workers_tasks[wt] = nt
}

type LocalWorker struct {
	WorkerType  string
	wg          *sync.WaitGroup
	NumRequests int64
	task        WorkerTask
	rs          ResultSaver
}

func (lw *LocalWorker) runWorker(id string) {
	defer lw.wg.Done()
	var i int64

	for i = 0; i < lw.NumRequests; i++ {
		startTime := time.Now()
		results, err := lw.task.Work()
		duration := time.Since(startTime)
		if err != nil {
			// TODO: report this back to UI in a better way? Convert to ErrorResult?
			panic(err)
			return
		}

		if results.Id == "" {
			results.Id = fmt.Sprintf("%s-%d", id, i)
		}
		results.Duration = duration
		// TODO: results storage
	}
}

func (lw *LocalWorker) Start(wg *sync.WaitGroup, numRequests int64) error {
	lw.NumRequests = numRequests
	lw.wg = wg
	lw.wg.Add(1)
	go lw.runWorker(uniuri.New())

	return nil
}

func (lw *LocalWorker) Halt() error {
	return nil
}

func Run(task string, conf common.ConfigGetter) error {
	//	if clusterConf != "" {
	//		// TODO: add ClusterWorker
	//		return nil, fmt.Errorf("TODO: Cluster support")
	//	}
	wt, ok := g_workers_tasks[task]
	if !ok {
		return fmt.Errorf("unknown worker type: %s", task)
	}

	bconf := conf.GetBasicConfig()
	workers := make([]Worker, bconf.Concurrency)
	for index, _ := range workers {
		workers[index] = &LocalWorker{task: wt(conf)}
	}

	defer func() {
		for _, worker := range workers {
			worker.Halt()
		}
	}()

	var wg sync.WaitGroup
	requestsPerWorker := int64(float64(bconf.NumRequests) / float64(bconf.Concurrency))

	for _, worker := range workers {
		err := worker.Start(&wg, requestsPerWorker)
		if err != nil {
			return err
		}
	}

	wg.Wait()

	return nil
}
