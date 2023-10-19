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

type Mplex struct {
	Element *gst.Element
}

func NewMplex() (*Mplex, error) {
	e, err := gst.NewElement("mplex")
	if err != nil {
		return nil, err
	}
	return &Mplex{Element: e}, nil
}

func NewMplexWithName(name string) (*Mplex, error) {
	e, err := gst.NewElementWithName("mplex", name)
	if err != nil {
		return nil, err
	}
	return &Mplex{Element: e}, nil
}

// ----- Properties -----

// bufsize (int)
//
// Target decoders video buffer size (kB) [default determined by format if not explicitly set]

func (e *Mplex) GetBufsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bufsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mplex) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// format (GstMplexFormat)
//
// Encoding profile format

func (e *Mplex) GetFormat() (GstMplexFormat, error) {
	var value GstMplexFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMplexFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMplexFormat")
	}
	return value, nil
}

func (e *Mplex) SetFormat(value GstMplexFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

// mux-bitrate (int)
//
// Bitrate of output stream in kbps (0 = autodetect)

func (e *Mplex) GetMuxBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mux-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mplex) SetMuxBitrate(value int) error {
	return e.Element.SetProperty("mux-bitrate", value)
}

// packets-per-pack (int)
//
// Number of packets per pack for generic formats

func (e *Mplex) GetPacketsPerPack() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("packets-per-pack")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mplex) SetPacketsPerPack(value int) error {
	return e.Element.SetProperty("packets-per-pack", value)
}

// sector-size (int)
//
// Specify sector size in bytes for generic formats

func (e *Mplex) GetSectorSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sector-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Mplex) SetSectorSize(value int) error {
	return e.Element.SetProperty("sector-size", value)
}

// system-headers (bool)
//
// Create system header in every pack for generic formats

func (e *Mplex) GetSystemHeaders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("system-headers")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mplex) SetSystemHeaders(value bool) error {
	return e.Element.SetProperty("system-headers", value)
}

// vbr (bool)
//
// Whether the input video stream is variable bitrate

func (e *Mplex) GetVbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Mplex) SetVbr(value bool) error {
	return e.Element.SetProperty("vbr", value)
}

// ----- Constants -----

type GstMplexFormat string

const (
	GstMplexFormatMpeg1 GstMplexFormat = "mpeg-1" // mpeg-1 (0) – Generic MPEG-1
	GstMplexFormatVcd GstMplexFormat = "vcd" // vcd (1) – Standard VCD
	GstMplexFormatVcdNsr GstMplexFormat = "vcd-nsr" // vcd-nsr (2) – User VCD
	GstMplexFormatMpeg2 GstMplexFormat = "mpeg-2" // mpeg-2 (3) – Generic MPEG-2
	GstMplexFormatSvcd GstMplexFormat = "svcd" // svcd (4) – Standard SVCD
	GstMplexFormatSvcdNsr GstMplexFormat = "svcd-nsr" // svcd-nsr (5) – User SVCD
	GstMplexFormatVcdStill GstMplexFormat = "vcd-still" // vcd-still (6) – VCD Stills sequences
	GstMplexFormatSvcdStill GstMplexFormat = "svcd-still" // svcd-still (7) – SVCD Stills sequences
	GstMplexFormatDvdNav GstMplexFormat = "dvd-nav" // dvd-nav (8) – DVD MPEG-2 for dvdauthor
	GstMplexFormatDvd GstMplexFormat = "dvd" // dvd (9) – DVD MPEG-2
	GstMplexFormatAtsc480I GstMplexFormat = "atsc-480i" // atsc-480i (10) – ATSC 480i
	GstMplexFormatAtsc480P GstMplexFormat = "atsc-480p" // atsc-480p (11) – ATSC 480p
	GstMplexFormatAtsc720P GstMplexFormat = "atsc-720p" // atsc-720p (12) – ATSC 720p
	GstMplexFormatAtsc1080I GstMplexFormat = "atsc-1080i" // atsc-1080i (13) – ATSC 1080i
)

