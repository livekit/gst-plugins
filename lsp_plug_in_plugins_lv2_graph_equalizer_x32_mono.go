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

type LspPlugInPluginsLv2GraphEqualizerX32Mono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GraphEqualizerX32Mono() (*LspPlugInPluginsLv2GraphEqualizerX32Mono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-graph-equalizer-x32-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX32Mono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GraphEqualizerX32MonoWithName(name string) (*LspPlugInPluginsLv2GraphEqualizerX32Mono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-graph-equalizer-x32-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX32Mono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2GraphEqualizerX32Monofft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fsel (GstLspPlugInPluginsLv2GraphEqualizerX32Monofsel)
//
// Band select

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFsel() (interface{}, error) {
	return e.Element.GetProperty("fsel")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetFsel(value interface{}) error {
	return e.Element.SetProperty("fsel", value)
}

// fv-0 (bool)
//
// Filter visibility  16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-1 (bool)
//
// Filter visibility  20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-10 (bool)
//
// Filter visibility  160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-11 (bool)
//
// Filter visibility  200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-12 (bool)
//
// Filter visibility  250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-13 (bool)
//
// Filter visibility  315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-14 (bool)
//
// Filter visibility  400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-15 (bool)
//
// Filter visibility  500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-16 (bool)
//
// Filter visibility  630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-17 (bool)
//
// Filter visibility  800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-18 (bool)
//
// Filter visibility  1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-19 (bool)
//
// Filter visibility  1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-2 (bool)
//
// Filter visibility  25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-20 (bool)
//
// Filter visibility  1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-21 (bool)
//
// Filter visibility  2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-22 (bool)
//
// Filter visibility  2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-23 (bool)
//
// Filter visibility  3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-24 (bool)
//
// Filter visibility  4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-25 (bool)
//
// Filter visibility  5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-26 (bool)
//
// Filter visibility  6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-27 (bool)
//
// Filter visibility  8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-28 (bool)
//
// Filter visibility  10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-29 (bool)
//
// Filter visibility  12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-3 (bool)
//
// Filter visibility  31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-30 (bool)
//
// Filter visibility  16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-31 (bool)
//
// Filter visibility  20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-4 (bool)
//
// Filter visibility  40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-5 (bool)
//
// Filter visibility  50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-6 (bool)
//
// Filter visibility  63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-7 (bool)
//
// Filter visibility  80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-8 (bool)
//
// Filter visibility  100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// fv-9 (bool)
//
// Filter visibility  125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetFv9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fv-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// g-0 (float32)
//
// Band gain 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG0(value float32) error {
	return e.Element.SetProperty("g-0", value)
}

// g-1 (float32)
//
// Band gain 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG1(value float32) error {
	return e.Element.SetProperty("g-1", value)
}

// g-10 (float32)
//
// Band gain 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG10(value float32) error {
	return e.Element.SetProperty("g-10", value)
}

// g-11 (float32)
//
// Band gain 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG11(value float32) error {
	return e.Element.SetProperty("g-11", value)
}

// g-12 (float32)
//
// Band gain 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG12(value float32) error {
	return e.Element.SetProperty("g-12", value)
}

// g-13 (float32)
//
// Band gain 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG13(value float32) error {
	return e.Element.SetProperty("g-13", value)
}

// g-14 (float32)
//
// Band gain 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG14(value float32) error {
	return e.Element.SetProperty("g-14", value)
}

// g-15 (float32)
//
// Band gain 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG15(value float32) error {
	return e.Element.SetProperty("g-15", value)
}

// g-16 (float32)
//
// Band gain 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG16(value float32) error {
	return e.Element.SetProperty("g-16", value)
}

// g-17 (float32)
//
// Band gain 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG17(value float32) error {
	return e.Element.SetProperty("g-17", value)
}

// g-18 (float32)
//
// Band gain 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG18(value float32) error {
	return e.Element.SetProperty("g-18", value)
}

// g-19 (float32)
//
// Band gain 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG19(value float32) error {
	return e.Element.SetProperty("g-19", value)
}

// g-2 (float32)
//
// Band gain 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG2(value float32) error {
	return e.Element.SetProperty("g-2", value)
}

// g-20 (float32)
//
// Band gain 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG20(value float32) error {
	return e.Element.SetProperty("g-20", value)
}

// g-21 (float32)
//
// Band gain 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG21(value float32) error {
	return e.Element.SetProperty("g-21", value)
}

// g-22 (float32)
//
// Band gain 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG22(value float32) error {
	return e.Element.SetProperty("g-22", value)
}

