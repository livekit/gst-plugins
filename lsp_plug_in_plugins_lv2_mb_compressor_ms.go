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

type LspPlugInPluginsLv2MbCompressorMs struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MbCompressorMs() (*LspPlugInPluginsLv2MbCompressorMs, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-mb-compressor-ms")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MbCompressorMsWithName(name string) (*LspPlugInPluginsLv2MbCompressorMs, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-mb-compressor-ms", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorMs{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-0m (float32)
//
// Attack level 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl0M(value float32) error {
	return e.Element.SetProperty("al-0m", value)
}

// al-0s (float32)
//
// Attack level 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl0S(value float32) error {
	return e.Element.SetProperty("al-0s", value)
}

// al-1m (float32)
//
// Attack level 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl1M(value float32) error {
	return e.Element.SetProperty("al-1m", value)
}

// al-1s (float32)
//
// Attack level 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl1S(value float32) error {
	return e.Element.SetProperty("al-1s", value)
}

// al-2m (float32)
//
// Attack level 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl2M(value float32) error {
	return e.Element.SetProperty("al-2m", value)
}

// al-2s (float32)
//
// Attack level 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl2S(value float32) error {
	return e.Element.SetProperty("al-2s", value)
}

// al-3m (float32)
//
// Attack level 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl3M(value float32) error {
	return e.Element.SetProperty("al-3m", value)
}

// al-3s (float32)
//
// Attack level 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl3S(value float32) error {
	return e.Element.SetProperty("al-3s", value)
}

// al-4m (float32)
//
// Attack level 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl4M(value float32) error {
	return e.Element.SetProperty("al-4m", value)
}

// al-4s (float32)
//
// Attack level 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl4S(value float32) error {
	return e.Element.SetProperty("al-4s", value)
}

// al-5m (float32)
//
// Attack level 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl5M(value float32) error {
	return e.Element.SetProperty("al-5m", value)
}

// al-5s (float32)
//
// Attack level 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl5S(value float32) error {
	return e.Element.SetProperty("al-5s", value)
}

// al-6m (float32)
//
// Attack level 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl6M(value float32) error {
	return e.Element.SetProperty("al-6m", value)
}

// al-6s (float32)
//
// Attack level 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl6S(value float32) error {
	return e.Element.SetProperty("al-6s", value)
}

// al-7m (float32)
//
// Attack level 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl7M(value float32) error {
	return e.Element.SetProperty("al-7m", value)
}

// al-7s (float32)
//
// Attack level 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAl7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAl7S(value float32) error {
	return e.Element.SetProperty("al-7s", value)
}

// at-0m (float32)
//
// Attack time 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt0M(value float32) error {
	return e.Element.SetProperty("at-0m", value)
}

// at-0s (float32)
//
// Attack time 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt0S(value float32) error {
	return e.Element.SetProperty("at-0s", value)
}

// at-1m (float32)
//
// Attack time 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt1M(value float32) error {
	return e.Element.SetProperty("at-1m", value)
}

// at-1s (float32)
//
// Attack time 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt1S(value float32) error {
	return e.Element.SetProperty("at-1s", value)
}

// at-2m (float32)
//
// Attack time 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt2M(value float32) error {
	return e.Element.SetProperty("at-2m", value)
}

// at-2s (float32)
//
// Attack time 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt2S(value float32) error {
	return e.Element.SetProperty("at-2s", value)
}

// at-3m (float32)
//
// Attack time 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt3M(value float32) error {
	return e.Element.SetProperty("at-3m", value)
}

// at-3s (float32)
//
// Attack time 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt3S(value float32) error {
	return e.Element.SetProperty("at-3s", value)
}

// at-4m (float32)
//
// Attack time 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt4M(value float32) error {
	return e.Element.SetProperty("at-4m", value)
}

// at-4s (float32)
//
// Attack time 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt4S(value float32) error {
	return e.Element.SetProperty("at-4s", value)
}

// at-5m (float32)
//
// Attack time 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt5M(value float32) error {
	return e.Element.SetProperty("at-5m", value)
}

// at-5s (float32)
//
// Attack time 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt5S(value float32) error {
	return e.Element.SetProperty("at-5s", value)
}

// at-6m (float32)
//
// Attack time 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt6M(value float32) error {
	return e.Element.SetProperty("at-6m", value)
}

// at-6s (float32)
//
// Attack time 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt6S(value float32) error {
	return e.Element.SetProperty("at-6s", value)
}

// at-7m (float32)
//
// Attack time 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt7M(value float32) error {
	return e.Element.SetProperty("at-7m", value)
}

// at-7s (float32)
//
// Attack time 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetAt7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetAt7S(value float32) error {
	return e.Element.SetProperty("at-7s", value)
}

// bm-0m (bool)
//
// Mute band 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm0M(value bool) error {
	return e.Element.SetProperty("bm-0m", value)
}

// bm-0s (bool)
//
// Mute band 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm0S(value bool) error {
	return e.Element.SetProperty("bm-0s", value)
}

// bm-1m (bool)
//
// Mute band 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm1M(value bool) error {
	return e.Element.SetProperty("bm-1m", value)
}

// bm-1s (bool)
//
// Mute band 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm1S(value bool) error {
	return e.Element.SetProperty("bm-1s", value)
}

// bm-2m (bool)
//
// Mute band 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm2M(value bool) error {
	return e.Element.SetProperty("bm-2m", value)
}

// bm-2s (bool)
//
// Mute band 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm2S(value bool) error {
	return e.Element.SetProperty("bm-2s", value)
}

// bm-3m (bool)
//
// Mute band 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm3M(value bool) error {
	return e.Element.SetProperty("bm-3m", value)
}

// bm-3s (bool)
//
// Mute band 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm3S(value bool) error {
	return e.Element.SetProperty("bm-3s", value)
}

// bm-4m (bool)
//
// Mute band 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm4M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm4M(value bool) error {
	return e.Element.SetProperty("bm-4m", value)
}

// bm-4s (bool)
//
// Mute band 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm4S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm4S(value bool) error {
	return e.Element.SetProperty("bm-4s", value)
}

// bm-5m (bool)
//
// Mute band 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm5M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm5M(value bool) error {
	return e.Element.SetProperty("bm-5m", value)
}

// bm-5s (bool)
//
// Mute band 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm5S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm5S(value bool) error {
	return e.Element.SetProperty("bm-5s", value)
}

// bm-6m (bool)
//
// Mute band 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm6M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm6M(value bool) error {
	return e.Element.SetProperty("bm-6m", value)
}

// bm-6s (bool)
//
// Mute band 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm6S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm6S(value bool) error {
	return e.Element.SetProperty("bm-6s", value)
}

// bm-7m (bool)
//
// Mute band 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm7M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm7M(value bool) error {
	return e.Element.SetProperty("bm-7m", value)
}

// bm-7s (bool)
//
// Mute band 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBm7S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBm7S(value bool) error {
	return e.Element.SetProperty("bm-7s", value)
}

// bs-0m (bool)
//
// Solo band 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs0M(value bool) error {
	return e.Element.SetProperty("bs-0m", value)
}

// bs-0s (bool)
//
// Solo band 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs0S(value bool) error {
	return e.Element.SetProperty("bs-0s", value)
}

// bs-1m (bool)
//
// Solo band 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs1M(value bool) error {
	return e.Element.SetProperty("bs-1m", value)
}

// bs-1s (bool)
//
// Solo band 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs1S(value bool) error {
	return e.Element.SetProperty("bs-1s", value)
}

// bs-2m (bool)
//
// Solo band 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs2M(value bool) error {
	return e.Element.SetProperty("bs-2m", value)
}

// bs-2s (bool)
//
// Solo band 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs2S(value bool) error {
	return e.Element.SetProperty("bs-2s", value)
}

// bs-3m (bool)
//
// Solo band 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs3M(value bool) error {
	return e.Element.SetProperty("bs-3m", value)
}

// bs-3s (bool)
//
// Solo band 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs3S(value bool) error {
	return e.Element.SetProperty("bs-3s", value)
}

// bs-4m (bool)
//
// Solo band 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs4M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs4M(value bool) error {
	return e.Element.SetProperty("bs-4m", value)
}

// bs-4s (bool)
//
// Solo band 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs4S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs4S(value bool) error {
	return e.Element.SetProperty("bs-4s", value)
}

// bs-5m (bool)
//
// Solo band 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs5M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs5M(value bool) error {
	return e.Element.SetProperty("bs-5m", value)
}

// bs-5s (bool)
//
// Solo band 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs5S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs5S(value bool) error {
	return e.Element.SetProperty("bs-5s", value)
}

// bs-6m (bool)
//
// Solo band 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs6M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs6M(value bool) error {
	return e.Element.SetProperty("bs-6m", value)
}

// bs-6s (bool)
//
// Solo band 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs6S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs6S(value bool) error {
	return e.Element.SetProperty("bs-6s", value)
}

// bs-7m (bool)
//
// Solo band 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs7M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs7M(value bool) error {
	return e.Element.SetProperty("bs-7m", value)
}

// bs-7s (bool)
//
// Solo band 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBs7S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBs7S(value bool) error {
	return e.Element.SetProperty("bs-7s", value)
}

// bsel (GstLspPlugInPluginsLv2MbCompressorMsbsel)
//
// Band selection

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBsel() (interface{}, error) {
	return e.Element.GetProperty("bsel")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBsel(value interface{}) error {
	return e.Element.SetProperty("bsel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MbCompressorMs) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cbe-1m (bool)
//
// Compression band enable 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe1M(value bool) error {
	return e.Element.SetProperty("cbe-1m", value)
}

// cbe-1s (bool)
//
// Compression band enable 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe1S(value bool) error {
	return e.Element.SetProperty("cbe-1s", value)
}

// cbe-2m (bool)
//
// Compression band enable 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe2M(value bool) error {
	return e.Element.SetProperty("cbe-2m", value)
}

// cbe-2s (bool)
//
// Compression band enable 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe2S(value bool) error {
	return e.Element.SetProperty("cbe-2s", value)
}

// cbe-3m (bool)
//
// Compression band enable 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe3M(value bool) error {
	return e.Element.SetProperty("cbe-3m", value)
}

// cbe-3s (bool)
//
// Compression band enable 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe3S(value bool) error {
	return e.Element.SetProperty("cbe-3s", value)
}

// cbe-4m (bool)
//
// Compression band enable 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe4M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe4M(value bool) error {
	return e.Element.SetProperty("cbe-4m", value)
}

