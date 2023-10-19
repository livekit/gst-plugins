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

type LspPlugInPluginsLv2MbCompressorMono struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MbCompressorMono() (*LspPlugInPluginsLv2MbCompressorMono, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-mb-compressor-mono")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MbCompressorMonoWithName(name string) (*LspPlugInPluginsLv2MbCompressorMono, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-mb-compressor-mono", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorMono{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-0 (float32)
//
// Attack level 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl0(value float32) error {
	return e.Element.SetProperty("al-0", value)
}

// al-1 (float32)
//
// Attack level 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl1(value float32) error {
	return e.Element.SetProperty("al-1", value)
}

// al-2 (float32)
//
// Attack level 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl2(value float32) error {
	return e.Element.SetProperty("al-2", value)
}

// al-3 (float32)
//
// Attack level 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl3(value float32) error {
	return e.Element.SetProperty("al-3", value)
}

// al-4 (float32)
//
// Attack level 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl4(value float32) error {
	return e.Element.SetProperty("al-4", value)
}

// al-5 (float32)
//
// Attack level 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl5(value float32) error {
	return e.Element.SetProperty("al-5", value)
}

// al-6 (float32)
//
// Attack level 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl6(value float32) error {
	return e.Element.SetProperty("al-6", value)
}

// al-7 (float32)
//
// Attack level 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAl7(value float32) error {
	return e.Element.SetProperty("al-7", value)
}

// at-0 (float32)
//
// Attack time 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt0(value float32) error {
	return e.Element.SetProperty("at-0", value)
}

// at-1 (float32)
//
// Attack time 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt1(value float32) error {
	return e.Element.SetProperty("at-1", value)
}

// at-2 (float32)
//
// Attack time 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt2(value float32) error {
	return e.Element.SetProperty("at-2", value)
}

// at-3 (float32)
//
// Attack time 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt3(value float32) error {
	return e.Element.SetProperty("at-3", value)
}

// at-4 (float32)
//
// Attack time 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt4(value float32) error {
	return e.Element.SetProperty("at-4", value)
}

// at-5 (float32)
//
// Attack time 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt5(value float32) error {
	return e.Element.SetProperty("at-5", value)
}

// at-6 (float32)
//
// Attack time 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt6(value float32) error {
	return e.Element.SetProperty("at-6", value)
}

// at-7 (float32)
//
// Attack time 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetAt7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetAt7(value float32) error {
	return e.Element.SetProperty("at-7", value)
}

// bm-0 (bool)
//
// Mute band 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm0(value bool) error {
	return e.Element.SetProperty("bm-0", value)
}

// bm-1 (bool)
//
// Mute band 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm1(value bool) error {
	return e.Element.SetProperty("bm-1", value)
}

// bm-2 (bool)
//
// Mute band 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm2(value bool) error {
	return e.Element.SetProperty("bm-2", value)
}

// bm-3 (bool)
//
// Mute band 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm3(value bool) error {
	return e.Element.SetProperty("bm-3", value)
}

// bm-4 (bool)
//
// Mute band 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm4(value bool) error {
	return e.Element.SetProperty("bm-4", value)
}

// bm-5 (bool)
//
// Mute band 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm5(value bool) error {
	return e.Element.SetProperty("bm-5", value)
}

// bm-6 (bool)
//
// Mute band 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm6(value bool) error {
	return e.Element.SetProperty("bm-6", value)
}

// bm-7 (bool)
//
// Mute band 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBm7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBm7(value bool) error {
	return e.Element.SetProperty("bm-7", value)
}

// bs-0 (bool)
//
// Solo band 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs0(value bool) error {
	return e.Element.SetProperty("bs-0", value)
}

// bs-1 (bool)
//
// Solo band 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs1(value bool) error {
	return e.Element.SetProperty("bs-1", value)
}

// bs-2 (bool)
//
// Solo band 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs2(value bool) error {
	return e.Element.SetProperty("bs-2", value)
}

// bs-3 (bool)
//
// Solo band 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs3(value bool) error {
	return e.Element.SetProperty("bs-3", value)
}

// bs-4 (bool)
//
// Solo band 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs4(value bool) error {
	return e.Element.SetProperty("bs-4", value)
}

// bs-5 (bool)
//
// Solo band 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs5(value bool) error {
	return e.Element.SetProperty("bs-5", value)
}

// bs-6 (bool)
//
// Solo band 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs6(value bool) error {
	return e.Element.SetProperty("bs-6", value)
}

// bs-7 (bool)
//
// Solo band 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBs7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBs7(value bool) error {
	return e.Element.SetProperty("bs-7", value)
}

