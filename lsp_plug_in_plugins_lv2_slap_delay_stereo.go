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

type LspPlugInPluginsLv2SlapDelayStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2SlapDelayStereo() (*LspPlugInPluginsLv2SlapDelayStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-slap-delay-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SlapDelayStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2SlapDelayStereoWithName(name string) (*LspPlugInPluginsLv2SlapDelayStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-slap-delay-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SlapDelayStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// dd0 (float32)
//
// Delay 0 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd0(value float32) error {
	return e.Element.SetProperty("dd0", value)
}

// dd1 (float32)
//
// Delay 1 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd1(value float32) error {
	return e.Element.SetProperty("dd1", value)
}

// dd10 (float32)
//
// Delay 10 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd10(value float32) error {
	return e.Element.SetProperty("dd10", value)
}

// dd11 (float32)
//
// Delay 11 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd11(value float32) error {
	return e.Element.SetProperty("dd11", value)
}

// dd12 (float32)
//
// Delay 12 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd12(value float32) error {
	return e.Element.SetProperty("dd12", value)
}

// dd13 (float32)
//
// Delay 13 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd13(value float32) error {
	return e.Element.SetProperty("dd13", value)
}

// dd14 (float32)
//
// Delay 14 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd14(value float32) error {
	return e.Element.SetProperty("dd14", value)
}

// dd15 (float32)
//
// Delay 15 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd15(value float32) error {
	return e.Element.SetProperty("dd15", value)
}

// dd2 (float32)
//
// Delay 2 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd2(value float32) error {
	return e.Element.SetProperty("dd2", value)
}

// dd3 (float32)
//
// Delay 3 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd3(value float32) error {
	return e.Element.SetProperty("dd3", value)
}

// dd4 (float32)
//
// Delay 4 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd4(value float32) error {
	return e.Element.SetProperty("dd4", value)
}

// dd5 (float32)
//
// Delay 5 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd5(value float32) error {
	return e.Element.SetProperty("dd5", value)
}

// dd6 (float32)
//
// Delay 6 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd6(value float32) error {
	return e.Element.SetProperty("dd6", value)
}

// dd7 (float32)
//
// Delay 7 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd7(value float32) error {
	return e.Element.SetProperty("dd7", value)
}

// dd8 (float32)
//
// Delay 8 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd8(value float32) error {
	return e.Element.SetProperty("dd8", value)
}

// dd9 (float32)
//
// Delay 9 distance

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDd9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dd9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDd9(value float32) error {
	return e.Element.SetProperty("dd9", value)
}

// df0 (float32)
//
// Delay 0 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf0(value float32) error {
	return e.Element.SetProperty("df0", value)
}

// df1 (float32)
//
// Delay 1 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf1(value float32) error {
	return e.Element.SetProperty("df1", value)
}

// df10 (float32)
//
// Delay 10 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf10(value float32) error {
	return e.Element.SetProperty("df10", value)
}

// df11 (float32)
//
// Delay 11 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf11(value float32) error {
	return e.Element.SetProperty("df11", value)
}

// df12 (float32)
//
// Delay 12 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf12(value float32) error {
	return e.Element.SetProperty("df12", value)
}

// df13 (float32)
//
// Delay 13 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf13(value float32) error {
	return e.Element.SetProperty("df13", value)
}

// df14 (float32)
//
// Delay 14 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf14(value float32) error {
	return e.Element.SetProperty("df14", value)
}

// df15 (float32)
//
// Delay 15 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf15(value float32) error {
	return e.Element.SetProperty("df15", value)
}

// df2 (float32)
//
// Delay 2 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf2(value float32) error {
	return e.Element.SetProperty("df2", value)
}

// df3 (float32)
//
// Delay 3 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf3(value float32) error {
	return e.Element.SetProperty("df3", value)
}

// df4 (float32)
//
// Delay 4 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf4(value float32) error {
	return e.Element.SetProperty("df4", value)
}

// df5 (float32)
//
// Delay 5 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf5(value float32) error {
	return e.Element.SetProperty("df5", value)
}

// df6 (float32)
//
// Delay 6 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf6(value float32) error {
	return e.Element.SetProperty("df6", value)
}

// df7 (float32)
//
// Delay 7 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf7(value float32) error {
	return e.Element.SetProperty("df7", value)
}

// df8 (float32)
//
// Delay 8 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf8(value float32) error {
	return e.Element.SetProperty("df8", value)
}

// df9 (float32)
//
// Delay 9 fraction

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDf9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("df9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDf9(value float32) error {
	return e.Element.SetProperty("df9", value)
}

// dg0 (float32)
//
// Delay 0 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg0(value float32) error {
	return e.Element.SetProperty("dg0", value)
}

// dg1 (float32)
//
// Delay 1 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg1(value float32) error {
	return e.Element.SetProperty("dg1", value)
}

// dg10 (float32)
//
// Delay 10 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg10(value float32) error {
	return e.Element.SetProperty("dg10", value)
}

// dg11 (float32)
//
// Delay 11 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg11(value float32) error {
	return e.Element.SetProperty("dg11", value)
}

// dg12 (float32)
//
// Delay 12 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg12(value float32) error {
	return e.Element.SetProperty("dg12", value)
}

// dg13 (float32)
//
// Delay 13 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg13(value float32) error {
	return e.Element.SetProperty("dg13", value)
}

// dg14 (float32)
//
// Delay 14 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg14(value float32) error {
	return e.Element.SetProperty("dg14", value)
}

// dg15 (float32)
//
// Delay 15 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg15(value float32) error {
	return e.Element.SetProperty("dg15", value)
}

// dg2 (float32)
//
// Delay 2 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg2(value float32) error {
	return e.Element.SetProperty("dg2", value)
}

// dg3 (float32)
//
// Delay 3 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg3(value float32) error {
	return e.Element.SetProperty("dg3", value)
}

// dg4 (float32)
//
// Delay 4 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg4(value float32) error {
	return e.Element.SetProperty("dg4", value)
}

// dg5 (float32)
//
// Delay 5 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg5(value float32) error {
	return e.Element.SetProperty("dg5", value)
}

// dg6 (float32)
//
// Delay 6 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg6(value float32) error {
	return e.Element.SetProperty("dg6", value)
}

// dg7 (float32)
//
// Delay 7 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg7(value float32) error {
	return e.Element.SetProperty("dg7", value)
}

// dg8 (float32)
//
// Delay 8 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg8(value float32) error {
	return e.Element.SetProperty("dg8", value)
}

// dg9 (float32)
//
// Delay 9 gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDg9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dg9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDg9(value float32) error {
	return e.Element.SetProperty("dg9", value)
}

// dm (bool)
//
// Dry mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dm")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm(value bool) error {
	return e.Element.SetProperty("dm", value)
}

// dm0 (GstLspPlugInPluginsLv2SlapDelayStereodm0)
//
// Delay 0 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm0() (interface{}, error) {
	return e.Element.GetProperty("dm0")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm0(value interface{}) error {
	return e.Element.SetProperty("dm0", value)
}