// cbe-4s (bool)
//
// Compression band enable 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe4S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe4S(value bool) error {
	return e.Element.SetProperty("cbe-4s", value)
}

// cbe-5m (bool)
//
// Compression band enable 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe5M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe5M(value bool) error {
	return e.Element.SetProperty("cbe-5m", value)
}

// cbe-5s (bool)
//
// Compression band enable 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe5S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe5S(value bool) error {
	return e.Element.SetProperty("cbe-5s", value)
}

// cbe-6m (bool)
//
// Compression band enable 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe6M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe6M(value bool) error {
	return e.Element.SetProperty("cbe-6m", value)
}

// cbe-6s (bool)
//
// Compression band enable 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe6S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe6S(value bool) error {
	return e.Element.SetProperty("cbe-6s", value)
}

// cbe-7m (bool)
//
// Compression band enable 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe7M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe7M(value bool) error {
	return e.Element.SetProperty("cbe-7m", value)
}

// cbe-7s (bool)
//
// Compression band enable 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCbe7S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCbe7S(value bool) error {
	return e.Element.SetProperty("cbe-7s", value)
}

// ce-0m (bool)
//
// Compressor enable 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe0M(value bool) error {
	return e.Element.SetProperty("ce-0m", value)
}

// ce-0s (bool)
//
// Compressor enable 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe0S(value bool) error {
	return e.Element.SetProperty("ce-0s", value)
}

// ce-1m (bool)
//
// Compressor enable 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe1M(value bool) error {
	return e.Element.SetProperty("ce-1m", value)
}

// ce-1s (bool)
//
// Compressor enable 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe1S(value bool) error {
	return e.Element.SetProperty("ce-1s", value)
}

// ce-2m (bool)
//
// Compressor enable 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe2M(value bool) error {
	return e.Element.SetProperty("ce-2m", value)
}

// ce-2s (bool)
//
// Compressor enable 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe2S(value bool) error {
	return e.Element.SetProperty("ce-2s", value)
}

// ce-3m (bool)
//
// Compressor enable 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe3M(value bool) error {
	return e.Element.SetProperty("ce-3m", value)
}

// ce-3s (bool)
//
// Compressor enable 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe3S(value bool) error {
	return e.Element.SetProperty("ce-3s", value)
}

// ce-4m (bool)
//
// Compressor enable 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe4M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe4M(value bool) error {
	return e.Element.SetProperty("ce-4m", value)
}

// ce-4s (bool)
//
// Compressor enable 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe4S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe4S(value bool) error {
	return e.Element.SetProperty("ce-4s", value)
}

// ce-5m (bool)
//
// Compressor enable 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe5M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe5M(value bool) error {
	return e.Element.SetProperty("ce-5m", value)
}

// ce-5s (bool)
//
// Compressor enable 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe5S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe5S(value bool) error {
	return e.Element.SetProperty("ce-5s", value)
}

// ce-6m (bool)
//
// Compressor enable 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe6M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe6M(value bool) error {
	return e.Element.SetProperty("ce-6m", value)
}

// ce-6s (bool)
//
// Compressor enable 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe6S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe6S(value bool) error {
	return e.Element.SetProperty("ce-6s", value)
}

// ce-7m (bool)
//
// Compressor enable 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe7M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe7M(value bool) error {
	return e.Element.SetProperty("ce-7m", value)
}

// ce-7s (bool)
//
// Compressor enable 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCe7S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCe7S(value bool) error {
	return e.Element.SetProperty("ce-7s", value)
}

// clm-0m (float32)
//
// Curve level meter 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-0s (float32)
//
// Curve level meter 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-1m (float32)
//
// Curve level meter 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-1s (float32)
//
// Curve level meter 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-2m (float32)
//
// Curve level meter 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-2s (float32)
//
// Curve level meter 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-3m (float32)
//
// Curve level meter 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-3s (float32)
//
// Curve level meter 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-4m (float32)
//
// Curve level meter 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-4s (float32)
//
// Curve level meter 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-5m (float32)
//
// Curve level meter 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-5s (float32)
//
// Curve level meter 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-6m (float32)
//
// Curve level meter 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-6s (float32)
//
// Curve level meter 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-7m (float32)
//
// Curve level meter 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-7s (float32)
//
// Curve level meter 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetClm7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cm-0m (GstLspPlugInPluginsLv2MbCompressorMscm0M)
//
// Compression mode 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm0M() (interface{}, error) {
	return e.Element.GetProperty("cm-0m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm0M(value interface{}) error {
	return e.Element.SetProperty("cm-0m", value)
}

// cm-0s (GstLspPlugInPluginsLv2MbCompressorMscm0S)
//
// Compression mode 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm0S() (interface{}, error) {
	return e.Element.GetProperty("cm-0s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm0S(value interface{}) error {
	return e.Element.SetProperty("cm-0s", value)
}

// cm-1m (GstLspPlugInPluginsLv2MbCompressorMscm1M)
//
// Compression mode 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm1M() (interface{}, error) {
	return e.Element.GetProperty("cm-1m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm1M(value interface{}) error {
	return e.Element.SetProperty("cm-1m", value)
}

// cm-1s (GstLspPlugInPluginsLv2MbCompressorMscm1S)
//
// Compression mode 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm1S() (interface{}, error) {
	return e.Element.GetProperty("cm-1s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm1S(value interface{}) error {
	return e.Element.SetProperty("cm-1s", value)
}

// cm-2m (GstLspPlugInPluginsLv2MbCompressorMscm2M)
//
// Compression mode 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm2M() (interface{}, error) {
	return e.Element.GetProperty("cm-2m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm2M(value interface{}) error {
	return e.Element.SetProperty("cm-2m", value)
}

// cm-2s (GstLspPlugInPluginsLv2MbCompressorMscm2S)
//
// Compression mode 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm2S() (interface{}, error) {
	return e.Element.GetProperty("cm-2s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm2S(value interface{}) error {
	return e.Element.SetProperty("cm-2s", value)
}

// cm-3m (GstLspPlugInPluginsLv2MbCompressorMscm3M)
//
// Compression mode 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm3M() (interface{}, error) {
	return e.Element.GetProperty("cm-3m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm3M(value interface{}) error {
	return e.Element.SetProperty("cm-3m", value)
}

// cm-3s (GstLspPlugInPluginsLv2MbCompressorMscm3S)
//
// Compression mode 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm3S() (interface{}, error) {
	return e.Element.GetProperty("cm-3s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm3S(value interface{}) error {
	return e.Element.SetProperty("cm-3s", value)
}

// cm-4m (GstLspPlugInPluginsLv2MbCompressorMscm4M)
//
// Compression mode 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm4M() (interface{}, error) {
	return e.Element.GetProperty("cm-4m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm4M(value interface{}) error {
	return e.Element.SetProperty("cm-4m", value)
}

// cm-4s (GstLspPlugInPluginsLv2MbCompressorMscm4S)
//
// Compression mode 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm4S() (interface{}, error) {
	return e.Element.GetProperty("cm-4s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm4S(value interface{}) error {
	return e.Element.SetProperty("cm-4s", value)
}

// cm-5m (GstLspPlugInPluginsLv2MbCompressorMscm5M)
//
// Compression mode 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm5M() (interface{}, error) {
	return e.Element.GetProperty("cm-5m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm5M(value interface{}) error {
	return e.Element.SetProperty("cm-5m", value)
}

// cm-5s (GstLspPlugInPluginsLv2MbCompressorMscm5S)
//
// Compression mode 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm5S() (interface{}, error) {
	return e.Element.GetProperty("cm-5s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm5S(value interface{}) error {
	return e.Element.SetProperty("cm-5s", value)
}

// cm-6m (GstLspPlugInPluginsLv2MbCompressorMscm6M)
//
// Compression mode 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm6M() (interface{}, error) {
	return e.Element.GetProperty("cm-6m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm6M(value interface{}) error {
	return e.Element.SetProperty("cm-6m", value)
}

// cm-6s (GstLspPlugInPluginsLv2MbCompressorMscm6S)
//
// Compression mode 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm6S() (interface{}, error) {
	return e.Element.GetProperty("cm-6s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm6S(value interface{}) error {
	return e.Element.SetProperty("cm-6s", value)
}

// cm-7m (GstLspPlugInPluginsLv2MbCompressorMscm7M)
//
// Compression mode 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm7M() (interface{}, error) {
	return e.Element.GetProperty("cm-7m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm7M(value interface{}) error {
	return e.Element.SetProperty("cm-7m", value)
}

// cm-7s (GstLspPlugInPluginsLv2MbCompressorMscm7S)
//
// Compression mode 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCm7S() (interface{}, error) {
	return e.Element.GetProperty("cm-7s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCm7S(value interface{}) error {
	return e.Element.SetProperty("cm-7s", value)
}

// cr-0m (float32)
//
// Ratio 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr0M(value float32) error {
	return e.Element.SetProperty("cr-0m", value)
}

// cr-0s (float32)
//
// Ratio 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr0S(value float32) error {
	return e.Element.SetProperty("cr-0s", value)
}

// cr-1m (float32)
//
// Ratio 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr1M(value float32) error {
	return e.Element.SetProperty("cr-1m", value)
}

// cr-1s (float32)
//
// Ratio 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr1S(value float32) error {
	return e.Element.SetProperty("cr-1s", value)
}

// cr-2m (float32)
//
// Ratio 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr2M(value float32) error {
	return e.Element.SetProperty("cr-2m", value)
}

// cr-2s (float32)
//
// Ratio 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr2S(value float32) error {
	return e.Element.SetProperty("cr-2s", value)
}

// cr-3m (float32)
//
// Ratio 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr3M(value float32) error {
	return e.Element.SetProperty("cr-3m", value)
}

// cr-3s (float32)
//
// Ratio 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr3S(value float32) error {
	return e.Element.SetProperty("cr-3s", value)
}

// cr-4m (float32)
//
// Ratio 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr4M(value float32) error {
	return e.Element.SetProperty("cr-4m", value)
}

// cr-4s (float32)
//
// Ratio 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr4S(value float32) error {
	return e.Element.SetProperty("cr-4s", value)
}

// cr-5m (float32)
//
// Ratio 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr5M(value float32) error {
	return e.Element.SetProperty("cr-5m", value)
}

// cr-5s (float32)
//
// Ratio 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr5S(value float32) error {
	return e.Element.SetProperty("cr-5s", value)
}

// cr-6m (float32)
//
// Ratio 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr6M(value float32) error {
	return e.Element.SetProperty("cr-6m", value)
}

// cr-6s (float32)
//
// Ratio 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr6S(value float32) error {
	return e.Element.SetProperty("cr-6s", value)
}

