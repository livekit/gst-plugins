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

type Progressreport struct {
	*base.GstBaseTransform
}

func NewProgressreport() (*Progressreport, error) {
	e, err := gst.NewElement("progressreport")
	if err != nil {
		return nil, err
	}
	return &Progressreport{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewProgressreportWithName(name string) (*Progressreport, error) {
	e, err := gst.NewElementWithName("progressreport", name)
	if err != nil {
		return nil, err
	}
	return &Progressreport{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// do-query (bool)
//
// Use a query instead of buffer metadata to determine stream position

func (e *Progressreport) GetDoQuery() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-query")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Progressreport) SetDoQuery(value bool) error {
	return e.Element.SetProperty("do-query", value)
}

// format (string)
//
// Format to use for the querying

func (e *Progressreport) GetFormat() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Progressreport) SetFormat(value string) error {
	return e.Element.SetProperty("format", value)
}

// silent (bool)
//
// Do not print output to stdout

func (e *Progressreport) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Progressreport) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// update-freq (int)
//
// Number of seconds between reports when data is flowing

func (e *Progressreport) GetUpdateFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("update-freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Progressreport) SetUpdateFreq(value int) error {
	return e.Element.SetProperty("update-freq", value)
}

