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

type LspPlugInPluginsLv2MbCompressorLr struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MbCompressorLr() (*LspPlugInPluginsLv2MbCompressorLr, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-mb-compressor-lr")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MbCompressorLrWithName(name string) (*LspPlugInPluginsLv2MbCompressorLr, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-mb-compressor-lr", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MbCompressorLr{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// al-0l (float32)
//
// Attack level 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl0L(value float32) error {
	return e.Element.SetProperty("al-0l", value)
}

// al-0r (float32)
//
// Attack level 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl0R(value float32) error {
	return e.Element.SetProperty("al-0r", value)
}

// al-1l (float32)
//
// Attack level 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl1L(value float32) error {
	return e.Element.SetProperty("al-1l", value)
}

// al-1r (float32)
//
// Attack level 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl1R(value float32) error {
	return e.Element.SetProperty("al-1r", value)
}

// al-2l (float32)
//
// Attack level 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl2L(value float32) error {
	return e.Element.SetProperty("al-2l", value)
}

// al-2r (float32)
//
// Attack level 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl2R(value float32) error {
	return e.Element.SetProperty("al-2r", value)
}

// al-3l (float32)
//
// Attack level 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl3L(value float32) error {
	return e.Element.SetProperty("al-3l", value)
}

// al-3r (float32)
//
// Attack level 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl3R(value float32) error {
	return e.Element.SetProperty("al-3r", value)
}

// al-4l (float32)
//
// Attack level 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl4L(value float32) error {
	return e.Element.SetProperty("al-4l", value)
}

// al-4r (float32)
//
// Attack level 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl4R(value float32) error {
	return e.Element.SetProperty("al-4r", value)
}

// al-5l (float32)
//
// Attack level 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl5L(value float32) error {
	return e.Element.SetProperty("al-5l", value)
}

// al-5r (float32)
//
// Attack level 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl5R(value float32) error {
	return e.Element.SetProperty("al-5r", value)
}

// al-6l (float32)
//
// Attack level 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl6L(value float32) error {
	return e.Element.SetProperty("al-6l", value)
}

// al-6r (float32)
//
// Attack level 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl6R(value float32) error {
	return e.Element.SetProperty("al-6r", value)
}

// al-7l (float32)
//
// Attack level 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl7L(value float32) error {
	return e.Element.SetProperty("al-7l", value)
}

// al-7r (float32)
//
// Attack level 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAl7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("al-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAl7R(value float32) error {
	return e.Element.SetProperty("al-7r", value)
}

// at-0l (float32)
//
// Attack time 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt0L(value float32) error {
	return e.Element.SetProperty("at-0l", value)
}

// at-0r (float32)
//
// Attack time 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt0R(value float32) error {
	return e.Element.SetProperty("at-0r", value)
}

// at-1l (float32)
//
// Attack time 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt1L(value float32) error {
	return e.Element.SetProperty("at-1l", value)
}

// at-1r (float32)
//
// Attack time 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt1R(value float32) error {
	return e.Element.SetProperty("at-1r", value)
}

// at-2l (float32)
//
// Attack time 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt2L(value float32) error {
	return e.Element.SetProperty("at-2l", value)
}

// at-2r (float32)
//
// Attack time 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt2R(value float32) error {
	return e.Element.SetProperty("at-2r", value)
}

// at-3l (float32)
//
// Attack time 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt3L(value float32) error {
	return e.Element.SetProperty("at-3l", value)
}

// at-3r (float32)
//
// Attack time 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt3R(value float32) error {
	return e.Element.SetProperty("at-3r", value)
}

// at-4l (float32)
//
// Attack time 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt4L(value float32) error {
	return e.Element.SetProperty("at-4l", value)
}

// at-4r (float32)
//
// Attack time 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt4R(value float32) error {
	return e.Element.SetProperty("at-4r", value)
}

// at-5l (float32)
//
// Attack time 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt5L(value float32) error {
	return e.Element.SetProperty("at-5l", value)
}

// at-5r (float32)
//
// Attack time 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt5R(value float32) error {
	return e.Element.SetProperty("at-5r", value)
}

// at-6l (float32)
//
// Attack time 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt6L(value float32) error {
	return e.Element.SetProperty("at-6l", value)
}

// at-6r (float32)
//
// Attack time 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt6R(value float32) error {
	return e.Element.SetProperty("at-6r", value)
}

// at-7l (float32)
//
// Attack time 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt7L(value float32) error {
	return e.Element.SetProperty("at-7l", value)
}

// at-7r (float32)
//
// Attack time 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetAt7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("at-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetAt7R(value float32) error {
	return e.Element.SetProperty("at-7r", value)
}

// bm-0l (bool)
//
// Mute band 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm0L(value bool) error {
	return e.Element.SetProperty("bm-0l", value)
}

// bm-0r (bool)
//
// Mute band 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm0R(value bool) error {
	return e.Element.SetProperty("bm-0r", value)
}

// bm-1l (bool)
//
// Mute band 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm1L(value bool) error {
	return e.Element.SetProperty("bm-1l", value)
}

// bm-1r (bool)
//
// Mute band 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm1R(value bool) error {
	return e.Element.SetProperty("bm-1r", value)
}

// bm-2l (bool)
//
// Mute band 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm2L(value bool) error {
	return e.Element.SetProperty("bm-2l", value)
}

// bm-2r (bool)
//
// Mute band 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm2R(value bool) error {
	return e.Element.SetProperty("bm-2r", value)
}

// bm-3l (bool)
//
// Mute band 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm3L(value bool) error {
	return e.Element.SetProperty("bm-3l", value)
}

// bm-3r (bool)
//
// Mute band 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm3R(value bool) error {
	return e.Element.SetProperty("bm-3r", value)
}

// bm-4l (bool)
//
// Mute band 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm4L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm4L(value bool) error {
	return e.Element.SetProperty("bm-4l", value)
}

// bm-4r (bool)
//
// Mute band 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm4R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm4R(value bool) error {
	return e.Element.SetProperty("bm-4r", value)
}

// bm-5l (bool)
//
// Mute band 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm5L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm5L(value bool) error {
	return e.Element.SetProperty("bm-5l", value)
}

// bm-5r (bool)
//
// Mute band 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm5R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm5R(value bool) error {
	return e.Element.SetProperty("bm-5r", value)
}

// bm-6l (bool)
//
// Mute band 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm6L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm6L(value bool) error {
	return e.Element.SetProperty("bm-6l", value)
}

// bm-6r (bool)
//
// Mute band 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm6R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm6R(value bool) error {
	return e.Element.SetProperty("bm-6r", value)
}

// bm-7l (bool)
//
// Mute band 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm7L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm7L(value bool) error {
	return e.Element.SetProperty("bm-7l", value)
}

// bm-7r (bool)
//
// Mute band 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBm7R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bm-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBm7R(value bool) error {
	return e.Element.SetProperty("bm-7r", value)
}

// bs-0l (bool)
//
// Solo band 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs0L(value bool) error {
	return e.Element.SetProperty("bs-0l", value)
}

// bs-0r (bool)
//
// Solo band 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs0R(value bool) error {
	return e.Element.SetProperty("bs-0r", value)
}

// bs-1l (bool)
//
// Solo band 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs1L(value bool) error {
	return e.Element.SetProperty("bs-1l", value)
}

// bs-1r (bool)
//
// Solo band 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs1R(value bool) error {
	return e.Element.SetProperty("bs-1r", value)
}

// bs-2l (bool)
//
// Solo band 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs2L(value bool) error {
	return e.Element.SetProperty("bs-2l", value)
}

// bs-2r (bool)
//
// Solo band 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs2R(value bool) error {
	return e.Element.SetProperty("bs-2r", value)
}

// bs-3l (bool)
//
// Solo band 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs3L(value bool) error {
	return e.Element.SetProperty("bs-3l", value)
}

// bs-3r (bool)
//
// Solo band 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs3R(value bool) error {
	return e.Element.SetProperty("bs-3r", value)
}

// bs-4l (bool)
//
// Solo band 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs4L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs4L(value bool) error {
	return e.Element.SetProperty("bs-4l", value)
}

// bs-4r (bool)
//
// Solo band 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs4R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs4R(value bool) error {
	return e.Element.SetProperty("bs-4r", value)
}

// bs-5l (bool)
//
// Solo band 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs5L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs5L(value bool) error {
	return e.Element.SetProperty("bs-5l", value)
}

// bs-5r (bool)
//
// Solo band 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs5R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs5R(value bool) error {
	return e.Element.SetProperty("bs-5r", value)
}

// bs-6l (bool)
//
// Solo band 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs6L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs6L(value bool) error {
	return e.Element.SetProperty("bs-6l", value)
}

// bs-6r (bool)
//
// Solo band 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs6R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs6R(value bool) error {
	return e.Element.SetProperty("bs-6r", value)
}

// bs-7l (bool)
//
// Solo band 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs7L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs7L(value bool) error {
	return e.Element.SetProperty("bs-7l", value)
}

// bs-7r (bool)
//
// Solo band 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBs7R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bs-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBs7R(value bool) error {
	return e.Element.SetProperty("bs-7r", value)
}

// bsel (GstLspPlugInPluginsLv2MbCompressorLrbsel)
//
// Band selection

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBsel() (interface{}, error) {
	return e.Element.GetProperty("bsel")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBsel(value interface{}) error {
	return e.Element.SetProperty("bsel", value)
}

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MbCompressorLr) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// cbe-1l (bool)
//
// Compression band enable 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe1L(value bool) error {
	return e.Element.SetProperty("cbe-1l", value)
}

// cbe-1r (bool)
//
// Compression band enable 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe1R(value bool) error {
	return e.Element.SetProperty("cbe-1r", value)
}

// cbe-2l (bool)
//
// Compression band enable 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe2L(value bool) error {
	return e.Element.SetProperty("cbe-2l", value)
}

// cbe-2r (bool)
//
// Compression band enable 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe2R(value bool) error {
	return e.Element.SetProperty("cbe-2r", value)
}

// cbe-3l (bool)
//
// Compression band enable 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe3L(value bool) error {
	return e.Element.SetProperty("cbe-3l", value)
}

// cbe-3r (bool)
//
// Compression band enable 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe3R(value bool) error {
	return e.Element.SetProperty("cbe-3r", value)
}

// cbe-4l (bool)
//
// Compression band enable 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe4L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe4L(value bool) error {
	return e.Element.SetProperty("cbe-4l", value)
}

// cbe-4r (bool)
//
// Compression band enable 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe4R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe4R(value bool) error {
	return e.Element.SetProperty("cbe-4r", value)
}