// cr-7m (float32)
//
// Ratio 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr7M(value float32) error {
	return e.Element.SetProperty("cr-7m", value)
}

// cr-7s (float32)
//
// Ratio 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetCr7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetCr7S(value float32) error {
	return e.Element.SetProperty("cr-7s", value)
}

// elm-0m (float32)
//
// Envelope level meter 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-0s (float32)
//
// Envelope level meter 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-1m (float32)
//
// Envelope level meter 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-1s (float32)
//
// Envelope level meter 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-2m (float32)
//
// Envelope level meter 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-2s (float32)
//
// Envelope level meter 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-3m (float32)
//
// Envelope level meter 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-3s (float32)
//
// Envelope level meter 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-4m (float32)
//
// Envelope level meter 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-4s (float32)
//
// Envelope level meter 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-5m (float32)
//
// Envelope level meter 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-5s (float32)
//
// Envelope level meter 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-6m (float32)
//
// Envelope level meter 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-6s (float32)
//
// Envelope level meter 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-7m (float32)
//
// Envelope level meter 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-7s (float32)
//
// Envelope level meter 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetElm7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// envb (GstLspPlugInPluginsLv2MbCompressorMsenvb)
//
// Envelope boost

func (e *LspPlugInPluginsLv2MbCompressorMs) GetEnvb() (interface{}, error) {
	return e.Element.GetProperty("envb")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetEnvb(value interface{}) error {
	return e.Element.SetProperty("envb", value)
}

// flt-m (bool)
//
// Band filter curves Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFltM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flt-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetFltM(value bool) error {
	return e.Element.SetProperty("flt-m", value)
}

// flt-s (bool)
//
// Band filter curves Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFltS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flt-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetFltS(value bool) error {
	return e.Element.SetProperty("flt-s", value)
}

// fre-0m (float32)
//
// Frequency range end 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-0s (float32)
//
// Frequency range end 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-1m (float32)
//
// Frequency range end 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-1s (float32)
//
// Frequency range end 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-2m (float32)
//
// Frequency range end 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-2s (float32)
//
// Frequency range end 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-3m (float32)
//
// Frequency range end 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-3s (float32)
//
// Frequency range end 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-4m (float32)
//
// Frequency range end 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-4s (float32)
//
// Frequency range end 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-5m (float32)
//
// Frequency range end 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-5s (float32)
//
// Frequency range end 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-6m (float32)
//
// Frequency range end 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-6s (float32)
//
// Frequency range end 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-7m (float32)
//
// Frequency range end 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-7s (float32)
//
// Frequency range end 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetFre7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// g-dry (float32)
//
// Dry gain

func (e *LspPlugInPluginsLv2MbCompressorMs) GetGDry() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-dry")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetGDry(value float32) error {
	return e.Element.SetProperty("g-dry", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2MbCompressorMs) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MbCompressorMs) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// g-wet (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2MbCompressorMs) GetGWet() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("g-wet")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetGWet(value float32) error {
	return e.Element.SetProperty("g-wet", value)
}

// hue-0m (float32)
//
// Hue  0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue0M(value float32) error {
	return e.Element.SetProperty("hue-0m", value)
}

// hue-0s (float32)
//
// Hue  0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue0S(value float32) error {
	return e.Element.SetProperty("hue-0s", value)
}

// hue-1m (float32)
//
// Hue  1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue1M(value float32) error {
	return e.Element.SetProperty("hue-1m", value)
}

// hue-1s (float32)
//
// Hue  1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue1S(value float32) error {
	return e.Element.SetProperty("hue-1s", value)
}

// hue-2m (float32)
//
// Hue  2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue2M(value float32) error {
	return e.Element.SetProperty("hue-2m", value)
}

// hue-2s (float32)
//
// Hue  2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue2S(value float32) error {
	return e.Element.SetProperty("hue-2s", value)
}

// hue-3m (float32)
//
// Hue  3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue3M(value float32) error {
	return e.Element.SetProperty("hue-3m", value)
}

// hue-3s (float32)
//
// Hue  3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue3S(value float32) error {
	return e.Element.SetProperty("hue-3s", value)
}

// hue-4m (float32)
//
// Hue  4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue4M(value float32) error {
	return e.Element.SetProperty("hue-4m", value)
}

// hue-4s (float32)
//
// Hue  4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue4S(value float32) error {
	return e.Element.SetProperty("hue-4s", value)
}

// hue-5m (float32)
//
// Hue  5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue5M(value float32) error {
	return e.Element.SetProperty("hue-5m", value)
}

// hue-5s (float32)
//
// Hue  5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue5S(value float32) error {
	return e.Element.SetProperty("hue-5s", value)
}

// hue-6m (float32)
//
// Hue  6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue6M(value float32) error {
	return e.Element.SetProperty("hue-6m", value)
}

// hue-6s (float32)
//
// Hue  6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue6S(value float32) error {
	return e.Element.SetProperty("hue-6s", value)
}

// hue-7m (float32)
//
// Hue  7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue7M(value float32) error {
	return e.Element.SetProperty("hue-7m", value)
}

// hue-7s (float32)
//
// Hue  7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetHue7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetHue7S(value float32) error {
	return e.Element.SetProperty("hue-7s", value)
}

// ife-m (bool)
//
// Input FFT graph enable Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetIfeM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ife-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetIfeM(value bool) error {
	return e.Element.SetProperty("ife-m", value)
}

// ife-s (bool)
//
// Input FFT graph enable Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetIfeS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ife-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetIfeS(value bool) error {
	return e.Element.SetProperty("ife-s", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2MbCompressorMs) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) GetIlmR() (float32, error) {
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

// kn-0m (float32)
//
// Knee 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn0M(value float32) error {
	return e.Element.SetProperty("kn-0m", value)
}

// kn-0s (float32)
//
// Knee 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn0S(value float32) error {
	return e.Element.SetProperty("kn-0s", value)
}

// kn-1m (float32)
//
// Knee 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn1M(value float32) error {
	return e.Element.SetProperty("kn-1m", value)
}

// kn-1s (float32)
//
// Knee 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn1S(value float32) error {
	return e.Element.SetProperty("kn-1s", value)
}

// kn-2m (float32)
//
// Knee 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn2M(value float32) error {
	return e.Element.SetProperty("kn-2m", value)
}

// kn-2s (float32)
//
// Knee 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn2S(value float32) error {
	return e.Element.SetProperty("kn-2s", value)
}

// kn-3m (float32)
//
// Knee 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn3M(value float32) error {
	return e.Element.SetProperty("kn-3m", value)
}

// kn-3s (float32)
//
// Knee 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn3S(value float32) error {
	return e.Element.SetProperty("kn-3s", value)
}

// kn-4m (float32)
//
// Knee 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn4M(value float32) error {
	return e.Element.SetProperty("kn-4m", value)
}

// kn-4s (float32)
//
// Knee 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn4S(value float32) error {
	return e.Element.SetProperty("kn-4s", value)
}

// kn-5m (float32)
//
// Knee 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn5M(value float32) error {
	return e.Element.SetProperty("kn-5m", value)
}

// kn-5s (float32)
//
// Knee 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn5S(value float32) error {
	return e.Element.SetProperty("kn-5s", value)
}

// kn-6m (float32)
//
// Knee 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn6M(value float32) error {
	return e.Element.SetProperty("kn-6m", value)
}

// kn-6s (float32)
//
// Knee 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn6S(value float32) error {
	return e.Element.SetProperty("kn-6s", value)
}

// kn-7m (float32)
//
// Knee 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn7M(value float32) error {
	return e.Element.SetProperty("kn-7m", value)
}

// kn-7s (float32)
//
// Knee 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetKn7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetKn7S(value float32) error {
	return e.Element.SetProperty("kn-7s", value)
}

// mk-0m (float32)
//
// Makeup gain 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk0M(value float32) error {
	return e.Element.SetProperty("mk-0m", value)
}

// mk-0s (float32)
//
// Makeup gain 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk0S(value float32) error {
	return e.Element.SetProperty("mk-0s", value)
}

// mk-1m (float32)
//
// Makeup gain 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk1M(value float32) error {
	return e.Element.SetProperty("mk-1m", value)
}

// mk-1s (float32)
//
// Makeup gain 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk1S(value float32) error {
	return e.Element.SetProperty("mk-1s", value)
}

// mk-2m (float32)
//
// Makeup gain 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk2M(value float32) error {
	return e.Element.SetProperty("mk-2m", value)
}

// mk-2s (float32)
//
// Makeup gain 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk2S(value float32) error {
	return e.Element.SetProperty("mk-2s", value)
}

// mk-3m (float32)
//
// Makeup gain 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk3M(value float32) error {
	return e.Element.SetProperty("mk-3m", value)
}

// mk-3s (float32)
//
// Makeup gain 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk3S(value float32) error {
	return e.Element.SetProperty("mk-3s", value)
}

// mk-4m (float32)
//
// Makeup gain 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk4M(value float32) error {
	return e.Element.SetProperty("mk-4m", value)
}

// mk-4s (float32)
//
// Makeup gain 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk4S(value float32) error {
	return e.Element.SetProperty("mk-4s", value)
}

// mk-5m (float32)
//
// Makeup gain 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk5M(value float32) error {
	return e.Element.SetProperty("mk-5m", value)
}

// mk-5s (float32)
//
// Makeup gain 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk5S(value float32) error {
	return e.Element.SetProperty("mk-5s", value)
}

// mk-6m (float32)
//
// Makeup gain 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk6M(value float32) error {
	return e.Element.SetProperty("mk-6m", value)
}

// mk-6s (float32)
//
// Makeup gain 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk6S(value float32) error {
	return e.Element.SetProperty("mk-6s", value)
}

// mk-7m (float32)
//
// Makeup gain 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk7M(value float32) error {
	return e.Element.SetProperty("mk-7m", value)
}

// mk-7s (float32)
//
// Makeup gain 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMk7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMk7S(value float32) error {
	return e.Element.SetProperty("mk-7s", value)
}

// mode (GstLspPlugInPluginsLv2MbCompressorMsmode)
//
// Compressor mode

func (e *LspPlugInPluginsLv2MbCompressorMs) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ofe-m (bool)
//
// Output FFT graph enable Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetOfeM() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofe-m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetOfeM(value bool) error {
	return e.Element.SetProperty("ofe-m", value)
}

