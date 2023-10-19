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

type Watchdog struct {
	*base.GstBaseTransform
}

func NewWatchdog() (*Watchdog, error) {
	e, err := gst.NewElement("watchdog")
	if err != nil {
		return nil, err
	}
	return &Watchdog{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewWatchdogWithName(name string) (*Watchdog, error) {
	e, err := gst.NewElementWithName("watchdog", name)
	if err != nil {
		return nil, err
	}
	return &Watchdog{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// timeout (int)
//
// Timeout (in ms) after which an element error is sent to the bus if no buffers are received. 0 means disabled.

func (e *Watchdog) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Watchdog) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

