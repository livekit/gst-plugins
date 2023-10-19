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

type LspPlugInPluginsLv2SlapDelayMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2SlapDelayMono() (*LspPlugInPluginsLv2SlapDelayMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-slap-delay-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SlapDelayMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2SlapDelayMonoWithName(name string) (*LspPlugInPluginsLv2SlapDelayMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-slap-delay-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2SlapDelayMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// dd0 (float32)
//
// Delay 0 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd0(value float32) error {
	return e.Element.SetProperty("dd0", value)
}

// dd1 (float32)
//
// Delay 1 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd1(value float32) error {
	return e.Element.SetProperty("dd1", value)
}

// dd10 (float32)
//
// Delay 10 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd10(value float32) error {
	return e.Element.SetProperty("dd10", value)
}

// dd11 (float32)
//
// Delay 11 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd11(value float32) error {
	return e.Element.SetProperty("dd11", value)
}

// dd12 (float32)
//
// Delay 12 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd12(value float32) error {
	return e.Element.SetProperty("dd12", value)
}

// dd13 (float32)
//
// Delay 13 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd13(value float32) error {
	return e.Element.SetProperty("dd13", value)
}

// dd14 (float32)
//
// Delay 14 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd14(value float32) error {
	return e.Element.SetProperty("dd14", value)
}

// dd15 (float32)
//
// Delay 15 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd15(value float32) error {
	return e.Element.SetProperty("dd15", value)
}

// dd2 (float32)
//
// Delay 2 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd2(value float32) error {
	return e.Element.SetProperty("dd2", value)
}

// dd3 (float32)
//
// Delay 3 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd3(value float32) error {
	return e.Element.SetProperty("dd3", value)
}

// dd4 (float32)
//
// Delay 4 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd4(value float32) error {
	return e.Element.SetProperty("dd4", value)
}

// dd5 (float32)
//
// Delay 5 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd5(value float32) error {
	return e.Element.SetProperty("dd5", value)
}

// dd6 (float32)
//
// Delay 6 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd6(value float32) error {
	return e.Element.SetProperty("dd6", value)
}

// dd7 (float32)
//
// Delay 7 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd7(value float32) error {
	return e.Element.SetProperty("dd7", value)
}

// dd8 (float32)
//
// Delay 8 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd8(value float32) error {
	return e.Element.SetProperty("dd8", value)
}

// dd9 (float32)
//
// Delay 9 distance

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDd9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDd9(value float32) error {
	return e.Element.SetProperty("dd9", value)
}

// df0 (float32)
//
// Delay 0 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf0(value float32) error {
	return e.Element.SetProperty("df0", value)
}

// df1 (float32)
//
// Delay 1 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf1(value float32) error {
	return e.Element.SetProperty("df1", value)
}

// df10 (float32)
//
// Delay 10 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf10(value float32) error {
	return e.Element.SetProperty("df10", value)
}

// df11 (float32)
//
// Delay 11 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf11(value float32) error {
	return e.Element.SetProperty("df11", value)
}

// df12 (float32)
//
// Delay 12 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf12(value float32) error {
	return e.Element.SetProperty("df12", value)
}

// df13 (float32)
//
// Delay 13 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf13(value float32) error {
	return e.Element.SetProperty("df13", value)
}

// df14 (float32)
//
// Delay 14 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf14(value float32) error {
	return e.Element.SetProperty("df14", value)
}

// df15 (float32)
//
// Delay 15 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf15(value float32) error {
	return e.Element.SetProperty("df15", value)
}

// df2 (float32)
//
// Delay 2 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf2(value float32) error {
	return e.Element.SetProperty("df2", value)
}

// df3 (float32)
//
// Delay 3 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf3(value float32) error {
	return e.Element.SetProperty("df3", value)
}

// df4 (float32)
//
// Delay 4 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf4(value float32) error {
	return e.Element.SetProperty("df4", value)
}

// df5 (float32)
//
// Delay 5 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf5(value float32) error {
	return e.Element.SetProperty("df5", value)
}

// df6 (float32)
//
// Delay 6 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf6(value float32) error {
	return e.Element.SetProperty("df6", value)
}

// df7 (float32)
//
// Delay 7 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf7(value float32) error {
	return e.Element.SetProperty("df7", value)
}

// df8 (float32)
//
// Delay 8 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf8(value float32) error {
	return e.Element.SetProperty("df8", value)
}

// df9 (float32)
//
// Delay 9 fraction

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDf9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDf9(value float32) error {
	return e.Element.SetProperty("df9", value)
}

// dg0 (float32)
//
// Delay 0 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg0(value float32) error {
	return e.Element.SetProperty("dg0", value)
}

// dg1 (float32)
//
// Delay 1 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg1(value float32) error {
	return e.Element.SetProperty("dg1", value)
}

// dg10 (float32)
//
// Delay 10 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg10(value float32) error {
	return e.Element.SetProperty("dg10", value)
}

// dg11 (float32)
//
// Delay 11 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg11(value float32) error {
	return e.Element.SetProperty("dg11", value)
}

// dg12 (float32)
//
// Delay 12 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg12(value float32) error {
	return e.Element.SetProperty("dg12", value)
}

// dg13 (float32)
//
// Delay 13 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg13(value float32) error {
	return e.Element.SetProperty("dg13", value)
}

// dg14 (float32)
//
// Delay 14 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg14(value float32) error {
	return e.Element.SetProperty("dg14", value)
}

// dg15 (float32)
//
// Delay 15 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg15(value float32) error {
	return e.Element.SetProperty("dg15", value)
}

// dg2 (float32)
//
// Delay 2 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg2(value float32) error {
	return e.Element.SetProperty("dg2", value)
}

// dg3 (float32)
//
// Delay 3 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg3(value float32) error {
	return e.Element.SetProperty("dg3", value)
}

// dg4 (float32)
//
// Delay 4 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg4(value float32) error {
	return e.Element.SetProperty("dg4", value)
}

// dg5 (float32)
//
// Delay 5 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg5(value float32) error {
	return e.Element.SetProperty("dg5", value)
}

// dg6 (float32)
//
// Delay 6 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg6(value float32) error {
	return e.Element.SetProperty("dg6", value)
}

// dg7 (float32)
//
// Delay 7 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg7(value float32) error {
	return e.Element.SetProperty("dg7", value)
}

// dg8 (float32)
//
// Delay 8 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg8(value float32) error {
	return e.Element.SetProperty("dg8", value)
}

// dg9 (float32)
//
// Delay 9 gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDg9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDg9(value float32) error {
	return e.Element.SetProperty("dg9", value)
}

// dm (bool)
//
// Dry mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm(value bool) error {
	return e.Element.SetProperty("dm", value)
}

// dm0 (GstLspPlugInPluginsLv2SlapDelayMonodm0)
//
// Delay 0 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm0() (interface{}, error) {
	return e.Element.GetProperty("dm0")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm0(value interface{}) error {
	return e.Element.SetProperty("dm0", value)
}

