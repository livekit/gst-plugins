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
	"github.com/tinyzimmer/go-gst/gst"

	"github.com/livekit/gstplugins/base"
)

type Frei0RSrcNois0R struct {
	*base.GstPushSrc
}

func NewFrei0RSrcNois0R() (*Frei0RSrcNois0R, error) {
	e, err := gst.NewElement("frei0r-src-nois0r")
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcNois0R{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewFrei0RSrcNois0RWithName(name string) (*Frei0RSrcNois0R, error) {
	e, err := gst.NewElementWithName("frei0r-src-nois0r", name)
	if err != nil {
		return nil, err
	}
	return &Frei0RSrcNois0R{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

