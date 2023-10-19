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

type Rtphdrextntp64 struct {
	Element *gst.Element
}

func NewRtphdrextntp64() (*Rtphdrextntp64, error) {
	e, err := gst.NewElement("rtphdrextntp64")
	if err != nil {
		return nil, err
	}
	return &Rtphdrextntp64{Element: e}, nil
}

func NewRtphdrextntp64WithName(name string) (*Rtphdrextntp64, error) {
	e, err := gst.NewElementWithName("rtphdrextntp64", name)
	if err != nil {
		return nil, err
	}
	return &Rtphdrextntp64{Element: e}, nil
}

// ----- Properties -----

// every-packet (bool)
//
// If set to TRUE the header extension will be added to every packet,
// independent of its timestamp. By default only the first packet with a
// given timestamp will get the header extension added.

func (e *Rtphdrextntp64) GetEveryPacket() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("every-packet")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtphdrextntp64) SetEveryPacket(value bool) error {
	return e.Element.SetProperty("every-packet", value)
}

// interval (uint64)
//
// The minimum interval between packets that get the header extension added.