// dm1 (GstLspPlugInPluginsLv2SlapDelayStereodm1)
//
// Delay 1 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm1() (interface{}, error) {
	return e.Element.GetProperty("dm1")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm1(value interface{}) error {
	return e.Element.SetProperty("dm1", value)
}

// dm10 (GstLspPlugInPluginsLv2SlapDelayStereodm10)
//
// Delay 10 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm10() (interface{}, error) {
	return e.Element.GetProperty("dm10")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm10(value interface{}) error {
	return e.Element.SetProperty("dm10", value)
}

// dm11 (GstLspPlugInPluginsLv2SlapDelayStereodm11)
//
// Delay 11 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm11() (interface{}, error) {
	return e.Element.GetProperty("dm11")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm11(value interface{}) error {
	return e.Element.SetProperty("dm11", value)
}

// dm12 (GstLspPlugInPluginsLv2SlapDelayStereodm12)
//
// Delay 12 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm12() (interface{}, error) {
	return e.Element.GetProperty("dm12")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm12(value interface{}) error {
	return e.Element.SetProperty("dm12", value)
}

// dm13 (GstLspPlugInPluginsLv2SlapDelayStereodm13)
//
// Delay 13 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm13() (interface{}, error) {
	return e.Element.GetProperty("dm13")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm13(value interface{}) error {
	return e.Element.SetProperty("dm13", value)
}

// dm14 (GstLspPlugInPluginsLv2SlapDelayStereodm14)
//
// Delay 14 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm14() (interface{}, error) {
	return e.Element.GetProperty("dm14")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm14(value interface{}) error {
	return e.Element.SetProperty("dm14", value)
}

// dm15 (GstLspPlugInPluginsLv2SlapDelayStereodm15)
//
// Delay 15 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm15() (interface{}, error) {
	return e.Element.GetProperty("dm15")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm15(value interface{}) error {
	return e.Element.SetProperty("dm15", value)
}

// dm2 (GstLspPlugInPluginsLv2SlapDelayStereodm2)
//
// Delay 2 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm2() (interface{}, error) {
	return e.Element.GetProperty("dm2")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm2(value interface{}) error {
	return e.Element.SetProperty("dm2", value)
}

// dm3 (GstLspPlugInPluginsLv2SlapDelayStereodm3)
//
// Delay 3 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm3() (interface{}, error) {
	return e.Element.GetProperty("dm3")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm3(value interface{}) error {
	return e.Element.SetProperty("dm3", value)
}

// dm4 (GstLspPlugInPluginsLv2SlapDelayStereodm4)
//
// Delay 4 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm4() (interface{}, error) {
	return e.Element.GetProperty("dm4")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm4(value interface{}) error {
	return e.Element.SetProperty("dm4", value)
}

// dm5 (GstLspPlugInPluginsLv2SlapDelayStereodm5)
//
// Delay 5 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm5() (interface{}, error) {
	return e.Element.GetProperty("dm5")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm5(value interface{}) error {
	return e.Element.SetProperty("dm5", value)
}

// dm6 (GstLspPlugInPluginsLv2SlapDelayStereodm6)
//
// Delay 6 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm6() (interface{}, error) {
	return e.Element.GetProperty("dm6")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm6(value interface{}) error {
	return e.Element.SetProperty("dm6", value)
}

// dm7 (GstLspPlugInPluginsLv2SlapDelayStereodm7)
//
// Delay 7 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm7() (interface{}, error) {
	return e.Element.GetProperty("dm7")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm7(value interface{}) error {
	return e.Element.SetProperty("dm7", value)
}

// dm8 (GstLspPlugInPluginsLv2SlapDelayStereodm8)
//
// Delay 8 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm8() (interface{}, error) {
	return e.Element.GetProperty("dm8")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm8(value interface{}) error {
	return e.Element.SetProperty("dm8", value)
}

// dm9 (GstLspPlugInPluginsLv2SlapDelayStereodm9)
//
// Delay 9 mode

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDm9() (interface{}, error) {
	return e.Element.GetProperty("dm9")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDm9(value interface{}) error {
	return e.Element.SetProperty("dm9", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDry() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dry")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// ds0 (int)
//
// Delay 0 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs0() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds0")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs0(value int) error {
	return e.Element.SetProperty("ds0", value)
}

// ds1 (int)
//
// Delay 1 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs1() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds1")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs1(value int) error {
	return e.Element.SetProperty("ds1", value)
}

// ds10 (int)
//
// Delay 10 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs10() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds10")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs10(value int) error {
	return e.Element.SetProperty("ds10", value)
}

// ds11 (int)
//
// Delay 11 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs11() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds11")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs11(value int) error {
	return e.Element.SetProperty("ds11", value)
}

// ds12 (int)
//
// Delay 12 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs12() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds12")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs12(value int) error {
	return e.Element.SetProperty("ds12", value)
}

// ds13 (int)
//
// Delay 13 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs13() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds13")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs13(value int) error {
	return e.Element.SetProperty("ds13", value)
}

// ds14 (int)
//
// Delay 14 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs14() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds14")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs14(value int) error {
	return e.Element.SetProperty("ds14", value)
}

// ds15 (int)
//
// Delay 15 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs15() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds15")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs15(value int) error {
	return e.Element.SetProperty("ds15", value)
}

// ds2 (int)
//
// Delay 2 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs2() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds2")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs2(value int) error {
	return e.Element.SetProperty("ds2", value)
}

// ds3 (int)
//
// Delay 3 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs3() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds3")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs3(value int) error {
	return e.Element.SetProperty("ds3", value)
}

// ds4 (int)
//
// Delay 4 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs4() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds4")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs4(value int) error {
	return e.Element.SetProperty("ds4", value)
}

// ds5 (int)
//
// Delay 5 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs5() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds5")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs5(value int) error {
	return e.Element.SetProperty("ds5", value)
}

// ds6 (int)
//
// Delay 6 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs6() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds6")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs6(value int) error {
	return e.Element.SetProperty("ds6", value)
}

// ds7 (int)
//
// Delay 7 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs7() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds7")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs7(value int) error {
	return e.Element.SetProperty("ds7", value)
}

// ds8 (int)
//
// Delay 8 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs8() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds8")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs8(value int) error {
	return e.Element.SetProperty("ds8", value)
}

// ds9 (int)
//
// Delay 9 denominator

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDs9() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ds9")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDs9(value int) error {
	return e.Element.SetProperty("ds9", value)
}

// dt0 (float32)
//
// Delay 0 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt0(value float32) error {
	return e.Element.SetProperty("dt0", value)
}

// dt1 (float32)
//
// Delay 1 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt1(value float32) error {
	return e.Element.SetProperty("dt1", value)
}

// dt10 (float32)
//
// Delay 10 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt10(value float32) error {
	return e.Element.SetProperty("dt10", value)
}

