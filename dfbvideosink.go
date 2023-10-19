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

type Dfbvideosink struct {
	*base.GstBaseSink
}

func NewDfbvideosink() (*Dfbvideosink, error) {
	e, err := gst.NewElement("dfbvideosink")
	if err != nil {
		return nil, err
	}
	return &Dfbvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDfbvideosinkWithName(name string) (*Dfbvideosink, error) {
	e, err := gst.NewElementWithName("dfbvideosink", name)
	if err != nil {
		return nil, err
	}
	return &Dfbvideosink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Dfbvideosink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dfbvideosink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Dfbvideosink) GetBlocksize() (uint, error) {
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

func (e *Dfbvideosink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// brightness (int)
//
// The brightness of the video

func (e *Dfbvideosink) GetBrightness() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("brightness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dfbvideosink) SetBrightness(value int) error {
	return e.Element.SetProperty("brightness", value)
}

// contrast (int)
//
// The contrast of the video

func (e *Dfbvideosink) GetContrast() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("contrast")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dfbvideosink) SetContrast(value int) error {
	return e.Element.SetProperty("contrast", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Dfbvideosink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dfbvideosink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// hue (int)
//
// The hue of the video

func (e *Dfbvideosink) GetHue() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("hue")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dfbvideosink) SetHue(value int) error {
	return e.Element.SetProperty("hue", value)
}

// last-sample (GstSample)
//
// The last sample received in the sink

func (e *Dfbvideosink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// layer-mode (GstDfbVideoSinkLayerMode)
//
// The cooperative level handling the access permission (set this to 'administrative' when the cursor is required)

func (e *Dfbvideosink) GetLayerMode() (GstDfbVideoSinkLayerMode, error) {
	var value GstDfbVideoSinkLayerMode
	var ok bool
	v, err := e.Element.GetProperty("layer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDfbVideoSinkLayerMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDfbVideoSinkLayerMode")
	}
	return value, nil
}

func (e *Dfbvideosink) SetLayerMode(value GstDfbVideoSinkLayerMode) error {
	e.Element.SetArg("layer-mode", string(value))
	return nil
}

// max-bitrate (uint64)
//
// The maximum bits per second to render (0 = disabled)

func (e *Dfbvideosink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dfbvideosink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Dfbvideosink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Dfbvideosink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// pixel-aspect-ratio (string)
//
// The pixel aspect ratio of the device

func (e *Dfbvideosink) GetPixelAspectRatio() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pixel-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dfbvideosink) SetPixelAspectRatio(value string) error {
	return e.Element.SetProperty("pixel-aspect-ratio", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Dfbvideosink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dfbvideosink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Dfbvideosink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dfbvideosink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Dfbvideosink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dfbvideosink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// saturation (int)
//
// The saturation of the video

func (e *Dfbvideosink) GetSaturation() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("saturation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Dfbvideosink) SetSaturation(value int) error {
	return e.Element.SetProperty("saturation", value)
}

// show-preroll-frame (bool)
//
// Whether to render video frames during preroll

func (e *Dfbvideosink) GetShowPrerollFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("show-preroll-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dfbvideosink) SetShowPrerollFrame(value bool) error {
	return e.Element.SetProperty("show-preroll-frame", value)
}

// stats (GstStructure)
//
// Sink Statistics

func (e *Dfbvideosink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// surface (GstGpointer)
//
// The target surface for video

func (e *Dfbvideosink) GetSurface() (interface{}, error) {
	return e.Element.GetProperty("surface")
}

func (e *Dfbvideosink) SetSurface(value interface{}) error {
	return e.Element.SetProperty("surface", value)
}

// sync (bool)
//
// Sync on the clock

func (e *Dfbvideosink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dfbvideosink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Dfbvideosink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Dfbvideosink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Dfbvideosink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Dfbvideosink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

// vsync (bool)
//
// Wait for next vertical sync to draw frames

func (e *Dfbvideosink) GetVsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vsync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Dfbvideosink) SetVsync(value bool) error {
	return e.Element.SetProperty("vsync", value)
}

// ----- Constants -----

type GstDfbVideoSinkLayerMode string

const (
	GstDfbVideoSinkLayerModeNone GstDfbVideoSinkLayerMode = "none" // none (0) – NONE
	GstDfbVideoSinkLayerModeExclusive GstDfbVideoSinkLayerMode = "exclusive" // exclusive (1) – DLSCL_EXCLUSIVE
	GstDfbVideoSinkLayerModeAdministrative GstDfbVideoSinkLayerMode = "administrative" // administrative (2) – DLSCL_ADMINISTRATIVE
)

