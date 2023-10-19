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

type Teletextdec struct {
	Element *gst.Element
}

func NewTeletextdec() (*Teletextdec, error) {
	e, err := gst.NewElement("teletextdec")
	if err != nil {
		return nil, err
	}
	return &Teletextdec{Element: e}, nil
}

func NewTeletextdecWithName(name string) (*Teletextdec, error) {
	e, err := gst.NewElementWithName("teletextdec", name)
	if err != nil {
		return nil, err
	}
	return &Teletextdec{Element: e}, nil
}

// ----- Properties -----

// font-description (string)
//
// Font description used for the pango output.

func (e *Teletextdec) GetFontDescription() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-description")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Teletextdec) SetFontDescription(value string) error {
	return e.Element.SetProperty("font-description", value)
}

// page (int)
//
// Number of page that should displayed

func (e *Teletextdec) GetPage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("page")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Teletextdec) SetPage(value int) error {
	return e.Element.SetProperty("page", value)
}

// subpage (int)
//
// Number of sub-page that should displayed (-1 for all)

func (e *Teletextdec) GetSubpage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("subpage")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Teletextdec) SetSubpage(value int) error {
	return e.Element.SetProperty("subpage", value)
}

// subtitles-mode (bool)
//
// Enables subtitles mode for text output stripping the blank lines and the teletext state lines

func (e *Teletextdec) GetSubtitlesMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("subtitles-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Teletextdec) SetSubtitlesMode(value bool) error {
	return e.Element.SetProperty("subtitles-mode", value)
}

// subtitles-template (string)
//
// Output template used to print each one of the subtitles lines

func (e *Teletextdec) GetSubtitlesTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitles-template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Teletextdec) SetSubtitlesTemplate(value string) error {
	return e.Element.SetProperty("subtitles-template", value)
}

