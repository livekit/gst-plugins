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

type Zbar struct {
	*base.GstBaseTransform
}

func NewZbar() (*Zbar, error) {
	e, err := gst.NewElement("zbar")
	if err != nil {
		return nil, err
	}
	return &Zbar{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewZbarWithName(name string) (*Zbar, error) {
	e, err := gst.NewElementWithName("zbar", name)
	if err != nil {
		return nil, err
	}
	return &Zbar{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// attach-frame (bool)
//
// Attach a frame dump to each barcode message

func (e *Zbar) GetAttachFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("attach-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Zbar) SetAttachFrame(value bool) error {
	return e.Element.SetProperty("attach-frame", value)
}

// cache (bool)
//
// Enable or disable the inter-image result cache

func (e *Zbar) GetCache() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cache")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Zbar) SetCache(value bool) error {
	return e.Element.SetProperty("cache", value)
}

// message (bool)
//
// Post a barcode message for each detected code

func (e *Zbar) GetMessage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("message")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Zbar) SetMessage(value bool) error {
	return e.Element.SetProperty("message", value)
}