// ofe-s (bool)
//
// Output FFT graph enable Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetOfeS() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofe-s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetOfeS(value bool) error {
	return e.Element.SetProperty("ofe-s", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2MbCompressorMs) GetOlmL() (float32, error) {
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
// Output level meter Right

func (e *LspPlugInPluginsLv2MbCompressorMs) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl-0m (float32)
//
// Release level 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-0s (float32)
//
// Release level 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-1m (float32)
//
// Release level 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-1s (float32)
//
// Release level 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-2m (float32)
//
// Release level 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-2s (float32)
//
// Release level 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-3m (float32)
//
// Release level 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-3s (float32)
//
// Release level 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-4m (float32)
//
// Release level 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-4s (float32)
//
// Release level 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-5m (float32)
//
// Release level 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-5s (float32)
//
// Release level 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-6m (float32)
//
// Release level 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-6s (float32)
//
// Release level 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-7m (float32)
//
// Release level 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-7s (float32)
//
// Release level 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRl7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-0m (float32)
//
// Reduction level meter 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-0s (float32)
//
// Reduction level meter 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-1m (float32)
//
// Reduction level meter 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-1s (float32)
//
// Reduction level meter 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-2m (float32)
//
// Reduction level meter 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-2s (float32)
//
// Reduction level meter 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-3m (float32)
//
// Reduction level meter 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-3s (float32)
//
// Reduction level meter 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-4m (float32)
//
// Reduction level meter 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-4s (float32)
//
// Reduction level meter 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-5m (float32)
//
// Reduction level meter 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-5s (float32)
//
// Reduction level meter 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-6m (float32)
//
// Reduction level meter 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-6s (float32)
//
// Reduction level meter 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-7m (float32)
//
// Reduction level meter 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-7s (float32)
//
// Reduction level meter 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRlm7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rrl-0m (float32)
//
// Relative release level 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl0M(value float32) error {
	return e.Element.SetProperty("rrl-0m", value)
}

// rrl-0s (float32)
//
// Relative release level 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl0S(value float32) error {
	return e.Element.SetProperty("rrl-0s", value)
}

// rrl-1m (float32)
//
// Relative release level 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl1M(value float32) error {
	return e.Element.SetProperty("rrl-1m", value)
}

// rrl-1s (float32)
//
// Relative release level 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl1S(value float32) error {
	return e.Element.SetProperty("rrl-1s", value)
}

// rrl-2m (float32)
//
// Relative release level 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl2M(value float32) error {
	return e.Element.SetProperty("rrl-2m", value)
}

// rrl-2s (float32)
//
// Relative release level 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl2S(value float32) error {
	return e.Element.SetProperty("rrl-2s", value)
}

// rrl-3m (float32)
//
// Relative release level 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl3M(value float32) error {
	return e.Element.SetProperty("rrl-3m", value)
}

// rrl-3s (float32)
//
// Relative release level 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl3S(value float32) error {
	return e.Element.SetProperty("rrl-3s", value)
}

// rrl-4m (float32)
//
// Relative release level 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl4M(value float32) error {
	return e.Element.SetProperty("rrl-4m", value)
}

// rrl-4s (float32)
//
// Relative release level 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl4S(value float32) error {
	return e.Element.SetProperty("rrl-4s", value)
}

// rrl-5m (float32)
//
// Relative release level 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl5M(value float32) error {
	return e.Element.SetProperty("rrl-5m", value)
}

// rrl-5s (float32)
//
// Relative release level 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl5S(value float32) error {
	return e.Element.SetProperty("rrl-5s", value)
}

// rrl-6m (float32)
//
// Relative release level 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl6M(value float32) error {
	return e.Element.SetProperty("rrl-6m", value)
}

// rrl-6s (float32)
//
// Relative release level 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl6S(value float32) error {
	return e.Element.SetProperty("rrl-6s", value)
}

// rrl-7m (float32)
//
// Relative release level 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl7M(value float32) error {
	return e.Element.SetProperty("rrl-7m", value)
}

// rrl-7s (float32)
//
// Relative release level 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRrl7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRrl7S(value float32) error {
	return e.Element.SetProperty("rrl-7s", value)
}

// rt-0m (float32)
//
// Release time 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt0M(value float32) error {
	return e.Element.SetProperty("rt-0m", value)
}

// rt-0s (float32)
//
// Release time 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt0S(value float32) error {
	return e.Element.SetProperty("rt-0s", value)
}

// rt-1m (float32)
//
// Release time 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt1M(value float32) error {
	return e.Element.SetProperty("rt-1m", value)
}

// rt-1s (float32)
//
// Release time 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt1S(value float32) error {
	return e.Element.SetProperty("rt-1s", value)
}

// rt-2m (float32)
//
// Release time 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt2M(value float32) error {
	return e.Element.SetProperty("rt-2m", value)
}

// rt-2s (float32)
//
// Release time 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt2S(value float32) error {
	return e.Element.SetProperty("rt-2s", value)
}

// rt-3m (float32)
//
// Release time 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt3M(value float32) error {
	return e.Element.SetProperty("rt-3m", value)
}

// rt-3s (float32)
//
// Release time 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt3S(value float32) error {
	return e.Element.SetProperty("rt-3s", value)
}

// rt-4m (float32)
//
// Release time 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt4M(value float32) error {
	return e.Element.SetProperty("rt-4m", value)
}

// rt-4s (float32)
//
// Release time 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt4S(value float32) error {
	return e.Element.SetProperty("rt-4s", value)
}

// rt-5m (float32)
//
// Release time 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt5M(value float32) error {
	return e.Element.SetProperty("rt-5m", value)
}

// rt-5s (float32)
//
// Release time 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt5S(value float32) error {
	return e.Element.SetProperty("rt-5s", value)
}

// rt-6m (float32)
//
// Release time 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt6M(value float32) error {
	return e.Element.SetProperty("rt-6m", value)
}

// rt-6s (float32)
//
// Release time 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt6S(value float32) error {
	return e.Element.SetProperty("rt-6s", value)
}

// rt-7m (float32)
//
// Release time 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt7M(value float32) error {
	return e.Element.SetProperty("rt-7m", value)
}

// rt-7s (float32)
//
// Release time 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetRt7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetRt7S(value float32) error {
	return e.Element.SetProperty("rt-7s", value)
}

// schc-0m (bool)
//
// Sidechain custom hi-cut 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc0M(value bool) error {
	return e.Element.SetProperty("schc-0m", value)
}

// schc-0s (bool)
//
// Sidechain custom hi-cut 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc0S(value bool) error {
	return e.Element.SetProperty("schc-0s", value)
}

// schc-1m (bool)
//
// Sidechain custom hi-cut 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc1M(value bool) error {
	return e.Element.SetProperty("schc-1m", value)
}

// schc-1s (bool)
//
// Sidechain custom hi-cut 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc1S(value bool) error {
	return e.Element.SetProperty("schc-1s", value)
}

// schc-2m (bool)
//
// Sidechain custom hi-cut 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc2M(value bool) error {
	return e.Element.SetProperty("schc-2m", value)
}

// schc-2s (bool)
//
// Sidechain custom hi-cut 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc2S(value bool) error {
	return e.Element.SetProperty("schc-2s", value)
}

// schc-3m (bool)
//
// Sidechain custom hi-cut 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc3M(value bool) error {
	return e.Element.SetProperty("schc-3m", value)
}

// schc-3s (bool)
//
// Sidechain custom hi-cut 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc3S(value bool) error {
	return e.Element.SetProperty("schc-3s", value)
}

// schc-4m (bool)
//
// Sidechain custom hi-cut 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc4M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc4M(value bool) error {
	return e.Element.SetProperty("schc-4m", value)
}

// schc-4s (bool)
//
// Sidechain custom hi-cut 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc4S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc4S(value bool) error {
	return e.Element.SetProperty("schc-4s", value)
}

// schc-5m (bool)
//
// Sidechain custom hi-cut 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc5M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc5M(value bool) error {
	return e.Element.SetProperty("schc-5m", value)
}

// schc-5s (bool)
//
// Sidechain custom hi-cut 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc5S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc5S(value bool) error {
	return e.Element.SetProperty("schc-5s", value)
}

// schc-6m (bool)
//
// Sidechain custom hi-cut 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc6M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc6M(value bool) error {
	return e.Element.SetProperty("schc-6m", value)
}

// schc-6s (bool)
//
// Sidechain custom hi-cut 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc6S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc6S(value bool) error {
	return e.Element.SetProperty("schc-6s", value)
}

// schc-7m (bool)
//
// Sidechain custom hi-cut 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc7M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc7M(value bool) error {
	return e.Element.SetProperty("schc-7m", value)
}

// schc-7s (bool)
//
// Sidechain custom hi-cut 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchc7S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchc7S(value bool) error {
	return e.Element.SetProperty("schc-7s", value)
}

// schf-0m (float32)
//
// Sidechain hi-cut frequency 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf0M(value float32) error {
	return e.Element.SetProperty("schf-0m", value)
}

// schf-0s (float32)
//
// Sidechain hi-cut frequency 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf0S(value float32) error {
	return e.Element.SetProperty("schf-0s", value)
}

// schf-1m (float32)
//
// Sidechain hi-cut frequency 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf1M(value float32) error {
	return e.Element.SetProperty("schf-1m", value)
}

// schf-1s (float32)
//
// Sidechain hi-cut frequency 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf1S(value float32) error {
	return e.Element.SetProperty("schf-1s", value)
}

// schf-2m (float32)
//
// Sidechain hi-cut frequency 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf2M(value float32) error {
	return e.Element.SetProperty("schf-2m", value)
}

// schf-2s (float32)
//
// Sidechain hi-cut frequency 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf2S(value float32) error {
	return e.Element.SetProperty("schf-2s", value)
}

// schf-3m (float32)
//
// Sidechain hi-cut frequency 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf3M(value float32) error {
	return e.Element.SetProperty("schf-3m", value)
}

// schf-3s (float32)
//
// Sidechain hi-cut frequency 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf3S(value float32) error {
	return e.Element.SetProperty("schf-3s", value)
}

// schf-4m (float32)
//
// Sidechain hi-cut frequency 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf4M(value float32) error {
	return e.Element.SetProperty("schf-4m", value)
}

// schf-4s (float32)
//
// Sidechain hi-cut frequency 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf4S(value float32) error {
	return e.Element.SetProperty("schf-4s", value)
}

// schf-5m (float32)
//
// Sidechain hi-cut frequency 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf5M(value float32) error {
	return e.Element.SetProperty("schf-5m", value)
}

// schf-5s (float32)
//
// Sidechain hi-cut frequency 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf5S(value float32) error {
	return e.Element.SetProperty("schf-5s", value)
}