// g-23 (float32)
//
// Band gain 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG23(value float32) error {
	return e.Element.SetProperty("g-23", value)
}

// g-24 (float32)
//
// Band gain 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG24(value float32) error {
	return e.Element.SetProperty("g-24", value)
}

// g-25 (float32)
//
// Band gain 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG25(value float32) error {
	return e.Element.SetProperty("g-25", value)
}

// g-26 (float32)
//
// Band gain 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG26(value float32) error {
	return e.Element.SetProperty("g-26", value)
}

// g-27 (float32)
//
// Band gain 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG27(value float32) error {
	return e.Element.SetProperty("g-27", value)
}

// g-28 (float32)
//
// Band gain 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG28(value float32) error {
	return e.Element.SetProperty("g-28", value)
}

// g-29 (float32)
//
// Band gain 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG29(value float32) error {
	return e.Element.SetProperty("g-29", value)
}

// g-3 (float32)
//
// Band gain 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG3(value float32) error {
	return e.Element.SetProperty("g-3", value)
}

// g-30 (float32)
//
// Band gain 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG30(value float32) error {
	return e.Element.SetProperty("g-30", value)
}

// g-31 (float32)
//
// Band gain 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG31(value float32) error {
	return e.Element.SetProperty("g-31", value)
}

// g-4 (float32)
//
// Band gain 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG4(value float32) error {
	return e.Element.SetProperty("g-4", value)
}

// g-5 (float32)
//
// Band gain 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG5(value float32) error {
	return e.Element.SetProperty("g-5", value)
}

// g-6 (float32)
//
// Band gain 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG6(value float32) error {
	return e.Element.SetProperty("g-6", value)
}

// g-7 (float32)
//
// Band gain 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG7(value float32) error {
	return e.Element.SetProperty("g-7", value)
}

// g-8 (float32)
//
// Band gain 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG8(value float32) error {
	return e.Element.SetProperty("g-8", value)
}

// g-9 (float32)
//
// Band gain 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetG9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetG9(value float32) error {
	return e.Element.SetProperty("g-9", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// im (float32)
//
// Input signal meter

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetIm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("im")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// mode (GstLspPlugInPluginsLv2GraphEqualizerX32Monomode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetOutLatency() (int, error) {
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

// react (float32)
//
// FFT reactivity

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetReact() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("react")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetShift() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("shift")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// slope (GstLspPlugInPluginsLv2GraphEqualizerX32Monoslope)
//
// Filter slope

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetSlope() (interface{}, error) {
	return e.Element.GetProperty("slope")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetSlope(value interface{}) error {
	return e.Element.SetProperty("slope", value)
}

// sm (float32)
//
// Output signal meter

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetSm() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sm")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// xe-0 (bool)
//
// Band on 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe0(value bool) error {
	return e.Element.SetProperty("xe-0", value)
}

// xe-1 (bool)
//
// Band on 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe1(value bool) error {
	return e.Element.SetProperty("xe-1", value)
}

// xe-10 (bool)
//
// Band on 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe10(value bool) error {
	return e.Element.SetProperty("xe-10", value)
}

// xe-11 (bool)
//
// Band on 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe11(value bool) error {
	return e.Element.SetProperty("xe-11", value)
}

// xe-12 (bool)
//
// Band on 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe12(value bool) error {
	return e.Element.SetProperty("xe-12", value)
}

// xe-13 (bool)
//
// Band on 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe13(value bool) error {
	return e.Element.SetProperty("xe-13", value)
}

// xe-14 (bool)
//
// Band on 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe14(value bool) error {
	return e.Element.SetProperty("xe-14", value)
}

// xe-15 (bool)
//
// Band on 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe15(value bool) error {
	return e.Element.SetProperty("xe-15", value)
}

// xe-16 (bool)
//
// Band on 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe16(value bool) error {
	return e.Element.SetProperty("xe-16", value)
}

// xe-17 (bool)
//
// Band on 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe17(value bool) error {
	return e.Element.SetProperty("xe-17", value)
}

// xe-18 (bool)
//
// Band on 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe18(value bool) error {
	return e.Element.SetProperty("xe-18", value)
}

// xe-19 (bool)
//
// Band on 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe19(value bool) error {
	return e.Element.SetProperty("xe-19", value)
}

// xe-2 (bool)
//
// Band on 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe2(value bool) error {
	return e.Element.SetProperty("xe-2", value)
}

// xe-20 (bool)
//
// Band on 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe20(value bool) error {
	return e.Element.SetProperty("xe-20", value)
}