// bsel (GstLspPlugInPluginsLv2MbCompressorMonobsel)
//
// Band selection

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBsel() (interface{}, error) {
	return e.Element.GetProperty("bsel")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBsel(value interface{}) error {
	return e.Element.SetProperty("bsel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MbCompressorMono) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cbe-1 (bool)
//
// Compression band enable 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe1(value bool) error {
	return e.Element.SetProperty("cbe-1", value)
}

// cbe-2 (bool)
//
// Compression band enable 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe2(value bool) error {
	return e.Element.SetProperty("cbe-2", value)
}

// cbe-3 (bool)
//
// Compression band enable 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe3(value bool) error {
	return e.Element.SetProperty("cbe-3", value)
}

// cbe-4 (bool)
//
// Compression band enable 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe4(value bool) error {
	return e.Element.SetProperty("cbe-4", value)
}

// cbe-5 (bool)
//
// Compression band enable 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe5(value bool) error {
	return e.Element.SetProperty("cbe-5", value)
}

// cbe-6 (bool)
//
// Compression band enable 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe6(value bool) error {
	return e.Element.SetProperty("cbe-6", value)
}

// cbe-7 (bool)
//
// Compression band enable 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCbe7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCbe7(value bool) error {
	return e.Element.SetProperty("cbe-7", value)
}

// ce-0 (bool)
//
// Compressor enable 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe0(value bool) error {
	return e.Element.SetProperty("ce-0", value)
}

// ce-1 (bool)
//
// Compressor enable 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe1(value bool) error {
	return e.Element.SetProperty("ce-1", value)
}

// ce-2 (bool)
//
// Compressor enable 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe2(value bool) error {
	return e.Element.SetProperty("ce-2", value)
}

// ce-3 (bool)
//
// Compressor enable 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe3(value bool) error {
	return e.Element.SetProperty("ce-3", value)
}

// ce-4 (bool)
//
// Compressor enable 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe4(value bool) error {
	return e.Element.SetProperty("ce-4", value)
}

// ce-5 (bool)
//
// Compressor enable 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe5(value bool) error {
	return e.Element.SetProperty("ce-5", value)
}

// ce-6 (bool)
//
// Compressor enable 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe6(value bool) error {
	return e.Element.SetProperty("ce-6", value)
}

// ce-7 (bool)
//
// Compressor enable 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCe7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCe7(value bool) error {
	return e.Element.SetProperty("ce-7", value)
}

// clm-0 (float32)
//
// Curve level meter 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetClm7() (float32, error) {
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

// cm-0 (GstLspPlugInPluginsLv2MbCompressorMonocm0)
//
// Compression mode 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm0() (interface{}, error) {
	return e.Element.GetProperty("cm-0")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm0(value interface{}) error {
	return e.Element.SetProperty("cm-0", value)
}

// cm-1 (GstLspPlugInPluginsLv2MbCompressorMonocm1)
//
// Compression mode 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm1() (interface{}, error) {
	return e.Element.GetProperty("cm-1")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm1(value interface{}) error {
	return e.Element.SetProperty("cm-1", value)
}

// cm-2 (GstLspPlugInPluginsLv2MbCompressorMonocm2)
//
// Compression mode 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm2() (interface{}, error) {
	return e.Element.GetProperty("cm-2")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm2(value interface{}) error {
	return e.Element.SetProperty("cm-2", value)
}

// cm-3 (GstLspPlugInPluginsLv2MbCompressorMonocm3)
//
// Compression mode 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm3() (interface{}, error) {
	return e.Element.GetProperty("cm-3")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm3(value interface{}) error {
	return e.Element.SetProperty("cm-3", value)
}

// cm-4 (GstLspPlugInPluginsLv2MbCompressorMonocm4)
//
// Compression mode 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm4() (interface{}, error) {
	return e.Element.GetProperty("cm-4")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm4(value interface{}) error {
	return e.Element.SetProperty("cm-4", value)
}

// cm-5 (GstLspPlugInPluginsLv2MbCompressorMonocm5)
//
// Compression mode 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm5() (interface{}, error) {
	return e.Element.GetProperty("cm-5")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm5(value interface{}) error {
	return e.Element.SetProperty("cm-5", value)
}

// cm-6 (GstLspPlugInPluginsLv2MbCompressorMonocm6)
//
// Compression mode 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm6() (interface{}, error) {
	return e.Element.GetProperty("cm-6")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm6(value interface{}) error {
	return e.Element.SetProperty("cm-6", value)
}