// schf-6m (float32)
//
// Sidechain hi-cut frequency 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf6M(value float32) error {
	return e.Element.SetProperty("schf-6m", value)
}

// schf-6s (float32)
//
// Sidechain hi-cut frequency 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf6S(value float32) error {
	return e.Element.SetProperty("schf-6s", value)
}

// schf-7m (float32)
//
// Sidechain hi-cut frequency 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf7M(value float32) error {
	return e.Element.SetProperty("schf-7m", value)
}

// schf-7s (float32)
//
// Sidechain hi-cut frequency 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSchf7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSchf7S(value float32) error {
	return e.Element.SetProperty("schf-7s", value)
}

// sclc-0m (bool)
//
// Sidechain custom lo-cut 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc0M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc0M(value bool) error {
	return e.Element.SetProperty("sclc-0m", value)
}

// sclc-0s (bool)
//
// Sidechain custom lo-cut 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc0S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc0S(value bool) error {
	return e.Element.SetProperty("sclc-0s", value)
}

// sclc-1m (bool)
//
// Sidechain custom lo-cut 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc1M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc1M(value bool) error {
	return e.Element.SetProperty("sclc-1m", value)
}

// sclc-1s (bool)
//
// Sidechain custom lo-cut 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc1S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc1S(value bool) error {
	return e.Element.SetProperty("sclc-1s", value)
}

// sclc-2m (bool)
//
// Sidechain custom lo-cut 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc2M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc2M(value bool) error {
	return e.Element.SetProperty("sclc-2m", value)
}

// sclc-2s (bool)
//
// Sidechain custom lo-cut 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc2S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc2S(value bool) error {
	return e.Element.SetProperty("sclc-2s", value)
}

// sclc-3m (bool)
//
// Sidechain custom lo-cut 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc3M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc3M(value bool) error {
	return e.Element.SetProperty("sclc-3m", value)
}

// sclc-3s (bool)
//
// Sidechain custom lo-cut 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc3S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc3S(value bool) error {
	return e.Element.SetProperty("sclc-3s", value)
}

// sclc-4m (bool)
//
// Sidechain custom lo-cut 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc4M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc4M(value bool) error {
	return e.Element.SetProperty("sclc-4m", value)
}

// sclc-4s (bool)
//
// Sidechain custom lo-cut 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc4S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc4S(value bool) error {
	return e.Element.SetProperty("sclc-4s", value)
}

// sclc-5m (bool)
//
// Sidechain custom lo-cut 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc5M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc5M(value bool) error {
	return e.Element.SetProperty("sclc-5m", value)
}

// sclc-5s (bool)
//
// Sidechain custom lo-cut 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc5S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc5S(value bool) error {
	return e.Element.SetProperty("sclc-5s", value)
}

// sclc-6m (bool)
//
// Sidechain custom lo-cut 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc6M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc6M(value bool) error {
	return e.Element.SetProperty("sclc-6m", value)
}

// sclc-6s (bool)
//
// Sidechain custom lo-cut 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc6S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc6S(value bool) error {
	return e.Element.SetProperty("sclc-6s", value)
}

// sclc-7m (bool)
//
// Sidechain custom lo-cut 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc7M() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc7M(value bool) error {
	return e.Element.SetProperty("sclc-7m", value)
}

// sclc-7s (bool)
//
// Sidechain custom lo-cut 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclc7S() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclc7S(value bool) error {
	return e.Element.SetProperty("sclc-7s", value)
}

// sclf-0m (float32)
//
// Sidechain lo-cut frequency 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf0M(value float32) error {
	return e.Element.SetProperty("sclf-0m", value)
}

// sclf-0s (float32)
//
// Sidechain lo-cut frequency 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf0S(value float32) error {
	return e.Element.SetProperty("sclf-0s", value)
}

// sclf-1m (float32)
//
// Sidechain lo-cut frequency 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf1M(value float32) error {
	return e.Element.SetProperty("sclf-1m", value)
}

// sclf-1s (float32)
//
// Sidechain lo-cut frequency 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf1S(value float32) error {
	return e.Element.SetProperty("sclf-1s", value)
}

// sclf-2m (float32)
//
// Sidechain lo-cut frequency 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf2M(value float32) error {
	return e.Element.SetProperty("sclf-2m", value)
}

// sclf-2s (float32)
//
// Sidechain lo-cut frequency 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf2S(value float32) error {
	return e.Element.SetProperty("sclf-2s", value)
}

// sclf-3m (float32)
//
// Sidechain lo-cut frequency 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf3M(value float32) error {
	return e.Element.SetProperty("sclf-3m", value)
}

// sclf-3s (float32)
//
// Sidechain lo-cut frequency 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf3S(value float32) error {
	return e.Element.SetProperty("sclf-3s", value)
}

// sclf-4m (float32)
//
// Sidechain lo-cut frequency 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf4M(value float32) error {
	return e.Element.SetProperty("sclf-4m", value)
}

// sclf-4s (float32)
//
// Sidechain lo-cut frequency 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf4S(value float32) error {
	return e.Element.SetProperty("sclf-4s", value)
}

// sclf-5m (float32)
//
// Sidechain lo-cut frequency 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf5M(value float32) error {
	return e.Element.SetProperty("sclf-5m", value)
}

// sclf-5s (float32)
//
// Sidechain lo-cut frequency 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf5S(value float32) error {
	return e.Element.SetProperty("sclf-5s", value)
}

// sclf-6m (float32)
//
// Sidechain lo-cut frequency 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf6M(value float32) error {
	return e.Element.SetProperty("sclf-6m", value)
}

// sclf-6s (float32)
//
// Sidechain lo-cut frequency 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf6S(value float32) error {
	return e.Element.SetProperty("sclf-6s", value)
}

// sclf-7m (float32)
//
// Sidechain lo-cut frequency 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf7M(value float32) error {
	return e.Element.SetProperty("sclf-7m", value)
}

// sclf-7s (float32)
//
// Sidechain lo-cut frequency 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSclf7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSclf7S(value float32) error {
	return e.Element.SetProperty("sclf-7s", value)
}

// scm-0m (GstLspPlugInPluginsLv2MbCompressorMsscm0M)
//
// Sidechain mode 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm0M() (interface{}, error) {
	return e.Element.GetProperty("scm-0m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm0M(value interface{}) error {
	return e.Element.SetProperty("scm-0m", value)
}

// scm-0s (GstLspPlugInPluginsLv2MbCompressorMsscm0S)
//
// Sidechain mode 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm0S() (interface{}, error) {
	return e.Element.GetProperty("scm-0s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm0S(value interface{}) error {
	return e.Element.SetProperty("scm-0s", value)
}

// scm-1m (GstLspPlugInPluginsLv2MbCompressorMsscm1M)
//
// Sidechain mode 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm1M() (interface{}, error) {
	return e.Element.GetProperty("scm-1m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm1M(value interface{}) error {
	return e.Element.SetProperty("scm-1m", value)
}

// scm-1s (GstLspPlugInPluginsLv2MbCompressorMsscm1S)
//
// Sidechain mode 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm1S() (interface{}, error) {
	return e.Element.GetProperty("scm-1s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm1S(value interface{}) error {
	return e.Element.SetProperty("scm-1s", value)
}

// scm-2m (GstLspPlugInPluginsLv2MbCompressorMsscm2M)
//
// Sidechain mode 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm2M() (interface{}, error) {
	return e.Element.GetProperty("scm-2m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm2M(value interface{}) error {
	return e.Element.SetProperty("scm-2m", value)
}

// scm-2s (GstLspPlugInPluginsLv2MbCompressorMsscm2S)
//
// Sidechain mode 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm2S() (interface{}, error) {
	return e.Element.GetProperty("scm-2s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm2S(value interface{}) error {
	return e.Element.SetProperty("scm-2s", value)
}

// scm-3m (GstLspPlugInPluginsLv2MbCompressorMsscm3M)
//
// Sidechain mode 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm3M() (interface{}, error) {
	return e.Element.GetProperty("scm-3m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm3M(value interface{}) error {
	return e.Element.SetProperty("scm-3m", value)
}

// scm-3s (GstLspPlugInPluginsLv2MbCompressorMsscm3S)
//
// Sidechain mode 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm3S() (interface{}, error) {
	return e.Element.GetProperty("scm-3s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm3S(value interface{}) error {
	return e.Element.SetProperty("scm-3s", value)
}

// scm-4m (GstLspPlugInPluginsLv2MbCompressorMsscm4M)
//
// Sidechain mode 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm4M() (interface{}, error) {
	return e.Element.GetProperty("scm-4m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm4M(value interface{}) error {
	return e.Element.SetProperty("scm-4m", value)
}

// scm-4s (GstLspPlugInPluginsLv2MbCompressorMsscm4S)
//
// Sidechain mode 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm4S() (interface{}, error) {
	return e.Element.GetProperty("scm-4s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm4S(value interface{}) error {
	return e.Element.SetProperty("scm-4s", value)
}

// scm-5m (GstLspPlugInPluginsLv2MbCompressorMsscm5M)
//
// Sidechain mode 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm5M() (interface{}, error) {
	return e.Element.GetProperty("scm-5m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm5M(value interface{}) error {
	return e.Element.SetProperty("scm-5m", value)
}

// scm-5s (GstLspPlugInPluginsLv2MbCompressorMsscm5S)
//
// Sidechain mode 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm5S() (interface{}, error) {
	return e.Element.GetProperty("scm-5s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm5S(value interface{}) error {
	return e.Element.SetProperty("scm-5s", value)
}

// scm-6m (GstLspPlugInPluginsLv2MbCompressorMsscm6M)
//
// Sidechain mode 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm6M() (interface{}, error) {
	return e.Element.GetProperty("scm-6m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm6M(value interface{}) error {
	return e.Element.SetProperty("scm-6m", value)
}

// scm-6s (GstLspPlugInPluginsLv2MbCompressorMsscm6S)
//
// Sidechain mode 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm6S() (interface{}, error) {
	return e.Element.GetProperty("scm-6s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm6S(value interface{}) error {
	return e.Element.SetProperty("scm-6s", value)
}

// scm-7m (GstLspPlugInPluginsLv2MbCompressorMsscm7M)
//
// Sidechain mode 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm7M() (interface{}, error) {
	return e.Element.GetProperty("scm-7m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm7M(value interface{}) error {
	return e.Element.SetProperty("scm-7m", value)
}

// scm-7s (GstLspPlugInPluginsLv2MbCompressorMsscm7S)
//
// Sidechain mode 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScm7S() (interface{}, error) {
	return e.Element.GetProperty("scm-7s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScm7S(value interface{}) error {
	return e.Element.SetProperty("scm-7s", value)
}