// cbe-5l (bool)
//
// Compression band enable 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe5L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe5L(value bool) error {
	return e.Element.SetProperty("cbe-5l", value)
}

// cbe-5r (bool)
//
// Compression band enable 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe5R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe5R(value bool) error {
	return e.Element.SetProperty("cbe-5r", value)
}

// cbe-6l (bool)
//
// Compression band enable 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe6L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe6L(value bool) error {
	return e.Element.SetProperty("cbe-6l", value)
}

// cbe-6r (bool)
//
// Compression band enable 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe6R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe6R(value bool) error {
	return e.Element.SetProperty("cbe-6r", value)
}

// cbe-7l (bool)
//
// Compression band enable 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe7L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe7L(value bool) error {
	return e.Element.SetProperty("cbe-7l", value)
}

// cbe-7r (bool)
//
// Compression band enable 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCbe7R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cbe-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCbe7R(value bool) error {
	return e.Element.SetProperty("cbe-7r", value)
}

// ce-0l (bool)
//
// Compressor enable 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe0L(value bool) error {
	return e.Element.SetProperty("ce-0l", value)
}

// ce-0r (bool)
//
// Compressor enable 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe0R(value bool) error {
	return e.Element.SetProperty("ce-0r", value)
}

// ce-1l (bool)
//
// Compressor enable 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe1L(value bool) error {
	return e.Element.SetProperty("ce-1l", value)
}

// ce-1r (bool)
//
// Compressor enable 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe1R(value bool) error {
	return e.Element.SetProperty("ce-1r", value)
}

// ce-2l (bool)
//
// Compressor enable 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe2L(value bool) error {
	return e.Element.SetProperty("ce-2l", value)
}

// ce-2r (bool)
//
// Compressor enable 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe2R(value bool) error {
	return e.Element.SetProperty("ce-2r", value)
}

// ce-3l (bool)
//
// Compressor enable 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe3L(value bool) error {
	return e.Element.SetProperty("ce-3l", value)
}

// ce-3r (bool)
//
// Compressor enable 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe3R(value bool) error {
	return e.Element.SetProperty("ce-3r", value)
}

// ce-4l (bool)
//
// Compressor enable 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe4L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe4L(value bool) error {
	return e.Element.SetProperty("ce-4l", value)
}

// ce-4r (bool)
//
// Compressor enable 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe4R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe4R(value bool) error {
	return e.Element.SetProperty("ce-4r", value)
}

// ce-5l (bool)
//
// Compressor enable 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe5L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe5L(value bool) error {
	return e.Element.SetProperty("ce-5l", value)
}

// ce-5r (bool)
//
// Compressor enable 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe5R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe5R(value bool) error {
	return e.Element.SetProperty("ce-5r", value)
}

// ce-6l (bool)
//
// Compressor enable 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe6L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe6L(value bool) error {
	return e.Element.SetProperty("ce-6l", value)
}

// ce-6r (bool)
//
// Compressor enable 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe6R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe6R(value bool) error {
	return e.Element.SetProperty("ce-6r", value)
}

// ce-7l (bool)
//
// Compressor enable 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe7L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe7L(value bool) error {
	return e.Element.SetProperty("ce-7l", value)
}

// ce-7r (bool)
//
// Compressor enable 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCe7R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ce-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCe7R(value bool) error {
	return e.Element.SetProperty("ce-7r", value)
}

// clm-0l (float32)
//
// Curve level meter 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-0r (float32)
//
// Curve level meter 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-1l (float32)
//
// Curve level meter 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-1r (float32)
//
// Curve level meter 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-2l (float32)
//
// Curve level meter 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-2r (float32)
//
// Curve level meter 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-3l (float32)
//
// Curve level meter 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-3r (float32)
//
// Curve level meter 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-4l (float32)
//
// Curve level meter 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-4r (float32)
//
// Curve level meter 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-5l (float32)
//
// Curve level meter 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-5r (float32)
//
// Curve level meter 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-6l (float32)
//
// Curve level meter 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-6r (float32)
//
// Curve level meter 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-7l (float32)
//
// Curve level meter 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// clm-7r (float32)
//
// Curve level meter 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetClm7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("clm-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// cm-0l (GstLspPlugInPluginsLv2MbCompressorLrcm0L)
//
// Compression mode 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm0L() (interface{}, error) {
	return e.Element.GetProperty("cm-0l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm0L(value interface{}) error {
	return e.Element.SetProperty("cm-0l", value)
}

// cm-0r (GstLspPlugInPluginsLv2MbCompressorLrcm0R)
//
// Compression mode 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm0R() (interface{}, error) {
	return e.Element.GetProperty("cm-0r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm0R(value interface{}) error {
	return e.Element.SetProperty("cm-0r", value)
}

// cm-1l (GstLspPlugInPluginsLv2MbCompressorLrcm1L)
//
// Compression mode 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm1L() (interface{}, error) {
	return e.Element.GetProperty("cm-1l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm1L(value interface{}) error {
	return e.Element.SetProperty("cm-1l", value)
}

// cm-1r (GstLspPlugInPluginsLv2MbCompressorLrcm1R)
//
// Compression mode 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm1R() (interface{}, error) {
	return e.Element.GetProperty("cm-1r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm1R(value interface{}) error {
	return e.Element.SetProperty("cm-1r", value)
}

// cm-2l (GstLspPlugInPluginsLv2MbCompressorLrcm2L)
//
// Compression mode 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm2L() (interface{}, error) {
	return e.Element.GetProperty("cm-2l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm2L(value interface{}) error {
	return e.Element.SetProperty("cm-2l", value)
}

// cm-2r (GstLspPlugInPluginsLv2MbCompressorLrcm2R)
//
// Compression mode 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm2R() (interface{}, error) {
	return e.Element.GetProperty("cm-2r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm2R(value interface{}) error {
	return e.Element.SetProperty("cm-2r", value)
}

// cm-3l (GstLspPlugInPluginsLv2MbCompressorLrcm3L)
//
// Compression mode 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm3L() (interface{}, error) {
	return e.Element.GetProperty("cm-3l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm3L(value interface{}) error {
	return e.Element.SetProperty("cm-3l", value)
}

// cm-3r (GstLspPlugInPluginsLv2MbCompressorLrcm3R)
//
// Compression mode 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm3R() (interface{}, error) {
	return e.Element.GetProperty("cm-3r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm3R(value interface{}) error {
	return e.Element.SetProperty("cm-3r", value)
}

// cm-4l (GstLspPlugInPluginsLv2MbCompressorLrcm4L)
//
// Compression mode 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm4L() (interface{}, error) {
	return e.Element.GetProperty("cm-4l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm4L(value interface{}) error {
	return e.Element.SetProperty("cm-4l", value)
}

// cm-4r (GstLspPlugInPluginsLv2MbCompressorLrcm4R)
//
// Compression mode 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm4R() (interface{}, error) {
	return e.Element.GetProperty("cm-4r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm4R(value interface{}) error {
	return e.Element.SetProperty("cm-4r", value)
}

// cm-5l (GstLspPlugInPluginsLv2MbCompressorLrcm5L)
//
// Compression mode 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm5L() (interface{}, error) {
	return e.Element.GetProperty("cm-5l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm5L(value interface{}) error {
	return e.Element.SetProperty("cm-5l", value)
}

// cm-5r (GstLspPlugInPluginsLv2MbCompressorLrcm5R)
//
// Compression mode 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm5R() (interface{}, error) {
	return e.Element.GetProperty("cm-5r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm5R(value interface{}) error {
	return e.Element.SetProperty("cm-5r", value)
}

// cm-6l (GstLspPlugInPluginsLv2MbCompressorLrcm6L)
//
// Compression mode 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm6L() (interface{}, error) {
	return e.Element.GetProperty("cm-6l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm6L(value interface{}) error {
	return e.Element.SetProperty("cm-6l", value)
}

// cm-6r (GstLspPlugInPluginsLv2MbCompressorLrcm6R)
//
// Compression mode 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm6R() (interface{}, error) {
	return e.Element.GetProperty("cm-6r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm6R(value interface{}) error {
	return e.Element.SetProperty("cm-6r", value)
}

// cm-7l (GstLspPlugInPluginsLv2MbCompressorLrcm7L)
//
// Compression mode 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm7L() (interface{}, error) {
	return e.Element.GetProperty("cm-7l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm7L(value interface{}) error {
	return e.Element.SetProperty("cm-7l", value)
}

// cm-7r (GstLspPlugInPluginsLv2MbCompressorLrcm7R)
//
// Compression mode 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCm7R() (interface{}, error) {
	return e.Element.GetProperty("cm-7r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCm7R(value interface{}) error {
	return e.Element.SetProperty("cm-7r", value)
}

// cr-0l (float32)
//
// Ratio 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr0L(value float32) error {
	return e.Element.SetProperty("cr-0l", value)
}

// cr-0r (float32)
//
// Ratio 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr0R(value float32) error {
	return e.Element.SetProperty("cr-0r", value)
}

// cr-1l (float32)
//
// Ratio 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr1L(value float32) error {
	return e.Element.SetProperty("cr-1l", value)
}

// cr-1r (float32)
//
// Ratio 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr1R(value float32) error {
	return e.Element.SetProperty("cr-1r", value)
}

// cr-2l (float32)
//
// Ratio 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr2L(value float32) error {
	return e.Element.SetProperty("cr-2l", value)
}

// cr-2r (float32)
//
// Ratio 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr2R(value float32) error {
	return e.Element.SetProperty("cr-2r", value)
}

// cr-3l (float32)
//
// Ratio 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr3L(value float32) error {
	return e.Element.SetProperty("cr-3l", value)
}

// cr-3r (float32)
//
// Ratio 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr3R(value float32) error {
	return e.Element.SetProperty("cr-3r", value)
}

// cr-4l (float32)
//
// Ratio 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr4L(value float32) error {
	return e.Element.SetProperty("cr-4l", value)
}

// cr-4r (float32)
//
// Ratio 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr4R(value float32) error {
	return e.Element.SetProperty("cr-4r", value)
}

// cr-5l (float32)
//
// Ratio 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr5L(value float32) error {
	return e.Element.SetProperty("cr-5l", value)
}

// cr-5r (float32)
//
// Ratio 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr5R(value float32) error {
	return e.Element.SetProperty("cr-5r", value)
}

// cr-6l (float32)
//
// Ratio 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr6L(value float32) error {
	return e.Element.SetProperty("cr-6l", value)
}

// cr-6r (float32)
//
// Ratio 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr6R(value float32) error {
	return e.Element.SetProperty("cr-6r", value)
}

// cr-7l (float32)
//
// Ratio 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr7L(value float32) error {
	return e.Element.SetProperty("cr-7l", value)
}

