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

type LspPlugInPluginsLv2MbCompressorStereo struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MbCompressorStereo() (*LspPlugInPluginsLv2MbCompressorStereo, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-mb-compressor-stereo")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MbCompressorStereoWithName(name string) (*LspPlugInPluginsLv2MbCompressorStereo, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-mb-compressor-stereo", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorStereo{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-0 (float32)
//
// Attack level 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl0(value float32) error {
	return e.Element.SetProperty("al-0", value)
}

// al-1 (float32)
//
// Attack level 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl1(value float32) error {
	return e.Element.SetProperty("al-1", value)
}

// al-2 (float32)
//
// Attack level 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl2(value float32) error {
	return e.Element.SetProperty("al-2", value)
}

// al-3 (float32)
//
// Attack level 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl3(value float32) error {
	return e.Element.SetProperty("al-3", value)
}

// al-4 (float32)
//
// Attack level 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl4(value float32) error {
	return e.Element.SetProperty("al-4", value)
}

// al-5 (float32)
//
// Attack level 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl5(value float32) error {
	return e.Element.SetProperty("al-5", value)
}

// al-6 (float32)
//
// Attack level 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl6(value float32) error {
	return e.Element.SetProperty("al-6", value)
}

// al-7 (float32)
//
// Attack level 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAl7(value float32) error {
	return e.Element.SetProperty("al-7", value)
}

// at-0 (float32)
//
// Attack time 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt0(value float32) error {
	return e.Element.SetProperty("at-0", value)
}

// at-1 (float32)
//
// Attack time 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt1(value float32) error {
	return e.Element.SetProperty("at-1", value)
}

// at-2 (float32)
//
// Attack time 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt2(value float32) error {
	return e.Element.SetProperty("at-2", value)
}

// at-3 (float32)
//
// Attack time 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt3(value float32) error {
	return e.Element.SetProperty("at-3", value)
}

// at-4 (float32)
//
// Attack time 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt4(value float32) error {
	return e.Element.SetProperty("at-4", value)
}

// at-5 (float32)
//
// Attack time 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt5(value float32) error {
	return e.Element.SetProperty("at-5", value)
}

// at-6 (float32)
//
// Attack time 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt6(value float32) error {
	return e.Element.SetProperty("at-6", value)
}

// at-7 (float32)
//
// Attack time 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetAt7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetAt7(value float32) error {
	return e.Element.SetProperty("at-7", value)
}

// bm-0 (bool)
//
// Mute band 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm0(value bool) error {
	return e.Element.SetProperty("bm-0", value)
}

// bm-1 (bool)
//
// Mute band 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm1(value bool) error {
	return e.Element.SetProperty("bm-1", value)
}

// bm-2 (bool)
//
// Mute band 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm2(value bool) error {
	return e.Element.SetProperty("bm-2", value)
}

// bm-3 (bool)
//
// Mute band 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm3(value bool) error {
	return e.Element.SetProperty("bm-3", value)
}

// bm-4 (bool)
//
// Mute band 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm4(value bool) error {
	return e.Element.SetProperty("bm-4", value)
}

// bm-5 (bool)
//
// Mute band 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm5(value bool) error {
	return e.Element.SetProperty("bm-5", value)
}

// bm-6 (bool)
//
// Mute band 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm6(value bool) error {
	return e.Element.SetProperty("bm-6", value)
}

// bm-7 (bool)
//
// Mute band 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBm7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBm7(value bool) error {
	return e.Element.SetProperty("bm-7", value)
}

// bs-0 (bool)
//
// Solo band 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs0(value bool) error {
	return e.Element.SetProperty("bs-0", value)
}

// bs-1 (bool)
//
// Solo band 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs1(value bool) error {
	return e.Element.SetProperty("bs-1", value)
}

// bs-2 (bool)
//
// Solo band 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs2(value bool) error {
	return e.Element.SetProperty("bs-2", value)
}

// bs-3 (bool)
//
// Solo band 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs3(value bool) error {
	return e.Element.SetProperty("bs-3", value)
}

// bs-4 (bool)
//
// Solo band 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs4(value bool) error {
	return e.Element.SetProperty("bs-4", value)
}

// bs-5 (bool)
//
// Solo band 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs5(value bool) error {
	return e.Element.SetProperty("bs-5", value)
}

// bs-6 (bool)
//
// Solo band 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs6(value bool) error {
	return e.Element.SetProperty("bs-6", value)
}

// bs-7 (bool)
//
// Solo band 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBs7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBs7(value bool) error {
	return e.Element.SetProperty("bs-7", value)
}

// bsel (GstLspPlugInPluginsLv2MbCompressorStereobsel)
//
// Band selection

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBsel() (interface{}, error) {
	return e.Element.GetProperty("bsel")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBsel(value interface{}) error {
	return e.Element.SetProperty("bsel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cbe-1 (bool)
//
// Compression band enable 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe1(value bool) error {
	return e.Element.SetProperty("cbe-1", value)
}

// cbe-2 (bool)
//
// Compression band enable 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe2(value bool) error {
	return e.Element.SetProperty("cbe-2", value)
}

// cbe-3 (bool)
//
// Compression band enable 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe3(value bool) error {
	return e.Element.SetProperty("cbe-3", value)
}

// cbe-4 (bool)
//
// Compression band enable 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe4(value bool) error {
	return e.Element.SetProperty("cbe-4", value)
}

