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

type Urisourcebin struct {
	Element *gst.Element
}

func NewUrisourcebin() (*Urisourcebin, error) {
	e, err := gst.NewElement("urisourcebin")
	if err != nil {
		return nil, err
	}
	return &Urisourcebin{Element: e}, nil
}

func NewUrisourcebinWithName(name string) (*Urisourcebin, error) {
	e, err := gst.NewElementWithName("urisourcebin", name)
	if err != nil {
		return nil, err
	}
	return &Urisourcebin{Element: e}, nil
}

// ----- Properties -----

// buffer-duration (int64)
//
// Buffer duration when buffering streams (-1 default value)

func (e *Urisourcebin) GetBufferDuration() (int64, error) {
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

func (e *Urisourcebin) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

// buffer-size (int)
//
// Buffer size when buffering streams (-1 default value)

func (e *Urisourcebin) GetBufferSize() (int, error) {
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

func (e *Urisourcebin) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Urisourcebin) GetConnectionSpeed() (uint64, error) {
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

func (e *Urisourcebin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// download (bool)
//
// Attempt download buffering when buffering network streams

func (e *Urisourcebin) GetDownload() (bool, error) {
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

func (e *Urisourcebin) SetDownload(value bool) error {
	return e.Element.SetProperty("download", value)
}

// download-dir (string)
//
// The directory where buffers are downloaded to, if 'download' is enabled.
// If not set (default), the XDG cache directory is used.

func (e *Urisourcebin) GetDownloadDir() (string, error) {
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

func (e *Urisourcebin) SetDownloadDir(value string) error {
	return e.Element.SetProperty("download-dir", value)
}

// high-watermark (float64)
//
// High threshold for buffering to finish. Only used if use-buffering is True

func (e *Urisourcebin) GetHighWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Urisourcebin) SetHighWatermark(value float64) error {
	return e.Element.SetProperty("high-watermark", value)
}

// low-watermark (float64)
//
// Low threshold for buffering to start. Only used if use-buffering is True

func (e *Urisourcebin) GetLowWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Urisourcebin) SetLowWatermark(value float64) error {
	return e.Element.SetProperty("low-watermark", value)
}

// parse-streams (bool)
//
// A parsebin element will be used on all non-raw streams, and urisourcebin
// will output the elementary streams. Recommended when buffering is used
// since it will provide accurate buffering levels.

func (e *Urisourcebin) GetParseStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("parse-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Urisourcebin) SetParseStreams(value bool) error {
	return e.Element.SetProperty("parse-streams", value)
}

// ring-buffer-max-size (uint64)
//
// Max. amount of data in the ring buffer (bytes, 0 = ring buffer disabled)

func (e *Urisourcebin) GetRingBufferMaxSize() (uint64, error) {
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

func (e *Urisourcebin) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

// source (GstElement)
//
// Source object used

func (e *Urisourcebin) GetSource() (interface{}, error) {
	return e.Element.GetProperty("source")
}

// statistics (GstStructure)
//
// A set of statistics over all the queue-like elements contained in this element

func (e *Urisourcebin) GetStatistics() (interface{}, error) {
	return e.Element.GetProperty("statistics")
}

// uri (string)
//
// URI to decode

func (e *Urisourcebin) GetUri() (string, error) {
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

func (e *Urisourcebin) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// use-buffering (bool)
//
// Perform buffering on demuxed/parsed media

func (e *Urisourcebin) GetUseBuffering() (bool, error) {
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

func (e *Urisourcebin) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