// dt11 (float32)
//
// Delay 11 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt11(value float32) error {
	return e.Element.SetProperty("dt11", value)
}

// dt12 (float32)
//
// Delay 12 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt12(value float32) error {
	return e.Element.SetProperty("dt12", value)
}

// dt13 (float32)
//
// Delay 13 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt13(value float32) error {
	return e.Element.SetProperty("dt13", value)
}

// dt14 (float32)
//
// Delay 14 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt14(value float32) error {
	return e.Element.SetProperty("dt14", value)
}

// dt15 (float32)
//
// Delay 15 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt15(value float32) error {
	return e.Element.SetProperty("dt15", value)
}

// dt2 (float32)
//
// Delay 2 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt2(value float32) error {
	return e.Element.SetProperty("dt2", value)
}

// dt3 (float32)
//
// Delay 3 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt3(value float32) error {
	return e.Element.SetProperty("dt3", value)
}

// dt4 (float32)
//
// Delay 4 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt4(value float32) error {
	return e.Element.SetProperty("dt4", value)
}

// dt5 (float32)
//
// Delay 5 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt5(value float32) error {
	return e.Element.SetProperty("dt5", value)
}

// dt6 (float32)
//
// Delay 6 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt6(value float32) error {
	return e.Element.SetProperty("dt6", value)
}

// dt7 (float32)
//
// Delay 7 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt7(value float32) error {
	return e.Element.SetProperty("dt7", value)
}

// dt8 (float32)
//
// Delay 8 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt8(value float32) error {
	return e.Element.SetProperty("dt8", value)
}

// dt9 (float32)
//
// Delay 9 time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetDt9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("dt9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetDt9(value float32) error {
	return e.Element.SetProperty("dt9", value)
}

// eq0 (bool)
//
// Equalizer 0 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq0(value bool) error {
	return e.Element.SetProperty("eq0", value)
}

// eq1 (bool)
//
// Equalizer 1 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq1(value bool) error {
	return e.Element.SetProperty("eq1", value)
}

// eq10 (bool)
//
// Equalizer 10 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq10(value bool) error {
	return e.Element.SetProperty("eq10", value)
}

// eq11 (bool)
//
// Equalizer 11 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq11(value bool) error {
	return e.Element.SetProperty("eq11", value)
}

// eq12 (bool)
//
// Equalizer 12 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq12(value bool) error {
	return e.Element.SetProperty("eq12", value)
}

// eq13 (bool)
//
// Equalizer 13 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq13(value bool) error {
	return e.Element.SetProperty("eq13", value)
}

// eq14 (bool)
//
// Equalizer 14 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq14(value bool) error {
	return e.Element.SetProperty("eq14", value)
}

// eq15 (bool)
//
// Equalizer 15 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq15(value bool) error {
	return e.Element.SetProperty("eq15", value)
}

// eq2 (bool)
//
// Equalizer 2 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq2(value bool) error {
	return e.Element.SetProperty("eq2", value)
}

// eq3 (bool)
//
// Equalizer 3 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq3(value bool) error {
	return e.Element.SetProperty("eq3", value)
}

// eq4 (bool)
//
// Equalizer 4 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq4(value bool) error {
	return e.Element.SetProperty("eq4", value)
}

// eq5 (bool)
//
// Equalizer 5 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq5(value bool) error {
	return e.Element.SetProperty("eq5", value)
}

// eq6 (bool)
//
// Equalizer 6 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq6(value bool) error {
	return e.Element.SetProperty("eq6", value)
}

// eq7 (bool)
//
// Equalizer 7 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq7(value bool) error {
	return e.Element.SetProperty("eq7", value)
}

// eq8 (bool)
//
// Equalizer 8 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq8(value bool) error {
	return e.Element.SetProperty("eq8", value)
}

// eq9 (bool)
//
// Equalizer 9 on

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetEq9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("eq9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetEq9(value bool) error {
	return e.Element.SetProperty("eq9", value)
}

// fbb0 (float32)
//
// Delay 0 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb0(value float32) error {
	return e.Element.SetProperty("fbb0", value)
}

// fbb1 (float32)
//
// Delay 1 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb1(value float32) error {
	return e.Element.SetProperty("fbb1", value)
}

// fbb10 (float32)
//
// Delay 10 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb10(value float32) error {
	return e.Element.SetProperty("fbb10", value)
}

// fbb11 (float32)
//
// Delay 11 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb11(value float32) error {
	return e.Element.SetProperty("fbb11", value)
}

// fbb12 (float32)
//
// Delay 12 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb12(value float32) error {
	return e.Element.SetProperty("fbb12", value)
}

// fbb13 (float32)
//
// Delay 13 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb13(value float32) error {
	return e.Element.SetProperty("fbb13", value)
}

// fbb14 (float32)
//
// Delay 14 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb14(value float32) error {
	return e.Element.SetProperty("fbb14", value)
}

// fbb15 (float32)
//
// Delay 15 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb15(value float32) error {
	return e.Element.SetProperty("fbb15", value)
}

// fbb2 (float32)
//
// Delay 2 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb2(value float32) error {
	return e.Element.SetProperty("fbb2", value)
}

// fbb3 (float32)
//
// Delay 3 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb3(value float32) error {
	return e.Element.SetProperty("fbb3", value)
}

// fbb4 (float32)
//
// Delay 4 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb4(value float32) error {
	return e.Element.SetProperty("fbb4", value)
}

// fbb5 (float32)
//
// Delay 5 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb5(value float32) error {
	return e.Element.SetProperty("fbb5", value)
}

// fbb6 (float32)
//
// Delay 6 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb6(value float32) error {
	return e.Element.SetProperty("fbb6", value)
}

// fbb7 (float32)
//
// Delay 7 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb7(value float32) error {
	return e.Element.SetProperty("fbb7", value)
}

// fbb8 (float32)
//
// Delay 8 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb8(value float32) error {
	return e.Element.SetProperty("fbb8", value)
}

// fbb9 (float32)
//
// Delay 9 bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbb9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbb9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbb9(value float32) error {
	return e.Element.SetProperty("fbb9", value)
}

// fbm0 (float32)
//
// Delay 0 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm0(value float32) error {
	return e.Element.SetProperty("fbm0", value)
}

// fbm1 (float32)
//
// Delay 1 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm1(value float32) error {
	return e.Element.SetProperty("fbm1", value)
}

// fbm10 (float32)
//
// Delay 10 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm10(value float32) error {
	return e.Element.SetProperty("fbm10", value)
}

// fbm11 (float32)
//
// Delay 11 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm11(value float32) error {
	return e.Element.SetProperty("fbm11", value)
}

// fbm12 (float32)
//
// Delay 12 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm12(value float32) error {
	return e.Element.SetProperty("fbm12", value)
}

// fbm13 (float32)
//
// Delay 13 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm13(value float32) error {
	return e.Element.SetProperty("fbm13", value)
}

