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

type Splitmuxsink struct {
	Element *gst.Element
}

func NewSplitmuxsink() (*Splitmuxsink, error) {
	e, err := gst.NewElement("splitmuxsink")
	if err != nil {
		return nil, err
	}
	return &Splitmuxsink{Element: e}, nil
}

func NewSplitmuxsinkWithName(name string) (*Splitmuxsink, error) {
	e, err := gst.NewElementWithName("splitmuxsink", name)
	if err != nil {
		return nil, err
	}
	return &Splitmuxsink{Element: e}, nil
}

// ----- Properties -----

// alignment-threshold (uint64)
//
// Allow non-reference streams to be that many ns before the reference stream

func (e *Splitmuxsink) GetAlignmentThreshold() (uint64, error) {
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

func (e *Splitmuxsink) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

// async-finalize (bool)
//
// Finalize each fragment asynchronously and start a new one

func (e *Splitmuxsink) GetAsyncFinalize() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async-finalize")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Splitmuxsink) SetAsyncFinalize(value bool) error {
	return e.Element.SetProperty("async-finalize", value)
}

// location (string)
//
// Format string pattern for the location of the files to write (e.g. video%%05d.mp4)

func (e *Splitmuxsink) GetLocation() (string, error) {
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

func (e *Splitmuxsink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-files (uint)
//
// Maximum number of files to keep on disk. Once the maximum is reached,old files start to be deleted to make room for new ones.

func (e *Splitmuxsink) GetMaxFiles() (uint, error) {
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

func (e *Splitmuxsink) SetMaxFiles(value uint) error {
	return e.Element.SetProperty("max-files", value)
}

// max-size-bytes (uint64)
//
// Max. amount of data per file (in bytes, 0=disable)

func (e *Splitmuxsink) GetMaxSizeBytes() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Splitmuxsink) SetMaxSizeBytes(value uint64) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

// max-size-time (uint64)
//
// Max. amount of time per file (in ns, 0=disable)

func (e *Splitmuxsink) GetMaxSizeTime() (uint64, error) {
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

func (e *Splitmuxsink) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

// max-size-timecode (string)
//
// Maximum difference in timecode between first and last frame. Separator is assumed to be ":" everywhere (e.g. 01:00:00:00). Will only be effective if a timecode track is present.

func (e *Splitmuxsink) GetMaxSizeTimecode() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("max-size-timecode")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Splitmuxsink) SetMaxSizeTimecode(value string) error {
	return e.Element.SetProperty("max-size-timecode", value)
}

// mux-overhead (float64)
//
// Extra size overhead of muxing (0.02 = 2%%)

func (e *Splitmuxsink) GetMuxOverhead() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("mux-overhead")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Splitmuxsink) SetMuxOverhead(value float64) error {
	return e.Element.SetProperty("mux-overhead", value)
}

// muxer (GstElement)
//
// The muxer element to use (NULL = default mp4mux). Valid only for async-finalize = FALSE

func (e *Splitmuxsink) GetMuxer() (interface{}, error) {
	return e.Element.GetProperty("muxer")
}

func (e *Splitmuxsink) SetMuxer(value interface{}) error {
	return e.Element.SetProperty("muxer", value)
}

// muxer-factory (string)
//
// The muxer element factory to use (default = mp4mux). Valid only for async-finalize = TRUE

func (e *Splitmuxsink) GetMuxerFactory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("muxer-factory")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Splitmuxsink) SetMuxerFactory(value string) error {
	return e.Element.SetProperty("muxer-factory", value)
}

// muxer-pad-map (GstStructure)
//
// A GstStructure specifies the mapping from splitmuxsink sink pads to muxer pads

func (e *Splitmuxsink) GetMuxerPadMap() (interface{}, error) {
	return e.Element.GetProperty("muxer-pad-map")
}