// cr-7r (float32)
//
// Ratio 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetCr7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("cr-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetCr7R(value float32) error {
	return e.Element.SetProperty("cr-7r", value)
}

// elm-0l (float32)
//
// Envelope level meter 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-0r (float32)
//
// Envelope level meter 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-1l (float32)
//
// Envelope level meter 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-1r (float32)
//
// Envelope level meter 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-2l (float32)
//
// Envelope level meter 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-2r (float32)
//
// Envelope level meter 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-3l (float32)
//
// Envelope level meter 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-3r (float32)
//
// Envelope level meter 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-4l (float32)
//
// Envelope level meter 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-4r (float32)
//
// Envelope level meter 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-5l (float32)
//
// Envelope level meter 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-5r (float32)
//
// Envelope level meter 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-6l (float32)
//
// Envelope level meter 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-6r (float32)
//
// Envelope level meter 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-7l (float32)
//
// Envelope level meter 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// elm-7r (float32)
//
// Envelope level meter 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetElm7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("elm-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// envb (GstLspPlugInPluginsLv2MbCompressorLrenvb)
//
// Envelope boost

func (e *LspPlugInPluginsLv2MbCompressorLr) GetEnvb() (interface{}, error) {
	return e.Element.GetProperty("envb")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetEnvb(value interface{}) error {
	return e.Element.SetProperty("envb", value)
}

// flt-l (bool)
//
// Band filter curves Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFltL() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flt-l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetFltL(value bool) error {
	return e.Element.SetProperty("flt-l", value)
}

// flt-r (bool)
//
// Band filter curves Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFltR() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flt-r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetFltR(value bool) error {
	return e.Element.SetProperty("flt-r", value)
}

// fre-0l (float32)
//
// Frequency range end 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-0r (float32)
//
// Frequency range end 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-1l (float32)
//
// Frequency range end 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-1r (float32)
//
// Frequency range end 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-2l (float32)
//
// Frequency range end 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-2r (float32)
//
// Frequency range end 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-3l (float32)
//
// Frequency range end 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-3r (float32)
//
// Frequency range end 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-4l (float32)
//
// Frequency range end 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-4r (float32)
//
// Frequency range end 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-5l (float32)
//
// Frequency range end 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-5r (float32)
//
// Frequency range end 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-6l (float32)
//
// Frequency range end 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-6r (float32)
//
// Frequency range end 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-7l (float32)
//
// Frequency range end 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// fre-7r (float32)
//
// Frequency range end 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetFre7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fre-7r")
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

func (e *LspPlugInPluginsLv2MbCompressorLr) GetGDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetGDry(value float32) error {
	return e.Element.SetProperty("g-dry", value)
}

// g-in (float32)
//
// Input gain

func (e *LspPlugInPluginsLv2MbCompressorLr) GetGIn() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetGIn(value float32) error {
	return e.Element.SetProperty("g-in", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MbCompressorLr) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// g-wet (float32)
//
// Wet gain

func (e *LspPlugInPluginsLv2MbCompressorLr) GetGWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetGWet(value float32) error {
	return e.Element.SetProperty("g-wet", value)
}

// hue-0l (float32)
//
// Hue  0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue0L(value float32) error {
	return e.Element.SetProperty("hue-0l", value)
}

// hue-0r (float32)
//
// Hue  0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue0R(value float32) error {
	return e.Element.SetProperty("hue-0r", value)
}

// hue-1l (float32)
//
// Hue  1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue1L(value float32) error {
	return e.Element.SetProperty("hue-1l", value)
}

// hue-1r (float32)
//
// Hue  1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue1R(value float32) error {
	return e.Element.SetProperty("hue-1r", value)
}

// hue-2l (float32)
//
// Hue  2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue2L(value float32) error {
	return e.Element.SetProperty("hue-2l", value)
}

// hue-2r (float32)
//
// Hue  2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue2R(value float32) error {
	return e.Element.SetProperty("hue-2r", value)
}

// hue-3l (float32)
//
// Hue  3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue3L(value float32) error {
	return e.Element.SetProperty("hue-3l", value)
}

// hue-3r (float32)
//
// Hue  3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue3R(value float32) error {
	return e.Element.SetProperty("hue-3r", value)
}

// hue-4l (float32)
//
// Hue  4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue4L(value float32) error {
	return e.Element.SetProperty("hue-4l", value)
}

// hue-4r (float32)
//
// Hue  4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue4R(value float32) error {
	return e.Element.SetProperty("hue-4r", value)
}

// hue-5l (float32)
//
// Hue  5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue5L(value float32) error {
	return e.Element.SetProperty("hue-5l", value)
}

// hue-5r (float32)
//
// Hue  5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue5R(value float32) error {
	return e.Element.SetProperty("hue-5r", value)
}

// hue-6l (float32)
//
// Hue  6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue6L(value float32) error {
	return e.Element.SetProperty("hue-6l", value)
}

// hue-6r (float32)
//
// Hue  6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue6R(value float32) error {
	return e.Element.SetProperty("hue-6r", value)
}

// hue-7l (float32)
//
// Hue  7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue7L(value float32) error {
	return e.Element.SetProperty("hue-7l", value)
}

// hue-7r (float32)
//
// Hue  7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetHue7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("hue-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetHue7R(value float32) error {
	return e.Element.SetProperty("hue-7r", value)
}

// ife-l (bool)
//
// Input FFT graph enable Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetIfeL() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetIfeL(value bool) error {
	return e.Element.SetProperty("ife-l", value)
}

// ife-r (bool)
//
// Input FFT graph enable Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetIfeR() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetIfeR(value bool) error {
	return e.Element.SetProperty("ife-r", value)
}

// ilm-l (float32)
//
// Input level meter Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetIlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) GetIlmR() (float32, error) {
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

// kn-0l (float32)
//
// Knee 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn0L(value float32) error {
	return e.Element.SetProperty("kn-0l", value)
}

// kn-0r (float32)
//
// Knee 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn0R(value float32) error {
	return e.Element.SetProperty("kn-0r", value)
}

// kn-1l (float32)
//
// Knee 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn1L(value float32) error {
	return e.Element.SetProperty("kn-1l", value)
}

// kn-1r (float32)
//
// Knee 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn1R(value float32) error {
	return e.Element.SetProperty("kn-1r", value)
}

// kn-2l (float32)
//
// Knee 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn2L(value float32) error {
	return e.Element.SetProperty("kn-2l", value)
}

// kn-2r (float32)
//
// Knee 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn2R(value float32) error {
	return e.Element.SetProperty("kn-2r", value)
}

// kn-3l (float32)
//
// Knee 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn3L(value float32) error {
	return e.Element.SetProperty("kn-3l", value)
}

// kn-3r (float32)
//
// Knee 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn3R(value float32) error {
	return e.Element.SetProperty("kn-3r", value)
}

// kn-4l (float32)
//
// Knee 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn4L(value float32) error {
	return e.Element.SetProperty("kn-4l", value)
}

// kn-4r (float32)
//
// Knee 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn4R(value float32) error {
	return e.Element.SetProperty("kn-4r", value)
}

// kn-5l (float32)
//
// Knee 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn5L(value float32) error {
	return e.Element.SetProperty("kn-5l", value)
}

// kn-5r (float32)
//
// Knee 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn5R(value float32) error {
	return e.Element.SetProperty("kn-5r", value)
}

// kn-6l (float32)
//
// Knee 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn6L(value float32) error {
	return e.Element.SetProperty("kn-6l", value)
}

// kn-6r (float32)
//
// Knee 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn6R(value float32) error {
	return e.Element.SetProperty("kn-6r", value)
}

// kn-7l (float32)
//
// Knee 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn7L(value float32) error {
	return e.Element.SetProperty("kn-7l", value)
}

// kn-7r (float32)
//
// Knee 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetKn7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("kn-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetKn7R(value float32) error {
	return e.Element.SetProperty("kn-7r", value)
}

// mk-0l (float32)
//
// Makeup gain 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk0L(value float32) error {
	return e.Element.SetProperty("mk-0l", value)
}

// mk-0r (float32)
//
// Makeup gain 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk0R(value float32) error {
	return e.Element.SetProperty("mk-0r", value)
}

// mk-1l (float32)
//
// Makeup gain 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk1L(value float32) error {
	return e.Element.SetProperty("mk-1l", value)
}

// mk-1r (float32)
//
// Makeup gain 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk1R(value float32) error {
	return e.Element.SetProperty("mk-1r", value)
}

// mk-2l (float32)
//
// Makeup gain 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk2L(value float32) error {
	return e.Element.SetProperty("mk-2l", value)
}

// mk-2r (float32)
//
// Makeup gain 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk2R(value float32) error {
	return e.Element.SetProperty("mk-2r", value)
}

// mk-3l (float32)
//
// Makeup gain 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk3L(value float32) error {
	return e.Element.SetProperty("mk-3l", value)
}

// mk-3r (float32)
//
// Makeup gain 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk3R(value float32) error {
	return e.Element.SetProperty("mk-3r", value)
}

// mk-4l (float32)
//
// Makeup gain 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk4L(value float32) error {
	return e.Element.SetProperty("mk-4l", value)
}

// mk-4r (float32)
//
// Makeup gain 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk4R(value float32) error {
	return e.Element.SetProperty("mk-4r", value)
}

// mk-5l (float32)
//
// Makeup gain 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk5L(value float32) error {
	return e.Element.SetProperty("mk-5l", value)
}

// mk-5r (float32)
//
// Makeup gain 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk5R(value float32) error {
	return e.Element.SetProperty("mk-5r", value)
}

// mk-6l (float32)
//
// Makeup gain 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk6L(value float32) error {
	return e.Element.SetProperty("mk-6l", value)
}

// mk-6r (float32)
//
// Makeup gain 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk6R(value float32) error {
	return e.Element.SetProperty("mk-6r", value)
}

// mk-7l (float32)
//
// Makeup gain 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk7L(value float32) error {
	return e.Element.SetProperty("mk-7l", value)
}

// mk-7r (float32)
//
// Makeup gain 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMk7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("mk-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMk7R(value float32) error {
	return e.Element.SetProperty("mk-7r", value)
}

// mode (GstLspPlugInPluginsLv2MbCompressorLrmode)
//
// Compressor mode

func (e *LspPlugInPluginsLv2MbCompressorLr) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// ofe-l (bool)
//
// Output FFT graph enable Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetOfeL() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetOfeL(value bool) error {
	return e.Element.SetProperty("ofe-l", value)
}

// ofe-r (bool)
//
// Output FFT graph enable Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetOfeR() (bool, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetOfeR(value bool) error {
	return e.Element.SetProperty("ofe-r", value)
}