// fbm14 (float32)
//
// Delay 14 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm14(value float32) error {
	return e.Element.SetProperty("fbm14", value)
}

// fbm15 (float32)
//
// Delay 15 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm15(value float32) error {
	return e.Element.SetProperty("fbm15", value)
}

// fbm2 (float32)
//
// Delay 2 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm2(value float32) error {
	return e.Element.SetProperty("fbm2", value)
}

// fbm3 (float32)
//
// Delay 3 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm3(value float32) error {
	return e.Element.SetProperty("fbm3", value)
}

// fbm4 (float32)
//
// Delay 4 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm4(value float32) error {
	return e.Element.SetProperty("fbm4", value)
}

// fbm5 (float32)
//
// Delay 5 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm5(value float32) error {
	return e.Element.SetProperty("fbm5", value)
}

// fbm6 (float32)
//
// Delay 6 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm6(value float32) error {
	return e.Element.SetProperty("fbm6", value)
}

// fbm7 (float32)
//
// Delay 7 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm7(value float32) error {
	return e.Element.SetProperty("fbm7", value)
}

// fbm8 (float32)
//
// Delay 8 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm8(value float32) error {
	return e.Element.SetProperty("fbm8", value)
}

// fbm9 (float32)
//
// Delay 9 middle

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbm9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbm9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbm9(value float32) error {
	return e.Element.SetProperty("fbm9", value)
}

// fbp0 (float32)
//
// Delay 0 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp0(value float32) error {
	return e.Element.SetProperty("fbp0", value)
}

// fbp1 (float32)
//
// Delay 1 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp1(value float32) error {
	return e.Element.SetProperty("fbp1", value)
}

// fbp10 (float32)
//
// Delay 10 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp10(value float32) error {
	return e.Element.SetProperty("fbp10", value)
}

// fbp11 (float32)
//
// Delay 11 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp11(value float32) error {
	return e.Element.SetProperty("fbp11", value)
}

// fbp12 (float32)
//
// Delay 12 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp12(value float32) error {
	return e.Element.SetProperty("fbp12", value)
}

// fbp13 (float32)
//
// Delay 13 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp13(value float32) error {
	return e.Element.SetProperty("fbp13", value)
}

// fbp14 (float32)
//
// Delay 14 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp14(value float32) error {
	return e.Element.SetProperty("fbp14", value)
}

// fbp15 (float32)
//
// Delay 15 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp15(value float32) error {
	return e.Element.SetProperty("fbp15", value)
}

// fbp2 (float32)
//
// Delay 2 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp2(value float32) error {
	return e.Element.SetProperty("fbp2", value)
}

// fbp3 (float32)
//
// Delay 3 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp3(value float32) error {
	return e.Element.SetProperty("fbp3", value)
}

// fbp4 (float32)
//
// Delay 4 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp4(value float32) error {
	return e.Element.SetProperty("fbp4", value)
}

// fbp5 (float32)
//
// Delay 5 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp5(value float32) error {
	return e.Element.SetProperty("fbp5", value)
}

// fbp6 (float32)
//
// Delay 6 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp6(value float32) error {
	return e.Element.SetProperty("fbp6", value)
}

// fbp7 (float32)
//
// Delay 7 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp7(value float32) error {
	return e.Element.SetProperty("fbp7", value)
}

// fbp8 (float32)
//
// Delay 8 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp8(value float32) error {
	return e.Element.SetProperty("fbp8", value)
}

// fbp9 (float32)
//
// Delay 9 presence

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbp9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbp9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbp9(value float32) error {
	return e.Element.SetProperty("fbp9", value)
}

// fbs0 (float32)
//
// Delay 0 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs0(value float32) error {
	return e.Element.SetProperty("fbs0", value)
}

// fbs1 (float32)
//
// Delay 1 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs1(value float32) error {
	return e.Element.SetProperty("fbs1", value)
}

// fbs10 (float32)
//
// Delay 10 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs10(value float32) error {
	return e.Element.SetProperty("fbs10", value)
}

// fbs11 (float32)
//
// Delay 11 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs11(value float32) error {
	return e.Element.SetProperty("fbs11", value)
}

// fbs12 (float32)
//
// Delay 12 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs12(value float32) error {
	return e.Element.SetProperty("fbs12", value)
}

// fbs13 (float32)
//
// Delay 13 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs13(value float32) error {
	return e.Element.SetProperty("fbs13", value)
}

// fbs14 (float32)
//
// Delay 14 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs14(value float32) error {
	return e.Element.SetProperty("fbs14", value)
}

// fbs15 (float32)
//
// Delay 15 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs15(value float32) error {
	return e.Element.SetProperty("fbs15", value)
}

// fbs2 (float32)
//
// Delay 2 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs2(value float32) error {
	return e.Element.SetProperty("fbs2", value)
}

// fbs3 (float32)
//
// Delay 3 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs3(value float32) error {
	return e.Element.SetProperty("fbs3", value)
}

// fbs4 (float32)
//
// Delay 4 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs4(value float32) error {
	return e.Element.SetProperty("fbs4", value)
}

// fbs5 (float32)
//
// Delay 5 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs5(value float32) error {
	return e.Element.SetProperty("fbs5", value)
}

// fbs6 (float32)
//
// Delay 6 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs6(value float32) error {
	return e.Element.SetProperty("fbs6", value)
}

// fbs7 (float32)
//
// Delay 7 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs7(value float32) error {
	return e.Element.SetProperty("fbs7", value)
}

// fbs8 (float32)
//
// Delay 8 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs8(value float32) error {
	return e.Element.SetProperty("fbs8", value)
}

// fbs9 (float32)
//
// Delay 9 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbs9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbs9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbs9(value float32) error {
	return e.Element.SetProperty("fbs9", value)
}

// fbt0 (float32)
//
// Delay 0 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt0(value float32) error {
	return e.Element.SetProperty("fbt0", value)
}

// fbt1 (float32)
//
// Delay 1 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt1(value float32) error {
	return e.Element.SetProperty("fbt1", value)
}

// fbt10 (float32)
//
// Delay 10 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt10(value float32) error {
	return e.Element.SetProperty("fbt10", value)
}

// fbt11 (float32)
//
// Delay 11 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt11(value float32) error {
	return e.Element.SetProperty("fbt11", value)
}

// fbt12 (float32)
//
// Delay 12 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt12(value float32) error {
	return e.Element.SetProperty("fbt12", value)
}

// fbt13 (float32)
//
// Delay 13 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt13(value float32) error {
	return e.Element.SetProperty("fbt13", value)
}

// fbt14 (float32)
//
// Delay 14 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt14(value float32) error {
	return e.Element.SetProperty("fbt14", value)
}

// fbt15 (float32)
//
// Delay 15 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt15(value float32) error {
	return e.Element.SetProperty("fbt15", value)
}

// fbt2 (float32)
//
// Delay 2 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt2(value float32) error {
	return e.Element.SetProperty("fbt2", value)
}

