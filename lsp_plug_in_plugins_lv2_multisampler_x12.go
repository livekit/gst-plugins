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

type LspPlugInPluginsLv2MultisamplerX12 struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MultisamplerX12() (*LspPlugInPluginsLv2MultisamplerX12, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-multisampler-x12")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MultisamplerX12{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MultisamplerX12WithName(name string) (*LspPlugInPluginsLv2MultisamplerX12, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-multisampler-x12", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MultisamplerX12{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MultisamplerX12) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX12) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2MultisamplerX12) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX12) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// fout (float32)
//
// Note-off fadeout

func (e *LspPlugInPluginsLv2MultisamplerX12) GetFout() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("fout")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetFout(value float32) error {
	return e.Element.SetProperty("fout", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MultisamplerX12) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX12) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// iact-0 (bool)
//
// Instrument activity 0

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-1 (bool)
//
// Instrument activity 1

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-10 (bool)
//
// Instrument activity 10

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-11 (bool)
//
// Instrument activity 11

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-2 (bool)
//
// Instrument activity 2

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-3 (bool)
//
// Instrument activity 3

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-4 (bool)
//
// Instrument activity 4

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-5 (bool)
//
// Instrument activity 5

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-6 (bool)
//
// Instrument activity 6

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-7 (bool)
//
// Instrument activity 7

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-8 (bool)
//
// Instrument activity 8

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-9 (bool)
//
// Instrument activity 9

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIact9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// imix-0 (float32)
//
// Instrument mix gain 0

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix0(value float32) error {
	return e.Element.SetProperty("imix-0", value)
}

// imix-1 (float32)
//
// Instrument mix gain 1

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix1(value float32) error {
	return e.Element.SetProperty("imix-1", value)
}

// imix-10 (float32)
//
// Instrument mix gain 10

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix10(value float32) error {
	return e.Element.SetProperty("imix-10", value)
}

// imix-11 (float32)
//
// Instrument mix gain 11

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix11(value float32) error {
	return e.Element.SetProperty("imix-11", value)
}

// imix-2 (float32)
//
// Instrument mix gain 2

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix2(value float32) error {
	return e.Element.SetProperty("imix-2", value)
}

// imix-3 (float32)
//
// Instrument mix gain 3

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix3(value float32) error {
	return e.Element.SetProperty("imix-3", value)
}

// imix-4 (float32)
//
// Instrument mix gain 4

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix4(value float32) error {
	return e.Element.SetProperty("imix-4", value)
}

// imix-5 (float32)
//
// Instrument mix gain 5

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix5(value float32) error {
	return e.Element.SetProperty("imix-5", value)
}

// imix-6 (float32)
//
// Instrument mix gain 6

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix6(value float32) error {
	return e.Element.SetProperty("imix-6", value)
}

// imix-7 (float32)
//
// Instrument mix gain 7

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix7(value float32) error {
	return e.Element.SetProperty("imix-7", value)
}

// imix-8 (float32)
//
// Instrument mix gain 8

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix8(value float32) error {
	return e.Element.SetProperty("imix-8", value)
}

// imix-9 (float32)
//
// Instrument mix gain 9

func (e *LspPlugInPluginsLv2MultisamplerX12) GetImix9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetImix9(value float32) error {
	return e.Element.SetProperty("imix-9", value)
}

// ion-0 (bool)
//
// Instrument on 0

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon0() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon0(value bool) error {
	return e.Element.SetProperty("ion-0", value)
}

// ion-1 (bool)
//
// Instrument on 1

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon1() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon1(value bool) error {
	return e.Element.SetProperty("ion-1", value)
}

// ion-10 (bool)
//
// Instrument on 10

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon10() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon10(value bool) error {
	return e.Element.SetProperty("ion-10", value)
}

// ion-11 (bool)
//
// Instrument on 11

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon11() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon11(value bool) error {
	return e.Element.SetProperty("ion-11", value)
}

// ion-2 (bool)
//
// Instrument on 2

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon2(value bool) error {
	return e.Element.SetProperty("ion-2", value)
}

// ion-3 (bool)
//
// Instrument on 3

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon3() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon3(value bool) error {
	return e.Element.SetProperty("ion-3", value)
}

// ion-4 (bool)
//
// Instrument on 4

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon4() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon4(value bool) error {
	return e.Element.SetProperty("ion-4", value)
}

// ion-5 (bool)
//
// Instrument on 5

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon5() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon5(value bool) error {
	return e.Element.SetProperty("ion-5", value)
}

// ion-6 (bool)
//
// Instrument on 6

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon6() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon6(value bool) error {
	return e.Element.SetProperty("ion-6", value)
}

// ion-7 (bool)
//
// Instrument on 7

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon7() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon7(value bool) error {
	return e.Element.SetProperty("ion-7", value)
}

// ion-8 (bool)
//
// Instrument on 8

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon8() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon8(value bool) error {
	return e.Element.SetProperty("ion-8", value)
}

// ion-9 (bool)
//
// Instrument on 9

func (e *LspPlugInPluginsLv2MultisamplerX12) GetIon9() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetIon9(value bool) error {
	return e.Element.SetProperty("ion-9", value)
}

