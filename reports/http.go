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
	"github.com/rcrowley/go-metrics"
)

func init() {
	AddReporter(&HTTPResponseSize{HTTPReport{BaseReport{ReportPriority: 99}}, nil})
}

type HTTPReport struct {
	BaseReport
}

func (ht *HTTPReport) Interest(ui common.UI, taskType string) bool {
	if taskType == "http" {
		return true
	}
	return false
}

type HTTPResponseSize struct {
	HTTPReport
	h metrics.Histogram
}

func (hrs *HTTPResponseSize) ReadResults(rr *common.ResultArchiveReader) {
	hrs.h = metrics.NewHistogram(metrics.NewExpDecaySample(1028, 0.015))

	for rr.Scan() {
		rv := rr.Entry()
		hrs.h.Update(int64(rv.Metrics["BodyLength"]))
	}
}

func (hrs *HTTPResponseSize) ConsoleOutput() {
	if hrs.h.Min() != hrs.h.Max() {
		fmt.Printf("Document Length:\n")
		fmt.Printf("				Mean		%v\n", int(hrs.h.Mean()))
		fmt.Printf("				90%%		%v\n", int(hrs.h.Percentile(0.90)))
		fmt.Printf("				95%%		%v\n", int(hrs.h.Percentile(0.95)))
		fmt.Printf("				99%%		%v\n", int(hrs.h.Percentile(0.99)))
	} else {
		fmt.Printf("Document Length:	%d\n", int(hrs.h.Max()))
	}
}
