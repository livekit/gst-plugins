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

type Multifilesrc struct {
	*base.GstPushSrc
}

func NewMultifilesrc() (*Multifilesrc, error) {
	e, err := gst.NewElement("multifilesrc")
	if err != nil {
		return nil, err
	}
	return &Multifilesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewMultifilesrcWithName(name string) (*Multifilesrc, error) {
	e, err := gst.NewElementWithName("multifilesrc", name)
	if err != nil {
		return nil, err
	}
	return &Multifilesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// Caps describing the format of the data.

func (e *Multifilesrc) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Multifilesrc) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// index (int)
//
// Index to use with location property to create file names.  The index is incremented by one for each buffer read.

func (e *Multifilesrc) GetIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multifilesrc) SetIndex(value int) error {
	return e.Element.SetProperty("index", value)
}

// location (string)
//
// Pattern to create file names of input files.  File names are created by calling sprintf with the pattern and the current index.

func (e *Multifilesrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Multifilesrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// loop (bool)
//
// Whether to repeat from the beginning when all files have been read.

func (e *Multifilesrc) GetLoop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("loop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multifilesrc) SetLoop(value bool) error {
	return e.Element.SetProperty("loop", value)
}

// start-index (int)
//
// Start value of index.  The initial value of index can be set either by setting index or start-index.  When the end of the loop is reached, the index will be set to the value start-index.

func (e *Multifilesrc) GetStartIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("start-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multifilesrc) SetStartIndex(value int) error {
	return e.Element.SetProperty("start-index", value)
}

// stop-index (int)
//
// Stop value of index.  The special value -1 means no stop.

func (e *Multifilesrc) GetStopIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stop-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multifilesrc) SetStopIndex(value int) error {
	return e.Element.SetProperty("stop-index", value)
}