// fbt3 (float32)
//
// Delay 3 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt3(value float32) error {
	return e.Element.SetProperty("fbt3", value)
}

// fbt4 (float32)
//
// Delay 4 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt4(value float32) error {
	return e.Element.SetProperty("fbt4", value)
}

// fbt5 (float32)
//
// Delay 5 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt5(value float32) error {
	return e.Element.SetProperty("fbt5", value)
}

// fbt6 (float32)
//
// Delay 6 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt6(value float32) error {
	return e.Element.SetProperty("fbt6", value)
}

// fbt7 (float32)
//
// Delay 7 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt7(value float32) error {
	return e.Element.SetProperty("fbt7", value)
}

// fbt8 (float32)
//
// Delay 8 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt8(value float32) error {
	return e.Element.SetProperty("fbt8", value)
}

// fbt9 (float32)
//
// Delay 9 treble

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFbt9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fbt9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFbt9(value float32) error {
	return e.Element.SetProperty("fbt9", value)
}

// fhc0 (float32)
//
// Delay 0 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc0(value float32) error {
	return e.Element.SetProperty("fhc0", value)
}

// fhc1 (float32)
//
// Delay 1 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc1(value float32) error {
	return e.Element.SetProperty("fhc1", value)
}

// fhc10 (float32)
//
// Delay 10 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc10(value float32) error {
	return e.Element.SetProperty("fhc10", value)
}

// fhc11 (float32)
//
// Delay 11 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc11(value float32) error {
	return e.Element.SetProperty("fhc11", value)
}

// fhc12 (float32)
//
// Delay 12 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc12(value float32) error {
	return e.Element.SetProperty("fhc12", value)
}

// fhc13 (float32)
//
// Delay 13 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc13(value float32) error {
	return e.Element.SetProperty("fhc13", value)
}

// fhc14 (float32)
//
// Delay 14 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc14(value float32) error {
	return e.Element.SetProperty("fhc14", value)
}

// fhc15 (float32)
//
// Delay 15 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc15(value float32) error {
	return e.Element.SetProperty("fhc15", value)
}

// fhc2 (float32)
//
// Delay 2 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc2(value float32) error {
	return e.Element.SetProperty("fhc2", value)
}

// fhc3 (float32)
//
// Delay 3 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc3(value float32) error {
	return e.Element.SetProperty("fhc3", value)
}

// fhc4 (float32)
//
// Delay 4 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc4(value float32) error {
	return e.Element.SetProperty("fhc4", value)
}

// fhc5 (float32)
//
// Delay 5 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc5(value float32) error {
	return e.Element.SetProperty("fhc5", value)
}

// fhc6 (float32)
//
// Delay 6 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc6(value float32) error {
	return e.Element.SetProperty("fhc6", value)
}

// fhc7 (float32)
//
// Delay 7 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc7(value float32) error {
	return e.Element.SetProperty("fhc7", value)
}

// fhc8 (float32)
//
// Delay 8 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc8(value float32) error {
	return e.Element.SetProperty("fhc8", value)
}

// fhc9 (float32)
//
// Delay 9 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFhc9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fhc9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFhc9(value float32) error {
	return e.Element.SetProperty("fhc9", value)
}

// flc0 (float32)
//
// Delay 0 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc0(value float32) error {
	return e.Element.SetProperty("flc0", value)
}

// flc1 (float32)
//
// Delay 1 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc1(value float32) error {
	return e.Element.SetProperty("flc1", value)
}

// flc10 (float32)
//
// Delay 10 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc10(value float32) error {
	return e.Element.SetProperty("flc10", value)
}

// flc11 (float32)
//
// Delay 11 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc11(value float32) error {
	return e.Element.SetProperty("flc11", value)
}

// flc12 (float32)
//
// Delay 12 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc12(value float32) error {
	return e.Element.SetProperty("flc12", value)
}

// flc13 (float32)
//
// Delay 13 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc13(value float32) error {
	return e.Element.SetProperty("flc13", value)
}

// flc14 (float32)
//
// Delay 14 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc14(value float32) error {
	return e.Element.SetProperty("flc14", value)
}

// flc15 (float32)
//
// Delay 15 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc15(value float32) error {
	return e.Element.SetProperty("flc15", value)
}

// flc2 (float32)
//
// Delay 2 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc2(value float32) error {
	return e.Element.SetProperty("flc2", value)
}

// flc3 (float32)
//
// Delay 3 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc3(value float32) error {
	return e.Element.SetProperty("flc3", value)
}

// flc4 (float32)
//
// Delay 4 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc4(value float32) error {
	return e.Element.SetProperty("flc4", value)
}

// flc5 (float32)
//
// Delay 5 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc5(value float32) error {
	return e.Element.SetProperty("flc5", value)
}

// flc6 (float32)
//
// Delay 6 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc6(value float32) error {
	return e.Element.SetProperty("flc6", value)
}

// flc7 (float32)
//
// Delay 7 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc7(value float32) error {
	return e.Element.SetProperty("flc7", value)
}

// flc8 (float32)
//
// Delay 8 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc8(value float32) error {
	return e.Element.SetProperty("flc8", value)
}

// flc9 (float32)
//
// Delay 9 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetFlc9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("flc9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetFlc9(value float32) error {
	return e.Element.SetProperty("flc9", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hfc0 (bool)
//
// Delay 0 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc0(value bool) error {
	return e.Element.SetProperty("hfc0", value)
}

// hfc1 (bool)
//
// Delay 1 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc1(value bool) error {
	return e.Element.SetProperty("hfc1", value)
}

// hfc10 (bool)
//
// Delay 10 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc10(value bool) error {
	return e.Element.SetProperty("hfc10", value)
}

// hfc11 (bool)
//
// Delay 11 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc11(value bool) error {
	return e.Element.SetProperty("hfc11", value)
}

// hfc12 (bool)
//
// Delay 12 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc12(value bool) error {
	return e.Element.SetProperty("hfc12", value)
}

// hfc13 (bool)
//
// Delay 13 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc13(value bool) error {
	return e.Element.SetProperty("hfc13", value)
}

// hfc14 (bool)
//
// Delay 14 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc14(value bool) error {
	return e.Element.SetProperty("hfc14", value)
}

// hfc15 (bool)
//
// Delay 15 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc15(value bool) error {
	return e.Element.SetProperty("hfc15", value)
}

// hfc2 (bool)
//
// Delay 2 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc2(value bool) error {
	return e.Element.SetProperty("hfc2", value)
}

// hfc3 (bool)
//
// Delay 3 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc3(value bool) error {
	return e.Element.SetProperty("hfc3", value)
}

// hfc4 (bool)
//
// Delay 4 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc4(value bool) error {
	return e.Element.SetProperty("hfc4", value)
}

// hfc5 (bool)
//
// Delay 5 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc5(value bool) error {
	return e.Element.SetProperty("hfc5", value)
}

// hfc6 (bool)
//
// Delay 6 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc6(value bool) error {
	return e.Element.SetProperty("hfc6", value)
}