// scp-0m (float32)
//
// Sidechain preamp 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp0M(value float32) error {
	return e.Element.SetProperty("scp-0m", value)
}

// scp-0s (float32)
//
// Sidechain preamp 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp0S(value float32) error {
	return e.Element.SetProperty("scp-0s", value)
}

// scp-1m (float32)
//
// Sidechain preamp 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp1M(value float32) error {
	return e.Element.SetProperty("scp-1m", value)
}

// scp-1s (float32)
//
// Sidechain preamp 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp1S(value float32) error {
	return e.Element.SetProperty("scp-1s", value)
}

// scp-2m (float32)
//
// Sidechain preamp 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp2M(value float32) error {
	return e.Element.SetProperty("scp-2m", value)
}

// scp-2s (float32)
//
// Sidechain preamp 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp2S(value float32) error {
	return e.Element.SetProperty("scp-2s", value)
}

// scp-3m (float32)
//
// Sidechain preamp 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp3M(value float32) error {
	return e.Element.SetProperty("scp-3m", value)
}

// scp-3s (float32)
//
// Sidechain preamp 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp3S(value float32) error {
	return e.Element.SetProperty("scp-3s", value)
}

// scp-4m (float32)
//
// Sidechain preamp 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp4M(value float32) error {
	return e.Element.SetProperty("scp-4m", value)
}

// scp-4s (float32)
//
// Sidechain preamp 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp4S(value float32) error {
	return e.Element.SetProperty("scp-4s", value)
}

// scp-5m (float32)
//
// Sidechain preamp 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp5M(value float32) error {
	return e.Element.SetProperty("scp-5m", value)
}

// scp-5s (float32)
//
// Sidechain preamp 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp5S(value float32) error {
	return e.Element.SetProperty("scp-5s", value)
}

// scp-6m (float32)
//
// Sidechain preamp 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp6M(value float32) error {
	return e.Element.SetProperty("scp-6m", value)
}

// scp-6s (float32)
//
// Sidechain preamp 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp6S(value float32) error {
	return e.Element.SetProperty("scp-6s", value)
}

// scp-7m (float32)
//
// Sidechain preamp 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp7M(value float32) error {
	return e.Element.SetProperty("scp-7m", value)
}

// scp-7s (float32)
//
// Sidechain preamp 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScp7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScp7S(value float32) error {
	return e.Element.SetProperty("scp-7s", value)
}

// scr-0m (float32)
//
// Sidechain reactivity 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr0M(value float32) error {
	return e.Element.SetProperty("scr-0m", value)
}

// scr-0s (float32)
//
// Sidechain reactivity 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr0S(value float32) error {
	return e.Element.SetProperty("scr-0s", value)
}

// scr-1m (float32)
//
// Sidechain reactivity 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr1M(value float32) error {
	return e.Element.SetProperty("scr-1m", value)
}

// scr-1s (float32)
//
// Sidechain reactivity 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr1S(value float32) error {
	return e.Element.SetProperty("scr-1s", value)
}

// scr-2m (float32)
//
// Sidechain reactivity 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr2M(value float32) error {
	return e.Element.SetProperty("scr-2m", value)
}

// scr-2s (float32)
//
// Sidechain reactivity 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr2S(value float32) error {
	return e.Element.SetProperty("scr-2s", value)
}

// scr-3m (float32)
//
// Sidechain reactivity 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr3M(value float32) error {
	return e.Element.SetProperty("scr-3m", value)
}

// scr-3s (float32)
//
// Sidechain reactivity 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr3S(value float32) error {
	return e.Element.SetProperty("scr-3s", value)
}

// scr-4m (float32)
//
// Sidechain reactivity 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr4M(value float32) error {
	return e.Element.SetProperty("scr-4m", value)
}

// scr-4s (float32)
//
// Sidechain reactivity 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr4S(value float32) error {
	return e.Element.SetProperty("scr-4s", value)
}

// scr-5m (float32)
//
// Sidechain reactivity 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr5M(value float32) error {
	return e.Element.SetProperty("scr-5m", value)
}

// scr-5s (float32)
//
// Sidechain reactivity 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr5S(value float32) error {
	return e.Element.SetProperty("scr-5s", value)
}

// scr-6m (float32)
//
// Sidechain reactivity 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr6M(value float32) error {
	return e.Element.SetProperty("scr-6m", value)
}

// scr-6s (float32)
//
// Sidechain reactivity 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr6S(value float32) error {
	return e.Element.SetProperty("scr-6s", value)
}

// scr-7m (float32)
//
// Sidechain reactivity 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr7M(value float32) error {
	return e.Element.SetProperty("scr-7m", value)
}

// scr-7s (float32)
//
// Sidechain reactivity 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScr7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScr7S(value float32) error {
	return e.Element.SetProperty("scr-7s", value)
}

// scs-0m (GstLspPlugInPluginsLv2MbCompressorMsscs0M)
//
// Sidechain source 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs0M() (interface{}, error) {
	return e.Element.GetProperty("scs-0m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs0M(value interface{}) error {
	return e.Element.SetProperty("scs-0m", value)
}

// scs-0s (GstLspPlugInPluginsLv2MbCompressorMsscs0S)
//
// Sidechain source 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs0S() (interface{}, error) {
	return e.Element.GetProperty("scs-0s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs0S(value interface{}) error {
	return e.Element.SetProperty("scs-0s", value)
}

// scs-1m (GstLspPlugInPluginsLv2MbCompressorMsscs1M)
//
// Sidechain source 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs1M() (interface{}, error) {
	return e.Element.GetProperty("scs-1m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs1M(value interface{}) error {
	return e.Element.SetProperty("scs-1m", value)
}

// scs-1s (GstLspPlugInPluginsLv2MbCompressorMsscs1S)
//
// Sidechain source 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs1S() (interface{}, error) {
	return e.Element.GetProperty("scs-1s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs1S(value interface{}) error {
	return e.Element.SetProperty("scs-1s", value)
}

// scs-2m (GstLspPlugInPluginsLv2MbCompressorMsscs2M)
//
// Sidechain source 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs2M() (interface{}, error) {
	return e.Element.GetProperty("scs-2m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs2M(value interface{}) error {
	return e.Element.SetProperty("scs-2m", value)
}

// scs-2s (GstLspPlugInPluginsLv2MbCompressorMsscs2S)
//
// Sidechain source 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs2S() (interface{}, error) {
	return e.Element.GetProperty("scs-2s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs2S(value interface{}) error {
	return e.Element.SetProperty("scs-2s", value)
}

// scs-3m (GstLspPlugInPluginsLv2MbCompressorMsscs3M)
//
// Sidechain source 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs3M() (interface{}, error) {
	return e.Element.GetProperty("scs-3m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs3M(value interface{}) error {
	return e.Element.SetProperty("scs-3m", value)
}

// scs-3s (GstLspPlugInPluginsLv2MbCompressorMsscs3S)
//
// Sidechain source 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs3S() (interface{}, error) {
	return e.Element.GetProperty("scs-3s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs3S(value interface{}) error {
	return e.Element.SetProperty("scs-3s", value)
}

// scs-4m (GstLspPlugInPluginsLv2MbCompressorMsscs4M)
//
// Sidechain source 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs4M() (interface{}, error) {
	return e.Element.GetProperty("scs-4m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs4M(value interface{}) error {
	return e.Element.SetProperty("scs-4m", value)
}

// scs-4s (GstLspPlugInPluginsLv2MbCompressorMsscs4S)
//
// Sidechain source 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs4S() (interface{}, error) {
	return e.Element.GetProperty("scs-4s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs4S(value interface{}) error {
	return e.Element.SetProperty("scs-4s", value)
}

// scs-5m (GstLspPlugInPluginsLv2MbCompressorMsscs5M)
//
// Sidechain source 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs5M() (interface{}, error) {
	return e.Element.GetProperty("scs-5m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs5M(value interface{}) error {
	return e.Element.SetProperty("scs-5m", value)
}

// scs-5s (GstLspPlugInPluginsLv2MbCompressorMsscs5S)
//
// Sidechain source 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs5S() (interface{}, error) {
	return e.Element.GetProperty("scs-5s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs5S(value interface{}) error {
	return e.Element.SetProperty("scs-5s", value)
}

// scs-6m (GstLspPlugInPluginsLv2MbCompressorMsscs6M)
//
// Sidechain source 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs6M() (interface{}, error) {
	return e.Element.GetProperty("scs-6m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs6M(value interface{}) error {
	return e.Element.SetProperty("scs-6m", value)
}

// scs-6s (GstLspPlugInPluginsLv2MbCompressorMsscs6S)
//
// Sidechain source 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs6S() (interface{}, error) {
	return e.Element.GetProperty("scs-6s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs6S(value interface{}) error {
	return e.Element.SetProperty("scs-6s", value)
}

// scs-7m (GstLspPlugInPluginsLv2MbCompressorMsscs7M)
//
// Sidechain source 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs7M() (interface{}, error) {
	return e.Element.GetProperty("scs-7m")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs7M(value interface{}) error {
	return e.Element.SetProperty("scs-7m", value)
}

// scs-7s (GstLspPlugInPluginsLv2MbCompressorMsscs7S)
//
// Sidechain source 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetScs7S() (interface{}, error) {
	return e.Element.GetProperty("scs-7s")
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetScs7S(value interface{}) error {
	return e.Element.SetProperty("scs-7s", value)
}

// sf-1m (float32)
//
// Split frequency 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf1M(value float32) error {
	return e.Element.SetProperty("sf-1m", value)
}

// sf-1s (float32)
//
// Split frequency 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf1S(value float32) error {
	return e.Element.SetProperty("sf-1s", value)
}

// sf-2m (float32)
//
// Split frequency 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf2M(value float32) error {
	return e.Element.SetProperty("sf-2m", value)
}

// sf-2s (float32)
//
// Split frequency 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf2S(value float32) error {
	return e.Element.SetProperty("sf-2s", value)
}

// sf-3m (float32)
//
// Split frequency 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf3M(value float32) error {
	return e.Element.SetProperty("sf-3m", value)
}

// sf-3s (float32)
//
// Split frequency 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf3S(value float32) error {
	return e.Element.SetProperty("sf-3s", value)
}

// sf-4m (float32)
//
// Split frequency 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf4M(value float32) error {
	return e.Element.SetProperty("sf-4m", value)
}