// olm-l (float32)
//
// Output level meter Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetOlmL() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) GetOlmR() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) GetReact() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetReact(value float32) error {
	return e.Element.SetProperty("react", value)
}

// rl-0l (float32)
//
// Release level 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-0r (float32)
//
// Release level 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-1l (float32)
//
// Release level 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-1r (float32)
//
// Release level 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-2l (float32)
//
// Release level 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-2r (float32)
//
// Release level 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-3l (float32)
//
// Release level 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-3r (float32)
//
// Release level 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-4l (float32)
//
// Release level 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-4r (float32)
//
// Release level 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-5l (float32)
//
// Release level 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-5r (float32)
//
// Release level 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-6l (float32)
//
// Release level 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-6r (float32)
//
// Release level 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-7l (float32)
//
// Release level 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rl-7r (float32)
//
// Release level 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRl7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rl-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-0l (float32)
//
// Reduction level meter 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-0r (float32)
//
// Reduction level meter 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-1l (float32)
//
// Reduction level meter 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-1r (float32)
//
// Reduction level meter 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-2l (float32)
//
// Reduction level meter 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-2r (float32)
//
// Reduction level meter 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-3l (float32)
//
// Reduction level meter 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-3r (float32)
//
// Reduction level meter 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-4l (float32)
//
// Reduction level meter 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-4r (float32)
//
// Reduction level meter 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-5l (float32)
//
// Reduction level meter 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-5r (float32)
//
// Reduction level meter 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-6l (float32)
//
// Reduction level meter 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-6r (float32)
//
// Reduction level meter 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-7l (float32)
//
// Reduction level meter 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rlm-7r (float32)
//
// Reduction level meter 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRlm7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rlm-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// rrl-0l (float32)
//
// Relative release level 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl0L(value float32) error {
	return e.Element.SetProperty("rrl-0l", value)
}

// rrl-0r (float32)
//
// Relative release level 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl0R(value float32) error {
	return e.Element.SetProperty("rrl-0r", value)
}

// rrl-1l (float32)
//
// Relative release level 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl1L(value float32) error {
	return e.Element.SetProperty("rrl-1l", value)
}

// rrl-1r (float32)
//
// Relative release level 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl1R(value float32) error {
	return e.Element.SetProperty("rrl-1r", value)
}

// rrl-2l (float32)
//
// Relative release level 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl2L(value float32) error {
	return e.Element.SetProperty("rrl-2l", value)
}

// rrl-2r (float32)
//
// Relative release level 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl2R(value float32) error {
	return e.Element.SetProperty("rrl-2r", value)
}

// rrl-3l (float32)
//
// Relative release level 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl3L(value float32) error {
	return e.Element.SetProperty("rrl-3l", value)
}

// rrl-3r (float32)
//
// Relative release level 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl3R(value float32) error {
	return e.Element.SetProperty("rrl-3r", value)
}

// rrl-4l (float32)
//
// Relative release level 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl4L(value float32) error {
	return e.Element.SetProperty("rrl-4l", value)
}

// rrl-4r (float32)
//
// Relative release level 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl4R(value float32) error {
	return e.Element.SetProperty("rrl-4r", value)
}

// rrl-5l (float32)
//
// Relative release level 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl5L(value float32) error {
	return e.Element.SetProperty("rrl-5l", value)
}

// rrl-5r (float32)
//
// Relative release level 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl5R(value float32) error {
	return e.Element.SetProperty("rrl-5r", value)
}

// rrl-6l (float32)
//
// Relative release level 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl6L(value float32) error {
	return e.Element.SetProperty("rrl-6l", value)
}

// rrl-6r (float32)
//
// Relative release level 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl6R(value float32) error {
	return e.Element.SetProperty("rrl-6r", value)
}

// rrl-7l (float32)
//
// Relative release level 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl7L(value float32) error {
	return e.Element.SetProperty("rrl-7l", value)
}

// rrl-7r (float32)
//
// Relative release level 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRrl7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rrl-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRrl7R(value float32) error {
	return e.Element.SetProperty("rrl-7r", value)
}

// rt-0l (float32)
//
// Release time 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt0L(value float32) error {
	return e.Element.SetProperty("rt-0l", value)
}

// rt-0r (float32)
//
// Release time 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt0R(value float32) error {
	return e.Element.SetProperty("rt-0r", value)
}

// rt-1l (float32)
//
// Release time 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt1L(value float32) error {
	return e.Element.SetProperty("rt-1l", value)
}

// rt-1r (float32)
//
// Release time 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt1R(value float32) error {
	return e.Element.SetProperty("rt-1r", value)
}

// rt-2l (float32)
//
// Release time 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt2L(value float32) error {
	return e.Element.SetProperty("rt-2l", value)
}

// rt-2r (float32)
//
// Release time 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt2R(value float32) error {
	return e.Element.SetProperty("rt-2r", value)
}

// rt-3l (float32)
//
// Release time 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt3L(value float32) error {
	return e.Element.SetProperty("rt-3l", value)
}

// rt-3r (float32)
//
// Release time 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt3R(value float32) error {
	return e.Element.SetProperty("rt-3r", value)
}

// rt-4l (float32)
//
// Release time 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt4L(value float32) error {
	return e.Element.SetProperty("rt-4l", value)
}

// rt-4r (float32)
//
// Release time 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt4R(value float32) error {
	return e.Element.SetProperty("rt-4r", value)
}

// rt-5l (float32)
//
// Release time 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt5L(value float32) error {
	return e.Element.SetProperty("rt-5l", value)
}

// rt-5r (float32)
//
// Release time 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt5R(value float32) error {
	return e.Element.SetProperty("rt-5r", value)
}

// rt-6l (float32)
//
// Release time 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt6L(value float32) error {
	return e.Element.SetProperty("rt-6l", value)
}

// rt-6r (float32)
//
// Release time 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt6R(value float32) error {
	return e.Element.SetProperty("rt-6r", value)
}

// rt-7l (float32)
//
// Release time 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt7L(value float32) error {
	return e.Element.SetProperty("rt-7l", value)
}

// rt-7r (float32)
//
// Release time 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetRt7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("rt-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetRt7R(value float32) error {
	return e.Element.SetProperty("rt-7r", value)
}

// schc-0l (bool)
//
// Sidechain custom hi-cut 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc0L(value bool) error {
	return e.Element.SetProperty("schc-0l", value)
}

// schc-0r (bool)
//
// Sidechain custom hi-cut 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc0R(value bool) error {
	return e.Element.SetProperty("schc-0r", value)
}

// schc-1l (bool)
//
// Sidechain custom hi-cut 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc1L(value bool) error {
	return e.Element.SetProperty("schc-1l", value)
}

// schc-1r (bool)
//
// Sidechain custom hi-cut 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc1R(value bool) error {
	return e.Element.SetProperty("schc-1r", value)
}

// schc-2l (bool)
//
// Sidechain custom hi-cut 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc2L(value bool) error {
	return e.Element.SetProperty("schc-2l", value)
}

// schc-2r (bool)
//
// Sidechain custom hi-cut 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc2R(value bool) error {
	return e.Element.SetProperty("schc-2r", value)
}

// schc-3l (bool)
//
// Sidechain custom hi-cut 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc3L(value bool) error {
	return e.Element.SetProperty("schc-3l", value)
}

// schc-3r (bool)
//
// Sidechain custom hi-cut 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc3R(value bool) error {
	return e.Element.SetProperty("schc-3r", value)
}

// schc-4l (bool)
//
// Sidechain custom hi-cut 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc4L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc4L(value bool) error {
	return e.Element.SetProperty("schc-4l", value)
}

// schc-4r (bool)
//
// Sidechain custom hi-cut 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc4R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc4R(value bool) error {
	return e.Element.SetProperty("schc-4r", value)
}

// schc-5l (bool)
//
// Sidechain custom hi-cut 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc5L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc5L(value bool) error {
	return e.Element.SetProperty("schc-5l", value)
}

// schc-5r (bool)
//
// Sidechain custom hi-cut 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc5R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc5R(value bool) error {
	return e.Element.SetProperty("schc-5r", value)
}

// schc-6l (bool)
//
// Sidechain custom hi-cut 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc6L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc6L(value bool) error {
	return e.Element.SetProperty("schc-6l", value)
}

// schc-6r (bool)
//
// Sidechain custom hi-cut 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc6R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc6R(value bool) error {
	return e.Element.SetProperty("schc-6r", value)
}

// schc-7l (bool)
//
// Sidechain custom hi-cut 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc7L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc7L(value bool) error {
	return e.Element.SetProperty("schc-7l", value)
}

// schc-7r (bool)
//
// Sidechain custom hi-cut 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchc7R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("schc-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchc7R(value bool) error {
	return e.Element.SetProperty("schc-7r", value)
}

// schf-0l (float32)
//
// Sidechain hi-cut frequency 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf0L(value float32) error {
	return e.Element.SetProperty("schf-0l", value)
}

// schf-0r (float32)
//
// Sidechain hi-cut frequency 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf0R(value float32) error {
	return e.Element.SetProperty("schf-0r", value)
}

// schf-1l (float32)
//
// Sidechain hi-cut frequency 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf1L(value float32) error {
	return e.Element.SetProperty("schf-1l", value)
}

// schf-1r (float32)
//
// Sidechain hi-cut frequency 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf1R(value float32) error {
	return e.Element.SetProperty("schf-1r", value)
}

// schf-2l (float32)
//
// Sidechain hi-cut frequency 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf2L(value float32) error {
	return e.Element.SetProperty("schf-2l", value)
}

// schf-2r (float32)
//
// Sidechain hi-cut frequency 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf2R(value float32) error {
	return e.Element.SetProperty("schf-2r", value)
}

// schf-3l (float32)
//
// Sidechain hi-cut frequency 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf3L(value float32) error {
	return e.Element.SetProperty("schf-3l", value)
}

// schf-3r (float32)
//
// Sidechain hi-cut frequency 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf3R(value float32) error {
	return e.Element.SetProperty("schf-3r", value)
}

// schf-4l (float32)
//
// Sidechain hi-cut frequency 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf4L(value float32) error {
	return e.Element.SetProperty("schf-4l", value)
}

// schf-4r (float32)
//
// Sidechain hi-cut frequency 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf4R(value float32) error {
	return e.Element.SetProperty("schf-4r", value)
}

// schf-5l (float32)
//
// Sidechain hi-cut frequency 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf5L(value float32) error {
	return e.Element.SetProperty("schf-5l", value)
}

// schf-5r (float32)
//
// Sidechain hi-cut frequency 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf5R(value float32) error {
	return e.Element.SetProperty("schf-5r", value)
}