// dm1 (GstLspPlugInPluginsLv2SlapDelayMonodm1)
//
// Delay 1 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm1() (interface{}, error) {
	return e.Element.GetProperty("dm1")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm1(value interface{}) error {
	return e.Element.SetProperty("dm1", value)
}

// dm10 (GstLspPlugInPluginsLv2SlapDelayMonodm10)
//
// Delay 10 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm10() (interface{}, error) {
	return e.Element.GetProperty("dm10")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm10(value interface{}) error {
	return e.Element.SetProperty("dm10", value)
}

// dm11 (GstLspPlugInPluginsLv2SlapDelayMonodm11)
//
// Delay 11 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm11() (interface{}, error) {
	return e.Element.GetProperty("dm11")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm11(value interface{}) error {
	return e.Element.SetProperty("dm11", value)
}

// dm12 (GstLspPlugInPluginsLv2SlapDelayMonodm12)
//
// Delay 12 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm12() (interface{}, error) {
	return e.Element.GetProperty("dm12")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm12(value interface{}) error {
	return e.Element.SetProperty("dm12", value)
}

// dm13 (GstLspPlugInPluginsLv2SlapDelayMonodm13)
//
// Delay 13 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm13() (interface{}, error) {
	return e.Element.GetProperty("dm13")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm13(value interface{}) error {
	return e.Element.SetProperty("dm13", value)
}

// dm14 (GstLspPlugInPluginsLv2SlapDelayMonodm14)
//
// Delay 14 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm14() (interface{}, error) {
	return e.Element.GetProperty("dm14")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm14(value interface{}) error {
	return e.Element.SetProperty("dm14", value)
}

// dm15 (GstLspPlugInPluginsLv2SlapDelayMonodm15)
//
// Delay 15 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm15() (interface{}, error) {
	return e.Element.GetProperty("dm15")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm15(value interface{}) error {
	return e.Element.SetProperty("dm15", value)
}

// dm2 (GstLspPlugInPluginsLv2SlapDelayMonodm2)
//
// Delay 2 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm2() (interface{}, error) {
	return e.Element.GetProperty("dm2")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm2(value interface{}) error {
	return e.Element.SetProperty("dm2", value)
}

// dm3 (GstLspPlugInPluginsLv2SlapDelayMonodm3)
//
// Delay 3 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm3() (interface{}, error) {
	return e.Element.GetProperty("dm3")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm3(value interface{}) error {
	return e.Element.SetProperty("dm3", value)
}

// dm4 (GstLspPlugInPluginsLv2SlapDelayMonodm4)
//
// Delay 4 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm4() (interface{}, error) {
	return e.Element.GetProperty("dm4")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm4(value interface{}) error {
	return e.Element.SetProperty("dm4", value)
}

// dm5 (GstLspPlugInPluginsLv2SlapDelayMonodm5)
//
// Delay 5 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm5() (interface{}, error) {
	return e.Element.GetProperty("dm5")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm5(value interface{}) error {
	return e.Element.SetProperty("dm5", value)
}

// dm6 (GstLspPlugInPluginsLv2SlapDelayMonodm6)
//
// Delay 6 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm6() (interface{}, error) {
	return e.Element.GetProperty("dm6")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm6(value interface{}) error {
	return e.Element.SetProperty("dm6", value)
}

// dm7 (GstLspPlugInPluginsLv2SlapDelayMonodm7)
//
// Delay 7 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm7() (interface{}, error) {
	return e.Element.GetProperty("dm7")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm7(value interface{}) error {
	return e.Element.SetProperty("dm7", value)
}

// dm8 (GstLspPlugInPluginsLv2SlapDelayMonodm8)
//
// Delay 8 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm8() (interface{}, error) {
	return e.Element.GetProperty("dm8")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm8(value interface{}) error {
	return e.Element.SetProperty("dm8", value)
}

// dm9 (GstLspPlugInPluginsLv2SlapDelayMonodm9)
//
// Delay 9 mode

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDm9() (interface{}, error) {
	return e.Element.GetProperty("dm9")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDm9(value interface{}) error {
	return e.Element.SetProperty("dm9", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// ds0 (int)
//
// Delay 0 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs0() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs0(value int) error {
	return e.Element.SetProperty("ds0", value)
}

// ds1 (int)
//
// Delay 1 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs1() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs1(value int) error {
	return e.Element.SetProperty("ds1", value)
}

// ds10 (int)
//
// Delay 10 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs10() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs10(value int) error {
	return e.Element.SetProperty("ds10", value)
}

// ds11 (int)
//
// Delay 11 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs11() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs11(value int) error {
	return e.Element.SetProperty("ds11", value)
}

// ds12 (int)
//
// Delay 12 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs12() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs12(value int) error {
	return e.Element.SetProperty("ds12", value)
}

// ds13 (int)
//
// Delay 13 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs13() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs13(value int) error {
	return e.Element.SetProperty("ds13", value)
}

// ds14 (int)
//
// Delay 14 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs14() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs14(value int) error {
	return e.Element.SetProperty("ds14", value)
}

// ds15 (int)
//
// Delay 15 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs15() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs15(value int) error {
	return e.Element.SetProperty("ds15", value)
}

// ds2 (int)
//
// Delay 2 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs2() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs2(value int) error {
	return e.Element.SetProperty("ds2", value)
}

// ds3 (int)
//
// Delay 3 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs3() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs3(value int) error {
	return e.Element.SetProperty("ds3", value)
}

// ds4 (int)
//
// Delay 4 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs4() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs4(value int) error {
	return e.Element.SetProperty("ds4", value)
}

// ds5 (int)
//
// Delay 5 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs5() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs5(value int) error {
	return e.Element.SetProperty("ds5", value)
}

// ds6 (int)
//
// Delay 6 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs6() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs6(value int) error {
	return e.Element.SetProperty("ds6", value)
}

// ds7 (int)
//
// Delay 7 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs7() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs7(value int) error {
	return e.Element.SetProperty("ds7", value)
}

// ds8 (int)
//
// Delay 8 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs8() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs8(value int) error {
	return e.Element.SetProperty("ds8", value)
}

// ds9 (int)
//
// Delay 9 denominator

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDs9() (int, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDs9(value int) error {
	return e.Element.SetProperty("ds9", value)
}

// dt0 (float32)
//
// Delay 0 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt0(value float32) error {
	return e.Element.SetProperty("dt0", value)
}

// dt1 (float32)
//
// Delay 1 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt1(value float32) error {
	return e.Element.SetProperty("dt1", value)
}

// dt10 (float32)
//
// Delay 10 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt10(value float32) error {
	return e.Element.SetProperty("dt10", value)
}

// dt11 (float32)
//
// Delay 11 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt11(value float32) error {
	return e.Element.SetProperty("dt11", value)
}

