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

type Decklinkaudiosrc struct {
	*base.GstPushSrc
}

func NewDecklinkaudiosrc() (*Decklinkaudiosrc, error) {
	e, err := gst.NewElement("decklinkaudiosrc")
	if err != nil {
		return nil, err
	}
	return &Decklinkaudiosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDecklinkaudiosrcWithName(name string) (*Decklinkaudiosrc, error) {
	e, err := gst.NewElementWithName("decklinkaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &Decklinkaudiosrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// alignment-threshold (uint64)
//
// Timestamp alignment threshold in nanoseconds

func (e *Decklinkaudiosrc) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decklinkaudiosrc) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

// buffer-size (uint)
//
// Size of internal buffer in number of video frames

func (e *Decklinkaudiosrc) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Decklinkaudiosrc) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

// channels (GstDecklinkAudioChannels)
//
// Audio channels

func (e *Decklinkaudiosrc) GetChannels() (GstDecklinkAudioChannels, error) {
	var value GstDecklinkAudioChannels
	var ok bool
	v, err := e.Element.GetProperty("channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkAudioChannels)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkAudioChannels")
	}
	return value, nil
}

func (e *Decklinkaudiosrc) SetChannels(value GstDecklinkAudioChannels) error {
	e.Element.SetArg("channels", string(value))
	return nil
}

// connection (GstDecklinkAudioConnection)
//
// Audio input connection to use

func (e *Decklinkaudiosrc) GetConnection() (GstDecklinkAudioConnection, error) {
	var value GstDecklinkAudioConnection
	var ok bool
	v, err := e.Element.GetProperty("connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkAudioConnection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkAudioConnection")
	}
	return value, nil
}

func (e *Decklinkaudiosrc) SetConnection(value GstDecklinkAudioConnection) error {
	e.Element.SetArg("connection", string(value))
	return nil
}

// device-number (int)
//
// Output device instance to use

func (e *Decklinkaudiosrc) GetDeviceNumber() (int, error) {
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

func (e *Decklinkaudiosrc) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

// discont-wait (uint64)
//
// Window of time in nanoseconds to wait before creating a discontinuity

func (e *Decklinkaudiosrc) GetDiscontWait() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("discont-wait")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decklinkaudiosrc) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

// hw-serial-number (string)
//
// The serial number (hardware ID) of the Decklink card

func (e *Decklinkaudiosrc) GetHwSerialNumber() (string, error) {
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

// persistent-id (int64)
//
// Decklink device to use. Higher priority than "device-number".
// BMDDeckLinkPersistentID is a device speciﬁc, 32-bit unique identiﬁer.
// It is stable even when the device is plugged in a diļ¬erent connector,
// across reboots, and when plugged into diļ¬erent computers.

func (e *Decklinkaudiosrc) GetPersistentId() (int64, error) {
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

func (e *Decklinkaudiosrc) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

// ----- Constants -----

type GstDecklinkAudioChannels string

const (
	GstDecklinkAudioChannels2 GstDecklinkAudioChannels = "2" // 2 (2) – 2 Channels
	GstDecklinkAudioChannels8 GstDecklinkAudioChannels = "8" // 8 (8) – 8 Channels
	GstDecklinkAudioChannels16 GstDecklinkAudioChannels = "16" // 16 (16) – 16 Channels
	GstDecklinkAudioChannelsMax GstDecklinkAudioChannels = "max" // max (0) – Maximum channels supported
)

type GstDecklinkAudioConnection string

const (
	GstDecklinkAudioConnectionAuto GstDecklinkAudioConnection = "auto" // auto (0) – Automatic
	GstDecklinkAudioConnectionEmbedded GstDecklinkAudioConnection = "embedded" // embedded (1) – SDI/HDMI embedded audio
	GstDecklinkAudioConnectionAes GstDecklinkAudioConnection = "aes" // aes (2) – AES/EBU input
	GstDecklinkAudioConnectionAnalog GstDecklinkAudioConnection = "analog" // analog (3) – Analog input
	GstDecklinkAudioConnectionAnalogXlr GstDecklinkAudioConnection = "analog-xlr" // analog-xlr (4) – Analog input (XLR)
	GstDecklinkAudioConnectionAnalogRca GstDecklinkAudioConnection = "analog-rca" // analog-rca (5) – Analog input (RCA)
)