// schf-6l (float32)
//
// Sidechain hi-cut frequency 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf6L(value float32) error {
	return e.Element.SetProperty("schf-6l", value)
}

// schf-6r (float32)
//
// Sidechain hi-cut frequency 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf6R(value float32) error {
	return e.Element.SetProperty("schf-6r", value)
}

// schf-7l (float32)
//
// Sidechain hi-cut frequency 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf7L(value float32) error {
	return e.Element.SetProperty("schf-7l", value)
}

// schf-7r (float32)
//
// Sidechain hi-cut frequency 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSchf7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("schf-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSchf7R(value float32) error {
	return e.Element.SetProperty("schf-7r", value)
}

// sclc-0l (bool)
//
// Sidechain custom lo-cut 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc0L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc0L(value bool) error {
	return e.Element.SetProperty("sclc-0l", value)
}

// sclc-0r (bool)
//
// Sidechain custom lo-cut 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc0R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc0R(value bool) error {
	return e.Element.SetProperty("sclc-0r", value)
}

// sclc-1l (bool)
//
// Sidechain custom lo-cut 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc1L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc1L(value bool) error {
	return e.Element.SetProperty("sclc-1l", value)
}

// sclc-1r (bool)
//
// Sidechain custom lo-cut 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc1R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc1R(value bool) error {
	return e.Element.SetProperty("sclc-1r", value)
}

// sclc-2l (bool)
//
// Sidechain custom lo-cut 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc2L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc2L(value bool) error {
	return e.Element.SetProperty("sclc-2l", value)
}

// sclc-2r (bool)
//
// Sidechain custom lo-cut 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc2R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc2R(value bool) error {
	return e.Element.SetProperty("sclc-2r", value)
}

// sclc-3l (bool)
//
// Sidechain custom lo-cut 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc3L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc3L(value bool) error {
	return e.Element.SetProperty("sclc-3l", value)
}

// sclc-3r (bool)
//
// Sidechain custom lo-cut 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc3R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc3R(value bool) error {
	return e.Element.SetProperty("sclc-3r", value)
}

// sclc-4l (bool)
//
// Sidechain custom lo-cut 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc4L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc4L(value bool) error {
	return e.Element.SetProperty("sclc-4l", value)
}

// sclc-4r (bool)
//
// Sidechain custom lo-cut 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc4R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc4R(value bool) error {
	return e.Element.SetProperty("sclc-4r", value)
}

// sclc-5l (bool)
//
// Sidechain custom lo-cut 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc5L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc5L(value bool) error {
	return e.Element.SetProperty("sclc-5l", value)
}

// sclc-5r (bool)
//
// Sidechain custom lo-cut 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc5R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc5R(value bool) error {
	return e.Element.SetProperty("sclc-5r", value)
}

// sclc-6l (bool)
//
// Sidechain custom lo-cut 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc6L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc6L(value bool) error {
	return e.Element.SetProperty("sclc-6l", value)
}

// sclc-6r (bool)
//
// Sidechain custom lo-cut 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc6R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc6R(value bool) error {
	return e.Element.SetProperty("sclc-6r", value)
}

// sclc-7l (bool)
//
// Sidechain custom lo-cut 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc7L() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc7L(value bool) error {
	return e.Element.SetProperty("sclc-7l", value)
}

// sclc-7r (bool)
//
// Sidechain custom lo-cut 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclc7R() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sclc-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclc7R(value bool) error {
	return e.Element.SetProperty("sclc-7r", value)
}

// sclf-0l (float32)
//
// Sidechain lo-cut frequency 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf0L(value float32) error {
	return e.Element.SetProperty("sclf-0l", value)
}

// sclf-0r (float32)
//
// Sidechain lo-cut frequency 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf0R(value float32) error {
	return e.Element.SetProperty("sclf-0r", value)
}

// sclf-1l (float32)
//
// Sidechain lo-cut frequency 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf1L(value float32) error {
	return e.Element.SetProperty("sclf-1l", value)
}

// sclf-1r (float32)
//
// Sidechain lo-cut frequency 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf1R(value float32) error {
	return e.Element.SetProperty("sclf-1r", value)
}

// sclf-2l (float32)
//
// Sidechain lo-cut frequency 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf2L(value float32) error {
	return e.Element.SetProperty("sclf-2l", value)
}

// sclf-2r (float32)
//
// Sidechain lo-cut frequency 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf2R(value float32) error {
	return e.Element.SetProperty("sclf-2r", value)
}

// sclf-3l (float32)
//
// Sidechain lo-cut frequency 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf3L(value float32) error {
	return e.Element.SetProperty("sclf-3l", value)
}

// sclf-3r (float32)
//
// Sidechain lo-cut frequency 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf3R(value float32) error {
	return e.Element.SetProperty("sclf-3r", value)
}

// sclf-4l (float32)
//
// Sidechain lo-cut frequency 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf4L(value float32) error {
	return e.Element.SetProperty("sclf-4l", value)
}

// sclf-4r (float32)
//
// Sidechain lo-cut frequency 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf4R(value float32) error {
	return e.Element.SetProperty("sclf-4r", value)
}

// sclf-5l (float32)
//
// Sidechain lo-cut frequency 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf5L(value float32) error {
	return e.Element.SetProperty("sclf-5l", value)
}

// sclf-5r (float32)
//
// Sidechain lo-cut frequency 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf5R(value float32) error {
	return e.Element.SetProperty("sclf-5r", value)
}

// sclf-6l (float32)
//
// Sidechain lo-cut frequency 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf6L(value float32) error {
	return e.Element.SetProperty("sclf-6l", value)
}

// sclf-6r (float32)
//
// Sidechain lo-cut frequency 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf6R(value float32) error {
	return e.Element.SetProperty("sclf-6r", value)
}

// sclf-7l (float32)
//
// Sidechain lo-cut frequency 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf7L(value float32) error {
	return e.Element.SetProperty("sclf-7l", value)
}

// sclf-7r (float32)
//
// Sidechain lo-cut frequency 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSclf7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sclf-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSclf7R(value float32) error {
	return e.Element.SetProperty("sclf-7r", value)
}

// scm-0l (GstLspPlugInPluginsLv2MbCompressorLrscm0L)
//
// Sidechain mode 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm0L() (interface{}, error) {
	return e.Element.GetProperty("scm-0l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm0L(value interface{}) error {
	return e.Element.SetProperty("scm-0l", value)
}

// scm-0r (GstLspPlugInPluginsLv2MbCompressorLrscm0R)
//
// Sidechain mode 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm0R() (interface{}, error) {
	return e.Element.GetProperty("scm-0r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm0R(value interface{}) error {
	return e.Element.SetProperty("scm-0r", value)
}

// scm-1l (GstLspPlugInPluginsLv2MbCompressorLrscm1L)
//
// Sidechain mode 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm1L() (interface{}, error) {
	return e.Element.GetProperty("scm-1l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm1L(value interface{}) error {
	return e.Element.SetProperty("scm-1l", value)
}

// scm-1r (GstLspPlugInPluginsLv2MbCompressorLrscm1R)
//
// Sidechain mode 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm1R() (interface{}, error) {
	return e.Element.GetProperty("scm-1r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm1R(value interface{}) error {
	return e.Element.SetProperty("scm-1r", value)
}

// scm-2l (GstLspPlugInPluginsLv2MbCompressorLrscm2L)
//
// Sidechain mode 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm2L() (interface{}, error) {
	return e.Element.GetProperty("scm-2l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm2L(value interface{}) error {
	return e.Element.SetProperty("scm-2l", value)
}

// scm-2r (GstLspPlugInPluginsLv2MbCompressorLrscm2R)
//
// Sidechain mode 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm2R() (interface{}, error) {
	return e.Element.GetProperty("scm-2r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm2R(value interface{}) error {
	return e.Element.SetProperty("scm-2r", value)
}

// scm-3l (GstLspPlugInPluginsLv2MbCompressorLrscm3L)
//
// Sidechain mode 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm3L() (interface{}, error) {
	return e.Element.GetProperty("scm-3l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm3L(value interface{}) error {
	return e.Element.SetProperty("scm-3l", value)
}

// scm-3r (GstLspPlugInPluginsLv2MbCompressorLrscm3R)
//
// Sidechain mode 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm3R() (interface{}, error) {
	return e.Element.GetProperty("scm-3r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm3R(value interface{}) error {
	return e.Element.SetProperty("scm-3r", value)
}

// scm-4l (GstLspPlugInPluginsLv2MbCompressorLrscm4L)
//
// Sidechain mode 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm4L() (interface{}, error) {
	return e.Element.GetProperty("scm-4l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm4L(value interface{}) error {
	return e.Element.SetProperty("scm-4l", value)
}

// scm-4r (GstLspPlugInPluginsLv2MbCompressorLrscm4R)
//
// Sidechain mode 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm4R() (interface{}, error) {
	return e.Element.GetProperty("scm-4r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm4R(value interface{}) error {
	return e.Element.SetProperty("scm-4r", value)
}

// scm-5l (GstLspPlugInPluginsLv2MbCompressorLrscm5L)
//
// Sidechain mode 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm5L() (interface{}, error) {
	return e.Element.GetProperty("scm-5l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm5L(value interface{}) error {
	return e.Element.SetProperty("scm-5l", value)
}

// scm-5r (GstLspPlugInPluginsLv2MbCompressorLrscm5R)
//
// Sidechain mode 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm5R() (interface{}, error) {
	return e.Element.GetProperty("scm-5r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm5R(value interface{}) error {
	return e.Element.SetProperty("scm-5r", value)
}

// scm-6l (GstLspPlugInPluginsLv2MbCompressorLrscm6L)
//
// Sidechain mode 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm6L() (interface{}, error) {
	return e.Element.GetProperty("scm-6l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm6L(value interface{}) error {
	return e.Element.SetProperty("scm-6l", value)
}

// scm-6r (GstLspPlugInPluginsLv2MbCompressorLrscm6R)
//
// Sidechain mode 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm6R() (interface{}, error) {
	return e.Element.GetProperty("scm-6r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm6R(value interface{}) error {
	return e.Element.SetProperty("scm-6r", value)
}

// scm-7l (GstLspPlugInPluginsLv2MbCompressorLrscm7L)
//
// Sidechain mode 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm7L() (interface{}, error) {
	return e.Element.GetProperty("scm-7l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm7L(value interface{}) error {
	return e.Element.SetProperty("scm-7l", value)
}

// scm-7r (GstLspPlugInPluginsLv2MbCompressorLrscm7R)
//
// Sidechain mode 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScm7R() (interface{}, error) {
	return e.Element.GetProperty("scm-7r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScm7R(value interface{}) error {
	return e.Element.SetProperty("scm-7r", value)
}

