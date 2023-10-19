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

type Decklinkvideosink struct {
	*base.GstBaseSink
}

func NewDecklinkvideosink() (*Decklinkvideosink, error) {
	e, err := gst.NewElement("decklinkvideosink")
	if err != nil {
		return nil, err
	}
	return &Decklinkvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDecklinkvideosinkWithName(name string) (*Decklinkvideosink, error) {
	e, err := gst.NewElementWithName("decklinkvideosink", name)
	if err != nil {
		return nil, err
	}
	return &Decklinkvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// afd-bar-line (int)
//
// Line number to use for inserting AFD/Bar data (0 = disabled)

func (e *Decklinkvideosink) GetAfdBarLine() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("afd-bar-line")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetAfdBarLine(value int) error {
	return e.Element.SetProperty("afd-bar-line", value)
}

// cc-line (int)
//
// Line number to use for inserting closed captions (0 = disabled)

func (e *Decklinkvideosink) GetCcLine() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cc-line")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetCcLine(value int) error {
	return e.Element.SetProperty("cc-line", value)
}

// device-number (int)
//
// Output device instance to use

func (e *Decklinkvideosink) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

// hw-serial-number (string)
//
// The serial number (hardware ID) of the Decklink card

func (e *Decklinkvideosink) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// keyer-level (int)
//
// Keyer level

func (e *Decklinkvideosink) GetKeyerLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyer-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetKeyerLevel(value int) error {
	return e.Element.SetProperty("keyer-level", value)
}

// keyer-mode (GstDecklinkKeyerMode)
//
// Keyer mode to be enabled

func (e *Decklinkvideosink) GetKeyerMode() (GstDecklinkKeyerMode, error) {
	var value GstDecklinkKeyerMode
	var ok bool
	v, err := e.Element.GetProperty("keyer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkKeyerMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkKeyerMode")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetKeyerMode(value GstDecklinkKeyerMode) error {
	e.Element.SetArg("keyer-mode", string(value))
	return nil
}

// mapping-format (GstDecklinkMappingFormat)
//
// Specifies the 3G-SDI mapping format to use (SMPTE ST 425-1:2017).

func (e *Decklinkvideosink) GetMappingFormat() (GstDecklinkMappingFormat, error) {
	var value GstDecklinkMappingFormat
	var ok bool
	v, err := e.Element.GetProperty("mapping-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkMappingFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkMappingFormat")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetMappingFormat(value GstDecklinkMappingFormat) error {
	e.Element.SetArg("mapping-format", string(value))
	return nil
}

// mode (GstDecklinkModes)
//
// Video Mode to use for playback

func (e *Decklinkvideosink) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *Decklinkvideosink) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// persistent-id (int64)
//
// Decklink device to use. Higher priority than "device-number".
// BMDDeckLinkPersistentID is a device speciﬁc, 32-bit unique identiﬁer.
// It is stable even when the device is plugged in a diļ¬erent connector,
// across reboots, and when plugged into diļ¬erent computers.

func (e *Decklinkvideosink) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Decklinkvideosink) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

// profile (GstDecklinkProfileId)
//
// Specifies decklink profile to use.

func (e *Decklinkvideosink) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

func (e *Decklinkvideosink) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

// timecode-format (GstDecklinkTimecodeFormat)
//
// Timecode format type to use for playback

func (e *Decklinkvideosink) GetTimecodeFormat() (interface{}, error) {
	return e.Element.GetProperty("timecode-format")
}

func (e *Decklinkvideosink) SetTimecodeFormat(value interface{}) error {
	return e.Element.SetProperty("timecode-format", value)
}

// video-format (GstDecklinkVideoFormat)
//
// Video format type to use for playback

func (e *Decklinkvideosink) GetVideoFormat() (interface{}, error) {
	return e.Element.GetProperty("video-format")
}

func (e *Decklinkvideosink) SetVideoFormat(value interface{}) error {
	return e.Element.SetProperty("video-format", value)
}

// ----- Constants -----

type GstDecklinkKeyerMode string

const (
	GstDecklinkKeyerModeOff GstDecklinkKeyerMode = "off" // off (0) – Off
	GstDecklinkKeyerModeInternal GstDecklinkKeyerMode = "internal" // internal (1) – Internal
	GstDecklinkKeyerModeExternal GstDecklinkKeyerMode = "external" // external (2) – External
)

type GstDecklinkMappingFormat string

const (
	GstDecklinkMappingFormatDefault GstDecklinkMappingFormat = "default" // default (0) – Default, don't change mapping format
	GstDecklinkMappingFormatLevelA GstDecklinkMappingFormat = "level-a" // level-a (1) – Level A
	GstDecklinkMappingFormatLevelB GstDecklinkMappingFormat = "level-b" // level-b (2) – Level B
)

