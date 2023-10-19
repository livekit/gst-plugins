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

type H264Parse struct {
	*base.GstBaseParse
}

func NewH264Parse() (*H264Parse, error) {
	e, err := gst.NewElement("h264parse")
	if err != nil {
		return nil, err
	}
	return &H264Parse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

func NewH264ParseWithName(name string) (*H264Parse, error) {
	e, err := gst.NewElementWithName("h264parse", name)
	if err != nil {
		return nil, err
	}
	return &H264Parse{GstBaseParse: &base.GstBaseParse{Element: e}}, nil
}

// ----- Properties -----

// config-interval (int)
//
// Send SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)

func (e *H264Parse) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *H264Parse) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

// update-timecode (bool)
//
// If the stream contains Picture Timing SEI, update their timecode values
// using upstream GstVideoTimeCodeMeta. However, if there are no Picture
// Timing SEI in bitstream, this property will not insert the SEI into the
// bitstream - it only modifies existing ones.
// Moreover, even if both GstVideoTimeCodeMeta and Picture Timing SEI
// are present, if pic_struct_present_flag of VUI is equal to zero,
// timecode values will not updated as there is not enough information
// in the stream to do so.

func (e *H264Parse) GetUpdateTimecode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("update-timecode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *H264Parse) SetUpdateTimecode(value bool) error {
	return e.Element.SetProperty("update-timecode", value)
}