// cm-7 (GstLspPlugInPluginsLv2MbCompressorMonocm7)
//
// Compression mode 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCm7() (interface{}, error) {
	return e.Element.GetProperty("cm-7")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCm7(value interface{}) error {
	return e.Element.SetProperty("cm-7", value)
}

// cr-0 (float32)
//
// Ratio 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr0(value float32) error {
	return e.Element.SetProperty("cr-0", value)
}

// cr-1 (float32)
//
// Ratio 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr1(value float32) error {
	return e.Element.SetProperty("cr-1", value)
}

// cr-2 (float32)
//
// Ratio 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr2(value float32) error {
	return e.Element.SetProperty("cr-2", value)
}

// cr-3 (float32)
//
// Ratio 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr3(value float32) error {
	return e.Element.SetProperty("cr-3", value)
}

// cr-4 (float32)
//
// Ratio 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr4(value float32) error {
	return e.Element.SetProperty("cr-4", value)
}

// cr-5 (float32)
//
// Ratio 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr5(value float32) error {
	return e.Element.SetProperty("cr-5", value)
}

// cr-6 (float32)
//
// Ratio 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr6(value float32) error {
	return e.Element.SetProperty("cr-6", value)
}

// cr-7 (float32)
//
// Ratio 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetCr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetCr7(value float32) error {
	return e.Element.SetProperty("cr-7", value)
}

// elm-0 (float32)
//
// Envelope level meter 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetElm7() (float32, error) {
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

// envb (GstLspPlugInPluginsLv2MbCompressorMonoenvb)
//
// Envelope boost

func (e *LspPlugInPluginsLv2MbCompressorMono) GetEnvb() (interface{}, error) {
	return e.Element.GetProperty("envb")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetEnvb(value interface{}) error {
	return e.Element.SetProperty("envb", value)
}

// flt (bool)
//
// Band filter curves

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFlt() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetFlt(value bool) error {
	return e.Element.SetProperty("flt", value)
}

// fre-0 (float32)
//
// Frequency range end 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetFre7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetGDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetGDry(value float32) error {
	return e.Element.SetProperty("g-dry", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2MbCompressorMono) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MbCompressorMono) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// g-wet (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2MbCompressorMono) GetGWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetGWet(value float32) error {
	return e.Element.SetProperty("g-wet", value)
}

// hue-0 (float32)
//
// Hue  0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue0(value float32) error {
	return e.Element.SetProperty("hue-0", value)
}

// hue-1 (float32)
//
// Hue  1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue1(value float32) error {
	return e.Element.SetProperty("hue-1", value)
}

// hue-2 (float32)
//
// Hue  2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue2(value float32) error {
	return e.Element.SetProperty("hue-2", value)
}

// hue-3 (float32)
//
// Hue  3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue3(value float32) error {
	return e.Element.SetProperty("hue-3", value)
}

// hue-4 (float32)
//
// Hue  4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue4(value float32) error {
	return e.Element.SetProperty("hue-4", value)
}

// hue-5 (float32)
//
// Hue  5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue5(value float32) error {
	return e.Element.SetProperty("hue-5", value)
}

// hue-6 (float32)
//
// Hue  6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue6(value float32) error {
	return e.Element.SetProperty("hue-6", value)
}

// hue-7 (float32)
//
// Hue  7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetHue7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetHue7(value float32) error {
	return e.Element.SetProperty("hue-7", value)
}

// ife (bool)
//
// Input FFT graph enable

func (e *LspPlugInPluginsLv2MbCompressorMono) GetIfe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ife")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetIfe(value bool) error {
	return e.Element.SetProperty("ife", value)
}

// ilm (float32)
//
// Input level meter

func (e *LspPlugInPluginsLv2MbCompressorMono) GetIlm() (float32, error) {
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

// kn-0 (float32)
//
// Knee 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn0(value float32) error {
	return e.Element.SetProperty("kn-0", value)
}

// kn-1 (float32)
//
// Knee 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn1(value float32) error {
	return e.Element.SetProperty("kn-1", value)
}

// kn-2 (float32)
//
// Knee 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn2(value float32) error {
	return e.Element.SetProperty("kn-2", value)
}

// kn-3 (float32)
//
// Knee 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn3(value float32) error {
	return e.Element.SetProperty("kn-3", value)
}

// kn-4 (float32)
//
// Knee 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn4(value float32) error {
	return e.Element.SetProperty("kn-4", value)
}

// kn-5 (float32)
//
// Knee 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn5(value float32) error {
	return e.Element.SetProperty("kn-5", value)
}

// kn-6 (float32)
//
// Knee 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn6(value float32) error {
	return e.Element.SetProperty("kn-6", value)
}

