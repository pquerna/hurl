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
	"bufio"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Result struct {
	Type     string
	Id       string
	Error    bool
	Start    time.Time
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

type ResultArchiveWriter struct {
	Path    string
	fwriter *os.File
	gwriter *gzip.Writer
	writer  *bufio.Writer
}

func (raw *ResultArchiveWriter) Write(rv *Result) error {
	b, err := json.Marshal(rv)
	if err != nil {
		return err
	}
	raw.writer.Write(b)
	raw.writer.WriteString("\n")
	return nil
}

func (raw *ResultArchiveWriter) Close() error {
	raw.writer.Flush()
	raw.gwriter.Close()
	return raw.fwriter.Close()
}

func (raw *ResultArchiveWriter) Remove() error {
	return os.Remove(raw.Path)
}

func NewResultArchiveWriter() *ResultArchiveWriter {
	tfile, err := ioutil.TempFile("", "hurlgz")
	if err != nil {
		panic(err)
	}
	gwriter := gzip.NewWriter(tfile)
	return &ResultArchiveWriter{
		Path:    tfile.Name(),
		fwriter: tfile,
		gwriter: gwriter,
		writer:  bufio.NewWriter(gwriter),
	}
}

type ResultArchiveReader struct {
	Path    string
	rrfile  *os.File
	gfile   *gzip.Reader
	scanner *bufio.Scanner
}

func NewResultArchiveReader(path string) *ResultArchiveReader {
	rar := &ResultArchiveReader{Path: path}
	err := rar.open(path)
	if err != nil {
		panic(err)
	}

	return rar
}

func (rar *ResultArchiveReader) Close() error {
	rar.scanner = nil
	if rar.gfile != nil {
		rar.gfile.Close()
		rar.gfile = nil
	}
	if rar.rrfile != nil {
		rar.rrfile.Close()
		rar.rrfile = nil
	}
	return nil
}

func (rr *ResultArchiveReader) Reset() {
	rr.Close()
	rr.open(rr.Path)
}

func (rr *ResultArchiveReader) open(path string) error {
	var err error
	rr.rrfile, err = os.Open(path)
	if err != nil {
		return err
	}

	rr.gfile, err = gzip.NewReader(rr.rrfile)
	if err != nil {
		return err
	}

	rr.scanner = bufio.NewScanner(rr.gfile)
	rr.scanner.Split(bufio.ScanLines)
	return nil
}

func (rr *ResultArchiveReader) Entry() *Result {
	b := rr.scanner.Bytes()
	rv := &Result{}
	json.Unmarshal(b, rv)
	return rv
}

func (rr *ResultArchiveReader) Scan() bool {
	return rr.scanner.Scan()
}