func (e *Splitmuxsink) SetMuxerPadMap(value interface{}) error {
	return e.Element.SetProperty("muxer-pad-map", value)
}

// muxer-preset (string)
//
// An optional GstPreset name to use for the muxer. This only has an effect
// in async-finalize=TRUE mode.

func (e *Splitmuxsink) GetMuxerPreset() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("muxer-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Splitmuxsink) SetMuxerPreset(value string) error {
	return e.Element.SetProperty("muxer-preset", value)
}

// muxer-properties (GstStructure)
//
// The muxer element properties to use. Example: {properties,boolean-prop=true,string-prop="hi"}. Valid only for async-finalize = TRUE

func (e *Splitmuxsink) GetMuxerProperties() (interface{}, error) {
	return e.Element.GetProperty("muxer-properties")
}

func (e *Splitmuxsink) SetMuxerProperties(value interface{}) error {
	return e.Element.SetProperty("muxer-properties", value)
}

// reset-muxer (bool)
//
// Reset the muxer after each segment. Disabling this will not work for most muxers.

func (e *Splitmuxsink) GetResetMuxer() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reset-muxer")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Splitmuxsink) SetResetMuxer(value bool) error {
	return e.Element.SetProperty("reset-muxer", value)
}

// send-keyframe-requests (bool)
//
// Request a keyframe every max-size-time ns to try splitting at that point. Needs max-size-bytes to be 0 in order to be effective.

func (e *Splitmuxsink) GetSendKeyframeRequests() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-keyframe-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Splitmuxsink) SetSendKeyframeRequests(value bool) error {
	return e.Element.SetProperty("send-keyframe-requests", value)
}

// sink (GstElement)
//
// The sink element (or element chain) to use (NULL = default filesink). Valid only for async-finalize = FALSE

func (e *Splitmuxsink) GetSink() (interface{}, error) {
	return e.Element.GetProperty("sink")
}

func (e *Splitmuxsink) SetSink(value interface{}) error {
	return e.Element.SetProperty("sink", value)
}

// sink-factory (string)
//
// The sink element factory to use (default = filesink). Valid only for async-finalize = TRUE

func (e *Splitmuxsink) GetSinkFactory() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sink-factory")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Splitmuxsink) SetSinkFactory(value string) error {
	return e.Element.SetProperty("sink-factory", value)
}

// sink-preset (string)
//
// An optional GstPreset name to use for the sink. This only has an effect
// in async-finalize=TRUE mode.

func (e *Splitmuxsink) GetSinkPreset() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sink-preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Splitmuxsink) SetSinkPreset(value string) error {
	return e.Element.SetProperty("sink-preset", value)
}

// sink-properties (GstStructure)
//
// The sink element properties to use. Example: {properties,boolean-prop=true,string-prop="hi"}. Valid only for async-finalize = TRUE

func (e *Splitmuxsink) GetSinkProperties() (interface{}, error) {
	return e.Element.GetProperty("sink-properties")
}

func (e *Splitmuxsink) SetSinkProperties(value interface{}) error {
	return e.Element.SetProperty("sink-properties", value)
}

// start-index (int)
//
// Start value of fragment index.

func (e *Splitmuxsink) GetStartIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("start-index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Splitmuxsink) SetStartIndex(value int) error {
	return e.Element.SetProperty("start-index", value)
}

// use-robust-muxing (bool)
//
// Check if muxers support robust muxing via the reserved-max-duration and reserved-duration-remaining properties and use them if so. (Only present on qtmux and mp4mux for now). splitmuxsink may then also  create new fragments if the reserved header space is about to overflow. Note that for mp4mux and qtmux, reserved-moov-update-period must be set manually by the app to a non-zero value for robust muxing to have an effect.

func (e *Splitmuxsink) GetUseRobustMuxing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-robust-muxing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Splitmuxsink) SetUseRobustMuxing(value bool) error {
	return e.Element.SetProperty("use-robust-muxing", value)
}

