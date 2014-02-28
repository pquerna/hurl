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

package swift

import (
	"fmt"
	"github.com/ncw/swift"
	"github.com/pquerna/hurl/common"
	"github.com/pquerna/hurl/workers"
)

func init() {
	workers.Register("swift", NewTask)
}

type Task struct {
	conf *common.SwiftConfig
	conn *swift.Connection
}

func NewTask(ui common.UI) (workers.WorkerTask, error) {
	c := ui.ConfigGet()
	conf := c.GetSwiftConfig()
	if conf == nil {
		panic("Invalid Configuration object for swift worker")
	}

	conn := swift.Connection{
		UserName: conf.Username,
		ApiKey:   conf.ApiKey,
		AuthUrl:  conf.AuthUrl,
		// https://auth.api.rackspacecloud.com/v1.0
		Region:    conf.Region,
		UserAgent: fmt.Sprintf("hurl/1 http load tester; https://github.com/pquerna/hurl;  username=%s", conf.Username),
	}

	err := conn.Authenticate()
	if err != nil {
		return nil, err
	}

	return &Task{conf: conf, conn: &conn}, nil
}

func (t *Task) Work(rv *common.Result) error {
	return nil
}