// kn-7 (float32)
//
// Knee 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetKn7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetKn7(value float32) error {
	return e.Element.SetProperty("kn-7", value)
}

// mk-0 (float32)
//
// Makeup gain 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk0(value float32) error {
	return e.Element.SetProperty("mk-0", value)
}

// mk-1 (float32)
//
// Makeup gain 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk1(value float32) error {
	return e.Element.SetProperty("mk-1", value)
}

// mk-2 (float32)
//
// Makeup gain 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk2(value float32) error {
	return e.Element.SetProperty("mk-2", value)
}

// mk-3 (float32)
//
// Makeup gain 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk3(value float32) error {
	return e.Element.SetProperty("mk-3", value)
}

// mk-4 (float32)
//
// Makeup gain 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk4(value float32) error {
	return e.Element.SetProperty("mk-4", value)
}

// mk-5 (float32)
//
// Makeup gain 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk5(value float32) error {
	return e.Element.SetProperty("mk-5", value)
}

// mk-6 (float32)
//
// Makeup gain 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk6(value float32) error {
	return e.Element.SetProperty("mk-6", value)
}

// mk-7 (float32)
//
// Makeup gain 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMk7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMk7(value float32) error {
	return e.Element.SetProperty("mk-7", value)
}

// mode (GstLspPlugInPluginsLv2MbCompressorMonomode)
//
// Compressor mode

func (e *LspPlugInPluginsLv2MbCompressorMono) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ofe (bool)
//
// Output FFT graph enable

func (e *LspPlugInPluginsLv2MbCompressorMono) GetOfe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ofe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetOfe(value bool) error {
	return e.Element.SetProperty("ofe", value)
}

// olm (float32)
//
// Output level meter

func (e *LspPlugInPluginsLv2MbCompressorMono) GetOlm() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl-0 (float32)
//
// Release level 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRlm7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl0(value float32) error {
	return e.Element.SetProperty("rrl-0", value)
}

// rrl-1 (float32)
//
// Relative release level 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl1(value float32) error {
	return e.Element.SetProperty("rrl-1", value)
}

// rrl-2 (float32)
//
// Relative release level 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl2(value float32) error {
	return e.Element.SetProperty("rrl-2", value)
}

// rrl-3 (float32)
//
// Relative release level 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl3(value float32) error {
	return e.Element.SetProperty("rrl-3", value)
}

// rrl-4 (float32)
//
// Relative release level 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl4(value float32) error {
	return e.Element.SetProperty("rrl-4", value)
}

// rrl-5 (float32)
//
// Relative release level 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl5(value float32) error {
	return e.Element.SetProperty("rrl-5", value)
}

// rrl-6 (float32)
//
// Relative release level 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl6(value float32) error {
	return e.Element.SetProperty("rrl-6", value)
}

// rrl-7 (float32)
//
// Relative release level 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRrl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRrl7(value float32) error {
	return e.Element.SetProperty("rrl-7", value)
}

// rt-0 (float32)
//
// Release time 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt0(value float32) error {
	return e.Element.SetProperty("rt-0", value)
}

// rt-1 (float32)
//
// Release time 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt1(value float32) error {
	return e.Element.SetProperty("rt-1", value)
}

// rt-2 (float32)
//
// Release time 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt2(value float32) error {
	return e.Element.SetProperty("rt-2", value)
}

// rt-3 (float32)
//
// Release time 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt3(value float32) error {
	return e.Element.SetProperty("rt-3", value)
}

// rt-4 (float32)
//
// Release time 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt4(value float32) error {
	return e.Element.SetProperty("rt-4", value)
}

// rt-5 (float32)
//
// Release time 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt5(value float32) error {
	return e.Element.SetProperty("rt-5", value)
}

// rt-6 (float32)
//
// Release time 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt6(value float32) error {
	return e.Element.SetProperty("rt-6", value)
}

// rt-7 (float32)
//
// Release time 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetRt7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetRt7(value float32) error {
	return e.Element.SetProperty("rt-7", value)
}

// schc-0 (bool)
//
// Sidechain custom hi-cut 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc0(value bool) error {
	return e.Element.SetProperty("schc-0", value)
}

// schc-1 (bool)
//
// Sidechain custom hi-cut 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc1(value bool) error {
	return e.Element.SetProperty("schc-1", value)
}

// schc-2 (bool)
//
// Sidechain custom hi-cut 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc2(value bool) error {
	return e.Element.SetProperty("schc-2", value)
}

// schc-3 (bool)
//
// Sidechain custom hi-cut 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc3(value bool) error {
	return e.Element.SetProperty("schc-3", value)
}

