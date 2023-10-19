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

type Rtmp2SinkexampleLaunchLine struct {
	*base.GstBaseSink
}

func NewRtmp2SinkexampleLaunchLine() (*Rtmp2SinkexampleLaunchLine, error) {
	e, err := gst.NewElement("rtmp2sinkExample launch line")
	if err != nil {
		return nil, err
	}
	return &Rtmp2SinkexampleLaunchLine{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewRtmp2SinkexampleLaunchLineWithName(name string) (*Rtmp2SinkexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("rtmp2sinkExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &Rtmp2SinkexampleLaunchLine{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// async-connect (bool)
//
// Connect on READY, otherwise on first push

func (e *Rtmp2SinkexampleLaunchLine) GetAsyncConnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async-connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtmp2SinkexampleLaunchLine) SetAsyncConnect(value bool) error {
	return e.Element.SetProperty("async-connect", value)
}

// chunk-size (uint)
//
// RTMP chunk size

func (e *Rtmp2SinkexampleLaunchLine) GetChunkSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("chunk-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtmp2SinkexampleLaunchLine) SetChunkSize(value uint) error {
	return e.Element.SetProperty("chunk-size", value)
}

// peak-kbps (uint)
//
// Bitrate in kbit/sec to pace outgoing packets

func (e *Rtmp2SinkexampleLaunchLine) GetPeakKbps() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("peak-kbps")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtmp2SinkexampleLaunchLine) SetPeakKbps(value uint) error {
	return e.Element.SetProperty("peak-kbps", value)
}

// stats (GstStructure)
//
// Retrieve a statistics structure

func (e *Rtmp2SinkexampleLaunchLine) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// stop-commands (GstRtmpStopCommands)
//
// Which commands (if any) to send on EOS event before closing connection

func (e *Rtmp2SinkexampleLaunchLine) GetStopCommands() (GstRtmpStopCommands, error) {
	var value GstRtmpStopCommands
	var ok bool
	v, err := e.Element.GetProperty("stop-commands")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtmpStopCommands)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtmpStopCommands")
	}
	return value, nil
}

func (e *Rtmp2SinkexampleLaunchLine) SetStopCommands(value GstRtmpStopCommands) error {
	e.Element.SetArg("stop-commands", string(value))
	return nil
}

// ----- Constants -----

type GstRtmpStopCommands string

const (
	GstRtmpStopCommandsNone GstRtmpStopCommands = "none" // none (0x00000000) – No command
	GstRtmpStopCommandsFcunpublish GstRtmpStopCommands = "fcunpublish" // fcunpublish (0x00000001) – FCUnpublish
	GstRtmpStopCommandsClosestream GstRtmpStopCommands = "closestream" // closestream (0x00000002) – closeStream
	GstRtmpStopCommandsDeletestream GstRtmpStopCommands = "deletestream" // deletestream (0x00000004) – deleteStream
)

