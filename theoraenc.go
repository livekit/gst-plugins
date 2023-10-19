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

type Theoraenc struct {
	Element *gst.Element
}

func NewTheoraenc() (*Theoraenc, error) {
	e, err := gst.NewElement("theoraenc")
	if err != nil {
		return nil, err
	}
	return &Theoraenc{Element: e}, nil
}

func NewTheoraencWithName(name string) (*Theoraenc, error) {
	e, err := gst.NewElementWithName("theoraenc", name)
	if err != nil {
		return nil, err
	}
	return &Theoraenc{Element: e}, nil
}

// ----- Properties -----

// bitrate (int)
//
// Compressed video bitrate (kbps)

func (e *Theoraenc) GetBitrate() (int, error) {
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

func (e *Theoraenc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// cap-overflow (bool)
//
// Enable capping of bit reservoir overflows

func (e *Theoraenc) GetCapOverflow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cap-overflow")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Theoraenc) SetCapOverflow(value bool) error {
	return e.Element.SetProperty("cap-overflow", value)
}

// cap-underflow (bool)
//
// Enable capping of bit reservoir underflows

func (e *Theoraenc) GetCapUnderflow() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cap-underflow")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Theoraenc) SetCapUnderflow(value bool) error {
	return e.Element.SetProperty("cap-underflow", value)
}

// drop-frames (bool)
//
// Allow or disallow frame dropping

func (e *Theoraenc) GetDropFrames() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Theoraenc) SetDropFrames(value bool) error {
	return e.Element.SetProperty("drop-frames", value)
}

// keyframe-auto (bool)
//
// Automatic keyframe detection

func (e *Theoraenc) GetKeyframeAuto() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keyframe-auto")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Theoraenc) SetKeyframeAuto(value bool) error {
	return e.Element.SetProperty("keyframe-auto", value)
}

// keyframe-force (int)
//
// Force keyframe every N frames

func (e *Theoraenc) GetKeyframeForce() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-force")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoraenc) SetKeyframeForce(value int) error {
	return e.Element.SetProperty("keyframe-force", value)
}

// keyframe-freq (int)
//
// Keyframe frequency

func (e *Theoraenc) GetKeyframeFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyframe-freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoraenc) SetKeyframeFreq(value int) error {
	return e.Element.SetProperty("keyframe-freq", value)
}

// multipass-cache-file (string)
//
// Multipass cache file

func (e *Theoraenc) GetMultipassCacheFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("multipass-cache-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Theoraenc) SetMultipassCacheFile(value string) error {
	return e.Element.SetProperty("multipass-cache-file", value)
}

// multipass-mode (GstTheoraEncMultipassMode)
//
// Single pass or first/second pass

func (e *Theoraenc) GetMultipassMode() (GstTheoraEncMultipassMode, error) {
	var value GstTheoraEncMultipassMode
	var ok bool
	v, err := e.Element.GetProperty("multipass-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTheoraEncMultipassMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTheoraEncMultipassMode")
	}
	return value, nil
}

func (e *Theoraenc) SetMultipassMode(value GstTheoraEncMultipassMode) error {
	e.Element.SetArg("multipass-mode", string(value))
	return nil
}

// quality (int)
//
// Video quality

func (e *Theoraenc) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoraenc) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

// rate-buffer (int)
//
// Sets the size of the rate control buffer, in units of frames.  The default value of 0 instructs the encoder to automatically select an appropriate value

func (e *Theoraenc) GetRateBuffer() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rate-buffer")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoraenc) SetRateBuffer(value int) error {
	return e.Element.SetProperty("rate-buffer", value)
}

// speed-level (int)
//
// Controls the amount of motion vector searching done while encoding

func (e *Theoraenc) GetSpeedLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("speed-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Theoraenc) SetSpeedLevel(value int) error {
	return e.Element.SetProperty("speed-level", value)
}

// vp3-compatible (bool)
//
// Disables non-VP3 compatible features

func (e *Theoraenc) GetVp3Compatible() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vp3-compatible")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Theoraenc) SetVp3Compatible(value bool) error {
	return e.Element.SetProperty("vp3-compatible", value)
}

// ----- Constants -----

type GstTheoraEncMultipassMode string

const (
	GstTheoraEncMultipassModeSinglePass GstTheoraEncMultipassMode = "single-pass" // single-pass (0) – Single pass
	GstTheoraEncMultipassModeFirstPass GstTheoraEncMultipassMode = "first-pass" // first-pass (1) – First pass
	GstTheoraEncMultipassModeSecondPass GstTheoraEncMultipassMode = "second-pass" // second-pass (2) – Second pass
)