// dt12 (float32)
//
// Delay 12 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt12(value float32) error {
	return e.Element.SetProperty("dt12", value)
}

// dt13 (float32)
//
// Delay 13 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt13(value float32) error {
	return e.Element.SetProperty("dt13", value)
}

// dt14 (float32)
//
// Delay 14 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt14(value float32) error {
	return e.Element.SetProperty("dt14", value)
}

// dt15 (float32)
//
// Delay 15 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt15(value float32) error {
	return e.Element.SetProperty("dt15", value)
}

// dt2 (float32)
//
// Delay 2 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt2(value float32) error {
	return e.Element.SetProperty("dt2", value)
}

// dt3 (float32)
//
// Delay 3 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt3(value float32) error {
	return e.Element.SetProperty("dt3", value)
}

// dt4 (float32)
//
// Delay 4 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt4(value float32) error {
	return e.Element.SetProperty("dt4", value)
}

// dt5 (float32)
//
// Delay 5 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt5(value float32) error {
	return e.Element.SetProperty("dt5", value)
}

// dt6 (float32)
//
// Delay 6 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt6(value float32) error {
	return e.Element.SetProperty("dt6", value)
}

// dt7 (float32)
//
// Delay 7 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt7(value float32) error {
	return e.Element.SetProperty("dt7", value)
}

// dt8 (float32)
//
// Delay 8 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt8(value float32) error {
	return e.Element.SetProperty("dt8", value)
}

// dt9 (float32)
//
// Delay 9 time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetDt9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetDt9(value float32) error {
	return e.Element.SetProperty("dt9", value)
}

// eq0 (bool)
//
// Equalizer 0 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq0() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq0(value bool) error {
	return e.Element.SetProperty("eq0", value)
}

// eq1 (bool)
//
// Equalizer 1 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq1() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq1(value bool) error {
	return e.Element.SetProperty("eq1", value)
}

// eq10 (bool)
//
// Equalizer 10 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq10() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq10(value bool) error {
	return e.Element.SetProperty("eq10", value)
}

// eq11 (bool)
//
// Equalizer 11 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq11() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq11(value bool) error {
	return e.Element.SetProperty("eq11", value)
}

// eq12 (bool)
//
// Equalizer 12 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq12() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq12(value bool) error {
	return e.Element.SetProperty("eq12", value)
}

// eq13 (bool)
//
// Equalizer 13 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq13() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq13(value bool) error {
	return e.Element.SetProperty("eq13", value)
}

// eq14 (bool)
//
// Equalizer 14 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq14() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq14(value bool) error {
	return e.Element.SetProperty("eq14", value)
}

// eq15 (bool)
//
// Equalizer 15 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq15() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq15(value bool) error {
	return e.Element.SetProperty("eq15", value)
}

// eq2 (bool)
//
// Equalizer 2 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq2() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq2(value bool) error {
	return e.Element.SetProperty("eq2", value)
}

// eq3 (bool)
//
// Equalizer 3 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq3() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq3(value bool) error {
	return e.Element.SetProperty("eq3", value)
}

// eq4 (bool)
//
// Equalizer 4 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq4() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq4(value bool) error {
	return e.Element.SetProperty("eq4", value)
}

// eq5 (bool)
//
// Equalizer 5 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq5() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq5(value bool) error {
	return e.Element.SetProperty("eq5", value)
}

// eq6 (bool)
//
// Equalizer 6 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq6() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq6(value bool) error {
	return e.Element.SetProperty("eq6", value)
}

// eq7 (bool)
//
// Equalizer 7 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq7() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq7(value bool) error {
	return e.Element.SetProperty("eq7", value)
}

// eq8 (bool)
//
// Equalizer 8 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq8() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq8(value bool) error {
	return e.Element.SetProperty("eq8", value)
}

// eq9 (bool)
//
// Equalizer 9 on

func (e *LspPlugInPluginsLv2SlapDelayMono) GetEq9() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetEq9(value bool) error {
	return e.Element.SetProperty("eq9", value)
}

// fbb0 (float32)
//
// Delay 0 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb0(value float32) error {
	return e.Element.SetProperty("fbb0", value)
}

// fbb1 (float32)
//
// Delay 1 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb1(value float32) error {
	return e.Element.SetProperty("fbb1", value)
}

// fbb10 (float32)
//
// Delay 10 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb10(value float32) error {
	return e.Element.SetProperty("fbb10", value)
}

// fbb11 (float32)
//
// Delay 11 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb11(value float32) error {
	return e.Element.SetProperty("fbb11", value)
}

// fbb12 (float32)
//
// Delay 12 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb12(value float32) error {
	return e.Element.SetProperty("fbb12", value)
}

// fbb13 (float32)
//
// Delay 13 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb13(value float32) error {
	return e.Element.SetProperty("fbb13", value)
}

// fbb14 (float32)
//
// Delay 14 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb14(value float32) error {
	return e.Element.SetProperty("fbb14", value)
}

// fbb15 (float32)
//
// Delay 15 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb15(value float32) error {
	return e.Element.SetProperty("fbb15", value)
}

// fbb2 (float32)
//
// Delay 2 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb2(value float32) error {
	return e.Element.SetProperty("fbb2", value)
}

// fbb3 (float32)
//
// Delay 3 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb3(value float32) error {
	return e.Element.SetProperty("fbb3", value)
}

// fbb4 (float32)
//
// Delay 4 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb4(value float32) error {
	return e.Element.SetProperty("fbb4", value)
}

// fbb5 (float32)
//
// Delay 5 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb5(value float32) error {
	return e.Element.SetProperty("fbb5", value)
}

// fbb6 (float32)
//
// Delay 6 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb6(value float32) error {
	return e.Element.SetProperty("fbb6", value)
}

// fbb7 (float32)
//
// Delay 7 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb7(value float32) error {
	return e.Element.SetProperty("fbb7", value)
}

// fbb8 (float32)
//
// Delay 8 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb8(value float32) error {
	return e.Element.SetProperty("fbb8", value)
}

// fbb9 (float32)
//
// Delay 9 bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbb9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbb9(value float32) error {
	return e.Element.SetProperty("fbb9", value)
}

// fbm0 (float32)
//
// Delay 0 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm0(value float32) error {
	return e.Element.SetProperty("fbm0", value)
}

// fbm1 (float32)
//
// Delay 1 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm1(value float32) error {
	return e.Element.SetProperty("fbm1", value)
}

// fbm10 (float32)
//
// Delay 10 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm10(value float32) error {
	return e.Element.SetProperty("fbm10", value)
}

// fbm11 (float32)
//
// Delay 11 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm11(value float32) error {
	return e.Element.SetProperty("fbm11", value)
}

// fbm12 (float32)
//
// Delay 12 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm12(value float32) error {
	return e.Element.SetProperty("fbm12", value)
}

// fbm13 (float32)
//
// Delay 13 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm13(value float32) error {
	return e.Element.SetProperty("fbm13", value)
}