// cbe-5 (bool)
//
// Compression band enable 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe5(value bool) error {
	return e.Element.SetProperty("cbe-5", value)
}

// cbe-6 (bool)
//
// Compression band enable 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe6(value bool) error {
	return e.Element.SetProperty("cbe-6", value)
}

// cbe-7 (bool)
//
// Compression band enable 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCbe7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCbe7(value bool) error {
	return e.Element.SetProperty("cbe-7", value)
}

// ce-0 (bool)
//
// Compressor enable 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe0(value bool) error {
	return e.Element.SetProperty("ce-0", value)
}

// ce-1 (bool)
//
// Compressor enable 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe1(value bool) error {
	return e.Element.SetProperty("ce-1", value)
}

// ce-2 (bool)
//
// Compressor enable 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe2(value bool) error {
	return e.Element.SetProperty("ce-2", value)
}

// ce-3 (bool)
//
// Compressor enable 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe3(value bool) error {
	return e.Element.SetProperty("ce-3", value)
}

// ce-4 (bool)
//
// Compressor enable 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe4(value bool) error {
	return e.Element.SetProperty("ce-4", value)
}

// ce-5 (bool)
//
// Compressor enable 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe5(value bool) error {
	return e.Element.SetProperty("ce-5", value)
}

// ce-6 (bool)
//
// Compressor enable 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe6(value bool) error {
	return e.Element.SetProperty("ce-6", value)
}

// ce-7 (bool)
//
// Compressor enable 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCe7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCe7(value bool) error {
	return e.Element.SetProperty("ce-7", value)
}

// clm-0 (float32)
//
// Curve level meter 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-1 (float32)
//
// Curve level meter 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-2 (float32)
//
// Curve level meter 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-3 (float32)
//
// Curve level meter 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-4 (float32)
//
// Curve level meter 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-5 (float32)
//
// Curve level meter 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-6 (float32)
//
// Curve level meter 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-7 (float32)
//
// Curve level meter 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetClm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cm-0 (GstLspPlugInPluginsLv2MbCompressorStereocm0)
//
// Compression mode 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm0() (interface{}, error) {
	return e.Element.GetProperty("cm-0")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm0(value interface{}) error {
	return e.Element.SetProperty("cm-0", value)
}

// cm-1 (GstLspPlugInPluginsLv2MbCompressorStereocm1)
//
// Compression mode 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm1() (interface{}, error) {
	return e.Element.GetProperty("cm-1")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm1(value interface{}) error {
	return e.Element.SetProperty("cm-1", value)
}

// cm-2 (GstLspPlugInPluginsLv2MbCompressorStereocm2)
//
// Compression mode 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm2() (interface{}, error) {
	return e.Element.GetProperty("cm-2")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm2(value interface{}) error {
	return e.Element.SetProperty("cm-2", value)
}

// cm-3 (GstLspPlugInPluginsLv2MbCompressorStereocm3)
//
// Compression mode 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm3() (interface{}, error) {
	return e.Element.GetProperty("cm-3")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm3(value interface{}) error {
	return e.Element.SetProperty("cm-3", value)
}

// cm-4 (GstLspPlugInPluginsLv2MbCompressorStereocm4)
//
// Compression mode 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm4() (interface{}, error) {
	return e.Element.GetProperty("cm-4")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm4(value interface{}) error {
	return e.Element.SetProperty("cm-4", value)
}

// cm-5 (GstLspPlugInPluginsLv2MbCompressorStereocm5)
//
// Compression mode 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm5() (interface{}, error) {
	return e.Element.GetProperty("cm-5")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm5(value interface{}) error {
	return e.Element.SetProperty("cm-5", value)
}

// cm-6 (GstLspPlugInPluginsLv2MbCompressorStereocm6)
//
// Compression mode 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm6() (interface{}, error) {
	return e.Element.GetProperty("cm-6")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm6(value interface{}) error {
	return e.Element.SetProperty("cm-6", value)
}

// cm-7 (GstLspPlugInPluginsLv2MbCompressorStereocm7)
//
// Compression mode 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCm7() (interface{}, error) {
	return e.Element.GetProperty("cm-7")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCm7(value interface{}) error {
	return e.Element.SetProperty("cm-7", value)
}

// cr-0 (float32)
//
// Ratio 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr0(value float32) error {
	return e.Element.SetProperty("cr-0", value)
}

// cr-1 (float32)
//
// Ratio 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr1(value float32) error {
	return e.Element.SetProperty("cr-1", value)
}

// cr-2 (float32)
//
// Ratio 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr2(value float32) error {
	return e.Element.SetProperty("cr-2", value)
}

// cr-3 (float32)
//
// Ratio 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr3(value float32) error {
	return e.Element.SetProperty("cr-3", value)
}

// cr-4 (float32)
//
// Ratio 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr4(value float32) error {
	return e.Element.SetProperty("cr-4", value)
}

// cr-5 (float32)
//
// Ratio 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr5(value float32) error {
	return e.Element.SetProperty("cr-5", value)
}

// cr-6 (float32)
//
// Ratio 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr6(value float32) error {
	return e.Element.SetProperty("cr-6", value)
}

// cr-7 (float32)
//
// Ratio 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetCr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetCr7(value float32) error {
	return e.Element.SetProperty("cr-7", value)
}