// scp-0l (float32)
//
// Sidechain preamp 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp0L(value float32) error {
	return e.Element.SetProperty("scp-0l", value)
}

// scp-0r (float32)
//
// Sidechain preamp 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp0R(value float32) error {
	return e.Element.SetProperty("scp-0r", value)
}

// scp-1l (float32)
//
// Sidechain preamp 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp1L(value float32) error {
	return e.Element.SetProperty("scp-1l", value)
}

// scp-1r (float32)
//
// Sidechain preamp 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp1R(value float32) error {
	return e.Element.SetProperty("scp-1r", value)
}

// scp-2l (float32)
//
// Sidechain preamp 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp2L(value float32) error {
	return e.Element.SetProperty("scp-2l", value)
}

// scp-2r (float32)
//
// Sidechain preamp 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp2R(value float32) error {
	return e.Element.SetProperty("scp-2r", value)
}

// scp-3l (float32)
//
// Sidechain preamp 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp3L(value float32) error {
	return e.Element.SetProperty("scp-3l", value)
}

// scp-3r (float32)
//
// Sidechain preamp 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp3R(value float32) error {
	return e.Element.SetProperty("scp-3r", value)
}

// scp-4l (float32)
//
// Sidechain preamp 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp4L(value float32) error {
	return e.Element.SetProperty("scp-4l", value)
}

// scp-4r (float32)
//
// Sidechain preamp 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp4R(value float32) error {
	return e.Element.SetProperty("scp-4r", value)
}

// scp-5l (float32)
//
// Sidechain preamp 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp5L(value float32) error {
	return e.Element.SetProperty("scp-5l", value)
}

// scp-5r (float32)
//
// Sidechain preamp 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp5R(value float32) error {
	return e.Element.SetProperty("scp-5r", value)
}

// scp-6l (float32)
//
// Sidechain preamp 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp6L(value float32) error {
	return e.Element.SetProperty("scp-6l", value)
}

// scp-6r (float32)
//
// Sidechain preamp 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp6R(value float32) error {
	return e.Element.SetProperty("scp-6r", value)
}

// scp-7l (float32)
//
// Sidechain preamp 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp7L(value float32) error {
	return e.Element.SetProperty("scp-7l", value)
}

// scp-7r (float32)
//
// Sidechain preamp 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScp7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scp-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScp7R(value float32) error {
	return e.Element.SetProperty("scp-7r", value)
}

// scr-0l (float32)
//
// Sidechain reactivity 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr0L(value float32) error {
	return e.Element.SetProperty("scr-0l", value)
}

// scr-0r (float32)
//
// Sidechain reactivity 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr0R(value float32) error {
	return e.Element.SetProperty("scr-0r", value)
}

// scr-1l (float32)
//
// Sidechain reactivity 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr1L(value float32) error {
	return e.Element.SetProperty("scr-1l", value)
}

// scr-1r (float32)
//
// Sidechain reactivity 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr1R(value float32) error {
	return e.Element.SetProperty("scr-1r", value)
}

// scr-2l (float32)
//
// Sidechain reactivity 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr2L(value float32) error {
	return e.Element.SetProperty("scr-2l", value)
}

// scr-2r (float32)
//
// Sidechain reactivity 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr2R(value float32) error {
	return e.Element.SetProperty("scr-2r", value)
}

// scr-3l (float32)
//
// Sidechain reactivity 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr3L(value float32) error {
	return e.Element.SetProperty("scr-3l", value)
}

// scr-3r (float32)
//
// Sidechain reactivity 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr3R(value float32) error {
	return e.Element.SetProperty("scr-3r", value)
}

// scr-4l (float32)
//
// Sidechain reactivity 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr4L(value float32) error {
	return e.Element.SetProperty("scr-4l", value)
}

// scr-4r (float32)
//
// Sidechain reactivity 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr4R(value float32) error {
	return e.Element.SetProperty("scr-4r", value)
}

// scr-5l (float32)
//
// Sidechain reactivity 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr5L(value float32) error {
	return e.Element.SetProperty("scr-5l", value)
}

// scr-5r (float32)
//
// Sidechain reactivity 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr5R(value float32) error {
	return e.Element.SetProperty("scr-5r", value)
}

// scr-6l (float32)
//
// Sidechain reactivity 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr6L(value float32) error {
	return e.Element.SetProperty("scr-6l", value)
}

// scr-6r (float32)
//
// Sidechain reactivity 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr6R(value float32) error {
	return e.Element.SetProperty("scr-6r", value)
}

// scr-7l (float32)
//
// Sidechain reactivity 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr7L(value float32) error {
	return e.Element.SetProperty("scr-7l", value)
}

// scr-7r (float32)
//
// Sidechain reactivity 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScr7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("scr-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScr7R(value float32) error {
	return e.Element.SetProperty("scr-7r", value)
}

// scs-0l (GstLspPlugInPluginsLv2MbCompressorLrscs0L)
//
// Sidechain source 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs0L() (interface{}, error) {
	return e.Element.GetProperty("scs-0l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs0L(value interface{}) error {
	return e.Element.SetProperty("scs-0l", value)
}

// scs-0r (GstLspPlugInPluginsLv2MbCompressorLrscs0R)
//
// Sidechain source 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs0R() (interface{}, error) {
	return e.Element.GetProperty("scs-0r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs0R(value interface{}) error {
	return e.Element.SetProperty("scs-0r", value)
}

// scs-1l (GstLspPlugInPluginsLv2MbCompressorLrscs1L)
//
// Sidechain source 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs1L() (interface{}, error) {
	return e.Element.GetProperty("scs-1l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs1L(value interface{}) error {
	return e.Element.SetProperty("scs-1l", value)
}

// scs-1r (GstLspPlugInPluginsLv2MbCompressorLrscs1R)
//
// Sidechain source 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs1R() (interface{}, error) {
	return e.Element.GetProperty("scs-1r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs1R(value interface{}) error {
	return e.Element.SetProperty("scs-1r", value)
}

// scs-2l (GstLspPlugInPluginsLv2MbCompressorLrscs2L)
//
// Sidechain source 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs2L() (interface{}, error) {
	return e.Element.GetProperty("scs-2l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs2L(value interface{}) error {
	return e.Element.SetProperty("scs-2l", value)
}

// scs-2r (GstLspPlugInPluginsLv2MbCompressorLrscs2R)
//
// Sidechain source 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs2R() (interface{}, error) {
	return e.Element.GetProperty("scs-2r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs2R(value interface{}) error {
	return e.Element.SetProperty("scs-2r", value)
}

// scs-3l (GstLspPlugInPluginsLv2MbCompressorLrscs3L)
//
// Sidechain source 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs3L() (interface{}, error) {
	return e.Element.GetProperty("scs-3l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs3L(value interface{}) error {
	return e.Element.SetProperty("scs-3l", value)
}

// scs-3r (GstLspPlugInPluginsLv2MbCompressorLrscs3R)
//
// Sidechain source 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs3R() (interface{}, error) {
	return e.Element.GetProperty("scs-3r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs3R(value interface{}) error {
	return e.Element.SetProperty("scs-3r", value)
}

// scs-4l (GstLspPlugInPluginsLv2MbCompressorLrscs4L)
//
// Sidechain source 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs4L() (interface{}, error) {
	return e.Element.GetProperty("scs-4l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs4L(value interface{}) error {
	return e.Element.SetProperty("scs-4l", value)
}

// scs-4r (GstLspPlugInPluginsLv2MbCompressorLrscs4R)
//
// Sidechain source 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs4R() (interface{}, error) {
	return e.Element.GetProperty("scs-4r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs4R(value interface{}) error {
	return e.Element.SetProperty("scs-4r", value)
}

// scs-5l (GstLspPlugInPluginsLv2MbCompressorLrscs5L)
//
// Sidechain source 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs5L() (interface{}, error) {
	return e.Element.GetProperty("scs-5l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs5L(value interface{}) error {
	return e.Element.SetProperty("scs-5l", value)
}

// scs-5r (GstLspPlugInPluginsLv2MbCompressorLrscs5R)
//
// Sidechain source 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs5R() (interface{}, error) {
	return e.Element.GetProperty("scs-5r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs5R(value interface{}) error {
	return e.Element.SetProperty("scs-5r", value)
}

// scs-6l (GstLspPlugInPluginsLv2MbCompressorLrscs6L)
//
// Sidechain source 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs6L() (interface{}, error) {
	return e.Element.GetProperty("scs-6l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs6L(value interface{}) error {
	return e.Element.SetProperty("scs-6l", value)
}

// scs-6r (GstLspPlugInPluginsLv2MbCompressorLrscs6R)
//
// Sidechain source 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs6R() (interface{}, error) {
	return e.Element.GetProperty("scs-6r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs6R(value interface{}) error {
	return e.Element.SetProperty("scs-6r", value)
}

// scs-7l (GstLspPlugInPluginsLv2MbCompressorLrscs7L)
//
// Sidechain source 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs7L() (interface{}, error) {
	return e.Element.GetProperty("scs-7l")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs7L(value interface{}) error {
	return e.Element.SetProperty("scs-7l", value)
}

// scs-7r (GstLspPlugInPluginsLv2MbCompressorLrscs7R)
//
// Sidechain source 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetScs7R() (interface{}, error) {
	return e.Element.GetProperty("scs-7r")
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetScs7R(value interface{}) error {
	return e.Element.SetProperty("scs-7r", value)
}

// sf-1l (float32)
//
// Split frequency 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf1L(value float32) error {
	return e.Element.SetProperty("sf-1l", value)
}

// sf-1r (float32)
//
// Split frequency 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf1R(value float32) error {
	return e.Element.SetProperty("sf-1r", value)
}

// sf-2l (float32)
//
// Split frequency 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf2L(value float32) error {
	return e.Element.SetProperty("sf-2l", value)
}

// sf-2r (float32)
//
// Split frequency 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf2R(value float32) error {
	return e.Element.SetProperty("sf-2r", value)
}

// sf-3l (float32)
//
// Split frequency 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf3L(value float32) error {
	return e.Element.SetProperty("sf-3l", value)
}

// sf-3r (float32)
//
// Split frequency 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf3R(value float32) error {
	return e.Element.SetProperty("sf-3r", value)
}

// sf-4l (float32)
//
// Split frequency 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf4L(value float32) error {
	return e.Element.SetProperty("sf-4l", value)
}

// sf-4r (float32)
//
// Split frequency 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf4R(value float32) error {
	return e.Element.SetProperty("sf-4r", value)
}

// sf-5l (float32)
//
// Split frequency 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf5L(value float32) error {
	return e.Element.SetProperty("sf-5l", value)
}

