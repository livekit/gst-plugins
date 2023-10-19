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

type Irtspparse struct {
	*base.GstBaseParse
}

func NewIrtspparse() (*Irtspparse, error) {
	e, err := gst.NewElement("irtspparse")
	if err != nil {
		return nil, err
	}
	return &Irtspparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewIrtspparseWithName(name string) (*Irtspparse, error) {
	e, err := gst.NewElementWithName("irtspparse", name)
	if err != nil {
		return nil, err
	}
	return &Irtspparse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// channel-id (int)
//
// Channel Identifier

func (e *Irtspparse) GetChannelId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("channel-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Irtspparse) SetChannelId(value int) error {
	return e.Element.SetProperty("channel-id", value)
}