// elm-0 (float32)
//
// Envelope level meter 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-1 (float32)
//
// Envelope level meter 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-2 (float32)
//
// Envelope level meter 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-3 (float32)
//
// Envelope level meter 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-4 (float32)
//
// Envelope level meter 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-5 (float32)
//
// Envelope level meter 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-6 (float32)
//
// Envelope level meter 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-7 (float32)
//
// Envelope level meter 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetElm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// envb (GstLspPlugInPluginsLv2MbCompressorStereoenvb)
//
// Envelope boost

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetEnvb() (interface{}, error) {
	return e.Element.GetProperty("envb")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetEnvb(value interface{}) error {
	return e.Element.SetProperty("envb", value)
}

// flt (bool)
//
// Band filter curves

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFlt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetFlt(value bool) error {
	return e.Element.SetProperty("flt", value)
}

// fre-0 (float32)
//
// Frequency range end 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-1 (float32)
//
// Frequency range end 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-2 (float32)
//
// Frequency range end 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-3 (float32)
//
// Frequency range end 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-4 (float32)
//
// Frequency range end 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-5 (float32)
//
// Frequency range end 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-6 (float32)
//
// Frequency range end 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-7 (float32)
//
// Frequency range end 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetFre7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-7")
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetGDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetGDry(value float32) error {
	return e.Element.SetProperty("g-dry", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// g-wet (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetGWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetGWet(value float32) error {
	return e.Element.SetProperty("g-wet", value)
}

// hue-0 (float32)
//
// Hue  0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// hue-1 (float32)
//
// Hue  1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue1(value float32) error {
	return e.Element.SetProperty("hue-1", value)
}

// hue-2 (float32)
//
// Hue  2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue2(value float32) error {
	return e.Element.SetProperty("hue-2", value)
}

// hue-3 (float32)
//
// Hue  3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue3(value float32) error {
	return e.Element.SetProperty("hue-3", value)
}

// hue-4 (float32)
//
// Hue  4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue4(value float32) error {
	return e.Element.SetProperty("hue-4", value)
}

// hue-5 (float32)
//
// Hue  5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue5(value float32) error {
	return e.Element.SetProperty("hue-5", value)
}

// hue-6 (float32)
//
// Hue  6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue6(value float32) error {
	return e.Element.SetProperty("hue-6", value)
}

// hue-7 (float32)
//
// Hue  7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetHue7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetHue7(value float32) error {
	return e.Element.SetProperty("hue-7", value)
}

// ife-l (bool)
//
// Input FFT graph enable Left

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetIfeL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ife-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetIfeL(value bool) error {
	return e.Element.SetProperty("ife-l", value)
}

// ife-r (bool)
//
// Input FFT graph enable Right

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetIfeR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ife-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetIfeR(value bool) error {
	return e.Element.SetProperty("ife-r", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetIlmR() (float32, error) {
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

// kn-0 (float32)
//
// Knee 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn0(value float32) error {
	return e.Element.SetProperty("kn-0", value)
}

// kn-1 (float32)
//
// Knee 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn1(value float32) error {
	return e.Element.SetProperty("kn-1", value)
}

// kn-2 (float32)
//
// Knee 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn2(value float32) error {
	return e.Element.SetProperty("kn-2", value)
}

// kn-3 (float32)
//
// Knee 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn3(value float32) error {
	return e.Element.SetProperty("kn-3", value)
}

// kn-4 (float32)
//
// Knee 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn4(value float32) error {
	return e.Element.SetProperty("kn-4", value)
}

// kn-5 (float32)
//
// Knee 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn5(value float32) error {
	return e.Element.SetProperty("kn-5", value)
}

// kn-6 (float32)
//
// Knee 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn6(value float32) error {
	return e.Element.SetProperty("kn-6", value)
}

// kn-7 (float32)
//
// Knee 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetKn7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetKn7(value float32) error {
	return e.Element.SetProperty("kn-7", value)
}

// mk-0 (float32)
//
// Makeup gain 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk0(value float32) error {
	return e.Element.SetProperty("mk-0", value)
}

// mk-1 (float32)
//
// Makeup gain 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk1(value float32) error {
	return e.Element.SetProperty("mk-1", value)
}

// mk-2 (float32)
//
// Makeup gain 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk2(value float32) error {
	return e.Element.SetProperty("mk-2", value)
}

// mk-3 (float32)
//
// Makeup gain 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk3(value float32) error {
	return e.Element.SetProperty("mk-3", value)
}

// mk-4 (float32)
//
// Makeup gain 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk4(value float32) error {
	return e.Element.SetProperty("mk-4", value)
}

// mk-5 (float32)
//
// Makeup gain 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk5(value float32) error {
	return e.Element.SetProperty("mk-5", value)
}

// mk-6 (float32)
//
// Makeup gain 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk6(value float32) error {
	return e.Element.SetProperty("mk-6", value)
}

// mk-7 (float32)
//
// Makeup gain 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMk7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMk7(value float32) error {
	return e.Element.SetProperty("mk-7", value)
}

// mode (GstLspPlugInPluginsLv2MbCompressorStereomode)
//
// Compressor mode

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ofe-l (bool)
//
// Output FFT graph enable Left

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetOfeL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofe-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetOfeL(value bool) error {
	return e.Element.SetProperty("ofe-l", value)
}