// xe-21 (bool)
//
// Band on 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe21(value bool) error {
	return e.Element.SetProperty("xe-21", value)
}

// xe-22 (bool)
//
// Band on 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe22(value bool) error {
	return e.Element.SetProperty("xe-22", value)
}

// xe-23 (bool)
//
// Band on 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe23(value bool) error {
	return e.Element.SetProperty("xe-23", value)
}

// xe-24 (bool)
//
// Band on 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe24(value bool) error {
	return e.Element.SetProperty("xe-24", value)
}

// xe-25 (bool)
//
// Band on 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe25(value bool) error {
	return e.Element.SetProperty("xe-25", value)
}

// xe-26 (bool)
//
// Band on 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe26(value bool) error {
	return e.Element.SetProperty("xe-26", value)
}

// xe-27 (bool)
//
// Band on 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe27(value bool) error {
	return e.Element.SetProperty("xe-27", value)
}

// xe-28 (bool)
//
// Band on 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe28(value bool) error {
	return e.Element.SetProperty("xe-28", value)
}

// xe-29 (bool)
//
// Band on 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe29(value bool) error {
	return e.Element.SetProperty("xe-29", value)
}

// xe-3 (bool)
//
// Band on 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe3(value bool) error {
	return e.Element.SetProperty("xe-3", value)
}

// xe-30 (bool)
//
// Band on 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe30(value bool) error {
	return e.Element.SetProperty("xe-30", value)
}

// xe-31 (bool)
//
// Band on 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe31(value bool) error {
	return e.Element.SetProperty("xe-31", value)
}

// xe-4 (bool)
//
// Band on 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe4(value bool) error {
	return e.Element.SetProperty("xe-4", value)
}

// xe-5 (bool)
//
// Band on 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe5(value bool) error {
	return e.Element.SetProperty("xe-5", value)
}

// xe-6 (bool)
//
// Band on 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe6(value bool) error {
	return e.Element.SetProperty("xe-6", value)
}

// xe-7 (bool)
//
// Band on 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe7(value bool) error {
	return e.Element.SetProperty("xe-7", value)
}

// xe-8 (bool)
//
// Band on 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe8(value bool) error {
	return e.Element.SetProperty("xe-8", value)
}

// xe-9 (bool)
//
// Band on 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXe9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xe-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXe9(value bool) error {
	return e.Element.SetProperty("xe-9", value)
}

// xm-0 (bool)
//
// Band mute 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm0(value bool) error {
	return e.Element.SetProperty("xm-0", value)
}

// xm-1 (bool)
//
// Band mute 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm1(value bool) error {
	return e.Element.SetProperty("xm-1", value)
}

// xm-10 (bool)
//
// Band mute 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm10(value bool) error {
	return e.Element.SetProperty("xm-10", value)
}

// xm-11 (bool)
//
// Band mute 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm11(value bool) error {
	return e.Element.SetProperty("xm-11", value)
}

// xm-12 (bool)
//
// Band mute 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm12(value bool) error {
	return e.Element.SetProperty("xm-12", value)
}

// xm-13 (bool)
//
// Band mute 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm13(value bool) error {
	return e.Element.SetProperty("xm-13", value)
}

// xm-14 (bool)
//
// Band mute 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm14(value bool) error {
	return e.Element.SetProperty("xm-14", value)
}

// xm-15 (bool)
//
// Band mute 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm15(value bool) error {
	return e.Element.SetProperty("xm-15", value)
}

// xm-16 (bool)
//
// Band mute 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm16(value bool) error {
	return e.Element.SetProperty("xm-16", value)
}

// xm-17 (bool)
//
// Band mute 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm17(value bool) error {
	return e.Element.SetProperty("xm-17", value)
}

// xm-18 (bool)
//
// Band mute 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm18(value bool) error {
	return e.Element.SetProperty("xm-18", value)
}

// xm-19 (bool)
//
// Band mute 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm19(value bool) error {
	return e.Element.SetProperty("xm-19", value)
}

// xm-2 (bool)
//
// Band mute 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm2(value bool) error {
	return e.Element.SetProperty("xm-2", value)
}

// xm-20 (bool)
//
// Band mute 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm20(value bool) error {
	return e.Element.SetProperty("xm-20", value)
}

// xm-21 (bool)
//
// Band mute 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm21(value bool) error {
	return e.Element.SetProperty("xm-21", value)
}

// xm-22 (bool)
//
// Band mute 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm22(value bool) error {
	return e.Element.SetProperty("xm-22", value)
}

