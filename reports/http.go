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

func init() {
	AddReporter(&HTTPResponseSize{})
}

type HTTPReport struct{}

func (ht *HTTPReport) Interest(ui common.UI, taskType string) bool {
	if taskType == "http" {
		return true
	}
	return false
}

type HTTPResponseSize struct {
	HTTPReport
}

func (hrs *HTTPResponseSize) ReadResults(rr *common.ResultArchiveReader) {
	for rr.Scan() {
		rv := rr.Entry()
		fmt.Printf("ResponseSize: %v\n", rv.Metrics["BodyLength"])
	}
}

func (hrs *HTTPResponseSize) ConsoleOutput() {
	fmt.Println("HTTPResponseSize!!!!")
}
