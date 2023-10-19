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

type D3Dvideosink struct {
	*base.GstBaseSink
}

func NewD3Dvideosink() (*D3Dvideosink, error) {
	e, err := gst.NewElement("d3dvideosink")
	if err != nil {
		return nil, err
	}
	return &D3Dvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewD3DvideosinkWithName(name string) (*D3Dvideosink, error) {
	e, err := gst.NewElementWithName("d3dvideosink", name)
	if err != nil {
		return nil, err
	}
	return &D3Dvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// create-render-window (bool)
//
// If no window ID is given, a new render window is created

func (e *D3Dvideosink) GetCreateRenderWindow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-render-window")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3Dvideosink) SetCreateRenderWindow(value bool) error {
	return e.Element.SetProperty("create-render-window", value)
}

// enable-navigation-events (bool)
//
// When enabled, navigation events are sent upstream

func (e *D3Dvideosink) GetEnableNavigationEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-navigation-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3Dvideosink) SetEnableNavigationEvents(value bool) error {
	return e.Element.SetProperty("enable-navigation-events", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *D3Dvideosink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3Dvideosink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// stream-stop-on-close (bool)
//
// If the render window is closed stop stream

func (e *D3Dvideosink) GetStreamStopOnClose() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("stream-stop-on-close")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *D3Dvideosink) SetStreamStopOnClose(value bool) error {
	return e.Element.SetProperty("stream-stop-on-close", value)
}