// ofe-r (bool)
//
// Output FFT graph enable Right

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetOfeR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofe-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetOfeR(value bool) error {
	return e.Element.SetProperty("ofe-r", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl-0 (float32)
//
// Release level 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-1 (float32)
//
// Release level 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-2 (float32)
//
// Release level 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-3 (float32)
//
// Release level 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-4 (float32)
//
// Release level 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-5 (float32)
//
// Release level 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-6 (float32)
//
// Release level 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-7 (float32)
//
// Release level 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-0 (float32)
//
// Reduction level meter 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-1 (float32)
//
// Reduction level meter 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-2 (float32)
//
// Reduction level meter 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-3 (float32)
//
// Reduction level meter 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-4 (float32)
//
// Reduction level meter 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-5 (float32)
//
// Reduction level meter 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-6 (float32)
//
// Reduction level meter 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-7 (float32)
//
// Reduction level meter 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRlm7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rrl-0 (float32)
//
// Relative release level 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl0(value float32) error {
	return e.Element.SetProperty("rrl-0", value)
}

// rrl-1 (float32)
//
// Relative release level 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl1(value float32) error {
	return e.Element.SetProperty("rrl-1", value)
}

// rrl-2 (float32)
//
// Relative release level 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl2(value float32) error {
	return e.Element.SetProperty("rrl-2", value)
}

// rrl-3 (float32)
//
// Relative release level 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl3(value float32) error {
	return e.Element.SetProperty("rrl-3", value)
}

// rrl-4 (float32)
//
// Relative release level 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl4(value float32) error {
	return e.Element.SetProperty("rrl-4", value)
}

// rrl-5 (float32)
//
// Relative release level 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl5(value float32) error {
	return e.Element.SetProperty("rrl-5", value)
}

// rrl-6 (float32)
//
// Relative release level 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl6(value float32) error {
	return e.Element.SetProperty("rrl-6", value)
}

// rrl-7 (float32)
//
// Relative release level 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRrl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRrl7(value float32) error {
	return e.Element.SetProperty("rrl-7", value)
}

// rt-0 (float32)
//
// Release time 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt0(value float32) error {
	return e.Element.SetProperty("rt-0", value)
}

// rt-1 (float32)
//
// Release time 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt1(value float32) error {
	return e.Element.SetProperty("rt-1", value)
}

// rt-2 (float32)
//
// Release time 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt2(value float32) error {
	return e.Element.SetProperty("rt-2", value)
}

// rt-3 (float32)
//
// Release time 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt3(value float32) error {
	return e.Element.SetProperty("rt-3", value)
}

// rt-4 (float32)
//
// Release time 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt4(value float32) error {
	return e.Element.SetProperty("rt-4", value)
}

// rt-5 (float32)
//
// Release time 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt5(value float32) error {
	return e.Element.SetProperty("rt-5", value)
}

// rt-6 (float32)
//
// Release time 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt6(value float32) error {
	return e.Element.SetProperty("rt-6", value)
}

// rt-7 (float32)
//
// Release time 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetRt7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetRt7(value float32) error {
	return e.Element.SetProperty("rt-7", value)
}

// schc-0 (bool)
//
// Sidechain custom hi-cut 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc0(value bool) error {
	return e.Element.SetProperty("schc-0", value)
}

// schc-1 (bool)
//
// Sidechain custom hi-cut 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc1(value bool) error {
	return e.Element.SetProperty("schc-1", value)
}

// schc-2 (bool)
//
// Sidechain custom hi-cut 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc2(value bool) error {
	return e.Element.SetProperty("schc-2", value)
}

// schc-3 (bool)
//
// Sidechain custom hi-cut 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc3(value bool) error {
	return e.Element.SetProperty("schc-3", value)
}

// schc-4 (bool)
//
// Sidechain custom hi-cut 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc4(value bool) error {
	return e.Element.SetProperty("schc-4", value)
}

// schc-5 (bool)
//
// Sidechain custom hi-cut 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc5(value bool) error {
	return e.Element.SetProperty("schc-5", value)
}

// schc-6 (bool)
//
// Sidechain custom hi-cut 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc6(value bool) error {
	return e.Element.SetProperty("schc-6", value)
}

// schc-7 (bool)
//
// Sidechain custom hi-cut 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchc7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchc7(value bool) error {
	return e.Element.SetProperty("schc-7", value)
}

// schf-0 (float32)
//
// Sidechain hi-cut frequency 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf0(value float32) error {
	return e.Element.SetProperty("schf-0", value)
}

// schf-1 (float32)
//
// Sidechain hi-cut frequency 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf1(value float32) error {
	return e.Element.SetProperty("schf-1", value)
}

// schf-2 (float32)
//
// Sidechain hi-cut frequency 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf2(value float32) error {
	return e.Element.SetProperty("schf-2", value)
}

// schf-3 (float32)
//
// Sidechain hi-cut frequency 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf3(value float32) error {
	return e.Element.SetProperty("schf-3", value)
}

// schf-4 (float32)
//
// Sidechain hi-cut frequency 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf4(value float32) error {
	return e.Element.SetProperty("schf-4", value)
}

// schf-5 (float32)
//
// Sidechain hi-cut frequency 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf5(value float32) error {
	return e.Element.SetProperty("schf-5", value)
}

// schf-6 (float32)
//
// Sidechain hi-cut frequency 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf6(value float32) error {
	return e.Element.SetProperty("schf-6", value)
}