// fbm14 (float32)
//
// Delay 14 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm14(value float32) error {
	return e.Element.SetProperty("fbm14", value)
}

// fbm15 (float32)
//
// Delay 15 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm15(value float32) error {
	return e.Element.SetProperty("fbm15", value)
}

// fbm2 (float32)
//
// Delay 2 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm2(value float32) error {
	return e.Element.SetProperty("fbm2", value)
}

// fbm3 (float32)
//
// Delay 3 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm3(value float32) error {
	return e.Element.SetProperty("fbm3", value)
}

// fbm4 (float32)
//
// Delay 4 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm4(value float32) error {
	return e.Element.SetProperty("fbm4", value)
}

// fbm5 (float32)
//
// Delay 5 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm5(value float32) error {
	return e.Element.SetProperty("fbm5", value)
}

// fbm6 (float32)
//
// Delay 6 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm6(value float32) error {
	return e.Element.SetProperty("fbm6", value)
}

// fbm7 (float32)
//
// Delay 7 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm7(value float32) error {
	return e.Element.SetProperty("fbm7", value)
}

// fbm8 (float32)
//
// Delay 8 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm8(value float32) error {
	return e.Element.SetProperty("fbm8", value)
}

// fbm9 (float32)
//
// Delay 9 middle

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbm9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbm9(value float32) error {
	return e.Element.SetProperty("fbm9", value)
}

// fbp0 (float32)
//
// Delay 0 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp0(value float32) error {
	return e.Element.SetProperty("fbp0", value)
}

// fbp1 (float32)
//
// Delay 1 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp1(value float32) error {
	return e.Element.SetProperty("fbp1", value)
}

// fbp10 (float32)
//
// Delay 10 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp10(value float32) error {
	return e.Element.SetProperty("fbp10", value)
}

// fbp11 (float32)
//
// Delay 11 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp11(value float32) error {
	return e.Element.SetProperty("fbp11", value)
}

// fbp12 (float32)
//
// Delay 12 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp12(value float32) error {
	return e.Element.SetProperty("fbp12", value)
}

// fbp13 (float32)
//
// Delay 13 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp13(value float32) error {
	return e.Element.SetProperty("fbp13", value)
}

// fbp14 (float32)
//
// Delay 14 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp14(value float32) error {
	return e.Element.SetProperty("fbp14", value)
}

// fbp15 (float32)
//
// Delay 15 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp15(value float32) error {
	return e.Element.SetProperty("fbp15", value)
}

// fbp2 (float32)
//
// Delay 2 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp2(value float32) error {
	return e.Element.SetProperty("fbp2", value)
}

// fbp3 (float32)
//
// Delay 3 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp3(value float32) error {
	return e.Element.SetProperty("fbp3", value)
}

// fbp4 (float32)
//
// Delay 4 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp4(value float32) error {
	return e.Element.SetProperty("fbp4", value)
}

// fbp5 (float32)
//
// Delay 5 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp5(value float32) error {
	return e.Element.SetProperty("fbp5", value)
}

// fbp6 (float32)
//
// Delay 6 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp6(value float32) error {
	return e.Element.SetProperty("fbp6", value)
}

// fbp7 (float32)
//
// Delay 7 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp7(value float32) error {
	return e.Element.SetProperty("fbp7", value)
}

// fbp8 (float32)
//
// Delay 8 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp8(value float32) error {
	return e.Element.SetProperty("fbp8", value)
}

// fbp9 (float32)
//
// Delay 9 presence

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbp9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbp9(value float32) error {
	return e.Element.SetProperty("fbp9", value)
}

// fbs0 (float32)
//
// Delay 0 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs0(value float32) error {
	return e.Element.SetProperty("fbs0", value)
}

// fbs1 (float32)
//
// Delay 1 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs1(value float32) error {
	return e.Element.SetProperty("fbs1", value)
}

// fbs10 (float32)
//
// Delay 10 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs10(value float32) error {
	return e.Element.SetProperty("fbs10", value)
}

// fbs11 (float32)
//
// Delay 11 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs11(value float32) error {
	return e.Element.SetProperty("fbs11", value)
}

// fbs12 (float32)
//
// Delay 12 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs12(value float32) error {
	return e.Element.SetProperty("fbs12", value)
}

// fbs13 (float32)
//
// Delay 13 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs13(value float32) error {
	return e.Element.SetProperty("fbs13", value)
}

// fbs14 (float32)
//
// Delay 14 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs14(value float32) error {
	return e.Element.SetProperty("fbs14", value)
}

// fbs15 (float32)
//
// Delay 15 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs15(value float32) error {
	return e.Element.SetProperty("fbs15", value)
}

// fbs2 (float32)
//
// Delay 2 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs2(value float32) error {
	return e.Element.SetProperty("fbs2", value)
}

// fbs3 (float32)
//
// Delay 3 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs3(value float32) error {
	return e.Element.SetProperty("fbs3", value)
}

// fbs4 (float32)
//
// Delay 4 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs4(value float32) error {
	return e.Element.SetProperty("fbs4", value)
}

// fbs5 (float32)
//
// Delay 5 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs5(value float32) error {
	return e.Element.SetProperty("fbs5", value)
}

// fbs6 (float32)
//
// Delay 6 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs6(value float32) error {
	return e.Element.SetProperty("fbs6", value)
}

// fbs7 (float32)
//
// Delay 7 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs7(value float32) error {
	return e.Element.SetProperty("fbs7", value)
}

// fbs8 (float32)
//
// Delay 8 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs8(value float32) error {
	return e.Element.SetProperty("fbs8", value)
}

// fbs9 (float32)
//
// Delay 9 sub-bass

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbs9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbs9(value float32) error {
	return e.Element.SetProperty("fbs9", value)
}

// fbt0 (float32)
//
// Delay 0 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt0(value float32) error {
	return e.Element.SetProperty("fbt0", value)
}

// fbt1 (float32)
//
// Delay 1 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt1(value float32) error {
	return e.Element.SetProperty("fbt1", value)
}

// fbt10 (float32)
//
// Delay 10 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt10(value float32) error {
	return e.Element.SetProperty("fbt10", value)
}

// fbt11 (float32)
//
// Delay 11 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt11(value float32) error {
	return e.Element.SetProperty("fbt11", value)
}

// fbt12 (float32)
//
// Delay 12 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt12(value float32) error {
	return e.Element.SetProperty("fbt12", value)
}

// fbt13 (float32)
//
// Delay 13 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt13(value float32) error {
	return e.Element.SetProperty("fbt13", value)
}

// fbt14 (float32)
//
// Delay 14 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt14(value float32) error {
	return e.Element.SetProperty("fbt14", value)
}

// fbt15 (float32)
//
// Delay 15 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt15(value float32) error {
	return e.Element.SetProperty("fbt15", value)
}