// sf-4s (float32)
//
// Split frequency 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf4S(value float32) error {
	return e.Element.SetProperty("sf-4s", value)
}

// sf-5m (float32)
//
// Split frequency 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf5M(value float32) error {
	return e.Element.SetProperty("sf-5m", value)
}

// sf-5s (float32)
//
// Split frequency 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf5S(value float32) error {
	return e.Element.SetProperty("sf-5s", value)
}

// sf-6m (float32)
//
// Split frequency 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf6M(value float32) error {
	return e.Element.SetProperty("sf-6m", value)
}

// sf-6s (float32)
//
// Split frequency 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf6S(value float32) error {
	return e.Element.SetProperty("sf-6s", value)
}

// sf-7m (float32)
//
// Split frequency 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf7M(value float32) error {
	return e.Element.SetProperty("sf-7m", value)
}

// sf-7s (float32)
//
// Split frequency 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSf7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSf7S(value float32) error {
	return e.Element.SetProperty("sf-7s", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2MbCompressorMs) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sla-0m (float32)
//
// Sidechain lookahead 0 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla0M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-0m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla0M(value float32) error {
	return e.Element.SetProperty("sla-0m", value)
}

// sla-0s (float32)
//
// Sidechain lookahead 0 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla0S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-0s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla0S(value float32) error {
	return e.Element.SetProperty("sla-0s", value)
}

// sla-1m (float32)
//
// Sidechain lookahead 1 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla1M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-1m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla1M(value float32) error {
	return e.Element.SetProperty("sla-1m", value)
}

// sla-1s (float32)
//
// Sidechain lookahead 1 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla1S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-1s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla1S(value float32) error {
	return e.Element.SetProperty("sla-1s", value)
}

// sla-2m (float32)
//
// Sidechain lookahead 2 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla2M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-2m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla2M(value float32) error {
	return e.Element.SetProperty("sla-2m", value)
}

// sla-2s (float32)
//
// Sidechain lookahead 2 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla2S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-2s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla2S(value float32) error {
	return e.Element.SetProperty("sla-2s", value)
}

// sla-3m (float32)
//
// Sidechain lookahead 3 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla3M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-3m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla3M(value float32) error {
	return e.Element.SetProperty("sla-3m", value)
}

// sla-3s (float32)
//
// Sidechain lookahead 3 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla3S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-3s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla3S(value float32) error {
	return e.Element.SetProperty("sla-3s", value)
}

// sla-4m (float32)
//
// Sidechain lookahead 4 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla4M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-4m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla4M(value float32) error {
	return e.Element.SetProperty("sla-4m", value)
}

// sla-4s (float32)
//
// Sidechain lookahead 4 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla4S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-4s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla4S(value float32) error {
	return e.Element.SetProperty("sla-4s", value)
}

// sla-5m (float32)
//
// Sidechain lookahead 5 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla5M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-5m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla5M(value float32) error {
	return e.Element.SetProperty("sla-5m", value)
}

// sla-5s (float32)
//
// Sidechain lookahead 5 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla5S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-5s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla5S(value float32) error {
	return e.Element.SetProperty("sla-5s", value)
}

// sla-6m (float32)
//
// Sidechain lookahead 6 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla6M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-6m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla6M(value float32) error {
	return e.Element.SetProperty("sla-6m", value)
}

// sla-6s (float32)
//
// Sidechain lookahead 6 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla6S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-6s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla6S(value float32) error {
	return e.Element.SetProperty("sla-6s", value)
}

// sla-7m (float32)
//
// Sidechain lookahead 7 Mid

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla7M() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-7m")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla7M(value float32) error {
	return e.Element.SetProperty("sla-7m", value)
}

// sla-7s (float32)
//
// Sidechain lookahead 7 Side

func (e *LspPlugInPluginsLv2MbCompressorMs) GetSla7S() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-7s")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMs) SetSla7S(value float32) error {
	return e.Element.SetProperty("sla-7s", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2MbCompressorMs) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMs) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MbCompressorMsscs5S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs5SMid LspPlugInPluginsLv2MbCompressorMsscs5S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs5SSide LspPlugInPluginsLv2MbCompressorMsscs5S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs5SLeft LspPlugInPluginsLv2MbCompressorMsscs5S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs5SRight LspPlugInPluginsLv2MbCompressorMsscs5S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsbsel string

const (
	LspPlugInPluginsLv2MbCompressorMsbselSplitMid LspPlugInPluginsLv2MbCompressorMsbsel = "Split Mid" // Split Mid (0) – Split Mid
	LspPlugInPluginsLv2MbCompressorMsbselSplitSide LspPlugInPluginsLv2MbCompressorMsbsel = "Split Side" // Split Side (1) – Split Side
	LspPlugInPluginsLv2MbCompressorMsbselBand0 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 0" // Band 0 (2) – Band 0
	LspPlugInPluginsLv2MbCompressorMsbselBand1 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 1" // Band 1 (3) – Band 1
	LspPlugInPluginsLv2MbCompressorMsbselBand2 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 2" // Band 2 (4) – Band 2
	LspPlugInPluginsLv2MbCompressorMsbselBand3 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 3" // Band 3 (5) – Band 3
	LspPlugInPluginsLv2MbCompressorMsbselBand4 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 4" // Band 4 (6) – Band 4
	LspPlugInPluginsLv2MbCompressorMsbselBand5 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 5" // Band 5 (7) – Band 5
	LspPlugInPluginsLv2MbCompressorMsbselBand6 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 6" // Band 6 (8) – Band 6
	LspPlugInPluginsLv2MbCompressorMsbselBand7 LspPlugInPluginsLv2MbCompressorMsbsel = "Band 7" // Band 7 (9) – Band 7
)

type LspPlugInPluginsLv2MbCompressorMscm2M string

