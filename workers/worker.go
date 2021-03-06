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
	"github.com/pquerna/hurl/reports"
	"sync"
	"time"
)

type WorkerTask interface {
	Work(rv *common.Result) error
}

type Worker interface {
	Start(taskType string, wg *sync.WaitGroup, reqChan chan int64, resChan chan *common.Result) error
	Halt() error
}

type newTask func(common.UI) (WorkerTask, error)

var g_workers_tasks map[string]newTask

func init() {
	g_workers_tasks = make(map[string]newTask)
}

func Register(wt string, nt newTask) {
	g_workers_tasks[wt] = nt
}

type LocalWorker struct {
	WorkerType string
	wg         *sync.WaitGroup
	reqChan    chan int64
	resChan    chan *common.Result
	task       WorkerTask
}

func (lw *LocalWorker) runWorker(taskType string, id string) {
	defer lw.wg.Done()

	for {
		reqNum, ok := <-lw.reqChan
		if !ok {
			return
		}
		rv := common.NewResult(taskType, fmt.Sprintf("%s-%d", id, reqNum))
		rv.Start = time.Now()
		err := lw.task.Work(rv)
		rv.Done()
		if err != nil {
			rv.Error = true
			rv.Meta["error"] = err.Error()
		}
		lw.resChan <- rv
	}
}

func (lw *LocalWorker) Start(taskType string, wg *sync.WaitGroup, reqChan chan int64, resChan chan *common.Result) error {
	lw.reqChan = reqChan
	lw.resChan = resChan
	lw.wg = wg
	lw.wg.Add(1)
	go lw.runWorker(taskType, uniuri.New())

	return nil
}

func (lw *LocalWorker) Halt() error {
	return nil
}

func resultHanlder(ui common.UI, wgres *sync.WaitGroup, resChan chan *common.Result, raw *common.ResultArchiveWriter) {
	defer func() { wgres.Done() }()
	var i int64 = 0
	for {
		rv, ok := <-resChan
		if !ok {
			return
		}
		i++
		ui.WorkStatus(i)
		raw.Write(rv)
	}
}

func Run(ui common.UI, taskType string) error {
	//	if clusterConf != "" {
	//		// TODO: add ClusterWorker
	//		return nil, fmt.Errorf("TODO: Cluster support")
	//	}
	wt, ok := g_workers_tasks[taskType]
	if !ok {
		return fmt.Errorf("unknown worker type: %s", taskType)
	}

	conf := ui.ConfigGet()
	bconf := conf.GetBasicConfig()
	workers := make([]Worker, bconf.Concurrency)
	for index, _ := range workers {
		w, err := wt(ui)
		if err != nil {
			return err
		}
		workers[index] = &LocalWorker{task: w}
	}

	ui.WorkStart(bconf.NumRequests)
	defer func() {
		for _, worker := range workers {
			worker.Halt()
		}
	}()

	var wg sync.WaitGroup
	var wgres sync.WaitGroup
	reqchan := make(chan int64, 1024*1024)
	// TODO: how big should this be?
	reschan := make(chan *common.Result)
	for _, worker := range workers {
		err := worker.Start(taskType, &wg, reqchan, reschan)
		if err != nil {
			return err
		}
	}

	var i int64

	for i = 0; i < bconf.NumRequests; i++ {
		reqchan <- i
		// TOOD: ui.WorkStatus(numDone int64)
	}
	wgres.Add(1)

	rw := common.NewResultArchiveWriter()
	defer rw.Remove()
	go resultHanlder(ui, &wgres, reschan, rw)

	close(reqchan)
	wg.Wait()
	close(reschan)
	wgres.Wait()
	ui.WorkEnd()

	rw.Close()

	rr := common.NewResultArchiveReader(rw.Path)
	return reports.Run(ui, taskType, conf, rr)
}