// fbt2 (float32)
//
// Delay 2 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt2(value float32) error {
	return e.Element.SetProperty("fbt2", value)
}

// fbt3 (float32)
//
// Delay 3 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt3(value float32) error {
	return e.Element.SetProperty("fbt3", value)
}

// fbt4 (float32)
//
// Delay 4 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt4(value float32) error {
	return e.Element.SetProperty("fbt4", value)
}

// fbt5 (float32)
//
// Delay 5 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt5(value float32) error {
	return e.Element.SetProperty("fbt5", value)
}

// fbt6 (float32)
//
// Delay 6 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt6(value float32) error {
	return e.Element.SetProperty("fbt6", value)
}

// fbt7 (float32)
//
// Delay 7 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt7(value float32) error {
	return e.Element.SetProperty("fbt7", value)
}

// fbt8 (float32)
//
// Delay 8 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt8(value float32) error {
	return e.Element.SetProperty("fbt8", value)
}

// fbt9 (float32)
//
// Delay 9 treble

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFbt9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFbt9(value float32) error {
	return e.Element.SetProperty("fbt9", value)
}

// fhc0 (float32)
//
// Delay 0 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc0(value float32) error {
	return e.Element.SetProperty("fhc0", value)
}

// fhc1 (float32)
//
// Delay 1 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc1(value float32) error {
	return e.Element.SetProperty("fhc1", value)
}

// fhc10 (float32)
//
// Delay 10 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc10(value float32) error {
	return e.Element.SetProperty("fhc10", value)
}

// fhc11 (float32)
//
// Delay 11 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc11(value float32) error {
	return e.Element.SetProperty("fhc11", value)
}

// fhc12 (float32)
//
// Delay 12 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc12(value float32) error {
	return e.Element.SetProperty("fhc12", value)
}

// fhc13 (float32)
//
// Delay 13 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc13(value float32) error {
	return e.Element.SetProperty("fhc13", value)
}

// fhc14 (float32)
//
// Delay 14 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc14(value float32) error {
	return e.Element.SetProperty("fhc14", value)
}

// fhc15 (float32)
//
// Delay 15 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc15(value float32) error {
	return e.Element.SetProperty("fhc15", value)
}

// fhc2 (float32)
//
// Delay 2 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc2(value float32) error {
	return e.Element.SetProperty("fhc2", value)
}

// fhc3 (float32)
//
// Delay 3 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc3(value float32) error {
	return e.Element.SetProperty("fhc3", value)
}

// fhc4 (float32)
//
// Delay 4 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc4(value float32) error {
	return e.Element.SetProperty("fhc4", value)
}

// fhc5 (float32)
//
// Delay 5 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc5(value float32) error {
	return e.Element.SetProperty("fhc5", value)
}

// fhc6 (float32)
//
// Delay 6 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc6(value float32) error {
	return e.Element.SetProperty("fhc6", value)
}

// fhc7 (float32)
//
// Delay 7 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc7(value float32) error {
	return e.Element.SetProperty("fhc7", value)
}

// fhc8 (float32)
//
// Delay 8 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc8(value float32) error {
	return e.Element.SetProperty("fhc8", value)
}

// fhc9 (float32)
//
// Delay 9 high-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFhc9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFhc9(value float32) error {
	return e.Element.SetProperty("fhc9", value)
}

// flc0 (float32)
//
// Delay 0 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc0() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc0(value float32) error {
	return e.Element.SetProperty("flc0", value)
}

// flc1 (float32)
//
// Delay 1 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc1() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc1(value float32) error {
	return e.Element.SetProperty("flc1", value)
}

// flc10 (float32)
//
// Delay 10 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc10() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc10(value float32) error {
	return e.Element.SetProperty("flc10", value)
}

// flc11 (float32)
//
// Delay 11 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc11() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc11(value float32) error {
	return e.Element.SetProperty("flc11", value)
}

// flc12 (float32)
//
// Delay 12 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc12() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc12(value float32) error {
	return e.Element.SetProperty("flc12", value)
}

// flc13 (float32)
//
// Delay 13 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc13() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc13(value float32) error {
	return e.Element.SetProperty("flc13", value)
}

// flc14 (float32)
//
// Delay 14 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc14() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc14(value float32) error {
	return e.Element.SetProperty("flc14", value)
}

// flc15 (float32)
//
// Delay 15 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc15() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc15(value float32) error {
	return e.Element.SetProperty("flc15", value)
}

// flc2 (float32)
//
// Delay 2 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc2() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc2(value float32) error {
	return e.Element.SetProperty("flc2", value)
}

// flc3 (float32)
//
// Delay 3 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc3() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc3(value float32) error {
	return e.Element.SetProperty("flc3", value)
}

// flc4 (float32)
//
// Delay 4 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc4() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc4(value float32) error {
	return e.Element.SetProperty("flc4", value)
}

// flc5 (float32)
//
// Delay 5 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc5() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc5(value float32) error {
	return e.Element.SetProperty("flc5", value)
}

// flc6 (float32)
//
// Delay 6 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc6() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc6(value float32) error {
	return e.Element.SetProperty("flc6", value)
}

// flc7 (float32)
//
// Delay 7 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc7() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc7(value float32) error {
	return e.Element.SetProperty("flc7", value)
}

// flc8 (float32)
//
// Delay 8 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc8() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc8(value float32) error {
	return e.Element.SetProperty("flc8", value)
}

// flc9 (float32)
//
// Delay 9 low-cut frequency

func (e *LspPlugInPluginsLv2SlapDelayMono) GetFlc9() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetFlc9(value float32) error {
	return e.Element.SetProperty("flc9", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2SlapDelayMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// hfc0 (bool)
//
// Delay 0 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc0() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc0(value bool) error {
	return e.Element.SetProperty("hfc0", value)
}

// hfc1 (bool)
//
// Delay 1 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc1() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc1(value bool) error {
	return e.Element.SetProperty("hfc1", value)
}

// hfc10 (bool)
//
// Delay 10 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc10() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc10(value bool) error {
	return e.Element.SetProperty("hfc10", value)
}

// hfc11 (bool)
//
// Delay 11 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc11() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc11(value bool) error {
	return e.Element.SetProperty("hfc11", value)
}

// hfc12 (bool)
//
// Delay 12 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc12() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc12(value bool) error {
	return e.Element.SetProperty("hfc12", value)
}

// hfc13 (bool)
//
// Delay 13 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc13() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc13(value bool) error {
	return e.Element.SetProperty("hfc13", value)
}

// hfc14 (bool)
//
// Delay 14 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc14() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc14(value bool) error {
	return e.Element.SetProperty("hfc14", value)
}

// hfc15 (bool)
//
// Delay 15 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc15() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc15(value bool) error {
	return e.Element.SetProperty("hfc15", value)
}

// hfc2 (bool)
//
// Delay 2 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc2() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc2(value bool) error {
	return e.Element.SetProperty("hfc2", value)
}

