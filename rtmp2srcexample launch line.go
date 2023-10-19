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

type Rtmp2SrcexampleLaunchLine struct {
	*base.GstPushSrc
}

func NewRtmp2SrcexampleLaunchLine() (*Rtmp2SrcexampleLaunchLine, error) {
	e, err := gst.NewElement("rtmp2srcExample launch line")
	if err != nil {
		return nil, err
	}
	return &Rtmp2SrcexampleLaunchLine{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewRtmp2SrcexampleLaunchLineWithName(name string) (*Rtmp2SrcexampleLaunchLine, error) {
	e, err := gst.NewElementWithName("rtmp2srcExample launch line", name)
	if err != nil {
		return nil, err
	}
	return &Rtmp2SrcexampleLaunchLine{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// async-connect (bool)
//
// Connect on READY, otherwise on first push

func (e *Rtmp2SrcexampleLaunchLine) GetAsyncConnect() (bool, error) {
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

func (e *Rtmp2SrcexampleLaunchLine) SetAsyncConnect(value bool) error {
	return e.Element.SetProperty("async-connect", value)
}

// idle-timeout (uint)
//
// The maximum allowed time in seconds for valid packets not to arrive from the peer (0 = no timeout)

func (e *Rtmp2SrcexampleLaunchLine) GetIdleTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("idle-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Rtmp2SrcexampleLaunchLine) SetIdleTimeout(value uint) error {
	return e.Element.SetProperty("idle-timeout", value)
}

// no-eof-is-error (bool)
//
// If set, an error is raised if the connection is closed without receiving an EOF RTMP message first.
// " If not set, those are reported using EOS.

func (e *Rtmp2SrcexampleLaunchLine) GetNoEofIsError() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("no-eof-is-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Rtmp2SrcexampleLaunchLine) SetNoEofIsError(value bool) error {
	return e.Element.SetProperty("no-eof-is-error", value)
}

// stats (GstStructure)
//
// Retrieve a statistics structure

func (e *Rtmp2SrcexampleLaunchLine) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

