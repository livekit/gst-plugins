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

type Dwritetextoverlay struct {
	*base.GstBaseTransform
}

func NewDwritetextoverlay() (*Dwritetextoverlay, error) {
	e, err := gst.NewElement("dwritetextoverlay")
	if err != nil {
		return nil, err
	}
	return &Dwritetextoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDwritetextoverlayWithName(name string) (*Dwritetextoverlay, error) {
	e, err := gst.NewElementWithName("dwritetextoverlay", name)
	if err != nil {
		return nil, err
	}
	return &Dwritetextoverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cc-field (int)
//
// The closed caption field to render when available, (-1 = automatic)

func (e *Dwritetextoverlay) GetCcField() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cc-field")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dwritetextoverlay) SetCcField(value int) error {
	return e.Element.SetProperty("cc-field", value)
}

// cc-timeout (uint64)
//
// Duration after which to erase overlay when no cc data has arrived for the selected field, in nanoseconds unit

func (e *Dwritetextoverlay) GetCcTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("cc-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dwritetextoverlay) SetCcTimeout(value uint64) error {
	return e.Element.SetProperty("cc-timeout", value)
}

// enable-cc (bool)
//
// Enable closed caption rendering

func (e *Dwritetextoverlay) GetEnableCc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritetextoverlay) SetEnableCc(value bool) error {
	return e.Element.SetProperty("enable-cc", value)
}

// remove-cc-meta (bool)
//
// Remove caption meta from output buffers when closed caption rendering is enabled

func (e *Dwritetextoverlay) GetRemoveCcMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remove-cc-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dwritetextoverlay) SetRemoveCcMeta(value bool) error {
	return e.Element.SetProperty("remove-cc-meta", value)
}

