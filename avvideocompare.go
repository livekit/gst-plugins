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

type Avvideocompare struct {
	Element *gst.Element
}

func NewAvvideocompare() (*Avvideocompare, error) {
	e, err := gst.NewElement("avvideocompare")
	if err != nil {
		return nil, err
	}
	return &Avvideocompare{Element: e}, nil
}

func NewAvvideocompareWithName(name string) (*Avvideocompare, error) {
	e, err := gst.NewElementWithName("avvideocompare", name)
	if err != nil {
		return nil, err
	}
	return &Avvideocompare{Element: e}, nil
}

// ----- Properties -----

// method (GstFfmpegVidCmpMethod)
//
// Method to compare video frames

func (e *Avvideocompare) GetMethod() (GstFfmpegVidCmpMethod, error) {
	var value GstFfmpegVidCmpMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFfmpegVidCmpMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFfmpegVidCmpMethod")
	}
	return value, nil
}

func (e *Avvideocompare) SetMethod(value GstFfmpegVidCmpMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

// stats-file (string)
//
// Set file where to store per-frame difference information, '-' for stdout

func (e *Avvideocompare) GetStatsFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("stats-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Avvideocompare) SetStatsFile(value string) error {
	return e.Element.SetProperty("stats-file", value)
}

// ----- Constants -----

type GstFfmpegVidCmpMethod string

const (
	GstFfmpegVidCmpMethodSsim GstFfmpegVidCmpMethod = "ssim" // ssim (0) – SSIM
	GstFfmpegVidCmpMethodPsnr GstFfmpegVidCmpMethod = "psnr" // psnr (1) – PSNR
)

