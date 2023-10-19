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

type Uridecodebin3 struct {
	Element *gst.Element
}

func NewUridecodebin3() (*Uridecodebin3, error) {
	e, err := gst.NewElement("uridecodebin3")
	if err != nil {
		return nil, err
	}
	return &Uridecodebin3{Element: e}, nil
}

func NewUridecodebin3WithName(name string) (*Uridecodebin3, error) {
	e, err := gst.NewElementWithName("uridecodebin3", name)
	if err != nil {
		return nil, err
	}
	return &Uridecodebin3{Element: e}, nil
}

// ----- Properties -----

// buffer-duration (int64)
//
// Buffer duration when buffering streams (-1 default value)

func (e *Uridecodebin3) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Uridecodebin3) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

// buffer-size (int)
//
// Buffer size when buffering streams (-1 default value)

func (e *Uridecodebin3) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Uridecodebin3) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// caps (GstCaps)
//
// The caps on which to stop decoding. (NULL = default)

func (e *Uridecodebin3) GetCaps() (*gst.Caps, error) {
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

func (e *Uridecodebin3) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Uridecodebin3) GetConnectionSpeed() (uint64, error) {
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

func (e *Uridecodebin3) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// current-suburi (string)
//
// The currently playing URI of a subtitle

func (e *Uridecodebin3) GetCurrentSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// current-uri (string)
//
// The currently playing URI

func (e *Uridecodebin3) GetCurrentUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// download (bool)
//
// Attempt download buffering when buffering network streams

func (e *Uridecodebin3) GetDownload() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("download")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Uridecodebin3) SetDownload(value bool) error {
	return e.Element.SetProperty("download", value)
}

// download-dir (string)
//
// The directory where buffers are downloaded to, if 'download' is enabled.
// If not set (default), the XDG cache directory is used.
// Will be applied to the next 'uri' played or until the element go back to
// the PAUSED state.

func (e *Uridecodebin3) GetDownloadDir() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("download-dir")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uridecodebin3) SetDownloadDir(value string) error {
	return e.Element.SetProperty("download-dir", value)
}

// instant-uri (bool)
//
// Changes to uri are applied immediately (instead of on EOS or when the
// element is set back to PLAYING).

func (e *Uridecodebin3) GetInstantUri() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("instant-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Uridecodebin3) SetInstantUri(value bool) error {
	return e.Element.SetProperty("instant-uri", value)
}

// ring-buffer-max-size (uint64)
//
// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)

func (e *Uridecodebin3) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Uridecodebin3) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

// source (GstElement)
//
// Source object used

func (e *Uridecodebin3) GetSource() (interface{}, error) {
	return e.Element.GetProperty("source")
}

// suburi (string)
//
// Optional URI of a subtitle

func (e *Uridecodebin3) GetSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uridecodebin3) SetSuburi(value string) error {
	return e.Element.SetProperty("suburi", value)
}

// uri (string)
//
// URI to decode

func (e *Uridecodebin3) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Uridecodebin3) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// use-buffering (bool)
//
// Perform buffering on demuxed/parsed media

func (e *Uridecodebin3) GetUseBuffering() (bool, error) {
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

func (e *Uridecodebin3) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

