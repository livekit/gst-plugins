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

type LspPlugInPluginsLv2GraphEqualizerX16Stereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2GraphEqualizerX16Stereo() (*LspPlugInPluginsLv2GraphEqualizerX16Stereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-graph-equalizer-x16-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX16Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2GraphEqualizerX16StereoWithName(name string) (*LspPlugInPluginsLv2GraphEqualizerX16Stereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-graph-equalizer-x16-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2GraphEqualizerX16Stereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bal (float32)
//
// Output balance

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetBal() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("bal")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetBal(value float32) error {
	return e.Element.SetProperty("bal", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// fft (GstLspPlugInPluginsLv2GraphEqualizerX16Stereofft)
//
// FFT analysis

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFft() (interface{}, error) {
	return e.Element.GetProperty("fft")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetFft(value interface{}) error {
	return e.Element.SetProperty("fft", value)
}

// fftv-l (bool)
//
// FFT visibility Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFftvL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fftv-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetFftvL(value bool) error {
	return e.Element.SetProperty("fftv-l", value)
}

// fftv-r (bool)
//
// FFT visibility Right

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFftvR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("fftv-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetFftvR(value bool) error {
	return e.Element.SetProperty("fftv-r", value)
}

// fv-0 (bool)
//
// Filter visibility  16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv0() (bool, error) {
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
// Filter visibility  25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv1() (bool, error) {
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
// Filter visibility  1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv10() (bool, error) {
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
// Filter visibility  2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv11() (bool, error) {
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
// Filter visibility  4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv12() (bool, error) {
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
// Filter visibility  6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv13() (bool, error) {
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
// Filter visibility  10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv14() (bool, error) {
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
// Filter visibility  16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv15() (bool, error) {
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

// fv-2 (bool)
//
// Filter visibility  40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv2() (bool, error) {
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

// fv-3 (bool)
//
// Filter visibility  63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv3() (bool, error) {
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

// fv-4 (bool)
//
// Filter visibility  100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv4() (bool, error) {
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
// Filter visibility  160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv5() (bool, error) {
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
// Filter visibility  250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv6() (bool, error) {
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
// Filter visibility  400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv7() (bool, error) {
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
// Filter visibility  630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv8() (bool, error) {
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
// Filter visibility  1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetFv9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG0() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG0(value float32) error {
	return e.Element.SetProperty("g-0", value)
}

// g-1 (float32)
//
// Band gain 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG1() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG1(value float32) error {
	return e.Element.SetProperty("g-1", value)
}

// g-10 (float32)
//
// Band gain 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG10() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG10(value float32) error {
	return e.Element.SetProperty("g-10", value)
}

// g-11 (float32)
//
// Band gain 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG11() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG11(value float32) error {
	return e.Element.SetProperty("g-11", value)
}

// g-12 (float32)
//
// Band gain 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG12() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG12(value float32) error {
	return e.Element.SetProperty("g-12", value)
}

// g-13 (float32)
//
// Band gain 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG13() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG13(value float32) error {
	return e.Element.SetProperty("g-13", value)
}

// g-14 (float32)
//
// Band gain 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG14() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG14(value float32) error {
	return e.Element.SetProperty("g-14", value)
}

// g-15 (float32)
//
// Band gain 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG15() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG15(value float32) error {
	return e.Element.SetProperty("g-15", value)
}

// g-2 (float32)
//
// Band gain 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG2() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG2(value float32) error {
	return e.Element.SetProperty("g-2", value)
}

// g-3 (float32)
//
// Band gain 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG3() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG3(value float32) error {
	return e.Element.SetProperty("g-3", value)
}

// g-4 (float32)
//
// Band gain 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG4() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG4(value float32) error {
	return e.Element.SetProperty("g-4", value)
}

// g-5 (float32)
//
// Band gain 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG5() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG5(value float32) error {
	return e.Element.SetProperty("g-5", value)
}

// g-6 (float32)
//
// Band gain 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG6() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG6(value float32) error {
	return e.Element.SetProperty("g-6", value)
}

// g-7 (float32)
//
// Band gain 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG7() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG7(value float32) error {
	return e.Element.SetProperty("g-7", value)
}

// g-8 (float32)
//
// Band gain 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG8() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG8(value float32) error {
	return e.Element.SetProperty("g-8", value)
}

// g-9 (float32)
//
// Band gain 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetG9() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetG9(value float32) error {
	return e.Element.SetProperty("g-9", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// iml (float32)
//
// Input signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetIml() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("iml")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// imr (float32)
//
// Input signal meter Right

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetImr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imr")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// mode (GstLspPlugInPluginsLv2GraphEqualizerX16Stereomode)
//
// Equalizer mode

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// slope (GstLspPlugInPluginsLv2GraphEqualizerX16Stereoslope)
//
// Filter slope

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetSlope() (interface{}, error) {
	return e.Element.GetProperty("slope")
}

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetSlope(value interface{}) error {
	return e.Element.SetProperty("slope", value)
}

// sml (float32)
//
// Output signal meter Left

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetSml() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sml")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// smr (float32)
//
// Output signal meter Right

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetSmr() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("smr")
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe0(value bool) error {
	return e.Element.SetProperty("xe-0", value)
}

// xe-1 (bool)
//
// Band on 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe1(value bool) error {
	return e.Element.SetProperty("xe-1", value)
}

