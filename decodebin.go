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

type Decodebin struct {
	Element *gst.Element
}

func NewDecodebin() (*Decodebin, error) {
	e, err := gst.NewElement("decodebin")
	if err != nil {
		return nil, err
	}
	return &Decodebin{Element: e}, nil
}

func NewDecodebinWithName(name string) (*Decodebin, error) {
	e, err := gst.NewElementWithName("decodebin", name)
	if err != nil {
		return nil, err
	}
	return &Decodebin{Element: e}, nil
}

// ----- Properties -----

// caps (GstCaps)
//
// The caps on which to stop decoding.

func (e *Decodebin) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Decodebin) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Decodebin) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decodebin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// expose-all-streams (bool)
//
// Expose all streams, including those of unknown type or that don't match the 'caps' property

func (e *Decodebin) GetExposeAllStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-all-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decodebin) SetExposeAllStreams(value bool) error {
	return e.Element.SetProperty("expose-all-streams", value)
}

// force-sw-decoders (bool)
//
// Use only sofware decoders to process streams

func (e *Decodebin) GetForceSwDecoders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-sw-decoders")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decodebin) SetForceSwDecoders(value bool) error {
	return e.Element.SetProperty("force-sw-decoders", value)
}

// high-percent (int)
//
// High threshold percent for buffering to finish.

func (e *Decodebin) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decodebin) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

// low-percent (int)
//
// Low threshold percent for buffering to start.

func (e *Decodebin) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Decodebin) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

// max-size-buffers (uint)
//
// Max amount of buffers in the queue (0=automatic).

func (e *Decodebin) GetMaxSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Decodebin) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

// max-size-bytes (uint)
//
// Max amount of bytes in the queue (0=automatic).

func (e *Decodebin) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Decodebin) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

// max-size-time (uint64)
//
// Max amount of time in the queue (in ns, 0=automatic).

func (e *Decodebin) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Decodebin) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

// post-stream-topology (bool)
//
// Post stream-topology messages

func (e *Decodebin) GetPostStreamTopology() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-stream-topology")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decodebin) SetPostStreamTopology(value bool) error {
	return e.Element.SetProperty("post-stream-topology", value)
}

// sink-caps (GstCaps)
//
// The caps of the input data. (NULL = use typefind element)

func (e *Decodebin) GetSinkCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("sink-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Decodebin) SetSinkCaps(value *gst.Caps) error {
	return e.Element.SetProperty("sink-caps", value)
}

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Decodebin) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Decodebin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// use-buffering (bool)
//
// Emit GST_MESSAGE_BUFFERING based on low-/high-percent thresholds

func (e *Decodebin) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Decodebin) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

