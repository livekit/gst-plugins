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

type Imagefreeze struct {
	Element *gst.Element
}

func NewImagefreeze() (*Imagefreeze, error) {
	e, err := gst.NewElement("imagefreeze")
	if err != nil {
		return nil, err
	}
	return &Imagefreeze{Element: e}, nil
}

func NewImagefreezeWithName(name string) (*Imagefreeze, error) {
	e, err := gst.NewElementWithName("imagefreeze", name)
	if err != nil {
		return nil, err
	}
	return &Imagefreeze{Element: e}, nil
}

// ----- Properties -----

// allow-replace (bool)
//
// Allow replacing the input buffer and always output the latest

func (e *Imagefreeze) GetAllowReplace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-replace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Imagefreeze) SetAllowReplace(value bool) error {
	return e.Element.SetProperty("allow-replace", value)
}

// is-live (bool)
//
// Selects whether the output stream should be a non-live stream based on
// the segment configured via a GST_EVENT_SEEK, or whether the output
// stream should be a live stream with the negotiated framerate.

func (e *Imagefreeze) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Imagefreeze) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

// num-buffers (int)
//
// Number of buffers to output before sending EOS (-1 = unlimited)

func (e *Imagefreeze) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Imagefreeze) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