// schc-4 (bool)
//
// Sidechain custom hi-cut 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc4(value bool) error {
	return e.Element.SetProperty("schc-4", value)
}

// schc-5 (bool)
//
// Sidechain custom hi-cut 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc5(value bool) error {
	return e.Element.SetProperty("schc-5", value)
}

// schc-6 (bool)
//
// Sidechain custom hi-cut 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc6(value bool) error {
	return e.Element.SetProperty("schc-6", value)
}

// schc-7 (bool)
//
// Sidechain custom hi-cut 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchc7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchc7(value bool) error {
	return e.Element.SetProperty("schc-7", value)
}

// schf-0 (float32)
//
// Sidechain hi-cut frequency 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf0(value float32) error {
	return e.Element.SetProperty("schf-0", value)
}

// schf-1 (float32)
//
// Sidechain hi-cut frequency 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf1(value float32) error {
	return e.Element.SetProperty("schf-1", value)
}

// schf-2 (float32)
//
// Sidechain hi-cut frequency 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf2(value float32) error {
	return e.Element.SetProperty("schf-2", value)
}

// schf-3 (float32)
//
// Sidechain hi-cut frequency 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf3(value float32) error {
	return e.Element.SetProperty("schf-3", value)
}

// schf-4 (float32)
//
// Sidechain hi-cut frequency 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf4(value float32) error {
	return e.Element.SetProperty("schf-4", value)
}

// schf-5 (float32)
//
// Sidechain hi-cut frequency 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf5(value float32) error {
	return e.Element.SetProperty("schf-5", value)
}

// schf-6 (float32)
//
// Sidechain hi-cut frequency 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf6(value float32) error {
	return e.Element.SetProperty("schf-6", value)
}

// schf-7 (float32)
//
// Sidechain hi-cut frequency 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSchf7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSchf7(value float32) error {
	return e.Element.SetProperty("schf-7", value)
}

// sclc-0 (bool)
//
// Sidechain custom lo-cut 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc0(value bool) error {
	return e.Element.SetProperty("sclc-0", value)
}

// sclc-1 (bool)
//
// Sidechain custom lo-cut 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc1(value bool) error {
	return e.Element.SetProperty("sclc-1", value)
}

// sclc-2 (bool)
//
// Sidechain custom lo-cut 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc2(value bool) error {
	return e.Element.SetProperty("sclc-2", value)
}

// sclc-3 (bool)
//
// Sidechain custom lo-cut 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc3(value bool) error {
	return e.Element.SetProperty("sclc-3", value)
}

// sclc-4 (bool)
//
// Sidechain custom lo-cut 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc4(value bool) error {
	return e.Element.SetProperty("sclc-4", value)
}

// sclc-5 (bool)
//
// Sidechain custom lo-cut 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc5(value bool) error {
	return e.Element.SetProperty("sclc-5", value)
}

// sclc-6 (bool)
//
// Sidechain custom lo-cut 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc6(value bool) error {
	return e.Element.SetProperty("sclc-6", value)
}

// sclc-7 (bool)
//
// Sidechain custom lo-cut 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclc7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclc7(value bool) error {
	return e.Element.SetProperty("sclc-7", value)
}

// sclf-0 (float32)
//
// Sidechain lo-cut frequency 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf0(value float32) error {
	return e.Element.SetProperty("sclf-0", value)
}

// sclf-1 (float32)
//
// Sidechain lo-cut frequency 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf1(value float32) error {
	return e.Element.SetProperty("sclf-1", value)
}

// sclf-2 (float32)
//
// Sidechain lo-cut frequency 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf2(value float32) error {
	return e.Element.SetProperty("sclf-2", value)
}

// sclf-3 (float32)
//
// Sidechain lo-cut frequency 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf3(value float32) error {
	return e.Element.SetProperty("sclf-3", value)
}

// sclf-4 (float32)
//
// Sidechain lo-cut frequency 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf4(value float32) error {
	return e.Element.SetProperty("sclf-4", value)
}

// sclf-5 (float32)
//
// Sidechain lo-cut frequency 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf5(value float32) error {
	return e.Element.SetProperty("sclf-5", value)
}

// sclf-6 (float32)
//
// Sidechain lo-cut frequency 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf6(value float32) error {
	return e.Element.SetProperty("sclf-6", value)
}

// sclf-7 (float32)
//
// Sidechain lo-cut frequency 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSclf7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSclf7(value float32) error {
	return e.Element.SetProperty("sclf-7", value)
}

// scm-0 (GstLspPlugInPluginsLv2MbCompressorMonoscm0)
//
// Sidechain mode 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm0() (interface{}, error) {
	return e.Element.GetProperty("scm-0")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm0(value interface{}) error {
	return e.Element.SetProperty("scm-0", value)
}