// sf-5r (float32)
//
// Split frequency 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf5R(value float32) error {
	return e.Element.SetProperty("sf-5r", value)
}

// sf-6l (float32)
//
// Split frequency 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf6L(value float32) error {
	return e.Element.SetProperty("sf-6l", value)
}

// sf-6r (float32)
//
// Split frequency 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf6R(value float32) error {
	return e.Element.SetProperty("sf-6r", value)
}

// sf-7l (float32)
//
// Split frequency 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf7L(value float32) error {
	return e.Element.SetProperty("sf-7l", value)
}

// sf-7r (float32)
//
// Split frequency 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSf7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sf-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSf7R(value float32) error {
	return e.Element.SetProperty("sf-7r", value)
}

// shift (float32)
//
// Shift gain

func (e *LspPlugInPluginsLv2MbCompressorLr) GetShift() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetShift(value float32) error {
	return e.Element.SetProperty("shift", value)
}

// sla-0l (float32)
//
// Sidechain lookahead 0 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla0L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-0l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla0L(value float32) error {
	return e.Element.SetProperty("sla-0l", value)
}

// sla-0r (float32)
//
// Sidechain lookahead 0 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla0R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-0r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla0R(value float32) error {
	return e.Element.SetProperty("sla-0r", value)
}

// sla-1l (float32)
//
// Sidechain lookahead 1 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla1L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-1l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla1L(value float32) error {
	return e.Element.SetProperty("sla-1l", value)
}

// sla-1r (float32)
//
// Sidechain lookahead 1 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla1R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-1r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla1R(value float32) error {
	return e.Element.SetProperty("sla-1r", value)
}

// sla-2l (float32)
//
// Sidechain lookahead 2 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla2L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-2l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla2L(value float32) error {
	return e.Element.SetProperty("sla-2l", value)
}

// sla-2r (float32)
//
// Sidechain lookahead 2 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla2R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-2r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla2R(value float32) error {
	return e.Element.SetProperty("sla-2r", value)
}

// sla-3l (float32)
//
// Sidechain lookahead 3 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla3L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-3l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla3L(value float32) error {
	return e.Element.SetProperty("sla-3l", value)
}

// sla-3r (float32)
//
// Sidechain lookahead 3 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla3R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-3r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla3R(value float32) error {
	return e.Element.SetProperty("sla-3r", value)
}

// sla-4l (float32)
//
// Sidechain lookahead 4 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla4L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-4l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla4L(value float32) error {
	return e.Element.SetProperty("sla-4l", value)
}

// sla-4r (float32)
//
// Sidechain lookahead 4 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla4R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-4r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla4R(value float32) error {
	return e.Element.SetProperty("sla-4r", value)
}

// sla-5l (float32)
//
// Sidechain lookahead 5 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla5L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-5l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla5L(value float32) error {
	return e.Element.SetProperty("sla-5l", value)
}

// sla-5r (float32)
//
// Sidechain lookahead 5 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla5R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-5r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla5R(value float32) error {
	return e.Element.SetProperty("sla-5r", value)
}

// sla-6l (float32)
//
// Sidechain lookahead 6 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla6L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-6l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla6L(value float32) error {
	return e.Element.SetProperty("sla-6l", value)
}

// sla-6r (float32)
//
// Sidechain lookahead 6 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla6R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-6r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla6R(value float32) error {
	return e.Element.SetProperty("sla-6r", value)
}

// sla-7l (float32)
//
// Sidechain lookahead 7 Left

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla7L() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-7l")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla7L(value float32) error {
	return e.Element.SetProperty("sla-7l", value)
}

// sla-7r (float32)
//
// Sidechain lookahead 7 Right

func (e *LspPlugInPluginsLv2MbCompressorLr) GetSla7R() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("sla-7r")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MbCompressorLr) SetSla7R(value float32) error {
	return e.Element.SetProperty("sla-7r", value)
}

// zoom (float32)
//
// Graph zoom

