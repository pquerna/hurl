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

package ui

import (
	"github.com/cheggaaa/pb"
	"github.com/pquerna/hurl/common"
)

type ConsoleUI struct {
	Bar    *pb.ProgressBar
	Config common.ConfigGetter
}

func NewConsoleUI() *ConsoleUI {
	return &ConsoleUI{}
}

func (cui *ConsoleUI) WorkStart(numTodo int64) {
	cui.Bar = pb.New(0)
	cui.Bar.Total = numTodo
	cui.Bar.Start()
}

func (cui *ConsoleUI) WorkStatus(numDone int64) {
	// TODO: Add Set64 to PB? super lame.
	cui.Bar.Set(int(numDone))
}

func (cui *ConsoleUI) WorkEnd() {
	cui.Bar.FinishPrint("Complete.")
}

func (cui *ConsoleUI) ConfigSet(config common.ConfigGetter) {
	cui.Config = config
}

func (cui *ConsoleUI) ConfigGet() common.ConfigGetter {
	return cui.Config
}