// xm-23 (bool)
//
// Band mute 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm23(value bool) error {
	return e.Element.SetProperty("xm-23", value)
}

// xm-24 (bool)
//
// Band mute 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm24(value bool) error {
	return e.Element.SetProperty("xm-24", value)
}

// xm-25 (bool)
//
// Band mute 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm25(value bool) error {
	return e.Element.SetProperty("xm-25", value)
}

// xm-26 (bool)
//
// Band mute 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm26(value bool) error {
	return e.Element.SetProperty("xm-26", value)
}

// xm-27 (bool)
//
// Band mute 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm27(value bool) error {
	return e.Element.SetProperty("xm-27", value)
}

// xm-28 (bool)
//
// Band mute 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm28(value bool) error {
	return e.Element.SetProperty("xm-28", value)
}

// xm-29 (bool)
//
// Band mute 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm29(value bool) error {
	return e.Element.SetProperty("xm-29", value)
}

// xm-3 (bool)
//
// Band mute 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm3(value bool) error {
	return e.Element.SetProperty("xm-3", value)
}

// xm-30 (bool)
//
// Band mute 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm30(value bool) error {
	return e.Element.SetProperty("xm-30", value)
}

// xm-31 (bool)
//
// Band mute 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm31(value bool) error {
	return e.Element.SetProperty("xm-31", value)
}

// xm-4 (bool)
//
// Band mute 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm4(value bool) error {
	return e.Element.SetProperty("xm-4", value)
}

// xm-5 (bool)
//
// Band mute 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm5(value bool) error {
	return e.Element.SetProperty("xm-5", value)
}

// xm-6 (bool)
//
// Band mute 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm6(value bool) error {
	return e.Element.SetProperty("xm-6", value)
}

// xm-7 (bool)
//
// Band mute 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm7(value bool) error {
	return e.Element.SetProperty("xm-7", value)
}

// xm-8 (bool)
//
// Band mute 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm8(value bool) error {
	return e.Element.SetProperty("xm-8", value)
}

// xm-9 (bool)
//
// Band mute 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXm9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xm-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXm9(value bool) error {
	return e.Element.SetProperty("xm-9", value)
}

// xs-0 (bool)
//
// Band solo 16

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs0(value bool) error {
	return e.Element.SetProperty("xs-0", value)
}

// xs-1 (bool)
//
// Band solo 20

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs1(value bool) error {
	return e.Element.SetProperty("xs-1", value)
}

// xs-10 (bool)
//
// Band solo 160

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs10(value bool) error {
	return e.Element.SetProperty("xs-10", value)
}

// xs-11 (bool)
//
// Band solo 200

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs11(value bool) error {
	return e.Element.SetProperty("xs-11", value)
}

// xs-12 (bool)
//
// Band solo 250

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs12(value bool) error {
	return e.Element.SetProperty("xs-12", value)
}

// xs-13 (bool)
//
// Band solo 315

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs13(value bool) error {
	return e.Element.SetProperty("xs-13", value)
}

// xs-14 (bool)
//
// Band solo 400

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs14(value bool) error {
	return e.Element.SetProperty("xs-14", value)
}

// xs-15 (bool)
//
// Band solo 500

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs15(value bool) error {
	return e.Element.SetProperty("xs-15", value)
}

// xs-16 (bool)
//
// Band solo 630

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs16(value bool) error {
	return e.Element.SetProperty("xs-16", value)
}

// xs-17 (bool)
//
// Band solo 800

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs17(value bool) error {
	return e.Element.SetProperty("xs-17", value)
}

// xs-18 (bool)
//
// Band solo 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs18(value bool) error {
	return e.Element.SetProperty("xs-18", value)
}

// xs-19 (bool)
//
// Band solo 1.25K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs19(value bool) error {
	return e.Element.SetProperty("xs-19", value)
}

// xs-2 (bool)
//
// Band solo 25

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs2(value bool) error {
	return e.Element.SetProperty("xs-2", value)
}

// xs-20 (bool)
//
// Band solo 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs20(value bool) error {
	return e.Element.SetProperty("xs-20", value)
}

// xs-21 (bool)
//
// Band solo 2K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs21(value bool) error {
	return e.Element.SetProperty("xs-21", value)
}

// xs-22 (bool)
//
// Band solo 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs22(value bool) error {
	return e.Element.SetProperty("xs-22", value)
}

// xs-23 (bool)
//
// Band solo 3.15K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs23(value bool) error {
	return e.Element.SetProperty("xs-23", value)
}

