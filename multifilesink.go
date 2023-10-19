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

type Multifilesink struct {
	*base.GstBaseSink
}

func NewMultifilesink() (*Multifilesink, error) {
	e, err := gst.NewElement("multifilesink")
	if err != nil {
		return nil, err
	}
	return &Multifilesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewMultifilesinkWithName(name string) (*Multifilesink, error) {
	e, err := gst.NewElementWithName("multifilesink", name)
	if err != nil {
		return nil, err
	}
	return &Multifilesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// aggregate-gops (bool)
//
// Whether to aggregate complete GOPs before doing any processing. Set this
// to TRUE to make sure each new file starts with a keyframe. This requires
// the upstream element to flag buffers containing key units and delta
// units correctly. At least the MPEG-PS and MPEG-TS muxers should be doing
// this.

func (e *Multifilesink) GetAggregateGops() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aggregate-gops")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multifilesink) SetAggregateGops(value bool) error {
	return e.Element.SetProperty("aggregate-gops", value)
}

// index (int)
//
// Index to use with location property to create file names.  The index is incremented by one for each buffer written.

func (e *Multifilesink) GetIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Multifilesink) SetIndex(value int) error {
	return e.Element.SetProperty("index", value)
}

// location (string)
//
// Location of the file to write

func (e *Multifilesink) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Multifilesink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-file-duration (uint64)
//
// Maximum file size before starting a new file in max-size mode.

func (e *Multifilesink) GetMaxFileDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-file-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Multifilesink) SetMaxFileDuration(value uint64) error {
	return e.Element.SetProperty("max-file-duration", value)
}

// max-file-size (uint64)
//
// Maximum file size before starting a new file in max-size mode.

func (e *Multifilesink) GetMaxFileSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-file-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Multifilesink) SetMaxFileSize(value uint64) error {
	return e.Element.SetProperty("max-file-size", value)
}

// max-files (uint)
//
// Maximum number of files to keep on disk. Once the maximum is reached, old
// files start to be deleted to make room for new ones.

func (e *Multifilesink) GetMaxFiles() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-files")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Multifilesink) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

// min-keyframe-distance (uint64)
//
// Minimum distance between keyframes in next-file=key-frame that causes a
// new file to be created. If two keyframes arrive closer to each other they
// will end up in the same file.

func (e *Multifilesink) GetMinKeyframeDistance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-keyframe-distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Multifilesink) SetMinKeyframeDistance(value uint64) error {
	return e.Element.SetProperty("min-keyframe-distance", value)
}

// next-file (GstMultiFileSinkNext)
//
// When to start a new file.

func (e *Multifilesink) GetNextFile() (GstMultiFileSinkNext, error) {
	var value GstMultiFileSinkNext
	var ok bool
	v, err := e.Element.GetProperty("next-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMultiFileSinkNext)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMultiFileSinkNext")
	}
	return value, nil
}

func (e *Multifilesink) SetNextFile(value GstMultiFileSinkNext) error {
	e.Element.SetArg("next-file", string(value))
	return nil
}

// post-messages (bool)
//
// Post a message on the GstBus for each file.

func (e *Multifilesink) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Multifilesink) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

// ----- Constants -----

type GstMultiFileSinkNext string

const (
	GstMultiFileSinkNextBuffer GstMultiFileSinkNext = "buffer" // buffer (0) – New file for each buffer
	GstMultiFileSinkNextDiscont GstMultiFileSinkNext = "discont" // discont (1) – New file after each discontinuity
	GstMultiFileSinkNextKeyFrame GstMultiFileSinkNext = "key-frame" // key-frame (2) – New file at each key frame (Useful for MPEG-TS segmenting)
	GstMultiFileSinkNextKeyUnitEvent GstMultiFileSinkNext = "key-unit-event" // key-unit-event (3) – New file after a force key unit event
	GstMultiFileSinkNextMaxSize GstMultiFileSinkNext = "max-size" // max-size (4) – New file when the configured maximum file size would be exceeded with the next buffer or buffer list
	GstMultiFileSinkNextMaxDuration GstMultiFileSinkNext = "max-duration" // max-duration (5) – New file when the configured maximum file duration would be exceeded with the next buffer or buffer list
)