// hfc3 (bool)
//
// Delay 3 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc3() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc3(value bool) error {
	return e.Element.SetProperty("hfc3", value)
}

// hfc4 (bool)
//
// Delay 4 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc4() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc4(value bool) error {
	return e.Element.SetProperty("hfc4", value)
}

// hfc5 (bool)
//
// Delay 5 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc5() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc5(value bool) error {
	return e.Element.SetProperty("hfc5", value)
}

// hfc6 (bool)
//
// Delay 6 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc6() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc6(value bool) error {
	return e.Element.SetProperty("hfc6", value)
}

// hfc7 (bool)
//
// Delay 7 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc7() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc7(value bool) error {
	return e.Element.SetProperty("hfc7", value)
}

// hfc8 (bool)
//
// Delay 8 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc8() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc8(value bool) error {
	return e.Element.SetProperty("hfc8", value)
}

// hfc9 (bool)
//
// Delay 9 high-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetHfc9() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetHfc9(value bool) error {
	return e.Element.SetProperty("hfc9", value)
}

// lfc0 (bool)
//
// Delay 0 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc0() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc0(value bool) error {
	return e.Element.SetProperty("lfc0", value)
}

// lfc1 (bool)
//
// Delay 1 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc1() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc1(value bool) error {
	return e.Element.SetProperty("lfc1", value)
}

// lfc10 (bool)
//
// Delay 10 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc10() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc10(value bool) error {
	return e.Element.SetProperty("lfc10", value)
}

// lfc11 (bool)
//
// Delay 11 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc11() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc11(value bool) error {
	return e.Element.SetProperty("lfc11", value)
}

// lfc12 (bool)
//
// Delay 12 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc12() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc12(value bool) error {
	return e.Element.SetProperty("lfc12", value)
}

// lfc13 (bool)
//
// Delay 13 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc13() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc13(value bool) error {
	return e.Element.SetProperty("lfc13", value)
}

// lfc14 (bool)
//
// Delay 14 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc14() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc14(value bool) error {
	return e.Element.SetProperty("lfc14", value)
}

// lfc15 (bool)
//
// Delay 15 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc15() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc15(value bool) error {
	return e.Element.SetProperty("lfc15", value)
}

// lfc2 (bool)
//
// Delay 2 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc2() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc2(value bool) error {
	return e.Element.SetProperty("lfc2", value)
}

// lfc3 (bool)
//
// Delay 3 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc3() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc3(value bool) error {
	return e.Element.SetProperty("lfc3", value)
}

// lfc4 (bool)
//
// Delay 4 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc4() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc4(value bool) error {
	return e.Element.SetProperty("lfc4", value)
}

// lfc5 (bool)
//
// Delay 5 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc5() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc5(value bool) error {
	return e.Element.SetProperty("lfc5", value)
}

// lfc6 (bool)
//
// Delay 6 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc6() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc6(value bool) error {
	return e.Element.SetProperty("lfc6", value)
}

// lfc7 (bool)
//
// Delay 7 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc7() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc7(value bool) error {
	return e.Element.SetProperty("lfc7", value)
}

// lfc8 (bool)
//
// Delay 8 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc8() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc8(value bool) error {
	return e.Element.SetProperty("lfc8", value)
}

// lfc9 (bool)
//
// Delay 9 low-cut

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLfc9() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLfc9(value bool) error {
	return e.Element.SetProperty("lfc9", value)
}

// lsel (GstLspPlugInPluginsLv2SlapDelayMonolsel)
//
// Delay line selector

func (e *LspPlugInPluginsLv2SlapDelayMono) GetLsel() (interface{}, error) {
	return e.Element.GetProperty("lsel")
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetLsel(value interface{}) error {
	return e.Element.SetProperty("lsel", value)
}

// m0 (bool)
//
// Delay 0 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM0() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM0(value bool) error {
	return e.Element.SetProperty("m0", value)
}

// m1 (bool)
//
// Delay 1 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM1() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM1(value bool) error {
	return e.Element.SetProperty("m1", value)
}

// m10 (bool)
//
// Delay 10 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM10() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM10(value bool) error {
	return e.Element.SetProperty("m10", value)
}

// m11 (bool)
//
// Delay 11 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM11() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM11(value bool) error {
	return e.Element.SetProperty("m11", value)
}

// m12 (bool)
//
// Delay 12 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM12() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM12(value bool) error {
	return e.Element.SetProperty("m12", value)
}

// m13 (bool)
//
// Delay 13 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM13() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM13(value bool) error {
	return e.Element.SetProperty("m13", value)
}

// m14 (bool)
//
// Delay 14 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM14() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM14(value bool) error {
	return e.Element.SetProperty("m14", value)
}

// m15 (bool)
//
// Delay 15 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM15() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM15(value bool) error {
	return e.Element.SetProperty("m15", value)
}

// m2 (bool)
//
// Delay 2 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM2() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM2(value bool) error {
	return e.Element.SetProperty("m2", value)
}

// m3 (bool)
//
// Delay 3 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM3() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM3(value bool) error {
	return e.Element.SetProperty("m3", value)
}

// m4 (bool)
//
// Delay 4 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM4() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM4(value bool) error {
	return e.Element.SetProperty("m4", value)
}

// m5 (bool)
//
// Delay 5 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM5() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM5(value bool) error {
	return e.Element.SetProperty("m5", value)
}

// m6 (bool)
//
// Delay 6 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM6() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM6(value bool) error {
	return e.Element.SetProperty("m6", value)
}

// m7 (bool)
//
// Delay 7 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM7() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM7(value bool) error {
	return e.Element.SetProperty("m7", value)
}

// m8 (bool)
//
// Delay 8 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM8() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM8(value bool) error {
	return e.Element.SetProperty("m8", value)
}

// m9 (bool)
//
// Delay 9 mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetM9() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetM9(value bool) error {
	return e.Element.SetProperty("m9", value)
}

// mono (bool)
//
// Mono output

func (e *LspPlugInPluginsLv2SlapDelayMono) GetMono() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetMono(value bool) error {
	return e.Element.SetProperty("mono", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2SlapDelayMono) GetOutLatency() (int, error) {
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

// p-in (float32)
//
// Input panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPIn() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p-in")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPIn(value float32) error {
	return e.Element.SetProperty("p-in", value)
}

// p0 (float32)
//
// Delay 0 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP0(value float32) error {
	return e.Element.SetProperty("p0", value)
}

// p1 (float32)
//
// Delay 1 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP1(value float32) error {
	return e.Element.SetProperty("p1", value)
}

// p10 (float32)
//
// Delay 10 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP10(value float32) error {
	return e.Element.SetProperty("p10", value)
}

// p11 (float32)
//
// Delay 11 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP11(value float32) error {
	return e.Element.SetProperty("p11", value)
}

// p12 (float32)
//
// Delay 12 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP12(value float32) error {
	return e.Element.SetProperty("p12", value)
}

