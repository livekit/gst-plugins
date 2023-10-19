// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Testsink struct {
	*base.GstBaseSink
}

func NewTestsink() (*Testsink, error) {
	e, err := gst.NewElement("testsink")
	if err != nil {
		return nil, err
	}
	return &Testsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewTestsinkWithName(name string) (*Testsink, error) {
	e, err := gst.NewElementWithName("testsink", name)
	if err != nil {
		return nil, err
	}
	return &Testsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// allowed-timestamp-deviation (int64)
//
// allowed average difference in usec between timestamp of next buffer and expected timestamp from analyzing last buffer

func (e *Testsink) GetAllowedTimestampDeviation() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("allowed-timestamp-deviation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Testsink) SetAllowedTimestampDeviation(value int64) error {
	return e.Element.SetProperty("allowed-timestamp-deviation", value)
}

// buffer-count (int64)
//
// number of buffers in stream

func (e *Testsink) GetBufferCount() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// expected-buffer-count (int64)
//
// expected number of buffers in stream

func (e *Testsink) GetExpectedBufferCount() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("expected-buffer-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Testsink) SetExpectedBufferCount(value int64) error {
	return e.Element.SetProperty("expected-buffer-count", value)
}

// expected-length (int64)
//
// expected length of stream

func (e *Testsink) GetExpectedLength() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("expected-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Testsink) SetExpectedLength(value int64) error {
	return e.Element.SetProperty("expected-length", value)
}

// expected-md5 (string)
//
// expected md5 of processing the whole data

func (e *Testsink) GetExpectedMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("expected-md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Testsink) SetExpectedMd5(value string) error {
	return e.Element.SetProperty("expected-md5", value)
}

// length (int64)
//
// length of stream

func (e *Testsink) GetLength() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// md5 (string)
//
// md5 of processing the whole data

func (e *Testsink) GetMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// timestamp-deviation (int64)
//
// average difference in usec between timestamp of next buffer and expected timestamp from analyzing last buffer

func (e *Testsink) GetTimestampDeviation() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp-deviation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

