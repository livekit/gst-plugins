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

type Deinterlace struct {
	Element *gst.Element
}

func NewDeinterlace() (*Deinterlace, error) {
	e, err := gst.NewElement("deinterlace")
	if err != nil {
		return nil, err
	}
	return &Deinterlace{Element: e}, nil
}

func NewDeinterlaceWithName(name string) (*Deinterlace, error) {
	e, err := gst.NewElementWithName("deinterlace", name)
	if err != nil {
		return nil, err
	}
	return &Deinterlace{Element: e}, nil
}

// ----- Properties -----

// drop-orphans (bool)
//
// This selects whether to drop orphan fields at the beginning of telecine
// patterns in active locking mode.

func (e *Deinterlace) GetDropOrphans() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-orphans")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Deinterlace) SetDropOrphans(value bool) error {
	return e.Element.SetProperty("drop-orphans", value)
}

// fields (GstDeinterlaceFields)
//
// This selects which fields should be output. If "all" is selected
// the output framerate will be double.

func (e *Deinterlace) GetFields() (GstDeinterlaceFields, error) {
	var value GstDeinterlaceFields
	var ok bool
	v, err := e.Element.GetProperty("fields")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceFields)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceFields")
	}
	return value, nil
}

func (e *Deinterlace) SetFields(value GstDeinterlaceFields) error {
	e.Element.SetArg("fields", string(value))
	return nil
}

// ignore-obscure (bool)
//
// This selects whether to ignore obscure/rare telecine patterns.
// NTSC 2:3 pulldown variants are the only really common patterns.

func (e *Deinterlace) GetIgnoreObscure() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-obscure")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Deinterlace) SetIgnoreObscure(value bool) error {
	return e.Element.SetProperty("ignore-obscure", value)
}

// locking (GstDeinterlaceLocking)
//
// This selects which approach to pattern locking is used which affects
// processing latency and accuracy of timestamp adjustment for telecine
// streams.

func (e *Deinterlace) GetLocking() (GstDeinterlaceLocking, error) {
	var value GstDeinterlaceLocking
	var ok bool
	v, err := e.Element.GetProperty("locking")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceLocking)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceLocking")
	}
	return value, nil
}

func (e *Deinterlace) SetLocking(value GstDeinterlaceLocking) error {
	e.Element.SetArg("locking", string(value))
	return nil
}

// method (GstDeinterlaceMethods)
//
// Selects the different deinterlacing algorithms that can be used.
// These provide different quality and CPU usage.

// mode (GstDeinterlaceModes)
//
// This selects whether the deinterlacing methods should
// always be applied or if they should only be applied
// on content that has the "interlaced" flag on the caps.

func (e *Deinterlace) GetMode() (GstDeinterlaceModes, error) {
	var value GstDeinterlaceModes
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceModes")
	}
	return value, nil
}

func (e *Deinterlace) SetMode(value GstDeinterlaceModes) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// tff (GstDeinterlaceFieldLayout)
//
// Deinterlace top field first

func (e *Deinterlace) GetTff() (GstDeinterlaceFieldLayout, error) {
	var value GstDeinterlaceFieldLayout
	var ok bool
	v, err := e.Element.GetProperty("tff")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDeinterlaceFieldLayout)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDeinterlaceFieldLayout")
	}
	return value, nil
}

func (e *Deinterlace) SetTff(value GstDeinterlaceFieldLayout) error {
	e.Element.SetArg("tff", string(value))
	return nil
}

// ----- Constants -----

type GstDeinterlaceFieldLayout string

const (
	GstDeinterlaceFieldLayoutAuto GstDeinterlaceFieldLayout = "auto" // auto (0) – Auto detection
	GstDeinterlaceFieldLayoutTff GstDeinterlaceFieldLayout = "tff" // tff (1) – Top field first
	GstDeinterlaceFieldLayoutBff GstDeinterlaceFieldLayout = "bff" // bff (2) – Bottom field first
)

type GstDeinterlaceFields string

const (
	GstDeinterlaceFieldsAll GstDeinterlaceFields = "all" // all (0) – All fields
	GstDeinterlaceFieldsTop GstDeinterlaceFields = "top" // top (1) – Top fields only
	GstDeinterlaceFieldsBottom GstDeinterlaceFields = "bottom" // bottom (2) – Bottom fields only
	GstDeinterlaceFieldsAuto GstDeinterlaceFields = "auto" // auto (3) – Automatically detect
)

type GstDeinterlaceLocking string

const (
	GstDeinterlaceLockingNone GstDeinterlaceLocking = "none" // none (0) – No pattern locking
	GstDeinterlaceLockingAuto GstDeinterlaceLocking = "auto" // auto (1) – Choose passive/active locking depending on whether upstream is live
	GstDeinterlaceLockingActive GstDeinterlaceLocking = "active" // active (2) – Block until pattern-locked. Use accurate timestamp interpolation within a pattern repeat.
	GstDeinterlaceLockingPassive GstDeinterlaceLocking = "passive" // passive (3) – Do not block. Use naïve timestamp adjustment until pattern-locked based on state history.
)

type GstDeinterlaceMethods string

const (
	GstDeinterlaceMethodsTomsmocomp GstDeinterlaceMethods = "tomsmocomp" // tomsmocomp (0) – Motion Adaptive: Motion Search
	GstDeinterlaceMethodsGreedyh GstDeinterlaceMethods = "greedyh" // greedyh (1) – Motion Adaptive: Advanced Detection
	GstDeinterlaceMethodsGreedyl GstDeinterlaceMethods = "greedyl" // greedyl (2) – Motion Adaptive: Simple Detection
	GstDeinterlaceMethodsVfir GstDeinterlaceMethods = "vfir" // vfir (3) – Blur Vertical
	GstDeinterlaceMethodsLinear GstDeinterlaceMethods = "linear" // linear (4) – Linear
	GstDeinterlaceMethodsLinearblend GstDeinterlaceMethods = "linearblend" // linearblend (5) – Blur: Temporal (Do Not Use)
	GstDeinterlaceMethodsScalerbob GstDeinterlaceMethods = "scalerbob" // scalerbob (6) – Double lines
	GstDeinterlaceMethodsWeave GstDeinterlaceMethods = "weave" // weave (7) – Weave (Do Not Use)
	GstDeinterlaceMethodsWeavetff GstDeinterlaceMethods = "weavetff" // weavetff (8) – Progressive: Top Field First (Do Not Use)
	GstDeinterlaceMethodsWeavebff GstDeinterlaceMethods = "weavebff" // weavebff (9) – Progressive: Bottom Field First (Do Not Use)
	GstDeinterlaceMethodsYadif GstDeinterlaceMethods = "yadif" // yadif (10) – YADIF Adaptive Deinterlacer
)

type GstDeinterlaceModes string

const (
	GstDeinterlaceModesAuto GstDeinterlaceModes = "auto" // auto (0) – Auto detection (best effort)
	GstDeinterlaceModesInterlaced GstDeinterlaceModes = "interlaced" // interlaced (1) – Force deinterlacing
	GstDeinterlaceModesDisabled GstDeinterlaceModes = "disabled" // disabled (2) – Run in passthrough mode
	GstDeinterlaceModesAutoStrict GstDeinterlaceModes = "auto-strict" // auto-strict (3) – Auto detection (strict)
)