// p13 (float32)
//
// Delay 13 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP13(value float32) error {
	return e.Element.SetProperty("p13", value)
}

// p14 (float32)
//
// Delay 14 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP14(value float32) error {
	return e.Element.SetProperty("p14", value)
}

// p15 (float32)
//
// Delay 15 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP15(value float32) error {
	return e.Element.SetProperty("p15", value)
}

// p2 (float32)
//
// Delay 2 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP2(value float32) error {
	return e.Element.SetProperty("p2", value)
}

// p3 (float32)
//
// Delay 3 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP3(value float32) error {
	return e.Element.SetProperty("p3", value)
}

// p4 (float32)
//
// Delay 4 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP4(value float32) error {
	return e.Element.SetProperty("p4", value)
}

// p5 (float32)
//
// Delay 5 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP5(value float32) error {
	return e.Element.SetProperty("p5", value)
}

// p6 (float32)
//
// Delay 6 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP6(value float32) error {
	return e.Element.SetProperty("p6", value)
}

// p7 (float32)
//
// Delay 7 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP7(value float32) error {
	return e.Element.SetProperty("p7", value)
}

// p8 (float32)
//
// Delay 8 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP8(value float32) error {
	return e.Element.SetProperty("p8", value)
}

// p9 (float32)
//
// Delay 9 panorama

func (e *LspPlugInPluginsLv2SlapDelayMono) GetP9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("p9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2SlapDelayMono) SetP9(value float32) error {
	return e.Element.SetProperty("p9", value)
}

// ph0 (bool)
//
// Delay 0 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh0() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh0(value bool) error {
	return e.Element.SetProperty("ph0", value)
}

// ph1 (bool)
//
// Delay 1 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh1() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh1(value bool) error {
	return e.Element.SetProperty("ph1", value)
}

// ph10 (bool)
//
// Delay 10 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh10() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh10(value bool) error {
	return e.Element.SetProperty("ph10", value)
}

// ph11 (bool)
//
// Delay 11 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh11() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh11(value bool) error {
	return e.Element.SetProperty("ph11", value)
}

// ph12 (bool)
//
// Delay 12 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh12() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh12(value bool) error {
	return e.Element.SetProperty("ph12", value)
}

// ph13 (bool)
//
// Delay 13 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh13() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh13(value bool) error {
	return e.Element.SetProperty("ph13", value)
}

// ph14 (bool)
//
// Delay 14 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh14() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh14(value bool) error {
	return e.Element.SetProperty("ph14", value)
}

// ph15 (bool)
//
// Delay 15 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh15() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh15(value bool) error {
	return e.Element.SetProperty("ph15", value)
}

// ph2 (bool)
//
// Delay 2 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh2() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh2(value bool) error {
	return e.Element.SetProperty("ph2", value)
}

// ph3 (bool)
//
// Delay 3 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh3() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh3(value bool) error {
	return e.Element.SetProperty("ph3", value)
}

// ph4 (bool)
//
// Delay 4 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh4() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh4(value bool) error {
	return e.Element.SetProperty("ph4", value)
}

// ph5 (bool)
//
// Delay 5 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh5() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh5(value bool) error {
	return e.Element.SetProperty("ph5", value)
}

// ph6 (bool)
//
// Delay 6 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh6() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh6(value bool) error {
	return e.Element.SetProperty("ph6", value)
}

// ph7 (bool)
//
// Delay 7 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh7() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh7(value bool) error {
	return e.Element.SetProperty("ph7", value)
}

// ph8 (bool)
//
// Delay 8 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh8() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh8(value bool) error {
	return e.Element.SetProperty("ph8", value)
}

// ph9 (bool)
//
// Delay 9 phase

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPh9() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPh9(value bool) error {
	return e.Element.SetProperty("ph9", value)
}

// pred (float32)
//
// Pre-delay

func (e *LspPlugInPluginsLv2SlapDelayMono) GetPred() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetPred(value float32) error {
	return e.Element.SetProperty("pred", value)
}

// ramp (bool)
//
// Ramping delay

func (e *LspPlugInPluginsLv2SlapDelayMono) GetRamp() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetRamp(value bool) error {
	return e.Element.SetProperty("ramp", value)
}

// s0 (bool)
//
// Delay 0 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS0() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS0(value bool) error {
	return e.Element.SetProperty("s0", value)
}

// s1 (bool)
//
// Delay 1 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS1() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS1(value bool) error {
	return e.Element.SetProperty("s1", value)
}

// s10 (bool)
//
// Delay 10 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS10() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS10(value bool) error {
	return e.Element.SetProperty("s10", value)
}

// s11 (bool)
//
// Delay 11 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS11() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS11(value bool) error {
	return e.Element.SetProperty("s11", value)
}

// s12 (bool)
//
// Delay 12 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS12() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS12(value bool) error {
	return e.Element.SetProperty("s12", value)
}

// s13 (bool)
//
// Delay 13 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS13() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS13(value bool) error {
	return e.Element.SetProperty("s13", value)
}

// s14 (bool)
//
// Delay 14 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS14() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS14(value bool) error {
	return e.Element.SetProperty("s14", value)
}

// s15 (bool)
//
// Delay 15 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS15() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS15(value bool) error {
	return e.Element.SetProperty("s15", value)
}

// s2 (bool)
//
// Delay 2 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS2() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS2(value bool) error {
	return e.Element.SetProperty("s2", value)
}

// s3 (bool)
//
// Delay 3 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS3() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS3(value bool) error {
	return e.Element.SetProperty("s3", value)
}

// s4 (bool)
//
// Delay 4 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS4() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS4(value bool) error {
	return e.Element.SetProperty("s4", value)
}

// s5 (bool)
//
// Delay 5 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS5() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS5(value bool) error {
	return e.Element.SetProperty("s5", value)
}

// s6 (bool)
//
// Delay 6 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS6() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS6(value bool) error {
	return e.Element.SetProperty("s6", value)
}

// s7 (bool)
//
// Delay 7 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS7() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS7(value bool) error {
	return e.Element.SetProperty("s7", value)
}

// s8 (bool)
//
// Delay 8 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS8() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS8(value bool) error {
	return e.Element.SetProperty("s8", value)
}

// s9 (bool)
//
// Delay 9 solo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetS9() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetS9(value bool) error {
	return e.Element.SetProperty("s9", value)
}

// strch (float32)
//
// Stretch time

func (e *LspPlugInPluginsLv2SlapDelayMono) GetStrch() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetStrch(value float32) error {
	return e.Element.SetProperty("strch", value)
}

// sync (bool)
//
// Tempo sync

func (e *LspPlugInPluginsLv2SlapDelayMono) GetSync() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// temp (float32)
//
// Tem

func (e *LspPlugInPluginsLv2SlapDelayMono) GetTemp() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetTemp(value float32) error {
	return e.Element.SetProperty("temp", value)
}