// xe-10 (bool)
//
// Band on 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe10(value bool) error {
	return e.Element.SetProperty("xe-10", value)
}

// xe-11 (bool)
//
// Band on 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe11(value bool) error {
	return e.Element.SetProperty("xe-11", value)
}

// xe-12 (bool)
//
// Band on 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe12(value bool) error {
	return e.Element.SetProperty("xe-12", value)
}

// xe-13 (bool)
//
// Band on 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe13(value bool) error {
	return e.Element.SetProperty("xe-13", value)
}

// xe-14 (bool)
//
// Band on 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe14(value bool) error {
	return e.Element.SetProperty("xe-14", value)
}

// xe-15 (bool)
//
// Band on 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe15(value bool) error {
	return e.Element.SetProperty("xe-15", value)
}

// xe-2 (bool)
//
// Band on 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe2(value bool) error {
	return e.Element.SetProperty("xe-2", value)
}

// xe-3 (bool)
//
// Band on 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe3(value bool) error {
	return e.Element.SetProperty("xe-3", value)
}

// xe-4 (bool)
//
// Band on 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe4(value bool) error {
	return e.Element.SetProperty("xe-4", value)
}

// xe-5 (bool)
//
// Band on 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe5(value bool) error {
	return e.Element.SetProperty("xe-5", value)
}

// xe-6 (bool)
//
// Band on 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe6(value bool) error {
	return e.Element.SetProperty("xe-6", value)
}

// xe-7 (bool)
//
// Band on 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe7(value bool) error {
	return e.Element.SetProperty("xe-7", value)
}

// xe-8 (bool)
//
// Band on 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe8(value bool) error {
	return e.Element.SetProperty("xe-8", value)
}

// xe-9 (bool)
//
// Band on 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXe9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXe9(value bool) error {
	return e.Element.SetProperty("xe-9", value)
}

// xm-0 (bool)
//
// Band mute 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm0(value bool) error {
	return e.Element.SetProperty("xm-0", value)
}

// xm-1 (bool)
//
// Band mute 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm1(value bool) error {
	return e.Element.SetProperty("xm-1", value)
}

// xm-10 (bool)
//
// Band mute 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm10(value bool) error {
	return e.Element.SetProperty("xm-10", value)
}

// xm-11 (bool)
//
// Band mute 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm11(value bool) error {
	return e.Element.SetProperty("xm-11", value)
}

// xm-12 (bool)
//
// Band mute 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm12(value bool) error {
	return e.Element.SetProperty("xm-12", value)
}

// xm-13 (bool)
//
// Band mute 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm13(value bool) error {
	return e.Element.SetProperty("xm-13", value)
}

// xm-14 (bool)
//
// Band mute 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm14(value bool) error {
	return e.Element.SetProperty("xm-14", value)
}

// xm-15 (bool)
//
// Band mute 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm15(value bool) error {
	return e.Element.SetProperty("xm-15", value)
}

// xm-2 (bool)
//
// Band mute 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm2(value bool) error {
	return e.Element.SetProperty("xm-2", value)
}

// xm-3 (bool)
//
// Band mute 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm3(value bool) error {
	return e.Element.SetProperty("xm-3", value)
}

// xm-4 (bool)
//
// Band mute 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm4(value bool) error {
	return e.Element.SetProperty("xm-4", value)
}

// xm-5 (bool)
//
// Band mute 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm5(value bool) error {
	return e.Element.SetProperty("xm-5", value)
}

// xm-6 (bool)
//
// Band mute 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm6(value bool) error {
	return e.Element.SetProperty("xm-6", value)
}

// xm-7 (bool)
//
// Band mute 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm7(value bool) error {
	return e.Element.SetProperty("xm-7", value)
}

// xm-8 (bool)
//
// Band mute 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm8(value bool) error {
	return e.Element.SetProperty("xm-8", value)
}

// xm-9 (bool)
//
// Band mute 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXm9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXm9(value bool) error {
	return e.Element.SetProperty("xm-9", value)
}

// xs-0 (bool)
//
// Band solo 16

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs0() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs0(value bool) error {
	return e.Element.SetProperty("xs-0", value)
}

// xs-1 (bool)
//
// Band solo 25

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs1() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs1(value bool) error {
	return e.Element.SetProperty("xs-1", value)
}