// schf-7 (float32)
//
// Sidechain hi-cut frequency 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSchf7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSchf7(value float32) error {
	return e.Element.SetProperty("schf-7", value)
}

// sclc-0 (bool)
//
// Sidechain custom lo-cut 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc0(value bool) error {
	return e.Element.SetProperty("sclc-0", value)
}

// sclc-1 (bool)
//
// Sidechain custom lo-cut 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc1(value bool) error {
	return e.Element.SetProperty("sclc-1", value)
}

// sclc-2 (bool)
//
// Sidechain custom lo-cut 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc2(value bool) error {
	return e.Element.SetProperty("sclc-2", value)
}

// sclc-3 (bool)
//
// Sidechain custom lo-cut 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc3(value bool) error {
	return e.Element.SetProperty("sclc-3", value)
}

// sclc-4 (bool)
//
// Sidechain custom lo-cut 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc4(value bool) error {
	return e.Element.SetProperty("sclc-4", value)
}

// sclc-5 (bool)
//
// Sidechain custom lo-cut 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc5(value bool) error {
	return e.Element.SetProperty("sclc-5", value)
}

// sclc-6 (bool)
//
// Sidechain custom lo-cut 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc6(value bool) error {
	return e.Element.SetProperty("sclc-6", value)
}

// sclc-7 (bool)
//
// Sidechain custom lo-cut 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclc7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclc7(value bool) error {
	return e.Element.SetProperty("sclc-7", value)
}

// sclf-0 (float32)
//
// Sidechain lo-cut frequency 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf0(value float32) error {
	return e.Element.SetProperty("sclf-0", value)
}

// sclf-1 (float32)
//
// Sidechain lo-cut frequency 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf1(value float32) error {
	return e.Element.SetProperty("sclf-1", value)
}

// sclf-2 (float32)
//
// Sidechain lo-cut frequency 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf2(value float32) error {
	return e.Element.SetProperty("sclf-2", value)
}

// sclf-3 (float32)
//
// Sidechain lo-cut frequency 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf3(value float32) error {
	return e.Element.SetProperty("sclf-3", value)
}

// sclf-4 (float32)
//
// Sidechain lo-cut frequency 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf4(value float32) error {
	return e.Element.SetProperty("sclf-4", value)
}

// sclf-5 (float32)
//
// Sidechain lo-cut frequency 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf5(value float32) error {
	return e.Element.SetProperty("sclf-5", value)
}

// sclf-6 (float32)
//
// Sidechain lo-cut frequency 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf6(value float32) error {
	return e.Element.SetProperty("sclf-6", value)
}

// sclf-7 (float32)
//
// Sidechain lo-cut frequency 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSclf7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSclf7(value float32) error {
	return e.Element.SetProperty("sclf-7", value)
}

// scm-0 (GstLspPlugInPluginsLv2MbCompressorStereoscm0)
//
// Sidechain mode 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm0() (interface{}, error) {
	return e.Element.GetProperty("scm-0")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm0(value interface{}) error {
	return e.Element.SetProperty("scm-0", value)
}

// scm-1 (GstLspPlugInPluginsLv2MbCompressorStereoscm1)
//
// Sidechain mode 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm1() (interface{}, error) {
	return e.Element.GetProperty("scm-1")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm1(value interface{}) error {
	return e.Element.SetProperty("scm-1", value)
}

// scm-2 (GstLspPlugInPluginsLv2MbCompressorStereoscm2)
//
// Sidechain mode 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm2() (interface{}, error) {
	return e.Element.GetProperty("scm-2")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm2(value interface{}) error {
	return e.Element.SetProperty("scm-2", value)
}

// scm-3 (GstLspPlugInPluginsLv2MbCompressorStereoscm3)
//
// Sidechain mode 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm3() (interface{}, error) {
	return e.Element.GetProperty("scm-3")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm3(value interface{}) error {
	return e.Element.SetProperty("scm-3", value)
}

// scm-4 (GstLspPlugInPluginsLv2MbCompressorStereoscm4)
//
// Sidechain mode 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm4() (interface{}, error) {
	return e.Element.GetProperty("scm-4")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm4(value interface{}) error {
	return e.Element.SetProperty("scm-4", value)
}

// scm-5 (GstLspPlugInPluginsLv2MbCompressorStereoscm5)
//
// Sidechain mode 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm5() (interface{}, error) {
	return e.Element.GetProperty("scm-5")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm5(value interface{}) error {
	return e.Element.SetProperty("scm-5", value)
}

// scm-6 (GstLspPlugInPluginsLv2MbCompressorStereoscm6)
//
// Sidechain mode 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm6() (interface{}, error) {
	return e.Element.GetProperty("scm-6")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm6(value interface{}) error {
	return e.Element.SetProperty("scm-6", value)
}

// scm-7 (GstLspPlugInPluginsLv2MbCompressorStereoscm7)
//
// Sidechain mode 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScm7() (interface{}, error) {
	return e.Element.GetProperty("scm-7")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScm7(value interface{}) error {
	return e.Element.SetProperty("scm-7", value)
}

// scp-0 (float32)
//
// Sidechain preamp 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp0(value float32) error {
	return e.Element.SetProperty("scp-0", value)
}

// scp-1 (float32)
//
// Sidechain preamp 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp1(value float32) error {
	return e.Element.SetProperty("scp-1", value)
}

