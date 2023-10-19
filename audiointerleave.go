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

type Audiointerleave struct {
	*base.GstAggregator
}

func NewAudiointerleave() (*Audiointerleave, error) {
	e, err := gst.NewElement("audiointerleave")
	if err != nil {
		return nil, err
	}
	return &Audiointerleave{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

func NewAudiointerleaveWithName(name string) (*Audiointerleave, error) {
	e, err := gst.NewElementWithName("audiointerleave", name)
	if err != nil {
		return nil, err
	}
	return &Audiointerleave{GstAggregator: &base.GstAggregator{Element: e}}, nil
}

// ----- Properties -----

// channel-positions (GstGvalueArray)
//
// Channel positions: This property controls the channel positions
// that are used on the src caps. The number of elements should be
// the same as the number of sink pads and the array should contain
// a valid list of channel positions. The n-th element of the array
// is the position of the n-th sink pad.

// channel-positions-from-input (bool)
//
// Channel positions from input: If this property is set to TRUE the channel
// positions will be taken from the input caps if valid channel positions for
// the output can be constructed from them. If this is set to TRUE setting the
// channel-positions property overwrites this property again.

func (e *Audiointerleave) GetChannelPositionsFromInput() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("channel-positions-from-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Audiointerleave) SetChannelPositionsFromInput(value bool) error {
	return e.Element.SetProperty("channel-positions-from-input", value)
}

