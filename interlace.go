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

type Interlace struct {
	Element *gst.Element
}

func NewInterlace() (*Interlace, error) {
	e, err := gst.NewElement("interlace")
	if err != nil {
		return nil, err
	}
	return &Interlace{Element: e}, nil
}

func NewInterlaceWithName(name string) (*Interlace, error) {
	e, err := gst.NewElementWithName("interlace", name)
	if err != nil {
		return nil, err
	}
	return &Interlace{Element: e}, nil
}

// ----- Properties -----

// allow-rff (bool)
//
// Allow generation of buffers with RFF flag set, i.e., duration of 3 fields

func (e *Interlace) GetAllowRff() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-rff")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Interlace) SetAllowRff(value bool) error {
	return e.Element.SetProperty("allow-rff", value)
}

// field-pattern (GstInterlacePattern)
//
// The output field pattern

func (e *Interlace) GetFieldPattern() (GstInterlacePattern, error) {
	var value GstInterlacePattern
	var ok bool
	v, err := e.Element.GetProperty("field-pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstInterlacePattern)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstInterlacePattern")
	}
	return value, nil
}

func (e *Interlace) SetFieldPattern(value GstInterlacePattern) error {
	e.Element.SetArg("field-pattern", string(value))
	return nil
}

// pattern-offset (uint)
//
// The initial field pattern offset. Counts from 0.

func (e *Interlace) GetPatternOffset() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pattern-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Interlace) SetPatternOffset(value uint) error {
	return e.Element.SetProperty("pattern-offset", value)
}

// top-field-first (bool)
//
// Interlaced stream should be top field first

func (e *Interlace) GetTopFieldFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("top-field-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Interlace) SetTopFieldFirst(value bool) error {
	return e.Element.SetProperty("top-field-first", value)
}

// ----- Constants -----

type GstInterlacePattern string

const (
	GstInterlacePattern11 GstInterlacePattern = "1:1" // 1:1 (0) – 1:1 (e.g. 60p -> 60i)
	GstInterlacePattern22 GstInterlacePattern = "2:2" // 2:2 (1) – 2:2 (e.g. 30p -> 60i)
	GstInterlacePattern23 GstInterlacePattern = "2:3" // 2:3 (2) – 2:3 (e.g. 24p -> 60i telecine)
	GstInterlacePattern2332 GstInterlacePattern = "2:3:3:2" // 2:3:3:2 (3) – 2:3:3:2 (e.g. 24p -> 60i telecine)
	GstInterlacePattern2113 GstInterlacePattern = "2-11:3" // 2-11:3 (4) – Euro 2-11:3 (e.g. 24p -> 50i telecine)
	GstInterlacePattern343 GstInterlacePattern = "3:4-3" // 3:4-3 (5) – 3:4-3 (e.g. 16p -> 60i telecine)
	GstInterlacePattern374 GstInterlacePattern = "3-7:4" // 3-7:4 (6) – 3-7:4 (e.g. 16p -> 50i telecine)
	GstInterlacePattern334 GstInterlacePattern = "3:3:4" // 3:3:4 (7) – 3:3:4 (e.g. 18p -> 60i telecine)
	GstInterlacePattern33 GstInterlacePattern = "3:3" // 3:3 (8) – 3:3 (e.g. 20p -> 60i telecine)
	GstInterlacePattern324 GstInterlacePattern = "3:2-4" // 3:2-4 (9) – 3:2-4 (e.g. 27.5p -> 60i telecine)
	GstInterlacePattern124 GstInterlacePattern = "1:2-4" // 1:2-4 (10) – 1:2-4 (e.g. 27.5p -> 50i telecine)
)

