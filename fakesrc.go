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

type Fakesrc struct {
	*base.GstBaseSrc
}

func NewFakesrc() (*Fakesrc, error) {
	e, err := gst.NewElement("fakesrc")
	if err != nil {
		return nil, err
	}
	return &Fakesrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewFakesrcWithName(name string) (*Fakesrc, error) {
	e, err := gst.NewElementWithName("fakesrc", name)
	if err != nil {
		return nil, err
	}
	return &Fakesrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// ----- Properties -----

// can-activate-pull (bool)
//
// Can activate in pull mode

func (e *Fakesrc) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakesrc) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

// can-activate-push (bool)
//
// Can activate in push mode

func (e *Fakesrc) GetCanActivatePush() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-push")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakesrc) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

// data (GstFakeSrcDataType)
//
// Data allocation method

func (e *Fakesrc) GetData() (GstFakeSrcDataType, error) {
	var value GstFakeSrcDataType
	var ok bool
	v, err := e.Element.GetProperty("data")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSrcDataType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSrcDataType")
	}
	return value, nil
}

func (e *Fakesrc) SetData(value GstFakeSrcDataType) error {
	e.Element.SetArg("data", string(value))
	return nil
}

// datarate (int)
//
// Timestamps buffers with number of bytes per second (0 = none)

func (e *Fakesrc) GetDatarate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("datarate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fakesrc) SetDatarate(value int) error {
	return e.Element.SetProperty("datarate", value)
}

// dump (bool)
//
// Dump buffer contents to stdout

func (e *Fakesrc) GetDump() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dump")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakesrc) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

// filltype (GstFakeSrcFillType)
//
// How to fill the buffer, if at all

func (e *Fakesrc) GetFilltype() (GstFakeSrcFillType, error) {
	var value GstFakeSrcFillType
	var ok bool
	v, err := e.Element.GetProperty("filltype")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSrcFillType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSrcFillType")
	}
	return value, nil
}

func (e *Fakesrc) SetFilltype(value GstFakeSrcFillType) error {
	e.Element.SetArg("filltype", string(value))
	return nil
}

// format (GstFormat)
//
// Set the format of the newsegment events to produce.

func (e *Fakesrc) GetFormat() (interface{}, error) {
	return e.Element.GetProperty("format")
}

func (e *Fakesrc) SetFormat(value interface{}) error {
	return e.Element.SetProperty("format", value)
}

// is-live (bool)
//
// True if the element cannot produce data in PAUSED

func (e *Fakesrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakesrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// last-message (string)
//
// The last status message

func (e *Fakesrc) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// parentsize (int)
//
// Size of parent buffer for sub-buffered allocation

func (e *Fakesrc) GetParentsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("parentsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fakesrc) SetParentsize(value int) error {
	return e.Element.SetProperty("parentsize", value)
}

// pattern (string)
//
// Set the pattern (unused)

func (e *Fakesrc) GetPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Fakesrc) SetPattern(value string) error {
	return e.Element.SetProperty("pattern", value)
}

// signal-handoffs (bool)
//
// Send a signal before pushing the buffer

func (e *Fakesrc) GetSignalHandoffs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal-handoffs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakesrc) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

// silent (bool)
//
// Don't produce last_message events

func (e *Fakesrc) GetSilent() (bool, error) {
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

func (e *Fakesrc) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

// sizemax (int)
//
// Maximum buffer size

func (e *Fakesrc) GetSizemax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sizemax")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fakesrc) SetSizemax(value int) error {
	return e.Element.SetProperty("sizemax", value)
}

// sizemin (int)
//
// Minimum buffer size

func (e *Fakesrc) GetSizemin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sizemin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Fakesrc) SetSizemin(value int) error {
	return e.Element.SetProperty("sizemin", value)
}

// sizetype (GstFakeSrcSizeType)
//
// How to determine buffer sizes

func (e *Fakesrc) GetSizetype() (GstFakeSrcSizeType, error) {
	var value GstFakeSrcSizeType
	var ok bool
	v, err := e.Element.GetProperty("sizetype")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSrcSizeType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSrcSizeType")
	}
	return value, nil
}

func (e *Fakesrc) SetSizetype(value GstFakeSrcSizeType) error {
	e.Element.SetArg("sizetype", string(value))
	return nil
}

// sync (bool)
//
// Sync to the clock to the datarate

func (e *Fakesrc) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Fakesrc) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// ----- Constants -----

type GstFakeSrcDataType string

const (
	GstFakeSrcDataTypeAllocate GstFakeSrcDataType = "allocate" // allocate (1) – Allocate data
	GstFakeSrcDataTypeSubbuffer GstFakeSrcDataType = "subbuffer" // subbuffer (2) – Subbuffer data
)

type GstFakeSrcFillType string

const (
	GstFakeSrcFillTypeNothing GstFakeSrcFillType = "nothing" // nothing (1) – Leave data as malloced
	GstFakeSrcFillTypeZero GstFakeSrcFillType = "zero" // zero (2) – Fill buffers with zeros
	GstFakeSrcFillTypeRandom GstFakeSrcFillType = "random" // random (3) – Fill buffers with random data
	GstFakeSrcFillTypePattern GstFakeSrcFillType = "pattern" // pattern (4) – Fill buffers with pattern 0x00 -> 0xff
	GstFakeSrcFillTypePatternSpan GstFakeSrcFillType = "pattern-span" // pattern-span (5) – Fill buffers with pattern 0x00 -> 0xff that spans buffers
)

type GstFakeSrcSizeType string

const (
	GstFakeSrcSizeTypeEmpty GstFakeSrcSizeType = "empty" // empty (1) – Send empty buffers
	GstFakeSrcSizeTypeFixed GstFakeSrcSizeType = "fixed" // fixed (2) – Fixed size buffers (sizemax sized)
	GstFakeSrcSizeTypeRandom GstFakeSrcSizeType = "random" // random (3) – Random sized buffers (sizemin <= size <= sizemax)
)

