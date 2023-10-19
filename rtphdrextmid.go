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

type Rtphdrextmid struct {
	Element *gst.Element
}

func NewRtphdrextmid() (*Rtphdrextmid, error) {
	e, err := gst.NewElement("rtphdrextmid")
	if err != nil {
		return nil, err
	}
	return &Rtphdrextmid{Element: e}, nil
}

func NewRtphdrextmidWithName(name string) (*Rtphdrextmid, error) {
	e, err := gst.NewElementWithName("rtphdrextmid", name)
	if err != nil {
		return nil, err
	}
	return &Rtphdrextmid{Element: e}, nil
}

// ----- Properties -----

// mid (string)
//
// The Media Identification (MID) value either last retrieved from the RTP
// Header extension, or to set on outgoing RTP packets.

func (e *Rtphdrextmid) GetMid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rtphdrextmid) SetMid(value string) error {
	return e.Element.SetProperty("mid", value)
}

