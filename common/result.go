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
	"time"
)

type Result struct {
	Type     string
	Id       string
	Error    bool
	Duration time.Duration
	// TODO: should we just do a map[string]interface{}?
	Meta    map[string]string
	Metrics map[string]float64
}

func NewResult(taskType string, id string) *Result {
	r := &Result{Id: id, Type: taskType}
	r.Meta = make(map[string]string)
	r.Metrics = make(map[string]float64)
	return r
}