// xs-10 (bool)
//
// Band solo 1.6K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs10() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs10(value bool) error {
	return e.Element.SetProperty("xs-10", value)
}

// xs-11 (bool)
//
// Band solo 2.5K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs11() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs11(value bool) error {
	return e.Element.SetProperty("xs-11", value)
}

// xs-12 (bool)
//
// Band solo 4K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs12() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs12(value bool) error {
	return e.Element.SetProperty("xs-12", value)
}

// xs-13 (bool)
//
// Band solo 6.3K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs13() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs13(value bool) error {
	return e.Element.SetProperty("xs-13", value)
}

// xs-14 (bool)
//
// Band solo 10K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs14() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs14(value bool) error {
	return e.Element.SetProperty("xs-14", value)
}

// xs-15 (bool)
//
// Band solo 16K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs15() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs15(value bool) error {
	return e.Element.SetProperty("xs-15", value)
}

// xs-2 (bool)
//
// Band solo 40

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs2() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs2(value bool) error {
	return e.Element.SetProperty("xs-2", value)
}

// xs-3 (bool)
//
// Band solo 63

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs3() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs3(value bool) error {
	return e.Element.SetProperty("xs-3", value)
}

// xs-4 (bool)
//
// Band solo 100

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs4() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs4(value bool) error {
	return e.Element.SetProperty("xs-4", value)
}

// xs-5 (bool)
//
// Band solo 160

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs5() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs5(value bool) error {
	return e.Element.SetProperty("xs-5", value)
}

// xs-6 (bool)
//
// Band solo 250

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs6() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs6(value bool) error {
	return e.Element.SetProperty("xs-6", value)
}

// xs-7 (bool)
//
// Band solo 400

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs7() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs7(value bool) error {
	return e.Element.SetProperty("xs-7", value)
}

// xs-8 (bool)
//
// Band solo 630

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs8() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs8(value bool) error {
	return e.Element.SetProperty("xs-8", value)
}

// xs-9 (bool)
//
// Band solo 1K

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetXs9() (bool, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetXs9(value bool) error {
	return e.Element.SetProperty("xs-9", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2GraphEqualizerX16Stereo) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2GraphEqualizerX16Stereofft string

const (
	LspPlugInPluginsLv2GraphEqualizerX16StereofftOff LspPlugInPluginsLv2GraphEqualizerX16Stereofft = "Off" // Off (0) – Off
	LspPlugInPluginsLv2GraphEqualizerX16StereofftPostEq LspPlugInPluginsLv2GraphEqualizerX16Stereofft = "Post-eq" // Post-eq (1) – Post-eq
	LspPlugInPluginsLv2GraphEqualizerX16StereofftPreEq LspPlugInPluginsLv2GraphEqualizerX16Stereofft = "Pre-eq" // Pre-eq (2) – Pre-eq
)

type LspPlugInPluginsLv2GraphEqualizerX16Stereomode string

const (
	LspPlugInPluginsLv2GraphEqualizerX16StereomodeIir LspPlugInPluginsLv2GraphEqualizerX16Stereomode = "IIR" // IIR (0) – IIR
	LspPlugInPluginsLv2GraphEqualizerX16StereomodeFir LspPlugInPluginsLv2GraphEqualizerX16Stereomode = "FIR" // FIR (1) – FIR
	LspPlugInPluginsLv2GraphEqualizerX16StereomodeFft LspPlugInPluginsLv2GraphEqualizerX16Stereomode = "FFT" // FFT (2) – FFT
)

type LspPlugInPluginsLv2GraphEqualizerX16Stereoslope string

const (
	LspPlugInPluginsLv2GraphEqualizerX16StereoslopeBt48 LspPlugInPluginsLv2GraphEqualizerX16Stereoslope = "BT48" // BT48 (0) – BT48
	LspPlugInPluginsLv2GraphEqualizerX16StereoslopeMt48 LspPlugInPluginsLv2GraphEqualizerX16Stereoslope = "MT48" // MT48 (1) – MT48
	LspPlugInPluginsLv2GraphEqualizerX16StereoslopeBt72 LspPlugInPluginsLv2GraphEqualizerX16Stereoslope = "BT72" // BT72 (2) – BT72
	LspPlugInPluginsLv2GraphEqualizerX16StereoslopeMt72 LspPlugInPluginsLv2GraphEqualizerX16Stereoslope = "MT72" // MT72 (3) – MT72
	LspPlugInPluginsLv2GraphEqualizerX16StereoslopeBt96 LspPlugInPluginsLv2GraphEqualizerX16Stereoslope = "BT96" // BT96 (4) – BT96
	LspPlugInPluginsLv2GraphEqualizerX16StereoslopeMt96 LspPlugInPluginsLv2GraphEqualizerX16Stereoslope = "MT96" // MT96 (5) – MT96
)

