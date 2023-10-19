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

type Cdiocddasrc struct {
	*base.GstPushSrc
}

func NewCdiocddasrc() (*Cdiocddasrc, error) {
	e, err := gst.NewElement("cdiocddasrc")
	if err != nil {
		return nil, err
	}
	return &Cdiocddasrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewCdiocddasrcWithName(name string) (*Cdiocddasrc, error) {
	e, err := gst.NewElementWithName("cdiocddasrc", name)
	if err != nil {
		return nil, err
	}
	return &Cdiocddasrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// read-speed (int)
//
// Read from device at the specified speed (-1 = default)

func (e *Cdiocddasrc) GetReadSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("read-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cdiocddasrc) SetReadSpeed(value int) error {
	return e.Element.SetProperty("read-speed", value)
}