const (
	LspPlugInPluginsLv2MbCompressorMscm2MDown LspPlugInPluginsLv2MbCompressorMscm2M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm2MUp LspPlugInPluginsLv2MbCompressorMscm2M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm2S string

const (
	LspPlugInPluginsLv2MbCompressorMscm2SDown LspPlugInPluginsLv2MbCompressorMscm2S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm2SUp LspPlugInPluginsLv2MbCompressorMscm2S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm4S string

const (
	LspPlugInPluginsLv2MbCompressorMscm4SDown LspPlugInPluginsLv2MbCompressorMscm4S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm4SUp LspPlugInPluginsLv2MbCompressorMscm4S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm7S string

const (
	LspPlugInPluginsLv2MbCompressorMscm7SDown LspPlugInPluginsLv2MbCompressorMscm7S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm7SUp LspPlugInPluginsLv2MbCompressorMscm7S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsscm2S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm2SPeak LspPlugInPluginsLv2MbCompressorMsscm2S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm2SRms LspPlugInPluginsLv2MbCompressorMsscm2S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm2SLpf LspPlugInPluginsLv2MbCompressorMsscm2S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm2SUniform LspPlugInPluginsLv2MbCompressorMsscm2S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMscm0M string

const (
	LspPlugInPluginsLv2MbCompressorMscm0MDown LspPlugInPluginsLv2MbCompressorMscm0M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm0MUp LspPlugInPluginsLv2MbCompressorMscm0M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsscm2M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm2MPeak LspPlugInPluginsLv2MbCompressorMsscm2M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm2MRms LspPlugInPluginsLv2MbCompressorMsscm2M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm2MLpf LspPlugInPluginsLv2MbCompressorMsscm2M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm2MUniform LspPlugInPluginsLv2MbCompressorMsscm2M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm4M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm4MPeak LspPlugInPluginsLv2MbCompressorMsscm4M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm4MRms LspPlugInPluginsLv2MbCompressorMsscm4M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm4MLpf LspPlugInPluginsLv2MbCompressorMsscm4M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm4MUniform LspPlugInPluginsLv2MbCompressorMsscm4M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscs1S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs1SMid LspPlugInPluginsLv2MbCompressorMsscs1S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs1SSide LspPlugInPluginsLv2MbCompressorMsscs1S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs1SLeft LspPlugInPluginsLv2MbCompressorMsscs1S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs1SRight LspPlugInPluginsLv2MbCompressorMsscs1S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs3S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs3SMid LspPlugInPluginsLv2MbCompressorMsscs3S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs3SSide LspPlugInPluginsLv2MbCompressorMsscs3S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs3SLeft LspPlugInPluginsLv2MbCompressorMsscs3S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs3SRight LspPlugInPluginsLv2MbCompressorMsscs3S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs7M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs7MMid LspPlugInPluginsLv2MbCompressorMsscs7M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs7MSide LspPlugInPluginsLv2MbCompressorMsscs7M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs7MLeft LspPlugInPluginsLv2MbCompressorMsscs7M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs7MRight LspPlugInPluginsLv2MbCompressorMsscs7M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsmode string

const (
	LspPlugInPluginsLv2MbCompressorMsmodeClassic LspPlugInPluginsLv2MbCompressorMsmode = "Classic" // Classic (0) – Classic
	LspPlugInPluginsLv2MbCompressorMsmodeModern LspPlugInPluginsLv2MbCompressorMsmode = "Modern" // Modern (1) – Modern
)

type LspPlugInPluginsLv2MbCompressorMsscm6M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm6MPeak LspPlugInPluginsLv2MbCompressorMsscm6M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm6MRms LspPlugInPluginsLv2MbCompressorMsscm6M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm6MLpf LspPlugInPluginsLv2MbCompressorMsscm6M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm6MUniform LspPlugInPluginsLv2MbCompressorMsscm6M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscs1M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs1MMid LspPlugInPluginsLv2MbCompressorMsscs1M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs1MSide LspPlugInPluginsLv2MbCompressorMsscs1M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs1MLeft LspPlugInPluginsLv2MbCompressorMsscs1M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs1MRight LspPlugInPluginsLv2MbCompressorMsscs1M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMscm6M string

const (
	LspPlugInPluginsLv2MbCompressorMscm6MDown LspPlugInPluginsLv2MbCompressorMscm6M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm6MUp LspPlugInPluginsLv2MbCompressorMscm6M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsscm4S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm4SPeak LspPlugInPluginsLv2MbCompressorMsscm4S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm4SRms LspPlugInPluginsLv2MbCompressorMsscm4S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm4SLpf LspPlugInPluginsLv2MbCompressorMsscm4S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm4SUniform LspPlugInPluginsLv2MbCompressorMsscm4S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm5M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm5MPeak LspPlugInPluginsLv2MbCompressorMsscm5M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm5MRms LspPlugInPluginsLv2MbCompressorMsscm5M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm5MLpf LspPlugInPluginsLv2MbCompressorMsscm5M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm5MUniform LspPlugInPluginsLv2MbCompressorMsscm5M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm5S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm5SPeak LspPlugInPluginsLv2MbCompressorMsscm5S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm5SRms LspPlugInPluginsLv2MbCompressorMsscm5S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm5SLpf LspPlugInPluginsLv2MbCompressorMsscm5S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm5SUniform LspPlugInPluginsLv2MbCompressorMsscm5S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm6S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm6SPeak LspPlugInPluginsLv2MbCompressorMsscm6S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm6SRms LspPlugInPluginsLv2MbCompressorMsscm6S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm6SLpf LspPlugInPluginsLv2MbCompressorMsscm6S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm6SUniform LspPlugInPluginsLv2MbCompressorMsscm6S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscs0S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs0SMid LspPlugInPluginsLv2MbCompressorMsscs0S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs0SSide LspPlugInPluginsLv2MbCompressorMsscs0S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs0SLeft LspPlugInPluginsLv2MbCompressorMsscs0S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs0SRight LspPlugInPluginsLv2MbCompressorMsscs0S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs2M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs2MMid LspPlugInPluginsLv2MbCompressorMsscs2M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs2MSide LspPlugInPluginsLv2MbCompressorMsscs2M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs2MLeft LspPlugInPluginsLv2MbCompressorMsscs2M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs2MRight LspPlugInPluginsLv2MbCompressorMsscs2M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMscm1M string

const (
	LspPlugInPluginsLv2MbCompressorMscm1MDown LspPlugInPluginsLv2MbCompressorMscm1M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm1MUp LspPlugInPluginsLv2MbCompressorMscm1M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm3M string

const (
	LspPlugInPluginsLv2MbCompressorMscm3MDown LspPlugInPluginsLv2MbCompressorMscm3M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm3MUp LspPlugInPluginsLv2MbCompressorMscm3M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm5M string

const (
	LspPlugInPluginsLv2MbCompressorMscm5MDown LspPlugInPluginsLv2MbCompressorMscm5M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm5MUp LspPlugInPluginsLv2MbCompressorMscm5M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsscm1M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm1MPeak LspPlugInPluginsLv2MbCompressorMsscm1M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm1MRms LspPlugInPluginsLv2MbCompressorMsscm1M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm1MLpf LspPlugInPluginsLv2MbCompressorMsscm1M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm1MUniform LspPlugInPluginsLv2MbCompressorMsscm1M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm3M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm3MPeak LspPlugInPluginsLv2MbCompressorMsscm3M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm3MRms LspPlugInPluginsLv2MbCompressorMsscm3M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm3MLpf LspPlugInPluginsLv2MbCompressorMsscm3M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm3MUniform LspPlugInPluginsLv2MbCompressorMsscm3M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm7S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm7SPeak LspPlugInPluginsLv2MbCompressorMsscm7S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm7SRms LspPlugInPluginsLv2MbCompressorMsscm7S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm7SLpf LspPlugInPluginsLv2MbCompressorMsscm7S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm7SUniform LspPlugInPluginsLv2MbCompressorMsscm7S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscs2S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs2SMid LspPlugInPluginsLv2MbCompressorMsscs2S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs2SSide LspPlugInPluginsLv2MbCompressorMsscs2S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs2SLeft LspPlugInPluginsLv2MbCompressorMsscs2S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs2SRight LspPlugInPluginsLv2MbCompressorMsscs2S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs3M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs3MMid LspPlugInPluginsLv2MbCompressorMsscs3M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs3MSide LspPlugInPluginsLv2MbCompressorMsscs3M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs3MLeft LspPlugInPluginsLv2MbCompressorMsscs3M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs3MRight LspPlugInPluginsLv2MbCompressorMsscs3M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs6S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs6SMid LspPlugInPluginsLv2MbCompressorMsscs6S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs6SSide LspPlugInPluginsLv2MbCompressorMsscs6S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs6SLeft LspPlugInPluginsLv2MbCompressorMsscs6S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs6SRight LspPlugInPluginsLv2MbCompressorMsscs6S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMscm3S string

const (
	LspPlugInPluginsLv2MbCompressorMscm3SDown LspPlugInPluginsLv2MbCompressorMscm3S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm3SUp LspPlugInPluginsLv2MbCompressorMscm3S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsscs4M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs4MMid LspPlugInPluginsLv2MbCompressorMsscs4M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs4MSide LspPlugInPluginsLv2MbCompressorMsscs4M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs4MLeft LspPlugInPluginsLv2MbCompressorMsscs4M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs4MRight LspPlugInPluginsLv2MbCompressorMsscs4M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs6M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs6MMid LspPlugInPluginsLv2MbCompressorMsscs6M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs6MSide LspPlugInPluginsLv2MbCompressorMsscs6M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs6MLeft LspPlugInPluginsLv2MbCompressorMsscs6M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs6MRight LspPlugInPluginsLv2MbCompressorMsscs6M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs4S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs4SMid LspPlugInPluginsLv2MbCompressorMsscs4S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs4SSide LspPlugInPluginsLv2MbCompressorMsscs4S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs4SLeft LspPlugInPluginsLv2MbCompressorMsscs4S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs4SRight LspPlugInPluginsLv2MbCompressorMsscs4S = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMscm4M string

const (
	LspPlugInPluginsLv2MbCompressorMscm4MDown LspPlugInPluginsLv2MbCompressorMscm4M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm4MUp LspPlugInPluginsLv2MbCompressorMscm4M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm5S string

const (
	LspPlugInPluginsLv2MbCompressorMscm5SDown LspPlugInPluginsLv2MbCompressorMscm5S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm5SUp LspPlugInPluginsLv2MbCompressorMscm5S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm7M string

const (
	LspPlugInPluginsLv2MbCompressorMscm7MDown LspPlugInPluginsLv2MbCompressorMscm7M = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm7MUp LspPlugInPluginsLv2MbCompressorMscm7M = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsscm0M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm0MPeak LspPlugInPluginsLv2MbCompressorMsscm0M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm0MRms LspPlugInPluginsLv2MbCompressorMsscm0M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm0MLpf LspPlugInPluginsLv2MbCompressorMsscm0M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm0MUniform LspPlugInPluginsLv2MbCompressorMsscm0M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm0S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm0SPeak LspPlugInPluginsLv2MbCompressorMsscm0S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm0SRms LspPlugInPluginsLv2MbCompressorMsscm0S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm0SLpf LspPlugInPluginsLv2MbCompressorMsscm0S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm0SUniform LspPlugInPluginsLv2MbCompressorMsscm0S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm1S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm1SPeak LspPlugInPluginsLv2MbCompressorMsscm1S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm1SRms LspPlugInPluginsLv2MbCompressorMsscm1S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm1SLpf LspPlugInPluginsLv2MbCompressorMsscm1S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm1SUniform LspPlugInPluginsLv2MbCompressorMsscm1S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscs0M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs0MMid LspPlugInPluginsLv2MbCompressorMsscs0M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs0MSide LspPlugInPluginsLv2MbCompressorMsscs0M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs0MLeft LspPlugInPluginsLv2MbCompressorMsscs0M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs0MRight LspPlugInPluginsLv2MbCompressorMsscs0M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMsscs5M string

const (
	LspPlugInPluginsLv2MbCompressorMsscs5MMid LspPlugInPluginsLv2MbCompressorMsscs5M = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs5MSide LspPlugInPluginsLv2MbCompressorMsscs5M = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs5MLeft LspPlugInPluginsLv2MbCompressorMsscs5M = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs5MRight LspPlugInPluginsLv2MbCompressorMsscs5M = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorMscm0S string

const (
	LspPlugInPluginsLv2MbCompressorMscm0SDown LspPlugInPluginsLv2MbCompressorMscm0S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm0SUp LspPlugInPluginsLv2MbCompressorMscm0S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm1S string

const (
	LspPlugInPluginsLv2MbCompressorMscm1SDown LspPlugInPluginsLv2MbCompressorMscm1S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm1SUp LspPlugInPluginsLv2MbCompressorMscm1S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMscm6S string

const (
	LspPlugInPluginsLv2MbCompressorMscm6SDown LspPlugInPluginsLv2MbCompressorMscm6S = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMscm6SUp LspPlugInPluginsLv2MbCompressorMscm6S = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMsenvb string

const (
	LspPlugInPluginsLv2MbCompressorMsenvbNone LspPlugInPluginsLv2MbCompressorMsenvb = "None" // None (0) – None
	LspPlugInPluginsLv2MbCompressorMsenvbPinkBt LspPlugInPluginsLv2MbCompressorMsenvb = "Pink BT" // Pink BT (1) – Pink BT
	LspPlugInPluginsLv2MbCompressorMsenvbPinkMt LspPlugInPluginsLv2MbCompressorMsenvb = "Pink MT" // Pink MT (2) – Pink MT
	LspPlugInPluginsLv2MbCompressorMsenvbBrownBt LspPlugInPluginsLv2MbCompressorMsenvb = "Brown BT" // Brown BT (3) – Brown BT
	LspPlugInPluginsLv2MbCompressorMsenvbBrownMt LspPlugInPluginsLv2MbCompressorMsenvb = "Brown MT" // Brown MT (4) – Brown MT
)

type LspPlugInPluginsLv2MbCompressorMsscm3S string

const (
	LspPlugInPluginsLv2MbCompressorMsscm3SPeak LspPlugInPluginsLv2MbCompressorMsscm3S = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm3SRms LspPlugInPluginsLv2MbCompressorMsscm3S = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm3SLpf LspPlugInPluginsLv2MbCompressorMsscm3S = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm3SUniform LspPlugInPluginsLv2MbCompressorMsscm3S = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscm7M string

const (
	LspPlugInPluginsLv2MbCompressorMsscm7MPeak LspPlugInPluginsLv2MbCompressorMsscm7M = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMsscm7MRms LspPlugInPluginsLv2MbCompressorMsscm7M = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMsscm7MLpf LspPlugInPluginsLv2MbCompressorMsscm7M = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMsscm7MUniform LspPlugInPluginsLv2MbCompressorMsscm7M = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMsscs7S string

const (
	LspPlugInPluginsLv2MbCompressorMsscs7SMid LspPlugInPluginsLv2MbCompressorMsscs7S = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorMsscs7SSide LspPlugInPluginsLv2MbCompressorMsscs7S = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorMsscs7SLeft LspPlugInPluginsLv2MbCompressorMsscs7S = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorMsscs7SRight LspPlugInPluginsLv2MbCompressorMsscs7S = "Right" // Right (3) – Right
)

