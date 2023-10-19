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

type Twolame struct {
	Element *gst.Element
}

func NewTwolame() (*Twolame, error) {
	e, err := gst.NewElement("twolame")
	if err != nil {
		return nil, err
	}
	return &Twolame{Element: e}, nil
}

func NewTwolameWithName(name string) (*Twolame, error) {
	e, err := gst.NewElementWithName("twolame", name)
	if err != nil {
		return nil, err
	}
	return &Twolame{Element: e}, nil
}

// ----- Properties -----

// ath-level (float32)
//
// ATH Level in dB

func (e *Twolame) GetAthLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ath-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Twolame) SetAthLevel(value float32) error {
	return e.Element.SetProperty("ath-level", value)
}

// bitrate (int)
//
// Bitrate in kbit/sec (8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, 192, 224, 256, 320, 384)

func (e *Twolame) GetBitrate() (int, error) {
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

func (e *Twolame) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// copyright (bool)
//
// Mark as copyright

func (e *Twolame) GetCopyright() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("copyright")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Twolame) SetCopyright(value bool) error {
	return e.Element.SetProperty("copyright", value)
}

// emphasis (GstTwoLameEmphasis)
//
// Pre-emphasis to apply to the decoded audio

func (e *Twolame) GetEmphasis() (GstTwoLameEmphasis, error) {
	var value GstTwoLameEmphasis
	var ok bool
	v, err := e.Element.GetProperty("emphasis")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTwoLameEmphasis)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTwoLameEmphasis")
	}
	return value, nil
}

func (e *Twolame) SetEmphasis(value GstTwoLameEmphasis) error {
	e.Element.SetArg("emphasis", string(value))
	return nil
}

// energy-level-extension (bool)
//
// Write peak PCM level to each frame

func (e *Twolame) GetEnergyLevelExtension() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("energy-level-extension")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Twolame) SetEnergyLevelExtension(value bool) error {
	return e.Element.SetProperty("energy-level-extension", value)
}

// error-protection (bool)
//
// Adds checksum to every frame

func (e *Twolame) GetErrorProtection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("error-protection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Twolame) SetErrorProtection(value bool) error {
	return e.Element.SetProperty("error-protection", value)
}

// mode (GstTwoLameMode)
//
// Encoding mode

func (e *Twolame) GetMode() (GstTwoLameMode, error) {
	var value GstTwoLameMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTwoLameMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTwoLameMode")
	}
	return value, nil
}

func (e *Twolame) SetMode(value GstTwoLameMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

// original (bool)
//
// Mark as original

func (e *Twolame) GetOriginal() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("original")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Twolame) SetOriginal(value bool) error {
	return e.Element.SetProperty("original", value)
}

// padding (GstTwoLamePadding)
//
// Padding type

func (e *Twolame) GetPadding() (GstTwoLamePadding, error) {
	var value GstTwoLamePadding
	var ok bool
	v, err := e.Element.GetProperty("padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTwoLamePadding)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTwoLamePadding")
	}
	return value, nil
}

func (e *Twolame) SetPadding(value GstTwoLamePadding) error {
	e.Element.SetArg("padding", string(value))
	return nil
}

// psymodel (int)
//
// Psychoacoustic model used to encode the audio

func (e *Twolame) GetPsymodel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("psymodel")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Twolame) SetPsymodel(value int) error {
	return e.Element.SetProperty("psymodel", value)
}

// quick-mode (bool)
//
// Calculate Psymodel every frames

func (e *Twolame) GetQuickMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("quick-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Twolame) SetQuickMode(value bool) error {
	return e.Element.SetProperty("quick-mode", value)
}

// quick-mode-count (int)
//
// Calculate Psymodel every n frames

func (e *Twolame) GetQuickModeCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quick-mode-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Twolame) SetQuickModeCount(value int) error {
	return e.Element.SetProperty("quick-mode-count", value)
}

// vbr (bool)
//
// Enable variable bitrate mode

func (e *Twolame) GetVbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Twolame) SetVbr(value bool) error {
	return e.Element.SetProperty("vbr", value)
}

// vbr-level (float32)
//
// VBR Level

func (e *Twolame) GetVbrLevel() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("vbr-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Twolame) SetVbrLevel(value float32) error {
	return e.Element.SetProperty("vbr-level", value)
}

// vbr-max-bitrate (int)
//
// Specify maximum VBR bitrate (0=off, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, 192, 224, 256, 320, 384)

func (e *Twolame) GetVbrMaxBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("vbr-max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Twolame) SetVbrMaxBitrate(value int) error {
	return e.Element.SetProperty("vbr-max-bitrate", value)
}

// ----- Constants -----

type GstTwoLameEmphasis string

const (
	GstTwoLameEmphasisNone GstTwoLameEmphasis = "none" // none (0) – No emphasis
	GstTwoLameEmphasis5 GstTwoLameEmphasis = "5" // 5 (1) – 50/15 ms
	GstTwoLameEmphasisCcit GstTwoLameEmphasis = "ccit" // ccit (3) – CCIT J.17
)

type GstTwoLameMode string

const (
	GstTwoLameModeAuto GstTwoLameMode = "auto" // auto (-1) – Auto
	GstTwoLameModeStereo GstTwoLameMode = "stereo" // stereo (0) – Stereo
	GstTwoLameModeJoint GstTwoLameMode = "joint" // joint (1) – Joint Stereo
	GstTwoLameModeDual GstTwoLameMode = "dual" // dual (2) – Dual Channel
	GstTwoLameModeMono GstTwoLameMode = "mono" // mono (3) – Mono
)

type GstTwoLamePadding string

const (
	GstTwoLamePaddingNever GstTwoLamePadding = "never" // never (0) – No Padding
	GstTwoLamePaddingAlways GstTwoLamePadding = "always" // always (1) – Always Pad
)