// scm-1 (GstLspPlugInPluginsLv2MbCompressorMonoscm1)
//
// Sidechain mode 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm1() (interface{}, error) {
	return e.Element.GetProperty("scm-1")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm1(value interface{}) error {
	return e.Element.SetProperty("scm-1", value)
}

// scm-2 (GstLspPlugInPluginsLv2MbCompressorMonoscm2)
//
// Sidechain mode 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm2() (interface{}, error) {
	return e.Element.GetProperty("scm-2")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm2(value interface{}) error {
	return e.Element.SetProperty("scm-2", value)
}

// scm-3 (GstLspPlugInPluginsLv2MbCompressorMonoscm3)
//
// Sidechain mode 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm3() (interface{}, error) {
	return e.Element.GetProperty("scm-3")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm3(value interface{}) error {
	return e.Element.SetProperty("scm-3", value)
}

// scm-4 (GstLspPlugInPluginsLv2MbCompressorMonoscm4)
//
// Sidechain mode 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm4() (interface{}, error) {
	return e.Element.GetProperty("scm-4")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm4(value interface{}) error {
	return e.Element.SetProperty("scm-4", value)
}

// scm-5 (GstLspPlugInPluginsLv2MbCompressorMonoscm5)
//
// Sidechain mode 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm5() (interface{}, error) {
	return e.Element.GetProperty("scm-5")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm5(value interface{}) error {
	return e.Element.SetProperty("scm-5", value)
}

// scm-6 (GstLspPlugInPluginsLv2MbCompressorMonoscm6)
//
// Sidechain mode 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm6() (interface{}, error) {
	return e.Element.GetProperty("scm-6")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm6(value interface{}) error {
	return e.Element.SetProperty("scm-6", value)
}

// scm-7 (GstLspPlugInPluginsLv2MbCompressorMonoscm7)
//
// Sidechain mode 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScm7() (interface{}, error) {
	return e.Element.GetProperty("scm-7")
}

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScm7(value interface{}) error {
	return e.Element.SetProperty("scm-7", value)
}

// scp-0 (float32)
//
// Sidechain preamp 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp0(value float32) error {
	return e.Element.SetProperty("scp-0", value)
}

// scp-1 (float32)
//
// Sidechain preamp 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp1(value float32) error {
	return e.Element.SetProperty("scp-1", value)
}

// scp-2 (float32)
//
// Sidechain preamp 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp2(value float32) error {
	return e.Element.SetProperty("scp-2", value)
}

// scp-3 (float32)
//
// Sidechain preamp 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp3(value float32) error {
	return e.Element.SetProperty("scp-3", value)
}

// scp-4 (float32)
//
// Sidechain preamp 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp4(value float32) error {
	return e.Element.SetProperty("scp-4", value)
}

// scp-5 (float32)
//
// Sidechain preamp 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp5(value float32) error {
	return e.Element.SetProperty("scp-5", value)
}

// scp-6 (float32)
//
// Sidechain preamp 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp6(value float32) error {
	return e.Element.SetProperty("scp-6", value)
}

// scp-7 (float32)
//
// Sidechain preamp 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScp7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScp7(value float32) error {
	return e.Element.SetProperty("scp-7", value)
}

// scr-0 (float32)
//
// Sidechain reactivity 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr0(value float32) error {
	return e.Element.SetProperty("scr-0", value)
}

// scr-1 (float32)
//
// Sidechain reactivity 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr1(value float32) error {
	return e.Element.SetProperty("scr-1", value)
}

// scr-2 (float32)
//
// Sidechain reactivity 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr2(value float32) error {
	return e.Element.SetProperty("scr-2", value)
}

// scr-3 (float32)
//
// Sidechain reactivity 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr3(value float32) error {
	return e.Element.SetProperty("scr-3", value)
}

// scr-4 (float32)
//
// Sidechain reactivity 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr4(value float32) error {
	return e.Element.SetProperty("scr-4", value)
}

// scr-5 (float32)
//
// Sidechain reactivity 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr5(value float32) error {
	return e.Element.SetProperty("scr-5", value)
}

// scr-6 (float32)
//
// Sidechain reactivity 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr6(value float32) error {
	return e.Element.SetProperty("scr-6", value)
}

// scr-7 (float32)
//
// Sidechain reactivity 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetScr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetScr7(value float32) error {
	return e.Element.SetProperty("scr-7", value)
}

// sf-1 (float32)
//
// Split frequency 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf1(value float32) error {
	return e.Element.SetProperty("sf-1", value)
}