// scp-2 (float32)
//
// Sidechain preamp 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp2(value float32) error {
	return e.Element.SetProperty("scp-2", value)
}

// scp-3 (float32)
//
// Sidechain preamp 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp3(value float32) error {
	return e.Element.SetProperty("scp-3", value)
}

// scp-4 (float32)
//
// Sidechain preamp 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp4(value float32) error {
	return e.Element.SetProperty("scp-4", value)
}

// scp-5 (float32)
//
// Sidechain preamp 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp5(value float32) error {
	return e.Element.SetProperty("scp-5", value)
}

// scp-6 (float32)
//
// Sidechain preamp 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp6(value float32) error {
	return e.Element.SetProperty("scp-6", value)
}

// scp-7 (float32)
//
// Sidechain preamp 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScp7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScp7(value float32) error {
	return e.Element.SetProperty("scp-7", value)
}

// scr-0 (float32)
//
// Sidechain reactivity 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr0(value float32) error {
	return e.Element.SetProperty("scr-0", value)
}

// scr-1 (float32)
//
// Sidechain reactivity 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr1(value float32) error {
	return e.Element.SetProperty("scr-1", value)
}

// scr-2 (float32)
//
// Sidechain reactivity 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr2(value float32) error {
	return e.Element.SetProperty("scr-2", value)
}

// scr-3 (float32)
//
// Sidechain reactivity 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr3(value float32) error {
	return e.Element.SetProperty("scr-3", value)
}

// scr-4 (float32)
//
// Sidechain reactivity 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr4(value float32) error {
	return e.Element.SetProperty("scr-4", value)
}

// scr-5 (float32)
//
// Sidechain reactivity 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr5(value float32) error {
	return e.Element.SetProperty("scr-5", value)
}

// scr-6 (float32)
//
// Sidechain reactivity 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr6(value float32) error {
	return e.Element.SetProperty("scr-6", value)
}

// scr-7 (float32)
//
// Sidechain reactivity 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScr7(value float32) error {
	return e.Element.SetProperty("scr-7", value)
}

// scs-0 (GstLspPlugInPluginsLv2MbCompressorStereoscs0)
//
// Sidechain source 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs0() (interface{}, error) {
	return e.Element.GetProperty("scs-0")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs0(value interface{}) error {
	return e.Element.SetProperty("scs-0", value)
}

// scs-1 (GstLspPlugInPluginsLv2MbCompressorStereoscs1)
//
// Sidechain source 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs1() (interface{}, error) {
	return e.Element.GetProperty("scs-1")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs1(value interface{}) error {
	return e.Element.SetProperty("scs-1", value)
}

// scs-2 (GstLspPlugInPluginsLv2MbCompressorStereoscs2)
//
// Sidechain source 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs2() (interface{}, error) {
	return e.Element.GetProperty("scs-2")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs2(value interface{}) error {
	return e.Element.SetProperty("scs-2", value)
}

// scs-3 (GstLspPlugInPluginsLv2MbCompressorStereoscs3)
//
// Sidechain source 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs3() (interface{}, error) {
	return e.Element.GetProperty("scs-3")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs3(value interface{}) error {
	return e.Element.SetProperty("scs-3", value)
}

// scs-4 (GstLspPlugInPluginsLv2MbCompressorStereoscs4)
//
// Sidechain source 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs4() (interface{}, error) {
	return e.Element.GetProperty("scs-4")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs4(value interface{}) error {
	return e.Element.SetProperty("scs-4", value)
}

// scs-5 (GstLspPlugInPluginsLv2MbCompressorStereoscs5)
//
// Sidechain source 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs5() (interface{}, error) {
	return e.Element.GetProperty("scs-5")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs5(value interface{}) error {
	return e.Element.SetProperty("scs-5", value)
}

// scs-6 (GstLspPlugInPluginsLv2MbCompressorStereoscs6)
//
// Sidechain source 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs6() (interface{}, error) {
	return e.Element.GetProperty("scs-6")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs6(value interface{}) error {
	return e.Element.SetProperty("scs-6", value)
}

// scs-7 (GstLspPlugInPluginsLv2MbCompressorStereoscs7)
//
// Sidechain source 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetScs7() (interface{}, error) {
	return e.Element.GetProperty("scs-7")
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetScs7(value interface{}) error {
	return e.Element.SetProperty("scs-7", value)
}

// sf-1 (float32)
//
// Split frequency 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf1(value float32) error {
	return e.Element.SetProperty("sf-1", value)
}

// sf-2 (float32)
//
// Split frequency 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf2(value float32) error {
	return e.Element.SetProperty("sf-2", value)
}

// sf-3 (float32)
//
// Split frequency 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf3(value float32) error {
	return e.Element.SetProperty("sf-3", value)
}

// sf-4 (float32)
//
// Split frequency 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf4(value float32) error {
	return e.Element.SetProperty("sf-4", value)
}

// sf-5 (float32)
//
// Split frequency 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf5(value float32) error {
	return e.Element.SetProperty("sf-5", value)
}

// sf-6 (float32)
//
// Split frequency 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf6(value float32) error {
	return e.Element.SetProperty("sf-6", value)
}

