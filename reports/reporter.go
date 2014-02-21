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

package reports

import (
	"fmt"
	"github.com/pquerna/hurl/common"
	"sort"
	"time"
)

type Reporter interface {
	Priority() int
	Interest(ui common.UI, taskType string) bool
	ReadResults(*common.ResultArchiveReader)
	ConsoleOutput()
}

type ByPriority []Reporter

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority() < a[j].Priority() }

var g_reporters []Reporter

func AddReporter(r Reporter) {
	if g_reporters == nil {
		g_reporters = make([]Reporter, 0)
	}
	g_reporters = append(g_reporters, r)
}

func init() {
	AddReporter(&overview{BaseReport: BaseReport{ReportPriority: 0}, ui: nil})
}

func Run(ui common.UI, taskType string, conf common.ConfigGetter, rr *common.ResultArchiveReader) error {
	reporters := make([]Reporter, 0, len(g_reporters))

	for _, r := range g_reporters {
		if r.Interest(ui, taskType) {
			reporters = append(reporters, r)
		}
	}

	sort.Sort(ByPriority(reporters))

	// TODO: Feed requests to all reporters in parallell, only reading the file from disk once.
	for _, r := range reporters {
		r.ReadResults(rr)
		rr.Reset()
	}

	fmt.Println()
	for _, r := range reporters {
		// TODO: other output types
		r.ConsoleOutput()
	}

	return nil
}

type BaseReport struct {
	ReportPriority int
}

func (br *BaseReport) Priority() int {
	return br.ReportPriority
}

type overview struct {
	BaseReport
	ui               common.UI
	timeTaken        time.Duration
	completeRequests int64
	failedRequests   int64
}

func (o *overview) Interest(ui common.UI, taskType string) bool {
	o.ui = ui
	return true
}

func (o *overview) ReadResults(rr *common.ResultArchiveReader) {
	var first *common.Result = nil
	var last *common.Result = nil
	for rr.Scan() {
		rv := rr.Entry()

		if last == nil || rv.Start.After(last.Start) {
			last = rv
		}

		if first == nil || rv.Start.Before(first.Start) {
			first = rv
		}

	}

	o.timeTaken = last.Start.Sub(first.Start)
}

func (o *overview) ConsoleOutput() {
	conf := o.ui.ConfigGet()
	bconf := conf.GetBasicConfig()
	fmt.Printf("Concurrency Level: %d\n", bconf.Concurrency)
	fmt.Printf("Time Taken: %v\n", o.timeTaken)
}
