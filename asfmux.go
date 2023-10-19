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

type Asfmux struct {
	Element *gst.Element
}

func NewAsfmux() (*Asfmux, error) {
	e, err := gst.NewElement("asfmux")
	if err != nil {
		return nil, err
	}
	return &Asfmux{Element: e}, nil
}

func NewAsfmuxWithName(name string) (*Asfmux, error) {
	e, err := gst.NewElementWithName("asfmux", name)
	if err != nil {
		return nil, err
	}
	return &Asfmux{Element: e}, nil
}

// ----- Properties -----

// merge-stream-tags (bool)
//
// If the stream metadata (received as events in the sink) should be merged to the main file metadata.

func (e *Asfmux) GetMergeStreamTags() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("merge-stream-tags")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Asfmux) SetMergeStreamTags(value bool) error {
	return e.Element.SetProperty("merge-stream-tags", value)
}

// packet-size (uint)
//
// The ASF packets size (bytes)

func (e *Asfmux) GetPacketSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("packet-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Asfmux) SetPacketSize(value uint) error {
	return e.Element.SetProperty("packet-size", value)
}

// padding (uint64)
//
// Size of the padding object to be added to the end of the header. If this less than 24 (the smaller size of an ASF object), no padding is added.

func (e *Asfmux) GetPadding() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Asfmux) SetPadding(value uint64) error {
	return e.Element.SetProperty("padding", value)
}

// preroll (uint64)
//
// The preroll time (milliseconds)

func (e *Asfmux) GetPreroll() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("preroll")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Asfmux) SetPreroll(value uint64) error {
	return e.Element.SetProperty("preroll", value)
}

// streamable (bool)
//
// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written.

func (e *Asfmux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Asfmux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