// xs-24 (bool)
//
// Band solo 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs24(value bool) error {
	return e.Element.SetProperty("xs-24", value)
}

// xs-25 (bool)
//
// Band solo 5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs25(value bool) error {
	return e.Element.SetProperty("xs-25", value)
}

// xs-26 (bool)
//
// Band solo 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs26(value bool) error {
	return e.Element.SetProperty("xs-26", value)
}

// xs-27 (bool)
//
// Band solo 8K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs27(value bool) error {
	return e.Element.SetProperty("xs-27", value)
}

// xs-28 (bool)
//
// Band solo 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs28(value bool) error {
	return e.Element.SetProperty("xs-28", value)
}

// xs-29 (bool)
//
// Band solo 12.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs29(value bool) error {
	return e.Element.SetProperty("xs-29", value)
}

// xs-3 (bool)
//
// Band solo 31.5

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs3(value bool) error {
	return e.Element.SetProperty("xs-3", value)
}

// xs-30 (bool)
//
// Band solo 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs30(value bool) error {
	return e.Element.SetProperty("xs-30", value)
}

// xs-31 (bool)
//
// Band solo 20K

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs31(value bool) error {
	return e.Element.SetProperty("xs-31", value)
}

// xs-4 (bool)
//
// Band solo 40

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs4(value bool) error {
	return e.Element.SetProperty("xs-4", value)
}

// xs-5 (bool)
//
// Band solo 50

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs5(value bool) error {
	return e.Element.SetProperty("xs-5", value)
}

// xs-6 (bool)
//
// Band solo 63

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs6(value bool) error {
	return e.Element.SetProperty("xs-6", value)
}

// xs-7 (bool)
//
// Band solo 80

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs7(value bool) error {
	return e.Element.SetProperty("xs-7", value)
}

// xs-8 (bool)
//
// Band solo 100

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs8(value bool) error {
	return e.Element.SetProperty("xs-8", value)
}

// xs-9 (bool)
//
// Band solo 125

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetXs9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("xs-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetXs9(value bool) error {
	return e.Element.SetProperty("xs-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) GetZoom() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX32Mono) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GraphEqualizerX32Monofft string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MonofftOff LspPlugInPluginsLv2GraphEqualizerX32Monofft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2GraphEqualizerX32MonofftPostEq LspPlugInPluginsLv2GraphEqualizerX32Monofft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2GraphEqualizerX32MonofftPreEq LspPlugInPluginsLv2GraphEqualizerX32Monofft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2GraphEqualizerX32Monofsel string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MonofselBands015 LspPlugInPluginsLv2GraphEqualizerX32Monofsel = "Bands 0-15" // Bands 0-15 (0) – Bands 0-15
	LspPlugInPluginsLv2GraphEqualizerX32MonofselBands1631 LspPlugInPluginsLv2GraphEqualizerX32Monofsel = "Bands 16-31" // Bands 16-31 (1) – Bands 16-31
)

type LspPlugInPluginsLv2GraphEqualizerX32Monomode string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MonomodeIir LspPlugInPluginsLv2GraphEqualizerX32Monomode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2GraphEqualizerX32MonomodeFir LspPlugInPluginsLv2GraphEqualizerX32Monomode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2GraphEqualizerX32MonomodeFft LspPlugInPluginsLv2GraphEqualizerX32Monomode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2GraphEqualizerX32Monoslope string

const (
	LspPlugInPluginsLv2GraphEqualizerX32MonoslopeBt48 LspPlugInPluginsLv2GraphEqualizerX32Monoslope = "BT48" // BT48 (0) – BT48
	LspPlugInPluginsLv2GraphEqualizerX32MonoslopeMt48 LspPlugInPluginsLv2GraphEqualizerX32Monoslope = "MT48" // MT48 (1) – MT48
	LspPlugInPluginsLv2GraphEqualizerX32MonoslopeBt72 LspPlugInPluginsLv2GraphEqualizerX32Monoslope = "BT72" // BT72 (2) – BT72
	LspPlugInPluginsLv2GraphEqualizerX32MonoslopeMt72 LspPlugInPluginsLv2GraphEqualizerX32Monoslope = "MT72" // MT72 (3) – MT72
	LspPlugInPluginsLv2GraphEqualizerX32MonoslopeBt96 LspPlugInPluginsLv2GraphEqualizerX32Monoslope = "BT96" // BT96 (4) – BT96
	LspPlugInPluginsLv2GraphEqualizerX32MonoslopeMt96 LspPlugInPluginsLv2GraphEqualizerX32Monoslope = "MT96" // MT96 (5) – MT96
)

