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

type Qmlglsrc struct {
	*base.GstPushSrc
}

func NewQmlglsrc() (*Qmlglsrc, error) {
	e, err := gst.NewElement("qmlglsrc")
	if err != nil {
		return nil, err
	}
	return &Qmlglsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewQmlglsrcWithName(name string) (*Qmlglsrc, error) {
	e, err := gst.NewElementWithName("qmlglsrc", name)
	if err != nil {
		return nil, err
	}
	return &Qmlglsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// blocksize (uint)
//
// Size in bytes to read per buffer (-1 = default)

func (e *Qmlglsrc) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Qmlglsrc) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// do-timestamp (bool)
//
// Apply current stream time to buffers

func (e *Qmlglsrc) GetDoTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsrc) SetDoTimestamp(value bool) error {
	return e.Element.SetProperty("do-timestamp", value)
}

// num-buffers (int)
//
// Number of buffers to output before sending EOS (-1 = unlimited)

func (e *Qmlglsrc) GetNumBuffers() (int, error) {
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

func (e *Qmlglsrc) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

// typefind (bool)
//
// Run typefind before negotiating (deprecated, non-functional)

func (e *Qmlglsrc) GetTypefind() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("typefind")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsrc) SetTypefind(value bool) error {
	return e.Element.SetProperty("typefind", value)
}

// use-default-fbo (bool)
//
// When set it will not create a new FBO for the QML render thread

func (e *Qmlglsrc) GetUseDefaultFbo() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-default-fbo")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qmlglsrc) SetUseDefaultFbo(value bool) error {
	return e.Element.SetProperty("use-default-fbo", value)
}

// window (GstGpointer)
//
// The QQuickWindow to place in the object hierarchy

func (e *Qmlglsrc) GetWindow() (interface{}, error) {
	return e.Element.GetProperty("window")
}

func (e *Qmlglsrc) SetWindow(value interface{}) error {
	return e.Element.SetProperty("window", value)
}