// hfc7 (bool)
//
// Delay 7 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc7(value bool) error {
	return e.Element.SetProperty("hfc7", value)
}

// hfc8 (bool)
//
// Delay 8 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc8(value bool) error {
	return e.Element.SetProperty("hfc8", value)
}

// hfc9 (bool)
//
// Delay 9 high-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetHfc9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hfc9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetHfc9(value bool) error {
	return e.Element.SetProperty("hfc9", value)
}

// lfc0 (bool)
//
// Delay 0 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc0(value bool) error {
	return e.Element.SetProperty("lfc0", value)
}

// lfc1 (bool)
//
// Delay 1 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc1(value bool) error {
	return e.Element.SetProperty("lfc1", value)
}

// lfc10 (bool)
//
// Delay 10 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc10(value bool) error {
	return e.Element.SetProperty("lfc10", value)
}

// lfc11 (bool)
//
// Delay 11 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc11(value bool) error {
	return e.Element.SetProperty("lfc11", value)
}

// lfc12 (bool)
//
// Delay 12 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc12(value bool) error {
	return e.Element.SetProperty("lfc12", value)
}

// lfc13 (bool)
//
// Delay 13 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc13(value bool) error {
	return e.Element.SetProperty("lfc13", value)
}

// lfc14 (bool)
//
// Delay 14 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc14(value bool) error {
	return e.Element.SetProperty("lfc14", value)
}

// lfc15 (bool)
//
// Delay 15 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc15(value bool) error {
	return e.Element.SetProperty("lfc15", value)
}

// lfc2 (bool)
//
// Delay 2 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc2(value bool) error {
	return e.Element.SetProperty("lfc2", value)
}

// lfc3 (bool)
//
// Delay 3 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc3(value bool) error {
	return e.Element.SetProperty("lfc3", value)
}

// lfc4 (bool)
//
// Delay 4 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc4(value bool) error {
	return e.Element.SetProperty("lfc4", value)
}

// lfc5 (bool)
//
// Delay 5 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc5(value bool) error {
	return e.Element.SetProperty("lfc5", value)
}

// lfc6 (bool)
//
// Delay 6 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc6(value bool) error {
	return e.Element.SetProperty("lfc6", value)
}

// lfc7 (bool)
//
// Delay 7 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc7(value bool) error {
	return e.Element.SetProperty("lfc7", value)
}

// lfc8 (bool)
//
// Delay 8 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc8(value bool) error {
	return e.Element.SetProperty("lfc8", value)
}

// lfc9 (bool)
//
// Delay 9 low-cut

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLfc9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfc9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLfc9(value bool) error {
	return e.Element.SetProperty("lfc9", value)
}

// lsel (GstLspPlugInPluginsLv2SlapDelayStereolsel)
//
// Delay line selector

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetLsel() (interface{}, error) {
	return e.Element.GetProperty("lsel")
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetLsel(value interface{}) error {
	return e.Element.SetProperty("lsel", value)
}

// m0 (bool)
//
// Delay 0 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM0(value bool) error {
	return e.Element.SetProperty("m0", value)
}

// m1 (bool)
//
// Delay 1 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM1(value bool) error {
	return e.Element.SetProperty("m1", value)
}

// m10 (bool)
//
// Delay 10 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM10(value bool) error {
	return e.Element.SetProperty("m10", value)
}

// m11 (bool)
//
// Delay 11 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM11(value bool) error {
	return e.Element.SetProperty("m11", value)
}

// m12 (bool)
//
// Delay 12 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM12(value bool) error {
	return e.Element.SetProperty("m12", value)
}

// m13 (bool)
//
// Delay 13 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM13(value bool) error {
	return e.Element.SetProperty("m13", value)
}

// m14 (bool)
//
// Delay 14 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM14(value bool) error {
	return e.Element.SetProperty("m14", value)
}

// m15 (bool)
//
// Delay 15 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM15(value bool) error {
	return e.Element.SetProperty("m15", value)
}

// m2 (bool)
//
// Delay 2 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM2(value bool) error {
	return e.Element.SetProperty("m2", value)
}

// m3 (bool)
//
// Delay 3 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM3(value bool) error {
	return e.Element.SetProperty("m3", value)
}

// m4 (bool)
//
// Delay 4 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM4(value bool) error {
	return e.Element.SetProperty("m4", value)
}

// m5 (bool)
//
// Delay 5 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM5(value bool) error {
	return e.Element.SetProperty("m5", value)
}

// m6 (bool)
//
// Delay 6 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM6(value bool) error {
	return e.Element.SetProperty("m6", value)
}

// m7 (bool)
//
// Delay 7 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM7(value bool) error {
	return e.Element.SetProperty("m7", value)
}

// m8 (bool)
//
// Delay 8 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM8(value bool) error {
	return e.Element.SetProperty("m8", value)
}

// m9 (bool)
//
// Delay 9 mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetM9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetM9(value bool) error {
	return e.Element.SetProperty("m9", value)
}

// mono (bool)
//
// Mono output

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetMono() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mono")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetMono(value bool) error {
	return e.Element.SetProperty("mono", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetOutLatency() (int, error) {
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

// ph0 (bool)
//
// Delay 0 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh0(value bool) error {
	return e.Element.SetProperty("ph0", value)
}

// ph1 (bool)
//
// Delay 1 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh1(value bool) error {
	return e.Element.SetProperty("ph1", value)
}

// ph10 (bool)
//
// Delay 10 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh10(value bool) error {
	return e.Element.SetProperty("ph10", value)
}

// ph11 (bool)
//
// Delay 11 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh11(value bool) error {
	return e.Element.SetProperty("ph11", value)
}

// ph12 (bool)
//
// Delay 12 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh12(value bool) error {
	return e.Element.SetProperty("ph12", value)
}

// ph13 (bool)
//
// Delay 13 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh13(value bool) error {
	return e.Element.SetProperty("ph13", value)
}

// ph14 (bool)
//
// Delay 14 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh14(value bool) error {
	return e.Element.SetProperty("ph14", value)
}

// ph15 (bool)
//
// Delay 15 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh15(value bool) error {
	return e.Element.SetProperty("ph15", value)
}

// ph2 (bool)
//
// Delay 2 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh2(value bool) error {
	return e.Element.SetProperty("ph2", value)
}

// ph3 (bool)
//
// Delay 3 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh3(value bool) error {
	return e.Element.SetProperty("ph3", value)
}

// ph4 (bool)
//
// Delay 4 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh4(value bool) error {
	return e.Element.SetProperty("ph4", value)
}

// ph5 (bool)
//
// Delay 5 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh5(value bool) error {
	return e.Element.SetProperty("ph5", value)
}

// ph6 (bool)
//
// Delay 6 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh6(value bool) error {
	return e.Element.SetProperty("ph6", value)
}

// ph7 (bool)
//
// Delay 7 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh7(value bool) error {
	return e.Element.SetProperty("ph7", value)
}

