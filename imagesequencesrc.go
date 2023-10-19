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

type Imagesequencesrc struct {
	*base.GstPushSrc
}

func NewImagesequencesrc() (*Imagesequencesrc, error) {
	e, err := gst.NewElement("imagesequencesrc")
	if err != nil {
		return nil, err
	}
	return &Imagesequencesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewImagesequencesrcWithName(name string) (*Imagesequencesrc, error) {
	e, err := gst.NewElementWithName("imagesequencesrc", name)
	if err != nil {
		return nil, err
	}
	return &Imagesequencesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// framerate (GstFraction)
//
// The output framerate.

func (e *Imagesequencesrc) GetFramerate() (interface{}, error) {
	return e.Element.GetProperty("framerate")
}

func (e *Imagesequencesrc) SetFramerate(value interface{}) error {
	return e.Element.SetProperty("framerate", value)
}

// location (string)
//
// Pattern to create file names of input files.  File names are created by calling sprintf with the pattern and the current index.

func (e *Imagesequencesrc) GetLocation() (string, error) {
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

func (e *Imagesequencesrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// start-index (int)
//
// Start value of index.  The initial value of index can be set either by setting index or start-index.  When the end of the loop is reached, the index will be set to the value start-index.

func (e *Imagesequencesrc) GetStartIndex() (int, error) {
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

func (e *Imagesequencesrc) SetStartIndex(value int) error {
	return e.Element.SetProperty("start-index", value)
}

// stop-index (int)
//
// Stop value of index.  The special value -1 means no stop.

func (e *Imagesequencesrc) GetStopIndex() (int, error) {
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

func (e *Imagesequencesrc) SetStopIndex(value int) error {
	return e.Element.SetProperty("stop-index", value)
}

