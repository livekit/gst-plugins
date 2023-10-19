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
)

type Rtpst20221Fecenc struct {
	Element *gst.Element
}

func NewRtpst20221Fecenc() (*Rtpst20221Fecenc, error) {
	e, err := gst.NewElement("rtpst2022-1-fecenc")
	if err != nil {
		return nil, err
	}
	return &Rtpst20221Fecenc{Element: e}, nil
}

func NewRtpst20221FecencWithName(name string) (*Rtpst20221Fecenc, error) {
	e, err := gst.NewElementWithName("rtpst2022-1-fecenc", name)
	if err != nil {
		return nil, err
	}
	return &Rtpst20221Fecenc{Element: e}, nil
}

// ----- Properties -----

// columns (uint)
//
// Number of columns to apply row FEC on, 0=disabled

func (e *Rtpst20221Fecenc) GetColumns() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpst20221Fecenc) SetColumns(value uint) error {
	return e.Element.SetProperty("columns", value)
}

// enable-column-fec (bool)
//
// Whether the encoder should compute and send column FEC

func (e *Rtpst20221Fecenc) GetEnableColumnFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-column-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpst20221Fecenc) SetEnableColumnFec(value bool) error {
	return e.Element.SetProperty("enable-column-fec", value)
}

// enable-row-fec (bool)
//
// Whether the encoder should compute and send row FEC

func (e *Rtpst20221Fecenc) GetEnableRowFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-row-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpst20221Fecenc) SetEnableRowFec(value bool) error {
	return e.Element.SetProperty("enable-row-fec", value)
}

// pt (int)
//
// The payload type of FEC packets

func (e *Rtpst20221Fecenc) GetPt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rtpst20221Fecenc) SetPt(value int) error {
	return e.Element.SetProperty("pt", value)
}

// rows (uint)
//
// Number of rows to apply column FEC on, 0=disabled

func (e *Rtpst20221Fecenc) GetRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtpst20221Fecenc) SetRows(value uint) error {
	return e.Element.SetProperty("rows", value)
}