// ph8 (bool)
//
// Delay 8 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh8(value bool) error {
	return e.Element.SetProperty("ph8", value)
}

// ph9 (bool)
//
// Delay 9 phase

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPh9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ph9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPh9(value bool) error {
	return e.Element.SetProperty("ph9", value)
}

// pl-in (float32)
//
// Input left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPlIn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl-in")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPlIn(value float32) error {
	return e.Element.SetProperty("pl-in", value)
}

// pl0 (float32)
//
// Delay 0 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl0(value float32) error {
	return e.Element.SetProperty("pl0", value)
}

// pl1 (float32)
//
// Delay 1 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl1(value float32) error {
	return e.Element.SetProperty("pl1", value)
}

// pl10 (float32)
//
// Delay 10 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl10(value float32) error {
	return e.Element.SetProperty("pl10", value)
}

// pl11 (float32)
//
// Delay 11 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl11(value float32) error {
	return e.Element.SetProperty("pl11", value)
}

// pl12 (float32)
//
// Delay 12 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl12(value float32) error {
	return e.Element.SetProperty("pl12", value)
}

// pl13 (float32)
//
// Delay 13 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl13(value float32) error {
	return e.Element.SetProperty("pl13", value)
}

// pl14 (float32)
//
// Delay 14 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl14(value float32) error {
	return e.Element.SetProperty("pl14", value)
}

// pl15 (float32)
//
// Delay 15 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl15(value float32) error {
	return e.Element.SetProperty("pl15", value)
}

// pl2 (float32)
//
// Delay 2 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl2(value float32) error {
	return e.Element.SetProperty("pl2", value)
}

// pl3 (float32)
//
// Delay 3 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl3(value float32) error {
	return e.Element.SetProperty("pl3", value)
}

// pl4 (float32)
//
// Delay 4 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl4(value float32) error {
	return e.Element.SetProperty("pl4", value)
}

// pl5 (float32)
//
// Delay 5 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl5(value float32) error {
	return e.Element.SetProperty("pl5", value)
}

// pl6 (float32)
//
// Delay 6 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl6(value float32) error {
	return e.Element.SetProperty("pl6", value)
}

// pl7 (float32)
//
// Delay 7 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl7(value float32) error {
	return e.Element.SetProperty("pl7", value)
}

// pl8 (float32)
//
// Delay 8 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl8(value float32) error {
	return e.Element.SetProperty("pl8", value)
}

// pl9 (float32)
//
// Delay 9 left channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPl9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pl9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPl9(value float32) error {
	return e.Element.SetProperty("pl9", value)
}

// pr-in (float32)
//
// Input right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPrIn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr-in")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPrIn(value float32) error {
	return e.Element.SetProperty("pr-in", value)
}

// pr0 (float32)
//
// Delay 0 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr0(value float32) error {
	return e.Element.SetProperty("pr0", value)
}

// pr1 (float32)
//
// Delay 1 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr1(value float32) error {
	return e.Element.SetProperty("pr1", value)
}

// pr10 (float32)
//
// Delay 10 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr10(value float32) error {
	return e.Element.SetProperty("pr10", value)
}

// pr11 (float32)
//
// Delay 11 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr11(value float32) error {
	return e.Element.SetProperty("pr11", value)
}

// pr12 (float32)
//
// Delay 12 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr12(value float32) error {
	return e.Element.SetProperty("pr12", value)
}

// pr13 (float32)
//
// Delay 13 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr13(value float32) error {
	return e.Element.SetProperty("pr13", value)
}

// pr14 (float32)
//
// Delay 14 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr14(value float32) error {
	return e.Element.SetProperty("pr14", value)
}

// pr15 (float32)
//
// Delay 15 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr15(value float32) error {
	return e.Element.SetProperty("pr15", value)
}

// pr2 (float32)
//
// Delay 2 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr2(value float32) error {
	return e.Element.SetProperty("pr2", value)
}

// pr3 (float32)
//
// Delay 3 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr3(value float32) error {
	return e.Element.SetProperty("pr3", value)
}

// pr4 (float32)
//
// Delay 4 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr4(value float32) error {
	return e.Element.SetProperty("pr4", value)
}

// pr5 (float32)
//
// Delay 5 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr5(value float32) error {
	return e.Element.SetProperty("pr5", value)
}

// pr6 (float32)
//
// Delay 6 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr6(value float32) error {
	return e.Element.SetProperty("pr6", value)
}

// pr7 (float32)
//
// Delay 7 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr7(value float32) error {
	return e.Element.SetProperty("pr7", value)
}

// pr8 (float32)
//
// Delay 8 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr8(value float32) error {
	return e.Element.SetProperty("pr8", value)
}

// pr9 (float32)
//
// Delay 9 right channel panorama

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPr9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pr9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPr9(value float32) error {
	return e.Element.SetProperty("pr9", value)
}

// pred (float32)
//
// Pre-delay

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetPred() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("pred")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetPred(value float32) error {
	return e.Element.SetProperty("pred", value)
}

// ramp (bool)
//
// Ramping delay

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetRamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ramp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetRamp(value bool) error {
	return e.Element.SetProperty("ramp", value)
}

// s0 (bool)
//
// Delay 0 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS0(value bool) error {
	return e.Element.SetProperty("s0", value)
}

// s1 (bool)
//
// Delay 1 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS1(value bool) error {
	return e.Element.SetProperty("s1", value)
}

// s10 (bool)
//
// Delay 10 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS10(value bool) error {
	return e.Element.SetProperty("s10", value)
}

// s11 (bool)
//
// Delay 11 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS11(value bool) error {
	return e.Element.SetProperty("s11", value)
}

// s12 (bool)
//
// Delay 12 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS12(value bool) error {
	return e.Element.SetProperty("s12", value)
}

// s13 (bool)
//
// Delay 13 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS13(value bool) error {
	return e.Element.SetProperty("s13", value)
}

// s14 (bool)
//
// Delay 14 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS14(value bool) error {
	return e.Element.SetProperty("s14", value)
}

// s15 (bool)
//
// Delay 15 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS15(value bool) error {
	return e.Element.SetProperty("s15", value)
}

// s2 (bool)
//
// Delay 2 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS2(value bool) error {
	return e.Element.SetProperty("s2", value)
}

// s3 (bool)
//
// Delay 3 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS3(value bool) error {
	return e.Element.SetProperty("s3", value)
}

// s4 (bool)
//
// Delay 4 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS4(value bool) error {
	return e.Element.SetProperty("s4", value)
}

// s5 (bool)
//
// Delay 5 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS5(value bool) error {
	return e.Element.SetProperty("s5", value)
}

// s6 (bool)
//
// Delay 6 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS6(value bool) error {
	return e.Element.SetProperty("s6", value)
}

// s7 (bool)
//
// Delay 7 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS7(value bool) error {
	return e.Element.SetProperty("s7", value)
}

