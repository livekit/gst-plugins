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

type Uridecodebin struct {
	Element *gst.Element
}

func NewUridecodebin() (*Uridecodebin, error) {
	e, err := gst.NewElement("uridecodebin")
	if err != nil {
		return nil, err
	}
	return &Uridecodebin{Element: e}, nil
}

func NewUridecodebinWithName(name string) (*Uridecodebin, error) {
	e, err := gst.NewElementWithName("uridecodebin", name)
	if err != nil {
		return nil, err
	}
	return &Uridecodebin{Element: e}, nil
}

// ----- Properties -----

// buffer-duration (int64)
//
// Buffer duration when buffering streams (-1 default value)

func (e *Uridecodebin) GetBufferDuration() (int64, error) {
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

func (e *Uridecodebin) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

// buffer-size (int)
//
// Buffer size when buffering streams (-1 default value)

func (e *Uridecodebin) GetBufferSize() (int, error) {
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

func (e *Uridecodebin) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// caps (GstCaps)
//
// The caps on which to stop decoding. (NULL = default)

func (e *Uridecodebin) GetCaps() (*gst.Caps, error) {
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

func (e *Uridecodebin) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Uridecodebin) GetConnectionSpeed() (uint64, error) {
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

func (e *Uridecodebin) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// download (bool)
//
// For certain media type, enable download buffering.

func (e *Uridecodebin) GetDownload() (bool, error) {
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

func (e *Uridecodebin) SetDownload(value bool) error {
	return e.Element.SetProperty("download", value)
}

// expose-all-streams (bool)
//
// Expose streams of unknown type.

// force-sw-decoders (bool)
//
// While auto-plugging, if set to TRUE, those decoders within
// "Hardware" klass will be ignored. Otherwise they will be tried.

func (e *Uridecodebin) GetForceSwDecoders() (bool, error) {
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

func (e *Uridecodebin) SetForceSwDecoders(value bool) error {
	return e.Element.SetProperty("force-sw-decoders", value)
}

// post-stream-topology (bool)
//
// Post stream-topology messages on the bus every time the topology changes.

func (e *Uridecodebin) GetPostStreamTopology() (bool, error) {
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

func (e *Uridecodebin) SetPostStreamTopology(value bool) error {
	return e.Element.SetProperty("post-stream-topology", value)
}

// ring-buffer-max-size (uint64)
//
// The maximum size of the ring buffer in kilobytes. If set to 0, the ring
// buffer is disabled. Default is 0.

func (e *Uridecodebin) GetRingBufferMaxSize() (uint64, error) {
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

func (e *Uridecodebin) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

// source (GstElement)
//
// Source object used

func (e *Uridecodebin) GetSource() (interface{}, error) {
	return e.Element.GetProperty("source")
}

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Uridecodebin) GetSubtitleEncoding() (string, error) {
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

func (e *Uridecodebin) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// uri (string)
//
// URI to decode

func (e *Uridecodebin) GetUri() (string, error) {
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

func (e *Uridecodebin) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// use-buffering (bool)
//
// Emit BUFFERING messages based on low-/high-percent thresholds of the
// demuxed or parsed data.
// When download buffering is activated and used for the current media
// type, this property does nothing. Otherwise perform buffering on the
// demuxed or parsed media.

func (e *Uridecodebin) GetUseBuffering() (bool, error) {
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

func (e *Uridecodebin) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

// ----- Constants -----

type GstAutoplugSelectResult string

const (
	GstAutoplugSelectResultTry GstAutoplugSelectResult = "try" // try (0) – GST_AUTOPLUG_SELECT_TRY
	GstAutoplugSelectResultExpose GstAutoplugSelectResult = "expose" // expose (1) – GST_AUTOPLUG_SELECT_EXPOSE
	GstAutoplugSelectResultSkip GstAutoplugSelectResult = "skip" // skip (2) – GST_AUTOPLUG_SELECT_SKIP
)

