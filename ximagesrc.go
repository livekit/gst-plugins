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

type Ximagesrc struct {
	*base.GstPushSrc
}

func NewXimagesrc() (*Ximagesrc, error) {
	e, err := gst.NewElement("ximagesrc")
	if err != nil {
		return nil, err
	}
	return &Ximagesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewXimagesrcWithName(name string) (*Ximagesrc, error) {
	e, err := gst.NewElementWithName("ximagesrc", name)
	if err != nil {
		return nil, err
	}
	return &Ximagesrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// display-name (string)
//
// X Display Name

func (e *Ximagesrc) GetDisplayName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("display-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Ximagesrc) SetDisplayName(value string) error {
	return e.Element.SetProperty("display-name", value)
}

// enable-navigation-events (bool)
//
// Enable navigation events

func (e *Ximagesrc) GetEnableNavigationEvents() (bool, error) {
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

func (e *Ximagesrc) SetEnableNavigationEvents(value bool) error {
	return e.Element.SetProperty("enable-navigation-events", value)
}

// endx (uint)
//
// X coordinate of bottom right corner of area to be recorded
// (0 for bottom right of screen)

func (e *Ximagesrc) GetEndx() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("endx")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ximagesrc) SetEndx(value uint) error {
	return e.Element.SetProperty("endx", value)
}

// endy (uint)
//
// Y coordinate of bottom right corner of area to be recorded
// (0 for bottom right of screen)

func (e *Ximagesrc) GetEndy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("endy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ximagesrc) SetEndy(value uint) error {
	return e.Element.SetProperty("endy", value)
}

// remote (bool)
//
// Whether the X display is remote. The element will try to use alternate calls
// known to work better with remote displays.

func (e *Ximagesrc) GetRemote() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("remote")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ximagesrc) SetRemote(value bool) error {
	return e.Element.SetProperty("remote", value)
}

// show-pointer (bool)
//
// Show mouse pointer (if XFixes extension enabled)

func (e *Ximagesrc) GetShowPointer() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-pointer")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ximagesrc) SetShowPointer(value bool) error {
	return e.Element.SetProperty("show-pointer", value)
}

// startx (uint)
//
// X coordinate of top left corner of area to be recorded
// (0 for top left of screen)

func (e *Ximagesrc) GetStartx() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("startx")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ximagesrc) SetStartx(value uint) error {
	return e.Element.SetProperty("startx", value)
}

// starty (uint)
//
// Y coordinate of top left corner of area to be recorded
// (0 for top left of screen)

func (e *Ximagesrc) GetStarty() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("starty")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Ximagesrc) SetStarty(value uint) error {
	return e.Element.SetProperty("starty", value)
}

// use-damage (bool)
//
// Use XDamage (if the XDamage extension is enabled)

func (e *Ximagesrc) GetUseDamage() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-damage")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Ximagesrc) SetUseDamage(value bool) error {
	return e.Element.SetProperty("use-damage", value)
}

// xid (uint64)
//
// The XID of the window to capture. 0 for the root window (default).

func (e *Ximagesrc) GetXid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("xid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Ximagesrc) SetXid(value uint64) error {
	return e.Element.SetProperty("xid", value)
}

// xname (string)
//
// The name of the window to capture, if any.

func (e *Ximagesrc) GetXname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("xname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Ximagesrc) SetXname(value string) error {
	return e.Element.SetProperty("xname", value)
}

