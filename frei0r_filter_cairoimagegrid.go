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

type Frei0RFilterCairoimagegrid struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterCairoimagegrid() (*Frei0RFilterCairoimagegrid, error) {
	e, err := gst.NewElement("frei0r-filter-cairoimagegrid")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCairoimagegrid{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterCairoimagegridWithName(name string) (*Frei0RFilterCairoimagegrid, error) {
	e, err := gst.NewElementWithName("frei0r-filter-cairoimagegrid", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterCairoimagegrid{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// columns (float64)
//
// Number of columns in the image grid. Input range 0 - 1 is interpreted as range 1 - 20

func (e *Frei0RFilterCairoimagegrid) GetColumns() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairoimagegrid) SetColumns(value float64) error {
	return e.Element.SetProperty("columns", value)
}

// rows (float64)
//
// Number of rows in the image grid. Input range 0 - 1 is interpreted as range 1 - 20

func (e *Frei0RFilterCairoimagegrid) GetRows() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterCairoimagegrid) SetRows(value float64) error {
	return e.Element.SetProperty("rows", value)
}