// tempo (float32)
//
// Tempo

func (e *LspPlugInPluginsLv2SlapDelayMono) GetTempo() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetTempo(value float32) error {
	return e.Element.SetProperty("tempo", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2SlapDelayMono) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// wm (bool)
//
// Wet mute

func (e *LspPlugInPluginsLv2SlapDelayMono) GetWm() (bool, error) {
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

func (e *LspPlugInPluginsLv2SlapDelayMono) SetWm(value bool) error {
	return e.Element.SetProperty("wm", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2SlapDelayMonodm0 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm0Off LspPlugInPluginsLv2SlapDelayMonodm0 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm0Time LspPlugInPluginsLv2SlapDelayMonodm0 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm0Distance LspPlugInPluginsLv2SlapDelayMonodm0 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm0Notes LspPlugInPluginsLv2SlapDelayMonodm0 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm1 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm1Off LspPlugInPluginsLv2SlapDelayMonodm1 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm1Time LspPlugInPluginsLv2SlapDelayMonodm1 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm1Distance LspPlugInPluginsLv2SlapDelayMonodm1 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm1Notes LspPlugInPluginsLv2SlapDelayMonodm1 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm15 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm15Off LspPlugInPluginsLv2SlapDelayMonodm15 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm15Time LspPlugInPluginsLv2SlapDelayMonodm15 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm15Distance LspPlugInPluginsLv2SlapDelayMonodm15 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm15Notes LspPlugInPluginsLv2SlapDelayMonodm15 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm9 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm9Off LspPlugInPluginsLv2SlapDelayMonodm9 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm9Time LspPlugInPluginsLv2SlapDelayMonodm9 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm9Distance LspPlugInPluginsLv2SlapDelayMonodm9 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm9Notes LspPlugInPluginsLv2SlapDelayMonodm9 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm11 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm11Off LspPlugInPluginsLv2SlapDelayMonodm11 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm11Time LspPlugInPluginsLv2SlapDelayMonodm11 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm11Distance LspPlugInPluginsLv2SlapDelayMonodm11 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm11Notes LspPlugInPluginsLv2SlapDelayMonodm11 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm4 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm4Off LspPlugInPluginsLv2SlapDelayMonodm4 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm4Time LspPlugInPluginsLv2SlapDelayMonodm4 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm4Distance LspPlugInPluginsLv2SlapDelayMonodm4 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm4Notes LspPlugInPluginsLv2SlapDelayMonodm4 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm5 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm5Off LspPlugInPluginsLv2SlapDelayMonodm5 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm5Time LspPlugInPluginsLv2SlapDelayMonodm5 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm5Distance LspPlugInPluginsLv2SlapDelayMonodm5 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm5Notes LspPlugInPluginsLv2SlapDelayMonodm5 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonolsel string

const (
	LspPlugInPluginsLv2SlapDelayMonolsel04 LspPlugInPluginsLv2SlapDelayMonolsel = "0-4" // 0-4 (0) – 0-4
	LspPlugInPluginsLv2SlapDelayMonolsel57 LspPlugInPluginsLv2SlapDelayMonolsel = "5-7" // 5-7 (1) – 5-7
	LspPlugInPluginsLv2SlapDelayMonolsel811 LspPlugInPluginsLv2SlapDelayMonolsel = "8-11" // 8-11 (2) – 8-11
	LspPlugInPluginsLv2SlapDelayMonolsel1215 LspPlugInPluginsLv2SlapDelayMonolsel = "12-15" // 12-15 (3) – 12-15
)

type LspPlugInPluginsLv2SlapDelayMonodm13 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm13Off LspPlugInPluginsLv2SlapDelayMonodm13 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm13Time LspPlugInPluginsLv2SlapDelayMonodm13 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm13Distance LspPlugInPluginsLv2SlapDelayMonodm13 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm13Notes LspPlugInPluginsLv2SlapDelayMonodm13 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm3 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm3Off LspPlugInPluginsLv2SlapDelayMonodm3 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm3Time LspPlugInPluginsLv2SlapDelayMonodm3 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm3Distance LspPlugInPluginsLv2SlapDelayMonodm3 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm3Notes LspPlugInPluginsLv2SlapDelayMonodm3 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm6 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm6Off LspPlugInPluginsLv2SlapDelayMonodm6 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm6Time LspPlugInPluginsLv2SlapDelayMonodm6 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm6Distance LspPlugInPluginsLv2SlapDelayMonodm6 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm6Notes LspPlugInPluginsLv2SlapDelayMonodm6 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm7 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm7Off LspPlugInPluginsLv2SlapDelayMonodm7 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm7Time LspPlugInPluginsLv2SlapDelayMonodm7 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm7Distance LspPlugInPluginsLv2SlapDelayMonodm7 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm7Notes LspPlugInPluginsLv2SlapDelayMonodm7 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm8 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm8Off LspPlugInPluginsLv2SlapDelayMonodm8 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm8Time LspPlugInPluginsLv2SlapDelayMonodm8 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm8Distance LspPlugInPluginsLv2SlapDelayMonodm8 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm8Notes LspPlugInPluginsLv2SlapDelayMonodm8 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm10 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm10Off LspPlugInPluginsLv2SlapDelayMonodm10 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm10Time LspPlugInPluginsLv2SlapDelayMonodm10 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm10Distance LspPlugInPluginsLv2SlapDelayMonodm10 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm10Notes LspPlugInPluginsLv2SlapDelayMonodm10 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm12 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm12Off LspPlugInPluginsLv2SlapDelayMonodm12 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm12Time LspPlugInPluginsLv2SlapDelayMonodm12 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm12Distance LspPlugInPluginsLv2SlapDelayMonodm12 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm12Notes LspPlugInPluginsLv2SlapDelayMonodm12 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm14 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm14Off LspPlugInPluginsLv2SlapDelayMonodm14 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm14Time LspPlugInPluginsLv2SlapDelayMonodm14 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm14Distance LspPlugInPluginsLv2SlapDelayMonodm14 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm14Notes LspPlugInPluginsLv2SlapDelayMonodm14 = "Notes" // Notes (3) – Notes
)

type LspPlugInPluginsLv2SlapDelayMonodm2 string

const (
	LspPlugInPluginsLv2SlapDelayMonodm2Off LspPlugInPluginsLv2SlapDelayMonodm2 = "Off" // Off (0) – Off
	LspPlugInPluginsLv2SlapDelayMonodm2Time LspPlugInPluginsLv2SlapDelayMonodm2 = "Time" // Time (1) – Time
	LspPlugInPluginsLv2SlapDelayMonodm2Distance LspPlugInPluginsLv2SlapDelayMonodm2 = "Distance" // Distance (2) – Distance
	LspPlugInPluginsLv2SlapDelayMonodm2Notes LspPlugInPluginsLv2SlapDelayMonodm2 = "Notes" // Notes (3) – Notes
)

