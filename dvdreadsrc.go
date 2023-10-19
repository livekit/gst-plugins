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

type Dvdreadsrc struct {
	*base.GstPushSrc
}

func NewDvdreadsrc() (*Dvdreadsrc, error) {
	e, err := gst.NewElement("dvdreadsrc")
	if err != nil {
		return nil, err
	}
	return &Dvdreadsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDvdreadsrcWithName(name string) (*Dvdreadsrc, error) {
	e, err := gst.NewElementWithName("dvdreadsrc", name)
	if err != nil {
		return nil, err
	}
	return &Dvdreadsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// angle (int)
//
// angle

func (e *Dvdreadsrc) GetAngle() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dvdreadsrc) SetAngle(value int) error {
	return e.Element.SetProperty("angle", value)
}

// chapter (int)
//
// chapter

func (e *Dvdreadsrc) GetChapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("chapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dvdreadsrc) SetChapter(value int) error {
	return e.Element.SetProperty("chapter", value)
}

// device (string)
//
// DVD device location

func (e *Dvdreadsrc) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dvdreadsrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

// title (int)
//
// title

func (e *Dvdreadsrc) GetTitle() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("title")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dvdreadsrc) SetTitle(value int) error {
	return e.Element.SetProperty("title", value)
}

