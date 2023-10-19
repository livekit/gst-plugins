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

type Rtpopuspay struct {
	Element *gst.Element
}

func NewRtpopuspay() (*Rtpopuspay, error) {
	e, err := gst.NewElement("rtpopuspay")
	if err != nil {
		return nil, err
	}
	return &Rtpopuspay{Element: e}, nil
}

func NewRtpopuspayWithName(name string) (*Rtpopuspay, error) {
	e, err := gst.NewElementWithName("rtpopuspay", name)
	if err != nil {
		return nil, err
	}
	return &Rtpopuspay{Element: e}, nil
}

// ----- Properties -----

// dtx (bool)
//
// If enabled, the payloader will not transmit empty packets.

func (e *Rtpopuspay) GetDtx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dtx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtpopuspay) SetDtx(value bool) error {
	return e.Element.SetProperty("dtx", value)
}

