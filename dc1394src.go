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

type Dc1394Src struct {
	*base.GstPushSrc
}

func NewDc1394Src() (*Dc1394Src, error) {
	e, err := gst.NewElement("dc1394src")
	if err != nil {
		return nil, err
	}
	return &Dc1394Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDc1394SrcWithName(name string) (*Dc1394Src, error) {
	e, err := gst.NewElementWithName("dc1394src", name)
	if err != nil {
		return nil, err
	}
	return &Dc1394Src{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// blocksize (uint)
//
// Size in bytes to read per buffer (-1 = default)

func (e *Dc1394Src) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dc1394Src) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// dma (uint)
//
// The number of frames in the Direct Memory Access ring buffer

func (e *Dc1394Src) GetDma() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dma")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Dc1394Src) SetDma(value uint) error {
	return e.Element.SetProperty("dma", value)
}

// do-timestamp (bool)
//
// Apply current stream time to buffers

func (e *Dc1394Src) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dc1394Src) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

// guid (string)
//
// The hexadecimal representation of the GUID of the camera (use first camera available if null)

func (e *Dc1394Src) GetGuid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("guid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dc1394Src) SetGuid(value string) error {
	return e.Element.SetProperty("guid", value)
}

// iso (GstDC1394ISOSpeed)
//
// The ISO bandwidth in Mbps

func (e *Dc1394Src) GetIso() (interface{}, error) {
	return e.Element.GetProperty("iso")
}

func (e *Dc1394Src) SetIso(value interface{}) error {
	return e.Element.SetProperty("iso", value)
}

// num-buffers (int)
//
// Number of buffers to output before sending EOS (-1 = unlimited)

func (e *Dc1394Src) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dc1394Src) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// typefind (bool)
//
// Run typefind before negotiating (deprecated, non-functional)

func (e *Dc1394Src) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dc1394Src) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

// unit (int)
//
// The unit number of the camera (-1 if no unit number is used)

func (e *Dc1394Src) GetUnit() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("unit")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dc1394Src) SetUnit(value int) error {
	return e.Element.SetProperty("unit", value)
}

// ----- Constants -----

type GstDc1394Isospeed string

const (
	GstDc1394Isospeed100 GstDc1394Isospeed = "100" // 100 (100) – DC1394 ISO speed 100
	GstDc1394Isospeed200 GstDc1394Isospeed = "200" // 200 (200) – DC1394 ISO speed 200
	GstDc1394Isospeed400 GstDc1394Isospeed = "400" // 400 (400) – DC1394 ISO speed 400
	GstDc1394Isospeed800 GstDc1394Isospeed = "800" // 800 (800) – DC1394 ISO speed 800
	GstDc1394Isospeed1600 GstDc1394Isospeed = "1600" // 1600 (1600) – DC1394 ISO speed 1600
	GstDc1394Isospeed3200 GstDc1394Isospeed = "3200" // 3200 (3200) – DC1394 ISO speed 3200
)