// msel (GstLspPlugInPluginsLv2MultisamplerX12Msel)
//
// Area selector

func (e *LspPlugInPluginsLv2MultisamplerX12) GetMsel() (interface{}, error) {
	return e.Element.GetProperty("msel")
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetMsel(value interface{}) error {
	return e.Element.SetProperty("msel", value)
}

// mute (bool)
//
// Forced mute

func (e *LspPlugInPluginsLv2MultisamplerX12) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// muting (bool)
//
// Mute on stop

func (e *LspPlugInPluginsLv2MultisamplerX12) GetMuting() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("muting")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetMuting(value bool) error {
	return e.Element.SetProperty("muting", value)
}

// noff (bool)
//
// Note-off handling

func (e *LspPlugInPluginsLv2MultisamplerX12) GetNoff() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("noff")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetNoff(value bool) error {
	return e.Element.SetProperty("noff", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2MultisamplerX12) GetOutLatency() (int, error) {
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

// panl-0 (float32)
//
// Instrument panorama left 0

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl0(value float32) error {
	return e.Element.SetProperty("panl-0", value)
}

// panl-1 (float32)
//
// Instrument panorama left 1

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl1(value float32) error {
	return e.Element.SetProperty("panl-1", value)
}

// panl-10 (float32)
//
// Instrument panorama left 10

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl10(value float32) error {
	return e.Element.SetProperty("panl-10", value)
}

// panl-11 (float32)
//
// Instrument panorama left 11

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl11(value float32) error {
	return e.Element.SetProperty("panl-11", value)
}

// panl-2 (float32)
//
// Instrument panorama left 2

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl2(value float32) error {
	return e.Element.SetProperty("panl-2", value)
}

// panl-3 (float32)
//
// Instrument panorama left 3

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl3(value float32) error {
	return e.Element.SetProperty("panl-3", value)
}

// panl-4 (float32)
//
// Instrument panorama left 4

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl4(value float32) error {
	return e.Element.SetProperty("panl-4", value)
}

// panl-5 (float32)
//
// Instrument panorama left 5

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl5(value float32) error {
	return e.Element.SetProperty("panl-5", value)
}

// panl-6 (float32)
//
// Instrument panorama left 6

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl6(value float32) error {
	return e.Element.SetProperty("panl-6", value)
}

// panl-7 (float32)
//
// Instrument panorama left 7

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl7(value float32) error {
	return e.Element.SetProperty("panl-7", value)
}

// panl-8 (float32)
//
// Instrument panorama left 8

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl8(value float32) error {
	return e.Element.SetProperty("panl-8", value)
}

// panl-9 (float32)
//
// Instrument panorama left 9

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanl9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanl9(value float32) error {
	return e.Element.SetProperty("panl-9", value)
}

// panr-0 (float32)
//
// Instrument manorama right 0

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr0() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-0")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr0(value float32) error {
	return e.Element.SetProperty("panr-0", value)
}

// panr-1 (float32)
//
// Instrument manorama right 1

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr1() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-1")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr1(value float32) error {
	return e.Element.SetProperty("panr-1", value)
}

// panr-10 (float32)
//
// Instrument manorama right 10

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr10() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-10")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr10(value float32) error {
	return e.Element.SetProperty("panr-10", value)
}

// panr-11 (float32)
//
// Instrument manorama right 11

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr11() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-11")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr11(value float32) error {
	return e.Element.SetProperty("panr-11", value)
}

// panr-2 (float32)
//
// Instrument manorama right 2

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr2() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-2")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr2(value float32) error {
	return e.Element.SetProperty("panr-2", value)
}

// panr-3 (float32)
//
// Instrument manorama right 3

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr3() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-3")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr3(value float32) error {
	return e.Element.SetProperty("panr-3", value)
}

// panr-4 (float32)
//
// Instrument manorama right 4

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr4() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-4")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr4(value float32) error {
	return e.Element.SetProperty("panr-4", value)
}

// panr-5 (float32)
//
// Instrument manorama right 5

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr5() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-5")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr5(value float32) error {
	return e.Element.SetProperty("panr-5", value)
}

// panr-6 (float32)
//
// Instrument manorama right 6

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr6() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-6")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr6(value float32) error {
	return e.Element.SetProperty("panr-6", value)
}

// panr-7 (float32)
//
// Instrument manorama right 7

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr7() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-7")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr7(value float32) error {
	return e.Element.SetProperty("panr-7", value)
}

// panr-8 (float32)
//
// Instrument manorama right 8

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr8() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-8")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr8(value float32) error {
	return e.Element.SetProperty("panr-8", value)
}

// panr-9 (float32)
//
// Instrument manorama right 9

func (e *LspPlugInPluginsLv2MultisamplerX12) GetPanr9() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-9")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX12) SetPanr9(value float32) error {
	return e.Element.SetProperty("panr-9", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2MultisamplerX12) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX12) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MultisamplerX12Msel string

const (
	LspPlugInPluginsLv2MultisamplerX12MselInstruments LspPlugInPluginsLv2MultisamplerX12Msel = "Instruments" // Instruments (0) – Instruments
	LspPlugInPluginsLv2MultisamplerX12MselMixer LspPlugInPluginsLv2MultisamplerX12Msel = "Mixer" // Mixer (1) – Mixer
)

