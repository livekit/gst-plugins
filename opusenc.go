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

type Opusenc struct {
	Element *gst.Element
}

func NewOpusenc() (*Opusenc, error) {
	e, err := gst.NewElement("opusenc")
	if err != nil {
		return nil, err
	}
	return &Opusenc{Element: e}, nil
}

func NewOpusencWithName(name string) (*Opusenc, error) {
	e, err := gst.NewElementWithName("opusenc", name)
	if err != nil {
		return nil, err
	}
	return &Opusenc{Element: e}, nil
}

// ----- Properties -----

// audio-type (GstOpusEncAudioType)
//
// What type of audio to optimize for

func (e *Opusenc) GetAudioType() (GstOpusEncAudioType, error) {
	var value GstOpusEncAudioType
	var ok bool
	v, err := e.Element.GetProperty("audio-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncAudioType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncAudioType")
	}
	return value, nil
}

func (e *Opusenc) SetAudioType(value GstOpusEncAudioType) error {
	e.Element.SetArg("audio-type", string(value))
	return nil
}

// bandwidth (GstOpusEncBandwidth)
//
// Audio Band Width

func (e *Opusenc) GetBandwidth() (GstOpusEncBandwidth, error) {
	var value GstOpusEncBandwidth
	var ok bool
	v, err := e.Element.GetProperty("bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncBandwidth)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncBandwidth")
	}
	return value, nil
}

func (e *Opusenc) SetBandwidth(value GstOpusEncBandwidth) error {
	e.Element.SetArg("bandwidth", string(value))
	return nil
}

// bitrate (int)
//
// Specify an encoding bit-rate (in bps).

func (e *Opusenc) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opusenc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bitrate-type (GstOpusEncBitrateType)
//
// Bitrate type

func (e *Opusenc) GetBitrateType() (GstOpusEncBitrateType, error) {
	var value GstOpusEncBitrateType
	var ok bool
	v, err := e.Element.GetProperty("bitrate-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncBitrateType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncBitrateType")
	}
	return value, nil
}

func (e *Opusenc) SetBitrateType(value GstOpusEncBitrateType) error {
	e.Element.SetArg("bitrate-type", string(value))
	return nil
}

// complexity (int)
//
// Complexity

func (e *Opusenc) GetComplexity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("complexity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opusenc) SetComplexity(value int) error {
	return e.Element.SetProperty("complexity", value)
}

// dtx (bool)
//
// DTX

func (e *Opusenc) GetDtx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dtx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Opusenc) SetDtx(value bool) error {
	return e.Element.SetProperty("dtx", value)
}

// frame-size (GstOpusEncFrameSize)
//
// The duration of an audio frame, in ms

func (e *Opusenc) GetFrameSize() (GstOpusEncFrameSize, error) {
	var value GstOpusEncFrameSize
	var ok bool
	v, err := e.Element.GetProperty("frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncFrameSize)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncFrameSize")
	}
	return value, nil
}

func (e *Opusenc) SetFrameSize(value GstOpusEncFrameSize) error {
	e.Element.SetArg("frame-size", string(value))
	return nil
}

// inband-fec (bool)
//
// Enable in-band forward error correction (use in combination with the packet-loss-percentage property)

func (e *Opusenc) GetInbandFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("inband-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Opusenc) SetInbandFec(value bool) error {
	return e.Element.SetProperty("inband-fec", value)
}

// max-payload-size (uint)
//
// Maximum payload size in bytes

func (e *Opusenc) GetMaxPayloadSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-payload-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Opusenc) SetMaxPayloadSize(value uint) error {
	return e.Element.SetProperty("max-payload-size", value)
}

// packet-loss-percentage (int)
//
// Packet loss percentage

func (e *Opusenc) GetPacketLossPercentage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("packet-loss-percentage")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Opusenc) SetPacketLossPercentage(value int) error {
	return e.Element.SetProperty("packet-loss-percentage", value)
}

// ----- Constants -----

type GstOpusEncFrameSize string

const (
	GstOpusEncFrameSize25 GstOpusEncFrameSize = "2.5" // 2.5 (2) – 2.5
	GstOpusEncFrameSize5 GstOpusEncFrameSize = "5" // 5 (5) – 5
	GstOpusEncFrameSize10 GstOpusEncFrameSize = "10" // 10 (10) – 10
	GstOpusEncFrameSize20 GstOpusEncFrameSize = "20" // 20 (20) – 20
	GstOpusEncFrameSize40 GstOpusEncFrameSize = "40" // 40 (40) – 40
	GstOpusEncFrameSize60 GstOpusEncFrameSize = "60" // 60 (60) – 60
)

type GstOpusEncAudioType string

const (
	GstOpusEncAudioTypeGeneric GstOpusEncAudioType = "generic" // generic (2049) – Generic audio
	GstOpusEncAudioTypeVoice GstOpusEncAudioType = "voice" // voice (2048) – Voice
	GstOpusEncAudioTypeRestrictedLowdelay GstOpusEncAudioType = "restricted-lowdelay" // restricted-lowdelay (2051) – Restricted low delay
)

type GstOpusEncBandwidth string

const (
	GstOpusEncBandwidthNarrowband GstOpusEncBandwidth = "narrowband" // narrowband (1101) – Narrow band
	GstOpusEncBandwidthMediumband GstOpusEncBandwidth = "mediumband" // mediumband (1102) – Medium band
	GstOpusEncBandwidthWideband GstOpusEncBandwidth = "wideband" // wideband (1103) – Wide band
	GstOpusEncBandwidthSuperwideband GstOpusEncBandwidth = "superwideband" // superwideband (1104) – Super wide band
	GstOpusEncBandwidthFullband GstOpusEncBandwidth = "fullband" // fullband (1105) – Full band
	GstOpusEncBandwidthAuto GstOpusEncBandwidth = "auto" // auto (-1000) – Auto
)

type GstOpusEncBitrateType string

const (
	GstOpusEncBitrateTypeCbr GstOpusEncBitrateType = "cbr" // cbr (0) – CBR
	GstOpusEncBitrateTypeVbr GstOpusEncBitrateType = "vbr" // vbr (1) – VBR
	GstOpusEncBitrateTypeConstrainedVbr GstOpusEncBitrateType = "constrained-vbr" // constrained-vbr (2) – Constrained VBR
)