// sf-7 (float32)
//
// Split frequency 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSf7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSf7(value float32) error {
	return e.Element.SetProperty("sf-7", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sla-0 (float32)
//
// Sidechain lookahead 0

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla0(value float32) error {
	return e.Element.SetProperty("sla-0", value)
}

// sla-1 (float32)
//
// Sidechain lookahead 1

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla1(value float32) error {
	return e.Element.SetProperty("sla-1", value)
}

// sla-2 (float32)
//
// Sidechain lookahead 2

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla2(value float32) error {
	return e.Element.SetProperty("sla-2", value)
}

// sla-3 (float32)
//
// Sidechain lookahead 3

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla3(value float32) error {
	return e.Element.SetProperty("sla-3", value)
}

// sla-4 (float32)
//
// Sidechain lookahead 4

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla4(value float32) error {
	return e.Element.SetProperty("sla-4", value)
}

// sla-5 (float32)
//
// Sidechain lookahead 5

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla5(value float32) error {
	return e.Element.SetProperty("sla-5", value)
}

// sla-6 (float32)
//
// Sidechain lookahead 6

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla6(value float32) error {
	return e.Element.SetProperty("sla-6", value)
}

// sla-7 (float32)
//
// Sidechain lookahead 7

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetSla7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetSla7(value float32) error {
	return e.Element.SetProperty("sla-7", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2MbCompressorStereo) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorStereo) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MbCompressorStereoscs6 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs6Mid LspPlugInPluginsLv2MbCompressorStereoscs6 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs6Side LspPlugInPluginsLv2MbCompressorStereoscs6 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs6Left LspPlugInPluginsLv2MbCompressorStereoscs6 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs6Right LspPlugInPluginsLv2MbCompressorStereoscs6 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereocm1 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm1Down LspPlugInPluginsLv2MbCompressorStereocm1 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm1Up LspPlugInPluginsLv2MbCompressorStereocm1 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereoscm2 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm2Peak LspPlugInPluginsLv2MbCompressorStereoscm2 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm2Rms LspPlugInPluginsLv2MbCompressorStereoscm2 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm2Lpf LspPlugInPluginsLv2MbCompressorStereoscm2 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm2Uniform LspPlugInPluginsLv2MbCompressorStereoscm2 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscm7 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm7Peak LspPlugInPluginsLv2MbCompressorStereoscm7 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm7Rms LspPlugInPluginsLv2MbCompressorStereoscm7 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm7Lpf LspPlugInPluginsLv2MbCompressorStereoscm7 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm7Uniform LspPlugInPluginsLv2MbCompressorStereoscm7 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscs5 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs5Mid LspPlugInPluginsLv2MbCompressorStereoscs5 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs5Side LspPlugInPluginsLv2MbCompressorStereoscs5 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs5Left LspPlugInPluginsLv2MbCompressorStereoscs5 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs5Right LspPlugInPluginsLv2MbCompressorStereoscs5 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereoscm0 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm0Peak LspPlugInPluginsLv2MbCompressorStereoscm0 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm0Rms LspPlugInPluginsLv2MbCompressorStereoscm0 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm0Lpf LspPlugInPluginsLv2MbCompressorStereoscm0 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm0Uniform LspPlugInPluginsLv2MbCompressorStereoscm0 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscs3 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs3Mid LspPlugInPluginsLv2MbCompressorStereoscs3 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs3Side LspPlugInPluginsLv2MbCompressorStereoscs3 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs3Left LspPlugInPluginsLv2MbCompressorStereoscs3 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs3Right LspPlugInPluginsLv2MbCompressorStereoscs3 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereoscs1 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs1Mid LspPlugInPluginsLv2MbCompressorStereoscs1 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs1Side LspPlugInPluginsLv2MbCompressorStereoscs1 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs1Left LspPlugInPluginsLv2MbCompressorStereoscs1 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs1Right LspPlugInPluginsLv2MbCompressorStereoscs1 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereoscs7 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs7Mid LspPlugInPluginsLv2MbCompressorStereoscs7 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs7Side LspPlugInPluginsLv2MbCompressorStereoscs7 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs7Left LspPlugInPluginsLv2MbCompressorStereoscs7 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs7Right LspPlugInPluginsLv2MbCompressorStereoscs7 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereobsel string

const (
	LspPlugInPluginsLv2MbCompressorStereobselSplit LspPlugInPluginsLv2MbCompressorStereobsel = "Split" // Split (0) – Split
	LspPlugInPluginsLv2MbCompressorStereobselBand0 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 0" // Band 0 (1) – Band 0
	LspPlugInPluginsLv2MbCompressorStereobselBand1 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 1" // Band 1 (2) – Band 1
	LspPlugInPluginsLv2MbCompressorStereobselBand2 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 2" // Band 2 (3) – Band 2
	LspPlugInPluginsLv2MbCompressorStereobselBand3 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 3" // Band 3 (4) – Band 3
	LspPlugInPluginsLv2MbCompressorStereobselBand4 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 4" // Band 4 (5) – Band 4
	LspPlugInPluginsLv2MbCompressorStereobselBand5 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 5" // Band 5 (6) – Band 5
	LspPlugInPluginsLv2MbCompressorStereobselBand6 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 6" // Band 6 (7) – Band 6
	LspPlugInPluginsLv2MbCompressorStereobselBand7 LspPlugInPluginsLv2MbCompressorStereobsel = "Band 7" // Band 7 (8) – Band 7
)

type LspPlugInPluginsLv2MbCompressorStereocm5 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm5Down LspPlugInPluginsLv2MbCompressorStereocm5 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm5Up LspPlugInPluginsLv2MbCompressorStereocm5 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereoscm5 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm5Peak LspPlugInPluginsLv2MbCompressorStereoscm5 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm5Rms LspPlugInPluginsLv2MbCompressorStereoscm5 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm5Lpf LspPlugInPluginsLv2MbCompressorStereoscm5 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm5Uniform LspPlugInPluginsLv2MbCompressorStereoscm5 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscm1 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm1Peak LspPlugInPluginsLv2MbCompressorStereoscm1 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm1Rms LspPlugInPluginsLv2MbCompressorStereoscm1 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm1Lpf LspPlugInPluginsLv2MbCompressorStereoscm1 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm1Uniform LspPlugInPluginsLv2MbCompressorStereoscm1 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscm6 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm6Peak LspPlugInPluginsLv2MbCompressorStereoscm6 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm6Rms LspPlugInPluginsLv2MbCompressorStereoscm6 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm6Lpf LspPlugInPluginsLv2MbCompressorStereoscm6 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm6Uniform LspPlugInPluginsLv2MbCompressorStereoscm6 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscs0 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs0Mid LspPlugInPluginsLv2MbCompressorStereoscs0 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs0Side LspPlugInPluginsLv2MbCompressorStereoscs0 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs0Left LspPlugInPluginsLv2MbCompressorStereoscs0 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs0Right LspPlugInPluginsLv2MbCompressorStereoscs0 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereocm7 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm7Down LspPlugInPluginsLv2MbCompressorStereocm7 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm7Up LspPlugInPluginsLv2MbCompressorStereocm7 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereomode string

const (
	LspPlugInPluginsLv2MbCompressorStereomodeClassic LspPlugInPluginsLv2MbCompressorStereomode = "Classic" // Classic (0) – Classic
	LspPlugInPluginsLv2MbCompressorStereomodeModern LspPlugInPluginsLv2MbCompressorStereomode = "Modern" // Modern (1) – Modern
)

type LspPlugInPluginsLv2MbCompressorStereoscm3 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm3Peak LspPlugInPluginsLv2MbCompressorStereoscm3 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm3Rms LspPlugInPluginsLv2MbCompressorStereoscm3 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm3Lpf LspPlugInPluginsLv2MbCompressorStereoscm3 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm3Uniform LspPlugInPluginsLv2MbCompressorStereoscm3 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscs4 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs4Mid LspPlugInPluginsLv2MbCompressorStereoscs4 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs4Side LspPlugInPluginsLv2MbCompressorStereoscs4 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs4Left LspPlugInPluginsLv2MbCompressorStereoscs4 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs4Right LspPlugInPluginsLv2MbCompressorStereoscs4 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereocm0 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm0Down LspPlugInPluginsLv2MbCompressorStereocm0 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm0Up LspPlugInPluginsLv2MbCompressorStereocm0 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereocm2 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm2Down LspPlugInPluginsLv2MbCompressorStereocm2 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm2Up LspPlugInPluginsLv2MbCompressorStereocm2 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereocm6 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm6Down LspPlugInPluginsLv2MbCompressorStereocm6 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm6Up LspPlugInPluginsLv2MbCompressorStereocm6 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereoscm4 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscm4Peak LspPlugInPluginsLv2MbCompressorStereoscm4 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorStereoscm4Rms LspPlugInPluginsLv2MbCompressorStereoscm4 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorStereoscm4Lpf LspPlugInPluginsLv2MbCompressorStereoscm4 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorStereoscm4Uniform LspPlugInPluginsLv2MbCompressorStereoscm4 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorStereoscs2 string

const (
	LspPlugInPluginsLv2MbCompressorStereoscs2Mid LspPlugInPluginsLv2MbCompressorStereoscs2 = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorStereoscs2Side LspPlugInPluginsLv2MbCompressorStereoscs2 = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorStereoscs2Left LspPlugInPluginsLv2MbCompressorStereoscs2 = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorStereoscs2Right LspPlugInPluginsLv2MbCompressorStereoscs2 = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorStereocm3 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm3Down LspPlugInPluginsLv2MbCompressorStereocm3 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm3Up LspPlugInPluginsLv2MbCompressorStereocm3 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereocm4 string

const (
	LspPlugInPluginsLv2MbCompressorStereocm4Down LspPlugInPluginsLv2MbCompressorStereocm4 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorStereocm4Up LspPlugInPluginsLv2MbCompressorStereocm4 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorStereoenvb string

const (
	LspPlugInPluginsLv2MbCompressorStereoenvbNone LspPlugInPluginsLv2MbCompressorStereoenvb = "None" // None (0) – None
	LspPlugInPluginsLv2MbCompressorStereoenvbPinkBt LspPlugInPluginsLv2MbCompressorStereoenvb = "Pink BT" // Pink BT (1) – Pink BT
	LspPlugInPluginsLv2MbCompressorStereoenvbPinkMt LspPlugInPluginsLv2MbCompressorStereoenvb = "Pink MT" // Pink MT (2) – Pink MT
	LspPlugInPluginsLv2MbCompressorStereoenvbBrownBt LspPlugInPluginsLv2MbCompressorStereoenvb = "Brown BT" // Brown BT (3) – Brown BT
	LspPlugInPluginsLv2MbCompressorStereoenvbBrownMt LspPlugInPluginsLv2MbCompressorStereoenvb = "Brown MT" // Brown MT (4) – Brown MT
)

