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

type LspPlugInPluginsLv2LimiterMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2LimiterMono() (*LspPlugInPluginsLv2LimiterMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-limiter-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2LimiterMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2LimiterMonoWithName(name string) (*LspPlugInPluginsLv2LimiterMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-limiter-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2LimiterMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// at (float32)
//
// Attack time

func (e *LspPlugInPluginsLv2LimiterMono) GetAt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// boost (bool)
//
// Gain boost

func (e *LspPlugInPluginsLv2LimiterMono) GetBoost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("boost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetBoost(value bool) error {
	return e.Element.SetProperty("boost", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2LimiterMono) GetBypass() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bypass")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2LimiterMono) GetClear() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clear")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// dith (GstLspPlugInPluginsLv2LimiterMonodith)
//
// Dithering

func (e *LspPlugInPluginsLv2LimiterMono) GetDith() (interface{}, error) {
	return e.Element.GetProperty("dith")
}

func (e *LspPlugInPluginsLv2LimiterMono) SetDith(value interface{}) error {
	return e.Element.SetProperty("dith", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2LimiterMono) GetGIn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-in")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2LimiterMono) GetGOut() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-out")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grgv (bool)
//
// Gain graph visibility

func (e *LspPlugInPluginsLv2LimiterMono) GetGrgv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grgv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetGrgv(value bool) error {
	return e.Element.SetProperty("grgv", value)
}

// grlm (float32)
//
// Gain reduction level meter

func (e *LspPlugInPluginsLv2LimiterMono) GetGrlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("grlm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// igv (bool)
//
// Input graph visibility

func (e *LspPlugInPluginsLv2LimiterMono) GetIgv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("igv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetIgv(value bool) error {
	return e.Element.SetProperty("igv", value)
}

// ilm (float32)
//
// Input level meter

func (e *LspPlugInPluginsLv2LimiterMono) GetIlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// knee (float32)
//
// Knee

func (e *LspPlugInPluginsLv2LimiterMono) GetKnee() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("knee")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetKnee(value float32) error {
	return e.Element.SetProperty("knee", value)
}

// lk (float32)
//
// Lookahead

func (e *LspPlugInPluginsLv2LimiterMono) GetLk() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("lk")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetLk(value float32) error {
	return e.Element.SetProperty("lk", value)
}

// mode (GstLspPlugInPluginsLv2LimiterMonomode)
//
// Operating mode

func (e *LspPlugInPluginsLv2LimiterMono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2LimiterMono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ogv (bool)
//
// Output graph visibility

func (e *LspPlugInPluginsLv2LimiterMono) GetOgv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ogv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetOgv(value bool) error {
	return e.Element.SetProperty("ogv", value)
}

// olm (float32)
//
// Outut level meter

func (e *LspPlugInPluginsLv2LimiterMono) GetOlm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2LimiterMono) GetOutLatency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("out-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// ovs (GstLspPlugInPluginsLv2LimiterMonoovs)
//
// Oversampling

func (e *LspPlugInPluginsLv2LimiterMono) GetOvs() (interface{}, error) {
	return e.Element.GetProperty("ovs")
}

func (e *LspPlugInPluginsLv2LimiterMono) SetOvs(value interface{}) error {
	return e.Element.SetProperty("ovs", value)
}

// pause (bool)
//
// Pause graph analysis

func (e *LspPlugInPluginsLv2LimiterMono) GetPause() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("pause")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2LimiterMono) GetRt() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scgv (bool)
//
// Sidechain graph visibility

func (e *LspPlugInPluginsLv2LimiterMono) GetScgv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scgv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetScgv(value bool) error {
	return e.Element.SetProperty("scgv", value)
}

// sclm (float32)
//
// Sidechain level meter

func (e *LspPlugInPluginsLv2LimiterMono) GetSclm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// scp (float32)
//
// Sidechain preamp

func (e *LspPlugInPluginsLv2LimiterMono) GetScp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// th (float32)
//
// Threshold

func (e *LspPlugInPluginsLv2LimiterMono) GetTh() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("th")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterMono) SetTh(value float32) error {
	return e.Element.SetProperty("th", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2LimiterMonodith string

const (
	LspPlugInPluginsLv2LimiterMonodithNone LspPlugInPluginsLv2LimiterMonodith = "None" // None (0) – None
	LspPlugInPluginsLv2LimiterMonodith7Bit LspPlugInPluginsLv2LimiterMonodith = "7bit" // 7bit (1) – 7bit
	LspPlugInPluginsLv2LimiterMonodith8Bit LspPlugInPluginsLv2LimiterMonodith = "8bit" // 8bit (2) – 8bit
	LspPlugInPluginsLv2LimiterMonodith11Bit LspPlugInPluginsLv2LimiterMonodith = "11bit" // 11bit (3) – 11bit
	LspPlugInPluginsLv2LimiterMonodith12Bit LspPlugInPluginsLv2LimiterMonodith = "12bit" // 12bit (4) – 12bit
	LspPlugInPluginsLv2LimiterMonodith15Bit LspPlugInPluginsLv2LimiterMonodith = "15bit" // 15bit (5) – 15bit
	LspPlugInPluginsLv2LimiterMonodith16Bit LspPlugInPluginsLv2LimiterMonodith = "16bit" // 16bit (6) – 16bit
	LspPlugInPluginsLv2LimiterMonodith23Bit LspPlugInPluginsLv2LimiterMonodith = "23bit" // 23bit (7) – 23bit
	LspPlugInPluginsLv2LimiterMonodith24Bit LspPlugInPluginsLv2LimiterMonodith = "24bit" // 24bit (8) – 24bit
)

type LspPlugInPluginsLv2LimiterMonomode string

const (
	LspPlugInPluginsLv2LimiterMonomodeClassic LspPlugInPluginsLv2LimiterMonomode = "Classic" // Classic (0) – Classic
	LspPlugInPluginsLv2LimiterMonomodeHermThin LspPlugInPluginsLv2LimiterMonomode = "Herm Thin" // Herm Thin (1) – Herm Thin
	LspPlugInPluginsLv2LimiterMonomodeHermWide LspPlugInPluginsLv2LimiterMonomode = "Herm Wide" // Herm Wide (2) – Herm Wide
	LspPlugInPluginsLv2LimiterMonomodeHermTail LspPlugInPluginsLv2LimiterMonomode = "Herm Tail" // Herm Tail (3) – Herm Tail
	LspPlugInPluginsLv2LimiterMonomodeHermDuck LspPlugInPluginsLv2LimiterMonomode = "Herm Duck" // Herm Duck (4) – Herm Duck
	LspPlugInPluginsLv2LimiterMonomodeExpThin LspPlugInPluginsLv2LimiterMonomode = "Exp Thin" // Exp Thin (5) – Exp Thin
	LspPlugInPluginsLv2LimiterMonomodeExpWide LspPlugInPluginsLv2LimiterMonomode = "Exp Wide" // Exp Wide (6) – Exp Wide
	LspPlugInPluginsLv2LimiterMonomodeExpTail LspPlugInPluginsLv2LimiterMonomode = "Exp Tail" // Exp Tail (7) – Exp Tail
	LspPlugInPluginsLv2LimiterMonomodeExpDuck LspPlugInPluginsLv2LimiterMonomode = "Exp Duck" // Exp Duck (8) – Exp Duck
	LspPlugInPluginsLv2LimiterMonomodeLineThin LspPlugInPluginsLv2LimiterMonomode = "Line Thin" // Line Thin (9) – Line Thin
	LspPlugInPluginsLv2LimiterMonomodeLineWide LspPlugInPluginsLv2LimiterMonomode = "Line Wide" // Line Wide (10) – Line Wide
	LspPlugInPluginsLv2LimiterMonomodeLineTail LspPlugInPluginsLv2LimiterMonomode = "Line Tail" // Line Tail (11) – Line Tail
	LspPlugInPluginsLv2LimiterMonomodeLineDuck LspPlugInPluginsLv2LimiterMonomode = "Line Duck" // Line Duck (12) – Line Duck
	LspPlugInPluginsLv2LimiterMonomodeMixedHerm LspPlugInPluginsLv2LimiterMonomode = "Mixed Herm" // Mixed Herm (13) – Mixed Herm
	LspPlugInPluginsLv2LimiterMonomodeMixedExp LspPlugInPluginsLv2LimiterMonomode = "Mixed Exp" // Mixed Exp (14) – Mixed Exp
	LspPlugInPluginsLv2LimiterMonomodeMixedLine LspPlugInPluginsLv2LimiterMonomode = "Mixed Line" // Mixed Line (15) – Mixed Line
)

type LspPlugInPluginsLv2LimiterMonoovs string

const (
	LspPlugInPluginsLv2LimiterMonoovsNone LspPlugInPluginsLv2LimiterMonoovs = "None" // None (0) – None
	LspPlugInPluginsLv2LimiterMonoovsHalfX22L LspPlugInPluginsLv2LimiterMonoovs = "Half x2(2L)" // Half x2(2L) (1) – Half x2(2L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX23L LspPlugInPluginsLv2LimiterMonoovs = "Half x2(3L)" // Half x2(3L) (2) – Half x2(3L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX32L LspPlugInPluginsLv2LimiterMonoovs = "Half x3(2L)" // Half x3(2L) (3) – Half x3(2L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX33L LspPlugInPluginsLv2LimiterMonoovs = "Half x3(3L)" // Half x3(3L) (4) – Half x3(3L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX42L LspPlugInPluginsLv2LimiterMonoovs = "Half x4(2L)" // Half x4(2L) (5) – Half x4(2L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX43L LspPlugInPluginsLv2LimiterMonoovs = "Half x4(3L)" // Half x4(3L) (6) – Half x4(3L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX62L LspPlugInPluginsLv2LimiterMonoovs = "Half x6(2L)" // Half x6(2L) (7) – Half x6(2L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX63L LspPlugInPluginsLv2LimiterMonoovs = "Half x6(3L)" // Half x6(3L) (8) – Half x6(3L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX82L LspPlugInPluginsLv2LimiterMonoovs = "Half x8(2L)" // Half x8(2L) (9) – Half x8(2L)
	LspPlugInPluginsLv2LimiterMonoovsHalfX83L LspPlugInPluginsLv2LimiterMonoovs = "Half x8(3L)" // Half x8(3L) (10) – Half x8(3L)
	LspPlugInPluginsLv2LimiterMonoovsFullX22L LspPlugInPluginsLv2LimiterMonoovs = "Full x2(2L)" // Full x2(2L) (11) – Full x2(2L)
	LspPlugInPluginsLv2LimiterMonoovsFullX23L LspPlugInPluginsLv2LimiterMonoovs = "Full x2(3L)" // Full x2(3L) (12) – Full x2(3L)
	LspPlugInPluginsLv2LimiterMonoovsFullX32L LspPlugInPluginsLv2LimiterMonoovs = "Full x3(2L)" // Full x3(2L) (13) – Full x3(2L)
	LspPlugInPluginsLv2LimiterMonoovsFullX33L LspPlugInPluginsLv2LimiterMonoovs = "Full x3(3L)" // Full x3(3L) (14) – Full x3(3L)
	LspPlugInPluginsLv2LimiterMonoovsFullX42L LspPlugInPluginsLv2LimiterMonoovs = "Full x4(2L)" // Full x4(2L) (15) – Full x4(2L)
	LspPlugInPluginsLv2LimiterMonoovsFullX43L LspPlugInPluginsLv2LimiterMonoovs = "Full x4(3L)" // Full x4(3L) (16) – Full x4(3L)
	LspPlugInPluginsLv2LimiterMonoovsFullX62L LspPlugInPluginsLv2LimiterMonoovs = "Full x6(2L)" // Full x6(2L) (17) – Full x6(2L)
	LspPlugInPluginsLv2LimiterMonoovsFullX63L LspPlugInPluginsLv2LimiterMonoovs = "Full x6(3L)" // Full x6(3L) (18) – Full x6(3L)
	LspPlugInPluginsLv2LimiterMonoovsFullX82L LspPlugInPluginsLv2LimiterMonoovs = "Full x8(2L)" // Full x8(2L) (19) – Full x8(2L)
	LspPlugInPluginsLv2LimiterMonoovsFullX83L LspPlugInPluginsLv2LimiterMonoovs = "Full x8(3L)" // Full x8(3L) (20) – Full x8(3L)
)