// s8 (bool)
//
// Delay 8 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS8(value bool) error {
	return e.Element.SetProperty("s8", value)
}

// s9 (bool)
//
// Delay 9 solo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetS9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("s9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetS9(value bool) error {
	return e.Element.SetProperty("s9", value)
}

// strch (float32)
//
// Stretch time

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetStrch() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("strch")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetStrch(value float32) error {
	return e.Element.SetProperty("strch", value)
}

// sync (bool)
//
// Tempo sync

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// temp (float32)
//
// Tem

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetTemp() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("temp")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetTemp(value float32) error {
	return e.Element.SetProperty("temp", value)
}

// tempo (float32)
//
// Tempo

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetTempo() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("tempo")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetTempo(value float32) error {
	return e.Element.SetProperty("tempo", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetWet() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("wet")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wm (bool)
//
// Wet mute

func (e *LspPlugInPluginsLv2SlapDelayStereo) GetWm() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wm")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayStereo) SetWm(value bool) error {
	return e.Element.SetProperty("wm", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2SlapDelayStereodm14 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm14Off LspPlugInPluginsLv2SlapDelayStereodm14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm14Time LspPlugInPluginsLv2SlapDelayStereodm14 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm14Distance LspPlugInPluginsLv2SlapDelayStereodm14 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm14Notes LspPlugInPluginsLv2SlapDelayStereodm14 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm2 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm2Off LspPlugInPluginsLv2SlapDelayStereodm2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm2Time LspPlugInPluginsLv2SlapDelayStereodm2 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm2Distance LspPlugInPluginsLv2SlapDelayStereodm2 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm2Notes LspPlugInPluginsLv2SlapDelayStereodm2 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm6 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm6Off LspPlugInPluginsLv2SlapDelayStereodm6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm6Time LspPlugInPluginsLv2SlapDelayStereodm6 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm6Distance LspPlugInPluginsLv2SlapDelayStereodm6 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm6Notes LspPlugInPluginsLv2SlapDelayStereodm6 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm9 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm9Off LspPlugInPluginsLv2SlapDelayStereodm9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm9Time LspPlugInPluginsLv2SlapDelayStereodm9 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm9Distance LspPlugInPluginsLv2SlapDelayStereodm9 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm9Notes LspPlugInPluginsLv2SlapDelayStereodm9 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm10 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm10Off LspPlugInPluginsLv2SlapDelayStereodm10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm10Time LspPlugInPluginsLv2SlapDelayStereodm10 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm10Distance LspPlugInPluginsLv2SlapDelayStereodm10 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm10Notes LspPlugInPluginsLv2SlapDelayStereodm10 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm13 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm13Off LspPlugInPluginsLv2SlapDelayStereodm13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm13Time LspPlugInPluginsLv2SlapDelayStereodm13 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm13Distance LspPlugInPluginsLv2SlapDelayStereodm13 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm13Notes LspPlugInPluginsLv2SlapDelayStereodm13 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm8 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm8Off LspPlugInPluginsLv2SlapDelayStereodm8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm8Time LspPlugInPluginsLv2SlapDelayStereodm8 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm8Distance LspPlugInPluginsLv2SlapDelayStereodm8 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm8Notes LspPlugInPluginsLv2SlapDelayStereodm8 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm4 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm4Off LspPlugInPluginsLv2SlapDelayStereodm4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm4Time LspPlugInPluginsLv2SlapDelayStereodm4 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm4Distance LspPlugInPluginsLv2SlapDelayStereodm4 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm4Notes LspPlugInPluginsLv2SlapDelayStereodm4 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm5 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm5Off LspPlugInPluginsLv2SlapDelayStereodm5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm5Time LspPlugInPluginsLv2SlapDelayStereodm5 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm5Distance LspPlugInPluginsLv2SlapDelayStereodm5 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm5Notes LspPlugInPluginsLv2SlapDelayStereodm5 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereolsel string

const (
	LspPlugInPluginsLv2SlapDelayStereolsel04 LspPlugInPluginsLv2SlapDelayStereolsel = "0-4" // 0-4 (0) – 0-4
	LspPlugInPluginsLv2SlapDelayStereolsel57 LspPlugInPluginsLv2SlapDelayStereolsel = "5-7" // 5-7 (1) – 5-7
	LspPlugInPluginsLv2SlapDelayStereolsel811 LspPlugInPluginsLv2SlapDelayStereolsel = "8-11" // 8-11 (2) – 8-11
	LspPlugInPluginsLv2SlapDelayStereolsel1215 LspPlugInPluginsLv2SlapDelayStereolsel = "12-15" // 12-15 (3) – 12-15
)

type LspPlugInPluginsLv2SlapDelayStereodm12 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm12Off LspPlugInPluginsLv2SlapDelayStereodm12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm12Time LspPlugInPluginsLv2SlapDelayStereodm12 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm12Distance LspPlugInPluginsLv2SlapDelayStereodm12 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm12Notes LspPlugInPluginsLv2SlapDelayStereodm12 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm15 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm15Off LspPlugInPluginsLv2SlapDelayStereodm15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm15Time LspPlugInPluginsLv2SlapDelayStereodm15 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm15Distance LspPlugInPluginsLv2SlapDelayStereodm15 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm15Notes LspPlugInPluginsLv2SlapDelayStereodm15 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm11 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm11Off LspPlugInPluginsLv2SlapDelayStereodm11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm11Time LspPlugInPluginsLv2SlapDelayStereodm11 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm11Distance LspPlugInPluginsLv2SlapDelayStereodm11 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm11Notes LspPlugInPluginsLv2SlapDelayStereodm11 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm3 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm3Off LspPlugInPluginsLv2SlapDelayStereodm3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm3Time LspPlugInPluginsLv2SlapDelayStereodm3 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm3Distance LspPlugInPluginsLv2SlapDelayStereodm3 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm3Notes LspPlugInPluginsLv2SlapDelayStereodm3 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm7 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm7Off LspPlugInPluginsLv2SlapDelayStereodm7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm7Time LspPlugInPluginsLv2SlapDelayStereodm7 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm7Distance LspPlugInPluginsLv2SlapDelayStereodm7 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm7Notes LspPlugInPluginsLv2SlapDelayStereodm7 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm0 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm0Off LspPlugInPluginsLv2SlapDelayStereodm0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm0Time LspPlugInPluginsLv2SlapDelayStereodm0 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm0Distance LspPlugInPluginsLv2SlapDelayStereodm0 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm0Notes LspPlugInPluginsLv2SlapDelayStereodm0 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayStereodm1 string

const (
	LspPlugInPluginsLv2SlapDelayStereodm1Off LspPlugInPluginsLv2SlapDelayStereodm1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayStereodm1Time LspPlugInPluginsLv2SlapDelayStereodm1 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayStereodm1Distance LspPlugInPluginsLv2SlapDelayStereodm1 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayStereodm1Notes LspPlugInPluginsLv2SlapDelayStereodm1 = "Notes" // Notes (3) – Notes
)