// sf-2 (float32)
//
// Split frequency 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf2(value float32) error {
	return e.Element.SetProperty("sf-2", value)
}

// sf-3 (float32)
//
// Split frequency 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf3(value float32) error {
	return e.Element.SetProperty("sf-3", value)
}

// sf-4 (float32)
//
// Split frequency 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf4(value float32) error {
	return e.Element.SetProperty("sf-4", value)
}

// sf-5 (float32)
//
// Split frequency 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf5(value float32) error {
	return e.Element.SetProperty("sf-5", value)
}

// sf-6 (float32)
//
// Split frequency 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf6(value float32) error {
	return e.Element.SetProperty("sf-6", value)
}

// sf-7 (float32)
//
// Split frequency 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSf7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSf7(value float32) error {
	return e.Element.SetProperty("sf-7", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2MbCompressorMono) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sla-0 (float32)
//
// Sidechain lookahead 0

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla0(value float32) error {
	return e.Element.SetProperty("sla-0", value)
}

// sla-1 (float32)
//
// Sidechain lookahead 1

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla1(value float32) error {
	return e.Element.SetProperty("sla-1", value)
}

// sla-2 (float32)
//
// Sidechain lookahead 2

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla2(value float32) error {
	return e.Element.SetProperty("sla-2", value)
}

// sla-3 (float32)
//
// Sidechain lookahead 3

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla3(value float32) error {
	return e.Element.SetProperty("sla-3", value)
}

// sla-4 (float32)
//
// Sidechain lookahead 4

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla4(value float32) error {
	return e.Element.SetProperty("sla-4", value)
}

// sla-5 (float32)
//
// Sidechain lookahead 5

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla5(value float32) error {
	return e.Element.SetProperty("sla-5", value)
}

// sla-6 (float32)
//
// Sidechain lookahead 6

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla6(value float32) error {
	return e.Element.SetProperty("sla-6", value)
}

// sla-7 (float32)
//
// Sidechain lookahead 7

func (e *LspPlugInPluginsLv2MbCompressorMono) GetSla7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetSla7(value float32) error {
	return e.Element.SetProperty("sla-7", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2MbCompressorMono) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorMono) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MbCompressorMonocm1 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm1Down LspPlugInPluginsLv2MbCompressorMonocm1 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm1Up LspPlugInPluginsLv2MbCompressorMonocm1 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonoenvb string

const (
	LspPlugInPluginsLv2MbCompressorMonoenvbNone LspPlugInPluginsLv2MbCompressorMonoenvb = "None" // None (0) – None
	LspPlugInPluginsLv2MbCompressorMonoenvbPinkBt LspPlugInPluginsLv2MbCompressorMonoenvb = "Pink BT" // Pink BT (1) – Pink BT
	LspPlugInPluginsLv2MbCompressorMonoenvbPinkMt LspPlugInPluginsLv2MbCompressorMonoenvb = "Pink MT" // Pink MT (2) – Pink MT
	LspPlugInPluginsLv2MbCompressorMonoenvbBrownBt LspPlugInPluginsLv2MbCompressorMonoenvb = "Brown BT" // Brown BT (3) – Brown BT
	LspPlugInPluginsLv2MbCompressorMonoenvbBrownMt LspPlugInPluginsLv2MbCompressorMonoenvb = "Brown MT" // Brown MT (4) – Brown MT
)

type LspPlugInPluginsLv2MbCompressorMonoscm1 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm1Peak LspPlugInPluginsLv2MbCompressorMonoscm1 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm1Rms LspPlugInPluginsLv2MbCompressorMonoscm1 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm1Lpf LspPlugInPluginsLv2MbCompressorMonoscm1 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm1Uniform LspPlugInPluginsLv2MbCompressorMonoscm1 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonoscm3 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm3Peak LspPlugInPluginsLv2MbCompressorMonoscm3 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm3Rms LspPlugInPluginsLv2MbCompressorMonoscm3 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm3Lpf LspPlugInPluginsLv2MbCompressorMonoscm3 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm3Uniform LspPlugInPluginsLv2MbCompressorMonoscm3 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonoscm5 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm5Peak LspPlugInPluginsLv2MbCompressorMonoscm5 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm5Rms LspPlugInPluginsLv2MbCompressorMonoscm5 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm5Lpf LspPlugInPluginsLv2MbCompressorMonoscm5 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm5Uniform LspPlugInPluginsLv2MbCompressorMonoscm5 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonoscm0 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm0Peak LspPlugInPluginsLv2MbCompressorMonoscm0 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm0Rms LspPlugInPluginsLv2MbCompressorMonoscm0 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm0Lpf LspPlugInPluginsLv2MbCompressorMonoscm0 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm0Uniform LspPlugInPluginsLv2MbCompressorMonoscm0 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonoscm2 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm2Peak LspPlugInPluginsLv2MbCompressorMonoscm2 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm2Rms LspPlugInPluginsLv2MbCompressorMonoscm2 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm2Lpf LspPlugInPluginsLv2MbCompressorMonoscm2 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm2Uniform LspPlugInPluginsLv2MbCompressorMonoscm2 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonobsel string