func (e *LspPlugInPluginsLv2MbCompressorLr) GetZoom() (float32, error) {
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

func (e *LspPlugInPluginsLv2MbCompressorLr) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MbCompressorLrscs6R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs6RMid LspPlugInPluginsLv2MbCompressorLrscs6R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs6RSide LspPlugInPluginsLv2MbCompressorLrscs6R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs6RLeft LspPlugInPluginsLv2MbCompressorLrscs6R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs6RRight LspPlugInPluginsLv2MbCompressorLrscs6R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrcm2R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm2RDown LspPlugInPluginsLv2MbCompressorLrcm2R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm2RUp LspPlugInPluginsLv2MbCompressorLrcm2R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm7R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm7RDown LspPlugInPluginsLv2MbCompressorLrcm7R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm7RUp LspPlugInPluginsLv2MbCompressorLrcm7R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrscm2L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm2LPeak LspPlugInPluginsLv2MbCompressorLrscm2L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm2LRms LspPlugInPluginsLv2MbCompressorLrscm2L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm2LLpf LspPlugInPluginsLv2MbCompressorLrscm2L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm2LUniform LspPlugInPluginsLv2MbCompressorLrscm2L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm4L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm4LPeak LspPlugInPluginsLv2MbCompressorLrscm4L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm4LRms LspPlugInPluginsLv2MbCompressorLrscm4L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm4LLpf LspPlugInPluginsLv2MbCompressorLrscm4L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm4LUniform LspPlugInPluginsLv2MbCompressorLrscm4L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm5R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm5RPeak LspPlugInPluginsLv2MbCompressorLrscm5R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm5RRms LspPlugInPluginsLv2MbCompressorLrscm5R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm5RLpf LspPlugInPluginsLv2MbCompressorLrscm5R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm5RUniform LspPlugInPluginsLv2MbCompressorLrscm5R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm7L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm7LPeak LspPlugInPluginsLv2MbCompressorLrscm7L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm7LRms LspPlugInPluginsLv2MbCompressorLrscm7L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm7LLpf LspPlugInPluginsLv2MbCompressorLrscm7L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm7LUniform LspPlugInPluginsLv2MbCompressorLrscm7L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs2R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs2RMid LspPlugInPluginsLv2MbCompressorLrscs2R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs2RSide LspPlugInPluginsLv2MbCompressorLrscs2R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs2RLeft LspPlugInPluginsLv2MbCompressorLrscs2R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs2RRight LspPlugInPluginsLv2MbCompressorLrscs2R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs7L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs7LMid LspPlugInPluginsLv2MbCompressorLrscs7L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs7LSide LspPlugInPluginsLv2MbCompressorLrscs7L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs7LLeft LspPlugInPluginsLv2MbCompressorLrscs7L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs7LRight LspPlugInPluginsLv2MbCompressorLrscs7L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrcm4L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm4LDown LspPlugInPluginsLv2MbCompressorLrcm4L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm4LUp LspPlugInPluginsLv2MbCompressorLrcm4L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrscm1L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm1LPeak LspPlugInPluginsLv2MbCompressorLrscm1L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm1LRms LspPlugInPluginsLv2MbCompressorLrscm1L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm1LLpf LspPlugInPluginsLv2MbCompressorLrscm1L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm1LUniform LspPlugInPluginsLv2MbCompressorLrscm1L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm5L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm5LPeak LspPlugInPluginsLv2MbCompressorLrscm5L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm5LRms LspPlugInPluginsLv2MbCompressorLrscm5L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm5LLpf LspPlugInPluginsLv2MbCompressorLrscm5L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm5LUniform LspPlugInPluginsLv2MbCompressorLrscm5L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs3L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs3LMid LspPlugInPluginsLv2MbCompressorLrscs3L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs3LSide LspPlugInPluginsLv2MbCompressorLrscs3L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs3LLeft LspPlugInPluginsLv2MbCompressorLrscs3L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs3LRight LspPlugInPluginsLv2MbCompressorLrscs3L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs7R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs7RMid LspPlugInPluginsLv2MbCompressorLrscs7R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs7RSide LspPlugInPluginsLv2MbCompressorLrscs7R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs7RLeft LspPlugInPluginsLv2MbCompressorLrscs7R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs7RRight LspPlugInPluginsLv2MbCompressorLrscs7R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscm3R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm3RPeak LspPlugInPluginsLv2MbCompressorLrscm3R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm3RRms LspPlugInPluginsLv2MbCompressorLrscm3R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm3RLpf LspPlugInPluginsLv2MbCompressorLrscm3R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm3RUniform LspPlugInPluginsLv2MbCompressorLrscm3R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm4R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm4RPeak LspPlugInPluginsLv2MbCompressorLrscm4R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm4RRms LspPlugInPluginsLv2MbCompressorLrscm4R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm4RLpf LspPlugInPluginsLv2MbCompressorLrscm4R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm4RUniform LspPlugInPluginsLv2MbCompressorLrscm4R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs2L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs2LMid LspPlugInPluginsLv2MbCompressorLrscs2L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs2LSide LspPlugInPluginsLv2MbCompressorLrscs2L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs2LLeft LspPlugInPluginsLv2MbCompressorLrscs2L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs2LRight LspPlugInPluginsLv2MbCompressorLrscs2L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs5R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs5RMid LspPlugInPluginsLv2MbCompressorLrscs5R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs5RSide LspPlugInPluginsLv2MbCompressorLrscs5R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs5RLeft LspPlugInPluginsLv2MbCompressorLrscs5R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs5RRight LspPlugInPluginsLv2MbCompressorLrscs5R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrbsel string

const (
	LspPlugInPluginsLv2MbCompressorLrbselSplitLeft LspPlugInPluginsLv2MbCompressorLrbsel = "Split Left" // Split Left (0) – Split Left
	LspPlugInPluginsLv2MbCompressorLrbselSplitRight LspPlugInPluginsLv2MbCompressorLrbsel = "Split Right" // Split Right (1) – Split Right
	LspPlugInPluginsLv2MbCompressorLrbselBand0 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 0" // Band 0 (2) – Band 0
	LspPlugInPluginsLv2MbCompressorLrbselBand1 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 1" // Band 1 (3) – Band 1
	LspPlugInPluginsLv2MbCompressorLrbselBand2 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 2" // Band 2 (4) – Band 2
	LspPlugInPluginsLv2MbCompressorLrbselBand3 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 3" // Band 3 (5) – Band 3
	LspPlugInPluginsLv2MbCompressorLrbselBand4 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 4" // Band 4 (6) – Band 4
	LspPlugInPluginsLv2MbCompressorLrbselBand5 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 5" // Band 5 (7) – Band 5
	LspPlugInPluginsLv2MbCompressorLrbselBand6 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 6" // Band 6 (8) – Band 6
	LspPlugInPluginsLv2MbCompressorLrbselBand7 LspPlugInPluginsLv2MbCompressorLrbsel = "Band 7" // Band 7 (9) – Band 7
)

type LspPlugInPluginsLv2MbCompressorLrcm4R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm4RDown LspPlugInPluginsLv2MbCompressorLrcm4R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm4RUp LspPlugInPluginsLv2MbCompressorLrcm4R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrscm3L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm3LPeak LspPlugInPluginsLv2MbCompressorLrscm3L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm3LRms LspPlugInPluginsLv2MbCompressorLrscm3L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm3LLpf LspPlugInPluginsLv2MbCompressorLrscm3L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm3LUniform LspPlugInPluginsLv2MbCompressorLrscm3L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm6L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm6LPeak LspPlugInPluginsLv2MbCompressorLrscm6L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm6LRms LspPlugInPluginsLv2MbCompressorLrscm6L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm6LLpf LspPlugInPluginsLv2MbCompressorLrscm6L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm6LUniform LspPlugInPluginsLv2MbCompressorLrscm6L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs4L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs4LMid LspPlugInPluginsLv2MbCompressorLrscs4L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs4LSide LspPlugInPluginsLv2MbCompressorLrscs4L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs4LLeft LspPlugInPluginsLv2MbCompressorLrscs4L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs4LRight LspPlugInPluginsLv2MbCompressorLrscs4L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrcm3R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm3RDown LspPlugInPluginsLv2MbCompressorLrcm3R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm3RUp LspPlugInPluginsLv2MbCompressorLrcm3R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm5R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm5RDown LspPlugInPluginsLv2MbCompressorLrcm5R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm5RUp LspPlugInPluginsLv2MbCompressorLrcm5R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrscm6R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm6RPeak LspPlugInPluginsLv2MbCompressorLrscm6R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm6RRms LspPlugInPluginsLv2MbCompressorLrscm6R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm6RLpf LspPlugInPluginsLv2MbCompressorLrscm6R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm6RUniform LspPlugInPluginsLv2MbCompressorLrscm6R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs3R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs3RMid LspPlugInPluginsLv2MbCompressorLrscs3R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs3RSide LspPlugInPluginsLv2MbCompressorLrscs3R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs3RLeft LspPlugInPluginsLv2MbCompressorLrscs3R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs3RRight LspPlugInPluginsLv2MbCompressorLrscs3R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrcm5L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm5LDown LspPlugInPluginsLv2MbCompressorLrcm5L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm5LUp LspPlugInPluginsLv2MbCompressorLrcm5L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm6R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm6RDown LspPlugInPluginsLv2MbCompressorLrcm6R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm6RUp LspPlugInPluginsLv2MbCompressorLrcm6R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrenvb string

const (
	LspPlugInPluginsLv2MbCompressorLrenvbNone LspPlugInPluginsLv2MbCompressorLrenvb = "None" // None (0) – None
	LspPlugInPluginsLv2MbCompressorLrenvbPinkBt LspPlugInPluginsLv2MbCompressorLrenvb = "Pink BT" // Pink BT (1) – Pink BT
	LspPlugInPluginsLv2MbCompressorLrenvbPinkMt LspPlugInPluginsLv2MbCompressorLrenvb = "Pink MT" // Pink MT (2) – Pink MT
	LspPlugInPluginsLv2MbCompressorLrenvbBrownBt LspPlugInPluginsLv2MbCompressorLrenvb = "Brown BT" // Brown BT (3) – Brown BT
	LspPlugInPluginsLv2MbCompressorLrenvbBrownMt LspPlugInPluginsLv2MbCompressorLrenvb = "Brown MT" // Brown MT (4) – Brown MT
)

type LspPlugInPluginsLv2MbCompressorLrscm2R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm2RPeak LspPlugInPluginsLv2MbCompressorLrscm2R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm2RRms LspPlugInPluginsLv2MbCompressorLrscm2R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm2RLpf LspPlugInPluginsLv2MbCompressorLrscm2R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm2RUniform LspPlugInPluginsLv2MbCompressorLrscm2R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs0R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs0RMid LspPlugInPluginsLv2MbCompressorLrscs0R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs0RSide LspPlugInPluginsLv2MbCompressorLrscs0R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs0RLeft LspPlugInPluginsLv2MbCompressorLrscs0R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs0RRight LspPlugInPluginsLv2MbCompressorLrscs0R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrcm0R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm0RDown LspPlugInPluginsLv2MbCompressorLrcm0R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm0RUp LspPlugInPluginsLv2MbCompressorLrcm0R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm1L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm1LDown LspPlugInPluginsLv2MbCompressorLrcm1L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm1LUp LspPlugInPluginsLv2MbCompressorLrcm1L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm1R string

const (
	LspPlugInPluginsLv2MbCompressorLrcm1RDown LspPlugInPluginsLv2MbCompressorLrcm1R = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm1RUp LspPlugInPluginsLv2MbCompressorLrcm1R = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm3L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm3LDown LspPlugInPluginsLv2MbCompressorLrcm3L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm3LUp LspPlugInPluginsLv2MbCompressorLrcm3L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm6L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm6LDown LspPlugInPluginsLv2MbCompressorLrcm6L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm6LUp LspPlugInPluginsLv2MbCompressorLrcm6L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrscm0R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm0RPeak LspPlugInPluginsLv2MbCompressorLrscm0R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm0RRms LspPlugInPluginsLv2MbCompressorLrscm0R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm0RLpf LspPlugInPluginsLv2MbCompressorLrscm0R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm0RUniform LspPlugInPluginsLv2MbCompressorLrscm0R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm1R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm1RPeak LspPlugInPluginsLv2MbCompressorLrscm1R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm1RRms LspPlugInPluginsLv2MbCompressorLrscm1R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm1RLpf LspPlugInPluginsLv2MbCompressorLrscm1R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm1RUniform LspPlugInPluginsLv2MbCompressorLrscm1R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs1L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs1LMid LspPlugInPluginsLv2MbCompressorLrscs1L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs1LSide LspPlugInPluginsLv2MbCompressorLrscs1L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs1LLeft LspPlugInPluginsLv2MbCompressorLrscs1L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs1LRight LspPlugInPluginsLv2MbCompressorLrscs1L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs1R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs1RMid LspPlugInPluginsLv2MbCompressorLrscs1R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs1RSide LspPlugInPluginsLv2MbCompressorLrscs1R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs1RLeft LspPlugInPluginsLv2MbCompressorLrscs1R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs1RRight LspPlugInPluginsLv2MbCompressorLrscs1R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs4R string

const (
	LspPlugInPluginsLv2MbCompressorLrscs4RMid LspPlugInPluginsLv2MbCompressorLrscs4R = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs4RSide LspPlugInPluginsLv2MbCompressorLrscs4R = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs4RLeft LspPlugInPluginsLv2MbCompressorLrscs4R = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs4RRight LspPlugInPluginsLv2MbCompressorLrscs4R = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs6L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs6LMid LspPlugInPluginsLv2MbCompressorLrscs6L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs6LSide LspPlugInPluginsLv2MbCompressorLrscs6L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs6LLeft LspPlugInPluginsLv2MbCompressorLrscs6L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs6LRight LspPlugInPluginsLv2MbCompressorLrscs6L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrscs5L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs5LMid LspPlugInPluginsLv2MbCompressorLrscs5L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs5LSide LspPlugInPluginsLv2MbCompressorLrscs5L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs5LLeft LspPlugInPluginsLv2MbCompressorLrscs5L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs5LRight LspPlugInPluginsLv2MbCompressorLrscs5L = "Right" // Right (3) – Right
)

type LspPlugInPluginsLv2MbCompressorLrcm0L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm0LDown LspPlugInPluginsLv2MbCompressorLrcm0L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm0LUp LspPlugInPluginsLv2MbCompressorLrcm0L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm2L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm2LDown LspPlugInPluginsLv2MbCompressorLrcm2L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm2LUp LspPlugInPluginsLv2MbCompressorLrcm2L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrcm7L string

const (
	LspPlugInPluginsLv2MbCompressorLrcm7LDown LspPlugInPluginsLv2MbCompressorLrcm7L = "Down" // Down (0) – Down
	LspPlugInPluginsLv2MbCompressorLrcm7LUp LspPlugInPluginsLv2MbCompressorLrcm7L = "Up" // Up (1) – Up
)

type LspPlugInPluginsLv2MbCompressorLrmode string

const (
	LspPlugInPluginsLv2MbCompressorLrmodeClassic LspPlugInPluginsLv2MbCompressorLrmode = "Classic" // Classic (0) – Classic
	LspPlugInPluginsLv2MbCompressorLrmodeModern LspPlugInPluginsLv2MbCompressorLrmode = "Modern" // Modern (1) – Modern
)

type LspPlugInPluginsLv2MbCompressorLrscm0L string

const (
	LspPlugInPluginsLv2MbCompressorLrscm0LPeak LspPlugInPluginsLv2MbCompressorLrscm0L = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm0LRms LspPlugInPluginsLv2MbCompressorLrscm0L = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm0LLpf LspPlugInPluginsLv2MbCompressorLrscm0L = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm0LUniform LspPlugInPluginsLv2MbCompressorLrscm0L = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscm7R string

const (
	LspPlugInPluginsLv2MbCompressorLrscm7RPeak LspPlugInPluginsLv2MbCompressorLrscm7R = "Peak" // Peak (0) – Peak
	LspPlugInPluginsLv2MbCompressorLrscm7RRms LspPlugInPluginsLv2MbCompressorLrscm7R = "RMS" // RMS (1) – RMS
	LspPlugInPluginsLv2MbCompressorLrscm7RLpf LspPlugInPluginsLv2MbCompressorLrscm7R = "LPF" // LPF (2) – LPF
	LspPlugInPluginsLv2MbCompressorLrscm7RUniform LspPlugInPluginsLv2MbCompressorLrscm7R = "Uniform" // Uniform (3) – Uniform
)

type LspPlugInPluginsLv2MbCompressorLrscs0L string

const (
	LspPlugInPluginsLv2MbCompressorLrscs0LMid LspPlugInPluginsLv2MbCompressorLrscs0L = "Mid" // Mid (0) – Mid
	LspPlugInPluginsLv2MbCompressorLrscs0LSide LspPlugInPluginsLv2MbCompressorLrscs0L = "Side" // Side (1) – Side
	LspPlugInPluginsLv2MbCompressorLrscs0LLeft LspPlugInPluginsLv2MbCompressorLrscs0L = "Left" // Left (2) – Left
	LspPlugInPluginsLv2MbCompressorLrscs0LRight LspPlugInPluginsLv2MbCompressorLrscs0L = "Right" // Right (3) – Right
)

