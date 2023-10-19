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

type Frei0RSrcPartik0L struct {
	*base.GstPushSrc
}

func NewFrei0RSrcPartik0L() (*Frei0RSrcPartik0L, error) {
	e, err := gst.NewElement("frei0r-src-partik0l")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcPartik0L{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcPartik0LWithName(name string) (*Frei0RSrcPartik0L, error) {
	e, err := gst.NewElementWithName("frei0r-src-partik0l", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcPartik0L{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// down (float64)
//
// blossom on a lower prime number

func (e *Frei0RSrcPartik0L) GetDown() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("down")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPartik0L) SetDown(value float64) error {
	return e.Element.SetProperty("down", value)
}

// up (float64)
//
// blossom on a higher prime number

func (e *Frei0RSrcPartik0L) GetUp() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("up")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Frei0RSrcPartik0L) SetUp(value float64) error {
	return e.Element.SetProperty("up", value)
}