const (
	LspPlugInPluginsLv2MbCompressorMonobselSplit LspPlugInPluginsLv2MbCompressorMonobsel = "Split" // Split (0) – Split
	LspPlugInPluginsLv2MbCompressorMonobselBand0 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 0" // Band 0 (1) – Band 0
	LspPlugInPluginsLv2MbCompressorMonobselBand1 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 1" // Band 1 (2) – Band 1
	LspPlugInPluginsLv2MbCompressorMonobselBand2 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 2" // Band 2 (3) – Band 2
	LspPlugInPluginsLv2MbCompressorMonobselBand3 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 3" // Band 3 (4) – Band 3
	LspPlugInPluginsLv2MbCompressorMonobselBand4 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 4" // Band 4 (5) – Band 4
	LspPlugInPluginsLv2MbCompressorMonobselBand5 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 5" // Band 5 (6) – Band 5
	LspPlugInPluginsLv2MbCompressorMonobselBand6 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 6" // Band 6 (7) – Band 6
	LspPlugInPluginsLv2MbCompressorMonobselBand7 LspPlugInPluginsLv2MbCompressorMonobsel = "Band 7" // Band 7 (8) – Band 7
)

type LspPlugInPluginsLv2MbCompressorMonocm4 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm4Down LspPlugInPluginsLv2MbCompressorMonocm4 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm4Up LspPlugInPluginsLv2MbCompressorMonocm4 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonocm7 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm7Down LspPlugInPluginsLv2MbCompressorMonocm7 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm7Up LspPlugInPluginsLv2MbCompressorMonocm7 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonoscm4 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm4Peak LspPlugInPluginsLv2MbCompressorMonoscm4 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm4Rms LspPlugInPluginsLv2MbCompressorMonoscm4 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm4Lpf LspPlugInPluginsLv2MbCompressorMonoscm4 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm4Uniform LspPlugInPluginsLv2MbCompressorMonoscm4 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonoscm6 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm6Peak LspPlugInPluginsLv2MbCompressorMonoscm6 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm6Rms LspPlugInPluginsLv2MbCompressorMonoscm6 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm6Lpf LspPlugInPluginsLv2MbCompressorMonoscm6 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm6Uniform LspPlugInPluginsLv2MbCompressorMonoscm6 = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorMonocm0 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm0Down LspPlugInPluginsLv2MbCompressorMonocm0 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm0Up LspPlugInPluginsLv2MbCompressorMonocm0 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonocm2 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm2Down LspPlugInPluginsLv2MbCompressorMonocm2 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm2Up LspPlugInPluginsLv2MbCompressorMonocm2 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonocm3 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm3Down LspPlugInPluginsLv2MbCompressorMonocm3 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm3Up LspPlugInPluginsLv2MbCompressorMonocm3 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonocm5 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm5Down LspPlugInPluginsLv2MbCompressorMonocm5 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm5Up LspPlugInPluginsLv2MbCompressorMonocm5 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonocm6 string

const (
	LspPlugInPluginsLv2MbCompressorMonocm6Down LspPlugInPluginsLv2MbCompressorMonocm6 = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorMonocm6Up LspPlugInPluginsLv2MbCompressorMonocm6 = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorMonomode string

const (
	LspPlugInPluginsLv2MbCompressorMonomodeClassic LspPlugInPluginsLv2MbCompressorMonomode = "Classic" // Classic (0) – Classic
	LspPlugInPluginsLv2MbCompressorMonomodeModern LspPlugInPluginsLv2MbCompressorMonomode = "Modern" // Modern (1) – Modern
)

type LspPlugInPluginsLv2MbCompressorMonoscm7 string

const (
	LspPlugInPluginsLv2MbCompressorMonoscm7Peak LspPlugInPluginsLv2MbCompressorMonoscm7 = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorMonoscm7Rms LspPlugInPluginsLv2MbCompressorMonoscm7 = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorMonoscm7Lpf LspPlugInPluginsLv2MbCompressorMonoscm7 = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorMonoscm7Uniform LspPlugInPluginsLv2MbCompressorMonoscm7 = "Uniform" // Uniform (3) – Uniform
)

