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
)

type Reporter interface {
	Interest(ui common.UI, taskType string) bool
	ReadResults(*common.ResultArchiveReader)
	ConsoleOutput()
}

var g_reporters []Reporter

func AddReporter(r Reporter) {
	if g_reporters == nil {
		g_reporters = make([]Reporter, 0)
	}
	g_reporters = append(g_reporters, r)
}

func init() {
	AddReporter(&overview{})
}

func Run(ui common.UI, taskType string, conf common.ConfigGetter, rr *common.ResultArchiveReader) error {
	reporters := make([]Reporter, 0, len(g_reporters))

	for _, r := range g_reporters {
		if r.Interest(ui, taskType) {
			reporters = append(reporters, r)
		}
	}

	for _, r := range reporters {
		r.ReadResults(rr)
		rr.Reset()
	}

	for _, r := range reporters {
		// TODO: other output types
		r.ConsoleOutput()
	}

	return nil
}

type overview struct {
	ui common.UI
}

func (o *overview) Interest(ui common.UI, taskType string) bool {
	o.ui = ui
	return true
}

func (o *overview) ReadResults(rr *common.ResultArchiveReader) {
}

func (o *overview) ConsoleOutput() {
	/*
	   Concurrency Level:      1
	   Time taken for tests:   0.010 seconds
	   Complete requests:      1
	   Failed requests:        0
	   Write errors:           0
	*/
	conf := o.ui.ConfigGet()
	bconf := conf.GetBasicConfig()
	fmt.Printf("Concurrency Level: %d\n", bconf.Concurrency)
}
