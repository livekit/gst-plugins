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

type Frei0RFilterFacebl0R struct {
	*base.GstBaseTransform
}

func NewFrei0RFilterFacebl0R() (*Frei0RFilterFacebl0R, error) {
	e, err := gst.NewElement("frei0r-filter-facebl0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterFacebl0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewFrei0RFilterFacebl0RWithName(name string) (*Frei0RFilterFacebl0R, error) {
	e, err := gst.NewElementWithName("frei0r-filter-facebl0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RFilterFacebl0R{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// classifier (string)
//
// Full path to the XML pattern model for recognition; look in /usr/share/opencv/haarcascades

func (e *Frei0RFilterFacebl0R) GetClassifier() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("classifier")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetClassifier(value string) error {
	return e.Element.SetProperty("classifier", value)
}

// ellipse (bool)
//
// Draw a red ellipse around the object

func (e *Frei0RFilterFacebl0R) GetEllipse() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ellipse")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetEllipse(value bool) error {
	return e.Element.SetProperty("ellipse", value)
}

// largest (float64)
//
// Maximum object size in pixels, divided by 10000

func (e *Frei0RFilterFacebl0R) GetLargest() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("largest")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetLargest(value float64) error {
	return e.Element.SetProperty("largest", value)
}

// neighbors (float64)
//
// Minimum number of rectangles that makes up an object, divided by 100

func (e *Frei0RFilterFacebl0R) GetNeighbors() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("neighbors")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetNeighbors(value float64) error {
	return e.Element.SetProperty("neighbors", value)
}

// recheck (float64)
//
// How often to detect an object in number of frames, divided by 1000

func (e *Frei0RFilterFacebl0R) GetRecheck() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("recheck")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetRecheck(value float64) error {
	return e.Element.SetProperty("recheck", value)
}

// search-scale (float64)
//
// The search window scale factor, divided by 10

func (e *Frei0RFilterFacebl0R) GetSearchScale() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("search-scale")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetSearchScale(value float64) error {
	return e.Element.SetProperty("search-scale", value)
}

// smallest (float64)
//
// Minimum window size in pixels, divided by 1000

func (e *Frei0RFilterFacebl0R) GetSmallest() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("smallest")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetSmallest(value float64) error {
	return e.Element.SetProperty("smallest", value)
}

// threads (float64)
//
// How many threads to use divided by 100; 0 uses CPU count

func (e *Frei0RFilterFacebl0R) GetThreads() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RFilterFacebl0R) SetThreads(value float64) error {
	return e.Element.SetProperty("threads", value)
}

