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

type Rfbsrc struct {
	*base.GstPushSrc
}

func NewRfbsrc() (*Rfbsrc, error) {
	e, err := gst.NewElement("rfbsrc")
	if err != nil {
		return nil, err
	}
	return &Rfbsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewRfbsrcWithName(name string) (*Rfbsrc, error) {
	e, err := gst.NewElementWithName("rfbsrc", name)
	if err != nil {
		return nil, err
	}
	return &Rfbsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// height (int)
//
// height of screen

func (e *Rfbsrc) GetHeight() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("height")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rfbsrc) SetHeight(value int) error {
	return e.Element.SetProperty("height", value)
}

// host (string)
//
// Host to connect to

func (e *Rfbsrc) GetHost() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("host")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rfbsrc) SetHost(value string) error {
	return e.Element.SetProperty("host", value)
}

// incremental (bool)
//
// Incremental updates

func (e *Rfbsrc) GetIncremental() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("incremental")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rfbsrc) SetIncremental(value bool) error {
	return e.Element.SetProperty("incremental", value)
}

// offset-x (int)
//
// x offset for screen scrapping

func (e *Rfbsrc) GetOffsetX() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-x")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rfbsrc) SetOffsetX(value int) error {
	return e.Element.SetProperty("offset-x", value)
}

// offset-y (int)
//
// y offset for screen scrapping

func (e *Rfbsrc) GetOffsetY() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("offset-y")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rfbsrc) SetOffsetY(value int) error {
	return e.Element.SetProperty("offset-y", value)
}

// password (string)
//
// Password for authentication

func (e *Rfbsrc) GetPassword() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("password")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rfbsrc) SetPassword(value string) error {
	return e.Element.SetProperty("password", value)
}

// port (int)
//
// Port

func (e *Rfbsrc) GetPort() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("port")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rfbsrc) SetPort(value int) error {
	return e.Element.SetProperty("port", value)
}

// shared (bool)
//
// Share desktop with other clients

func (e *Rfbsrc) GetShared() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("shared")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rfbsrc) SetShared(value bool) error {
	return e.Element.SetProperty("shared", value)
}

// uri (string)
//
// uri to an RFB from. All GStreamer parameters can be
// encoded in the URI, this URI format is RFC compliant.

func (e *Rfbsrc) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rfbsrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// use-copyrect (bool)
//
// Use copyrect encoding

func (e *Rfbsrc) GetUseCopyrect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-copyrect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rfbsrc) SetUseCopyrect(value bool) error {
	return e.Element.SetProperty("use-copyrect", value)
}

// version (string)
//
// RFB protocol version

func (e *Rfbsrc) GetVersion() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("version")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Rfbsrc) SetVersion(value string) error {
	return e.Element.SetProperty("version", value)
}

// view-only (bool)
//
// only view the desktop

func (e *Rfbsrc) GetViewOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("view-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rfbsrc) SetViewOnly(value bool) error {
	return e.Element.SetProperty("view-only", value)
}

// width (int)
//
// width of screen

func (e *Rfbsrc) GetWidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("width")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Rfbsrc) SetWidth(value int) error {
	return e.Element.SetProperty("width", value)
}

