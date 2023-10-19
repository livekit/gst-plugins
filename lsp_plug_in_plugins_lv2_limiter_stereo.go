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

type LspPlugInPluginsLv2LimiterStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2LimiterStereo() (*LspPlugInPluginsLv2LimiterStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-limiter-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2LimiterStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2LimiterStereoWithName(name string) (*LspPlugInPluginsLv2LimiterStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-limiter-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2LimiterStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// at (float32)
//
// Attack time

func (e *LspPlugInPluginsLv2LimiterStereo) GetAt() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetAt(value float32) error {
	return e.Element.SetProperty("at", value)
}

// boost (bool)
//
// Gain boost

func (e *LspPlugInPluginsLv2LimiterStereo) GetBoost() (bool, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetBoost(value bool) error {
	return e.Element.SetProperty("boost", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2LimiterStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// clear (bool)
//
// Clear graph analysis

func (e *LspPlugInPluginsLv2LimiterStereo) GetClear() (bool, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetClear(value bool) error {
	return e.Element.SetProperty("clear", value)
}

// dith (GstLspPlugInPluginsLv2LimiterStereodith)
//
// Dithering

func (e *LspPlugInPluginsLv2LimiterStereo) GetDith() (interface{}, error) {
	return e.Element.GetProperty("dith")
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetDith(value interface{}) error {
	return e.Element.SetProperty("dith", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2LimiterStereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2LimiterStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// grgv-l (bool)
//
// Gain graph visibility Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetGrgvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grgv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetGrgvL(value bool) error {
	return e.Element.SetProperty("grgv-l", value)
}

// grgv-r (bool)
//
// Gain graph visibility Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetGrgvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("grgv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetGrgvR(value bool) error {
	return e.Element.SetProperty("grgv-r", value)
}

// grlm-l (float32)
//
// Gain reduction level meter Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetGrlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("grlm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// grlm-r (float32)
//
// Gain reduction level meter Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetGrlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("grlm-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// igv-l (bool)
//
// Input graph visibility Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetIgvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("igv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetIgvL(value bool) error {
	return e.Element.SetProperty("igv-l", value)
}

// igv-r (bool)
//
// Input graph visibility Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetIgvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("igv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetIgvR(value bool) error {
	return e.Element.SetProperty("igv-r", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetIlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// ilm-r (float32)
//
// Input level meter Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetIlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ilm-r")
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

func (e *LspPlugInPluginsLv2LimiterStereo) GetKnee() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetKnee(value float32) error {
	return e.Element.SetProperty("knee", value)
}

// lk (float32)
//
// Lookahead

func (e *LspPlugInPluginsLv2LimiterStereo) GetLk() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetLk(value float32) error {
	return e.Element.SetProperty("lk", value)
}

// mode (GstLspPlugInPluginsLv2LimiterStereomode)
//
// Operating mode

func (e *LspPlugInPluginsLv2LimiterStereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ogv-l (bool)
//
// Output graph visibility Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetOgvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ogv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetOgvL(value bool) error {
	return e.Element.SetProperty("ogv-l", value)
}

// ogv-r (bool)
//
// Output graph visibility Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetOgvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ogv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetOgvR(value bool) error {
	return e.Element.SetProperty("ogv-r", value)
}

// olm-l (float32)
//
// Outut level meter Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetOlmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// olm-r (float32)
//
// Outut level meter Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetOlmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("olm-r")
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

func (e *LspPlugInPluginsLv2LimiterStereo) GetOutLatency() (int, error) {
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

// ovs (GstLspPlugInPluginsLv2LimiterStereoovs)
//
// Oversampling

func (e *LspPlugInPluginsLv2LimiterStereo) GetOvs() (interface{}, error) {
	return e.Element.GetProperty("ovs")
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetOvs(value interface{}) error {
	return e.Element.SetProperty("ovs", value)
}

// pause (bool)
//
// Pause graph analysis

func (e *LspPlugInPluginsLv2LimiterStereo) GetPause() (bool, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetPause(value bool) error {
	return e.Element.SetProperty("pause", value)
}

// rt (float32)
//
// Release time

func (e *LspPlugInPluginsLv2LimiterStereo) GetRt() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetRt(value float32) error {
	return e.Element.SetProperty("rt", value)
}

// scgv-l (bool)
//
// Sidechain graph visibility Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetScgvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scgv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetScgvL(value bool) error {
	return e.Element.SetProperty("scgv-l", value)
}

// scgv-r (bool)
//
// Sidechain graph visibility Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetScgvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("scgv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetScgvR(value bool) error {
	return e.Element.SetProperty("scgv-r", value)
}

// sclm-l (float32)
//
// Sidechain level meter Left

func (e *LspPlugInPluginsLv2LimiterStereo) GetSclmL() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclm-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// sclm-r (float32)
//
// Sidechain level meter Right

func (e *LspPlugInPluginsLv2LimiterStereo) GetSclmR() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclm-r")
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

func (e *LspPlugInPluginsLv2LimiterStereo) GetScp() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetScp(value float32) error {
	return e.Element.SetProperty("scp", value)
}

// slink (float32)
//
// Stereo linking

func (e *LspPlugInPluginsLv2LimiterStereo) GetSlink() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("slink")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2LimiterStereo) SetSlink(value float32) error {
	return e.Element.SetProperty("slink", value)
}

// th (float32)
//
// Threshold

func (e *LspPlugInPluginsLv2LimiterStereo) GetTh() (float32, error) {
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

func (e *LspPlugInPluginsLv2LimiterStereo) SetTh(value float32) error {
	return e.Element.SetProperty("th", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2LimiterStereomode string

const (
	LspPlugInPluginsLv2LimiterStereomodeClassic LspPlugInPluginsLv2LimiterStereomode = "Classic" // Classic (0) – Classic
	LspPlugInPluginsLv2LimiterStereomodeHermThin LspPlugInPluginsLv2LimiterStereomode = "Herm Thin" // Herm Thin (1) – Herm Thin
	LspPlugInPluginsLv2LimiterStereomodeHermWide LspPlugInPluginsLv2LimiterStereomode = "Herm Wide" // Herm Wide (2) – Herm Wide
	LspPlugInPluginsLv2LimiterStereomodeHermTail LspPlugInPluginsLv2LimiterStereomode = "Herm Tail" // Herm Tail (3) – Herm Tail
	LspPlugInPluginsLv2LimiterStereomodeHermDuck LspPlugInPluginsLv2LimiterStereomode = "Herm Duck" // Herm Duck (4) – Herm Duck
	LspPlugInPluginsLv2LimiterStereomodeExpThin LspPlugInPluginsLv2LimiterStereomode = "Exp Thin" // Exp Thin (5) – Exp Thin
	LspPlugInPluginsLv2LimiterStereomodeExpWide LspPlugInPluginsLv2LimiterStereomode = "Exp Wide" // Exp Wide (6) – Exp Wide
	LspPlugInPluginsLv2LimiterStereomodeExpTail LspPlugInPluginsLv2LimiterStereomode = "Exp Tail" // Exp Tail (7) – Exp Tail
	LspPlugInPluginsLv2LimiterStereomodeExpDuck LspPlugInPluginsLv2LimiterStereomode = "Exp Duck" // Exp Duck (8) – Exp Duck
	LspPlugInPluginsLv2LimiterStereomodeLineThin LspPlugInPluginsLv2LimiterStereomode = "Line Thin" // Line Thin (9) – Line Thin
	LspPlugInPluginsLv2LimiterStereomodeLineWide LspPlugInPluginsLv2LimiterStereomode = "Line Wide" // Line Wide (10) – Line Wide
	LspPlugInPluginsLv2LimiterStereomodeLineTail LspPlugInPluginsLv2LimiterStereomode = "Line Tail" // Line Tail (11) – Line Tail
	LspPlugInPluginsLv2LimiterStereomodeLineDuck LspPlugInPluginsLv2LimiterStereomode = "Line Duck" // Line Duck (12) – Line Duck
	LspPlugInPluginsLv2LimiterStereomodeMixedHerm LspPlugInPluginsLv2LimiterStereomode = "Mixed Herm" // Mixed Herm (13) – Mixed Herm
	LspPlugInPluginsLv2LimiterStereomodeMixedExp LspPlugInPluginsLv2LimiterStereomode = "Mixed Exp" // Mixed Exp (14) – Mixed Exp
	LspPlugInPluginsLv2LimiterStereomodeMixedLine LspPlugInPluginsLv2LimiterStereomode = "Mixed Line" // Mixed Line (15) – Mixed Line
)

type LspPlugInPluginsLv2LimiterStereoovs string

const (
	LspPlugInPluginsLv2LimiterStereoovsNone LspPlugInPluginsLv2LimiterStereoovs = "None" // None (0) – None
	LspPlugInPluginsLv2LimiterStereoovsHalfX22L LspPlugInPluginsLv2LimiterStereoovs = "Half x2(2L)" // Half x2(2L) (1) – Half x2(2L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX23L LspPlugInPluginsLv2LimiterStereoovs = "Half x2(3L)" // Half x2(3L) (2) – Half x2(3L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX32L LspPlugInPluginsLv2LimiterStereoovs = "Half x3(2L)" // Half x3(2L) (3) – Half x3(2L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX33L LspPlugInPluginsLv2LimiterStereoovs = "Half x3(3L)" // Half x3(3L) (4) – Half x3(3L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX42L LspPlugInPluginsLv2LimiterStereoovs = "Half x4(2L)" // Half x4(2L) (5) – Half x4(2L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX43L LspPlugInPluginsLv2LimiterStereoovs = "Half x4(3L)" // Half x4(3L) (6) – Half x4(3L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX62L LspPlugInPluginsLv2LimiterStereoovs = "Half x6(2L)" // Half x6(2L) (7) – Half x6(2L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX63L LspPlugInPluginsLv2LimiterStereoovs = "Half x6(3L)" // Half x6(3L) (8) – Half x6(3L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX82L LspPlugInPluginsLv2LimiterStereoovs = "Half x8(2L)" // Half x8(2L) (9) – Half x8(2L)
	LspPlugInPluginsLv2LimiterStereoovsHalfX83L LspPlugInPluginsLv2LimiterStereoovs = "Half x8(3L)" // Half x8(3L) (10) – Half x8(3L)
	LspPlugInPluginsLv2LimiterStereoovsFullX22L LspPlugInPluginsLv2LimiterStereoovs = "Full x2(2L)" // Full x2(2L) (11) – Full x2(2L)
	LspPlugInPluginsLv2LimiterStereoovsFullX23L LspPlugInPluginsLv2LimiterStereoovs = "Full x2(3L)" // Full x2(3L) (12) – Full x2(3L)
	LspPlugInPluginsLv2LimiterStereoovsFullX32L LspPlugInPluginsLv2LimiterStereoovs = "Full x3(2L)" // Full x3(2L) (13) – Full x3(2L)
	LspPlugInPluginsLv2LimiterStereoovsFullX33L LspPlugInPluginsLv2LimiterStereoovs = "Full x3(3L)" // Full x3(3L) (14) – Full x3(3L)
	LspPlugInPluginsLv2LimiterStereoovsFullX42L LspPlugInPluginsLv2LimiterStereoovs = "Full x4(2L)" // Full x4(2L) (15) – Full x4(2L)
	LspPlugInPluginsLv2LimiterStereoovsFullX43L LspPlugInPluginsLv2LimiterStereoovs = "Full x4(3L)" // Full x4(3L) (16) – Full x4(3L)
	LspPlugInPluginsLv2LimiterStereoovsFullX62L LspPlugInPluginsLv2LimiterStereoovs = "Full x6(2L)" // Full x6(2L) (17) – Full x6(2L)
	LspPlugInPluginsLv2LimiterStereoovsFullX63L LspPlugInPluginsLv2LimiterStereoovs = "Full x6(3L)" // Full x6(3L) (18) – Full x6(3L)
	LspPlugInPluginsLv2LimiterStereoovsFullX82L LspPlugInPluginsLv2LimiterStereoovs = "Full x8(2L)" // Full x8(2L) (19) – Full x8(2L)
	LspPlugInPluginsLv2LimiterStereoovsFullX83L LspPlugInPluginsLv2LimiterStereoovs = "Full x8(3L)" // Full x8(3L) (20) – Full x8(3L)
)

type LspPlugInPluginsLv2LimiterStereodith string

const (
	LspPlugInPluginsLv2LimiterStereodithNone LspPlugInPluginsLv2LimiterStereodith = "None" // None (0) – None
	LspPlugInPluginsLv2LimiterStereodith7Bit LspPlugInPluginsLv2LimiterStereodith = "7bit" // 7bit (1) – 7bit
	LspPlugInPluginsLv2LimiterStereodith8Bit LspPlugInPluginsLv2LimiterStereodith = "8bit" // 8bit (2) – 8bit
	LspPlugInPluginsLv2LimiterStereodith11Bit LspPlugInPluginsLv2LimiterStereodith = "11bit" // 11bit (3) – 11bit
	LspPlugInPluginsLv2LimiterStereodith12Bit LspPlugInPluginsLv2LimiterStereodith = "12bit" // 12bit (4) – 12bit
	LspPlugInPluginsLv2LimiterStereodith15Bit LspPlugInPluginsLv2LimiterStereodith = "15bit" // 15bit (5) – 15bit
	LspPlugInPluginsLv2LimiterStereodith16Bit LspPlugInPluginsLv2LimiterStereodith = "16bit" // 16bit (6) – 16bit
	LspPlugInPluginsLv2LimiterStereodith23Bit LspPlugInPluginsLv2LimiterStereodith = "23bit" // 23bit (7) – 23bit
	LspPlugInPluginsLv2LimiterStereodith24Bit LspPlugInPluginsLv2LimiterStereodith = "24bit" // 24bit (8) – 24bit
)

