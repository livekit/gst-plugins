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

type LspPlugInPluginsLv2MultisamplerX48 struct {
	*base.GstBaseTransform
}

func NewLspPlugInPluginsLv2MultisamplerX48() (*LspPlugInPluginsLv2MultisamplerX48, error) {
	e, err := gst.NewElement("lsp-plug-in-plugins-lv2-multisampler-x48")
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MultisamplerX48{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewLspPlugInPluginsLv2MultisamplerX48WithName(name string) (*LspPlugInPluginsLv2MultisamplerX48, error) {
	e, err := gst.NewElementWithName("lsp-plug-in-plugins-lv2-multisampler-x48", name)
	if err != nil {
		return nil, err
	}
	return &LspPlugInPluginsLv2MultisamplerX48{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// bypass (bool)
//
// Bypass

func (e *LspPlugInPluginsLv2MultisamplerX48) GetBypass() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetBypass(value bool) error {
	return e.Element.SetProperty("bypass", value)
}

// dry (float32)
//
// Dry amount

func (e *LspPlugInPluginsLv2MultisamplerX48) GetDry() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetDry(value float32) error {
	return e.Element.SetProperty("dry", value)
}

// fout (float32)
//
// Note-off fadeout

func (e *LspPlugInPluginsLv2MultisamplerX48) GetFout() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetFout(value float32) error {
	return e.Element.SetProperty("fout", value)
}

// g-out (float32)
//
// Output gain

func (e *LspPlugInPluginsLv2MultisamplerX48) GetGOut() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetGOut(value float32) error {
	return e.Element.SetProperty("g-out", value)
}

// iact-0 (bool)
//
// Instrument activity 0

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact10() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact11() (bool, error) {
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

// iact-12 (bool)
//
// Instrument activity 12

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-13 (bool)
//
// Instrument activity 13

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-14 (bool)
//
// Instrument activity 14

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-15 (bool)
//
// Instrument activity 15

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-16 (bool)
//
// Instrument activity 16

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-17 (bool)
//
// Instrument activity 17

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-18 (bool)
//
// Instrument activity 18

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-19 (bool)
//
// Instrument activity 19

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-19")
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact2() (bool, error) {
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

// iact-20 (bool)
//
// Instrument activity 20

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-21 (bool)
//
// Instrument activity 21

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-22 (bool)
//
// Instrument activity 22

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-23 (bool)
//
// Instrument activity 23

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-24 (bool)
//
// Instrument activity 24

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-25 (bool)
//
// Instrument activity 25

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-26 (bool)
//
// Instrument activity 26

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-27 (bool)
//
// Instrument activity 27

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-28 (bool)
//
// Instrument activity 28

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-29 (bool)
//
// Instrument activity 29

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-29")
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact3() (bool, error) {
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

// iact-30 (bool)
//
// Instrument activity 30

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-31 (bool)
//
// Instrument activity 31

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-32 (bool)
//
// Instrument activity 32

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact32() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-32")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-33 (bool)
//
// Instrument activity 33

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact33() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-33")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-34 (bool)
//
// Instrument activity 34

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact34() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-34")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-35 (bool)
//
// Instrument activity 35

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact35() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-35")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-36 (bool)
//
// Instrument activity 36

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact36() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-36")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-37 (bool)
//
// Instrument activity 37

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact37() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-37")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-38 (bool)
//
// Instrument activity 38

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact38() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-38")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-39 (bool)
//
// Instrument activity 39

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact39() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-39")
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact4() (bool, error) {
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

// iact-40 (bool)
//
// Instrument activity 40

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact40() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-40")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-41 (bool)
//
// Instrument activity 41

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact41() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-41")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-42 (bool)
//
// Instrument activity 42

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact42() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-42")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-43 (bool)
//
// Instrument activity 43

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact43() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-43")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-44 (bool)
//
// Instrument activity 44

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact44() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-44")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-45 (bool)
//
// Instrument activity 45

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact45() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-45")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-46 (bool)
//
// Instrument activity 46

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact46() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-46")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// iact-47 (bool)
//
// Instrument activity 47

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact47() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iact-47")
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact8() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIact9() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix0(value float32) error {
	return e.Element.SetProperty("imix-0", value)
}

// imix-1 (float32)
//
// Instrument mix gain 1

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix1(value float32) error {
	return e.Element.SetProperty("imix-1", value)
}

// imix-10 (float32)
//
// Instrument mix gain 10

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix10() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix10(value float32) error {
	return e.Element.SetProperty("imix-10", value)
}

// imix-11 (float32)
//
// Instrument mix gain 11

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix11() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix11(value float32) error {
	return e.Element.SetProperty("imix-11", value)
}

// imix-12 (float32)
//
// Instrument mix gain 12

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix12(value float32) error {
	return e.Element.SetProperty("imix-12", value)
}

// imix-13 (float32)
//
// Instrument mix gain 13

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix13(value float32) error {
	return e.Element.SetProperty("imix-13", value)
}

// imix-14 (float32)
//
// Instrument mix gain 14

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix14(value float32) error {
	return e.Element.SetProperty("imix-14", value)
}

// imix-15 (float32)
//
// Instrument mix gain 15

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix15(value float32) error {
	return e.Element.SetProperty("imix-15", value)
}

// imix-16 (float32)
//
// Instrument mix gain 16

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix16(value float32) error {
	return e.Element.SetProperty("imix-16", value)
}

// imix-17 (float32)
//
// Instrument mix gain 17

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix17(value float32) error {
	return e.Element.SetProperty("imix-17", value)
}

// imix-18 (float32)
//
// Instrument mix gain 18

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix18(value float32) error {
	return e.Element.SetProperty("imix-18", value)
}

// imix-19 (float32)
//
// Instrument mix gain 19

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix19(value float32) error {
	return e.Element.SetProperty("imix-19", value)
}

// imix-2 (float32)
//
// Instrument mix gain 2

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix2(value float32) error {
	return e.Element.SetProperty("imix-2", value)
}

// imix-20 (float32)
//
// Instrument mix gain 20

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix20(value float32) error {
	return e.Element.SetProperty("imix-20", value)
}

// imix-21 (float32)
//
// Instrument mix gain 21

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix21(value float32) error {
	return e.Element.SetProperty("imix-21", value)
}

// imix-22 (float32)
//
// Instrument mix gain 22

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix22(value float32) error {
	return e.Element.SetProperty("imix-22", value)
}

// imix-23 (float32)
//
// Instrument mix gain 23

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix23(value float32) error {
	return e.Element.SetProperty("imix-23", value)
}

// imix-24 (float32)
//
// Instrument mix gain 24

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix24(value float32) error {
	return e.Element.SetProperty("imix-24", value)
}

// imix-25 (float32)
//
// Instrument mix gain 25

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix25(value float32) error {
	return e.Element.SetProperty("imix-25", value)
}

// imix-26 (float32)
//
// Instrument mix gain 26

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix26(value float32) error {
	return e.Element.SetProperty("imix-26", value)
}

// imix-27 (float32)
//
// Instrument mix gain 27

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix27(value float32) error {
	return e.Element.SetProperty("imix-27", value)
}

// imix-28 (float32)
//
// Instrument mix gain 28

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix28(value float32) error {
	return e.Element.SetProperty("imix-28", value)
}

// imix-29 (float32)
//
// Instrument mix gain 29

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix29(value float32) error {
	return e.Element.SetProperty("imix-29", value)
}

// imix-3 (float32)
//
// Instrument mix gain 3

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix3(value float32) error {
	return e.Element.SetProperty("imix-3", value)
}

// imix-30 (float32)
//
// Instrument mix gain 30

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix30(value float32) error {
	return e.Element.SetProperty("imix-30", value)
}

// imix-31 (float32)
//
// Instrument mix gain 31

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix31(value float32) error {
	return e.Element.SetProperty("imix-31", value)
}

// imix-32 (float32)
//
// Instrument mix gain 32

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix32() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-32")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix32(value float32) error {
	return e.Element.SetProperty("imix-32", value)
}

// imix-33 (float32)
//
// Instrument mix gain 33

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix33() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-33")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix33(value float32) error {
	return e.Element.SetProperty("imix-33", value)
}

// imix-34 (float32)
//
// Instrument mix gain 34

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix34() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-34")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix34(value float32) error {
	return e.Element.SetProperty("imix-34", value)
}

// imix-35 (float32)
//
// Instrument mix gain 35

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix35() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-35")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix35(value float32) error {
	return e.Element.SetProperty("imix-35", value)
}

// imix-36 (float32)
//
// Instrument mix gain 36

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix36() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-36")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix36(value float32) error {
	return e.Element.SetProperty("imix-36", value)
}

// imix-37 (float32)
//
// Instrument mix gain 37

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix37() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-37")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix37(value float32) error {
	return e.Element.SetProperty("imix-37", value)
}

// imix-38 (float32)
//
// Instrument mix gain 38

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix38() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-38")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix38(value float32) error {
	return e.Element.SetProperty("imix-38", value)
}

// imix-39 (float32)
//
// Instrument mix gain 39

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix39() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-39")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix39(value float32) error {
	return e.Element.SetProperty("imix-39", value)
}

// imix-4 (float32)
//
// Instrument mix gain 4

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix4(value float32) error {
	return e.Element.SetProperty("imix-4", value)
}

// imix-40 (float32)
//
// Instrument mix gain 40

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix40() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-40")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix40(value float32) error {
	return e.Element.SetProperty("imix-40", value)
}

// imix-41 (float32)
//
// Instrument mix gain 41

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix41() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-41")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix41(value float32) error {
	return e.Element.SetProperty("imix-41", value)
}

// imix-42 (float32)
//
// Instrument mix gain 42

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix42() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-42")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix42(value float32) error {
	return e.Element.SetProperty("imix-42", value)
}

// imix-43 (float32)
//
// Instrument mix gain 43

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix43() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-43")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix43(value float32) error {
	return e.Element.SetProperty("imix-43", value)
}

// imix-44 (float32)
//
// Instrument mix gain 44

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix44() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-44")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix44(value float32) error {
	return e.Element.SetProperty("imix-44", value)
}

// imix-45 (float32)
//
// Instrument mix gain 45

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix45() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-45")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix45(value float32) error {
	return e.Element.SetProperty("imix-45", value)
}

// imix-46 (float32)
//
// Instrument mix gain 46

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix46() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-46")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix46(value float32) error {
	return e.Element.SetProperty("imix-46", value)
}

// imix-47 (float32)
//
// Instrument mix gain 47

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix47() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("imix-47")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix47(value float32) error {
	return e.Element.SetProperty("imix-47", value)
}

// imix-5 (float32)
//
// Instrument mix gain 5

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix5(value float32) error {
	return e.Element.SetProperty("imix-5", value)
}

// imix-6 (float32)
//
// Instrument mix gain 6

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix6(value float32) error {
	return e.Element.SetProperty("imix-6", value)
}

// imix-7 (float32)
//
// Instrument mix gain 7

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix7(value float32) error {
	return e.Element.SetProperty("imix-7", value)
}

// imix-8 (float32)
//
// Instrument mix gain 8

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix8() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix8(value float32) error {
	return e.Element.SetProperty("imix-8", value)
}

// imix-9 (float32)
//
// Instrument mix gain 9

func (e *LspPlugInPluginsLv2MultisamplerX48) GetImix9() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetImix9(value float32) error {
	return e.Element.SetProperty("imix-9", value)
}

// ion-0 (bool)
//
// Instrument on 0

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon0() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon0(value bool) error {
	return e.Element.SetProperty("ion-0", value)
}

// ion-1 (bool)
//
// Instrument on 1

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon1() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon1(value bool) error {
	return e.Element.SetProperty("ion-1", value)
}

// ion-10 (bool)
//
// Instrument on 10

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon10() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon10(value bool) error {
	return e.Element.SetProperty("ion-10", value)
}

// ion-11 (bool)
//
// Instrument on 11

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon11() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon11(value bool) error {
	return e.Element.SetProperty("ion-11", value)
}

// ion-12 (bool)
//
// Instrument on 12

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon12() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon12(value bool) error {
	return e.Element.SetProperty("ion-12", value)
}

// ion-13 (bool)
//
// Instrument on 13

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon13() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon13(value bool) error {
	return e.Element.SetProperty("ion-13", value)
}

// ion-14 (bool)
//
// Instrument on 14

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon14() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon14(value bool) error {
	return e.Element.SetProperty("ion-14", value)
}

// ion-15 (bool)
//
// Instrument on 15

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon15() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon15(value bool) error {
	return e.Element.SetProperty("ion-15", value)
}

// ion-16 (bool)
//
// Instrument on 16

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon16() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon16(value bool) error {
	return e.Element.SetProperty("ion-16", value)
}

// ion-17 (bool)
//
// Instrument on 17

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon17() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon17(value bool) error {
	return e.Element.SetProperty("ion-17", value)
}

// ion-18 (bool)
//
// Instrument on 18

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon18() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon18(value bool) error {
	return e.Element.SetProperty("ion-18", value)
}

// ion-19 (bool)
//
// Instrument on 19

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon19() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon19(value bool) error {
	return e.Element.SetProperty("ion-19", value)
}

// ion-2 (bool)
//
// Instrument on 2

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon2() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon2(value bool) error {
	return e.Element.SetProperty("ion-2", value)
}

// ion-20 (bool)
//
// Instrument on 20

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon20() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon20(value bool) error {
	return e.Element.SetProperty("ion-20", value)
}

// ion-21 (bool)
//
// Instrument on 21

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon21() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon21(value bool) error {
	return e.Element.SetProperty("ion-21", value)
}

// ion-22 (bool)
//
// Instrument on 22

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon22() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon22(value bool) error {
	return e.Element.SetProperty("ion-22", value)
}

// ion-23 (bool)
//
// Instrument on 23

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon23() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon23(value bool) error {
	return e.Element.SetProperty("ion-23", value)
}

// ion-24 (bool)
//
// Instrument on 24

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon24() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon24(value bool) error {
	return e.Element.SetProperty("ion-24", value)
}

// ion-25 (bool)
//
// Instrument on 25

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon25() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon25(value bool) error {
	return e.Element.SetProperty("ion-25", value)
}

// ion-26 (bool)
//
// Instrument on 26

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon26() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon26(value bool) error {
	return e.Element.SetProperty("ion-26", value)
}

// ion-27 (bool)
//
// Instrument on 27

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon27() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon27(value bool) error {
	return e.Element.SetProperty("ion-27", value)
}

// ion-28 (bool)
//
// Instrument on 28

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon28() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon28(value bool) error {
	return e.Element.SetProperty("ion-28", value)
}

// ion-29 (bool)
//
// Instrument on 29

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon29() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon29(value bool) error {
	return e.Element.SetProperty("ion-29", value)
}

// ion-3 (bool)
//
// Instrument on 3

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon3() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon3(value bool) error {
	return e.Element.SetProperty("ion-3", value)
}

// ion-30 (bool)
//
// Instrument on 30

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon30() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon30(value bool) error {
	return e.Element.SetProperty("ion-30", value)
}

// ion-31 (bool)
//
// Instrument on 31

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon31() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon31(value bool) error {
	return e.Element.SetProperty("ion-31", value)
}

// ion-32 (bool)
//
// Instrument on 32

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon32() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-32")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon32(value bool) error {
	return e.Element.SetProperty("ion-32", value)
}

// ion-33 (bool)
//
// Instrument on 33

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon33() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-33")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon33(value bool) error {
	return e.Element.SetProperty("ion-33", value)
}

// ion-34 (bool)
//
// Instrument on 34

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon34() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-34")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon34(value bool) error {
	return e.Element.SetProperty("ion-34", value)
}

// ion-35 (bool)
//
// Instrument on 35

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon35() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-35")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon35(value bool) error {
	return e.Element.SetProperty("ion-35", value)
}

// ion-36 (bool)
//
// Instrument on 36

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon36() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-36")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon36(value bool) error {
	return e.Element.SetProperty("ion-36", value)
}

// ion-37 (bool)
//
// Instrument on 37

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon37() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-37")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon37(value bool) error {
	return e.Element.SetProperty("ion-37", value)
}

// ion-38 (bool)
//
// Instrument on 38

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon38() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-38")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon38(value bool) error {
	return e.Element.SetProperty("ion-38", value)
}

// ion-39 (bool)
//
// Instrument on 39

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon39() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-39")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon39(value bool) error {
	return e.Element.SetProperty("ion-39", value)
}

// ion-4 (bool)
//
// Instrument on 4

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon4() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon4(value bool) error {
	return e.Element.SetProperty("ion-4", value)
}

// ion-40 (bool)
//
// Instrument on 40

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon40() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-40")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon40(value bool) error {
	return e.Element.SetProperty("ion-40", value)
}

// ion-41 (bool)
//
// Instrument on 41

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon41() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-41")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon41(value bool) error {
	return e.Element.SetProperty("ion-41", value)
}

// ion-42 (bool)
//
// Instrument on 42

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon42() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-42")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon42(value bool) error {
	return e.Element.SetProperty("ion-42", value)
}

// ion-43 (bool)
//
// Instrument on 43

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon43() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-43")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon43(value bool) error {
	return e.Element.SetProperty("ion-43", value)
}

// ion-44 (bool)
//
// Instrument on 44

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon44() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-44")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon44(value bool) error {
	return e.Element.SetProperty("ion-44", value)
}

// ion-45 (bool)
//
// Instrument on 45

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon45() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-45")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon45(value bool) error {
	return e.Element.SetProperty("ion-45", value)
}

// ion-46 (bool)
//
// Instrument on 46

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon46() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-46")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon46(value bool) error {
	return e.Element.SetProperty("ion-46", value)
}

// ion-47 (bool)
//
// Instrument on 47

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon47() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ion-47")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon47(value bool) error {
	return e.Element.SetProperty("ion-47", value)
}

// ion-5 (bool)
//
// Instrument on 5

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon5() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon5(value bool) error {
	return e.Element.SetProperty("ion-5", value)
}

// ion-6 (bool)
//
// Instrument on 6

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon6() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon6(value bool) error {
	return e.Element.SetProperty("ion-6", value)
}

// ion-7 (bool)
//
// Instrument on 7

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon7() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon7(value bool) error {
	return e.Element.SetProperty("ion-7", value)
}

// ion-8 (bool)
//
// Instrument on 8

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon8() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon8(value bool) error {
	return e.Element.SetProperty("ion-8", value)
}

// ion-9 (bool)
//
// Instrument on 9

func (e *LspPlugInPluginsLv2MultisamplerX48) GetIon9() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetIon9(value bool) error {
	return e.Element.SetProperty("ion-9", value)
}

// msel (GstLspPlugInPluginsLv2MultisamplerX48Msel)
//
// Area selector

func (e *LspPlugInPluginsLv2MultisamplerX48) GetMsel() (interface{}, error) {
	return e.Element.GetProperty("msel")
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetMsel(value interface{}) error {
	return e.Element.SetProperty("msel", value)
}

// mute (bool)
//
// Forced mute

func (e *LspPlugInPluginsLv2MultisamplerX48) GetMute() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// muting (bool)
//
// Mute on stop

func (e *LspPlugInPluginsLv2MultisamplerX48) GetMuting() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetMuting(value bool) error {
	return e.Element.SetProperty("muting", value)
}

// noff (bool)
//
// Note-off handling

func (e *LspPlugInPluginsLv2MultisamplerX48) GetNoff() (bool, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetNoff(value bool) error {
	return e.Element.SetProperty("noff", value)
}

// out-latency (int)
//
// Latency OUT

func (e *LspPlugInPluginsLv2MultisamplerX48) GetOutLatency() (int, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl0(value float32) error {
	return e.Element.SetProperty("panl-0", value)
}

// panl-1 (float32)
//
// Instrument panorama left 1

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl1(value float32) error {
	return e.Element.SetProperty("panl-1", value)
}

// panl-10 (float32)
//
// Instrument panorama left 10

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl10() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl10(value float32) error {
	return e.Element.SetProperty("panl-10", value)
}

// panl-11 (float32)
//
// Instrument panorama left 11

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl11() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl11(value float32) error {
	return e.Element.SetProperty("panl-11", value)
}

// panl-12 (float32)
//
// Instrument panorama left 12

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl12(value float32) error {
	return e.Element.SetProperty("panl-12", value)
}

// panl-13 (float32)
//
// Instrument panorama left 13

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl13(value float32) error {
	return e.Element.SetProperty("panl-13", value)
}

// panl-14 (float32)
//
// Instrument panorama left 14

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl14(value float32) error {
	return e.Element.SetProperty("panl-14", value)
}

// panl-15 (float32)
//
// Instrument panorama left 15

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl15(value float32) error {
	return e.Element.SetProperty("panl-15", value)
}

// panl-16 (float32)
//
// Instrument panorama left 16

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl16(value float32) error {
	return e.Element.SetProperty("panl-16", value)
}

// panl-17 (float32)
//
// Instrument panorama left 17

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl17(value float32) error {
	return e.Element.SetProperty("panl-17", value)
}

// panl-18 (float32)
//
// Instrument panorama left 18

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl18(value float32) error {
	return e.Element.SetProperty("panl-18", value)
}

// panl-19 (float32)
//
// Instrument panorama left 19

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl19(value float32) error {
	return e.Element.SetProperty("panl-19", value)
}

// panl-2 (float32)
//
// Instrument panorama left 2

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl2(value float32) error {
	return e.Element.SetProperty("panl-2", value)
}

// panl-20 (float32)
//
// Instrument panorama left 20

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl20(value float32) error {
	return e.Element.SetProperty("panl-20", value)
}

// panl-21 (float32)
//
// Instrument panorama left 21

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl21(value float32) error {
	return e.Element.SetProperty("panl-21", value)
}

// panl-22 (float32)
//
// Instrument panorama left 22

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl22(value float32) error {
	return e.Element.SetProperty("panl-22", value)
}

// panl-23 (float32)
//
// Instrument panorama left 23

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl23(value float32) error {
	return e.Element.SetProperty("panl-23", value)
}

// panl-24 (float32)
//
// Instrument panorama left 24

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl24(value float32) error {
	return e.Element.SetProperty("panl-24", value)
}

// panl-25 (float32)
//
// Instrument panorama left 25

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl25(value float32) error {
	return e.Element.SetProperty("panl-25", value)
}

// panl-26 (float32)
//
// Instrument panorama left 26

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl26(value float32) error {
	return e.Element.SetProperty("panl-26", value)
}

// panl-27 (float32)
//
// Instrument panorama left 27

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl27(value float32) error {
	return e.Element.SetProperty("panl-27", value)
}

// panl-28 (float32)
//
// Instrument panorama left 28

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl28(value float32) error {
	return e.Element.SetProperty("panl-28", value)
}

// panl-29 (float32)
//
// Instrument panorama left 29

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl29(value float32) error {
	return e.Element.SetProperty("panl-29", value)
}

// panl-3 (float32)
//
// Instrument panorama left 3

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl3(value float32) error {
	return e.Element.SetProperty("panl-3", value)
}

// panl-30 (float32)
//
// Instrument panorama left 30

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl30(value float32) error {
	return e.Element.SetProperty("panl-30", value)
}

// panl-31 (float32)
//
// Instrument panorama left 31

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl31(value float32) error {
	return e.Element.SetProperty("panl-31", value)
}

// panl-32 (float32)
//
// Instrument panorama left 32

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl32() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-32")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl32(value float32) error {
	return e.Element.SetProperty("panl-32", value)
}

// panl-33 (float32)
//
// Instrument panorama left 33

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl33() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-33")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl33(value float32) error {
	return e.Element.SetProperty("panl-33", value)
}

// panl-34 (float32)
//
// Instrument panorama left 34

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl34() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-34")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl34(value float32) error {
	return e.Element.SetProperty("panl-34", value)
}

// panl-35 (float32)
//
// Instrument panorama left 35

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl35() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-35")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl35(value float32) error {
	return e.Element.SetProperty("panl-35", value)
}

// panl-36 (float32)
//
// Instrument panorama left 36

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl36() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-36")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl36(value float32) error {
	return e.Element.SetProperty("panl-36", value)
}

// panl-37 (float32)
//
// Instrument panorama left 37

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl37() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-37")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl37(value float32) error {
	return e.Element.SetProperty("panl-37", value)
}

// panl-38 (float32)
//
// Instrument panorama left 38

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl38() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-38")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl38(value float32) error {
	return e.Element.SetProperty("panl-38", value)
}

// panl-39 (float32)
//
// Instrument panorama left 39

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl39() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-39")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl39(value float32) error {
	return e.Element.SetProperty("panl-39", value)
}

// panl-4 (float32)
//
// Instrument panorama left 4

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl4(value float32) error {
	return e.Element.SetProperty("panl-4", value)
}

// panl-40 (float32)
//
// Instrument panorama left 40

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl40() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-40")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl40(value float32) error {
	return e.Element.SetProperty("panl-40", value)
}

// panl-41 (float32)
//
// Instrument panorama left 41

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl41() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-41")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl41(value float32) error {
	return e.Element.SetProperty("panl-41", value)
}

// panl-42 (float32)
//
// Instrument panorama left 42

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl42() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-42")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl42(value float32) error {
	return e.Element.SetProperty("panl-42", value)
}

// panl-43 (float32)
//
// Instrument panorama left 43

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl43() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-43")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl43(value float32) error {
	return e.Element.SetProperty("panl-43", value)
}

// panl-44 (float32)
//
// Instrument panorama left 44

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl44() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-44")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl44(value float32) error {
	return e.Element.SetProperty("panl-44", value)
}

// panl-45 (float32)
//
// Instrument panorama left 45

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl45() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-45")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl45(value float32) error {
	return e.Element.SetProperty("panl-45", value)
}

// panl-46 (float32)
//
// Instrument panorama left 46

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl46() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-46")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl46(value float32) error {
	return e.Element.SetProperty("panl-46", value)
}

// panl-47 (float32)
//
// Instrument panorama left 47

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl47() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panl-47")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl47(value float32) error {
	return e.Element.SetProperty("panl-47", value)
}

// panl-5 (float32)
//
// Instrument panorama left 5

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl5(value float32) error {
	return e.Element.SetProperty("panl-5", value)
}

// panl-6 (float32)
//
// Instrument panorama left 6

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl6(value float32) error {
	return e.Element.SetProperty("panl-6", value)
}

// panl-7 (float32)
//
// Instrument panorama left 7

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl7(value float32) error {
	return e.Element.SetProperty("panl-7", value)
}

// panl-8 (float32)
//
// Instrument panorama left 8

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl8() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl8(value float32) error {
	return e.Element.SetProperty("panl-8", value)
}

// panl-9 (float32)
//
// Instrument panorama left 9

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanl9() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanl9(value float32) error {
	return e.Element.SetProperty("panl-9", value)
}

// panr-0 (float32)
//
// Instrument manorama right 0

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr0() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr0(value float32) error {
	return e.Element.SetProperty("panr-0", value)
}

// panr-1 (float32)
//
// Instrument manorama right 1

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr1() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr1(value float32) error {
	return e.Element.SetProperty("panr-1", value)
}

// panr-10 (float32)
//
// Instrument manorama right 10

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr10() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr10(value float32) error {
	return e.Element.SetProperty("panr-10", value)
}

// panr-11 (float32)
//
// Instrument manorama right 11

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr11() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr11(value float32) error {
	return e.Element.SetProperty("panr-11", value)
}

// panr-12 (float32)
//
// Instrument manorama right 12

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr12() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-12")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr12(value float32) error {
	return e.Element.SetProperty("panr-12", value)
}

// panr-13 (float32)
//
// Instrument manorama right 13

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr13() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-13")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr13(value float32) error {
	return e.Element.SetProperty("panr-13", value)
}

// panr-14 (float32)
//
// Instrument manorama right 14

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr14() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-14")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr14(value float32) error {
	return e.Element.SetProperty("panr-14", value)
}

// panr-15 (float32)
//
// Instrument manorama right 15

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr15() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-15")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr15(value float32) error {
	return e.Element.SetProperty("panr-15", value)
}

// panr-16 (float32)
//
// Instrument manorama right 16

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr16() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-16")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr16(value float32) error {
	return e.Element.SetProperty("panr-16", value)
}

// panr-17 (float32)
//
// Instrument manorama right 17

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr17() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-17")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr17(value float32) error {
	return e.Element.SetProperty("panr-17", value)
}

// panr-18 (float32)
//
// Instrument manorama right 18

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr18() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-18")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr18(value float32) error {
	return e.Element.SetProperty("panr-18", value)
}

// panr-19 (float32)
//
// Instrument manorama right 19

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr19() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-19")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr19(value float32) error {
	return e.Element.SetProperty("panr-19", value)
}

// panr-2 (float32)
//
// Instrument manorama right 2

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr2() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr2(value float32) error {
	return e.Element.SetProperty("panr-2", value)
}

// panr-20 (float32)
//
// Instrument manorama right 20

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr20() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-20")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr20(value float32) error {
	return e.Element.SetProperty("panr-20", value)
}

// panr-21 (float32)
//
// Instrument manorama right 21

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr21() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-21")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr21(value float32) error {
	return e.Element.SetProperty("panr-21", value)
}

// panr-22 (float32)
//
// Instrument manorama right 22

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr22() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-22")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr22(value float32) error {
	return e.Element.SetProperty("panr-22", value)
}

// panr-23 (float32)
//
// Instrument manorama right 23

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr23() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-23")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr23(value float32) error {
	return e.Element.SetProperty("panr-23", value)
}

// panr-24 (float32)
//
// Instrument manorama right 24

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr24() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-24")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr24(value float32) error {
	return e.Element.SetProperty("panr-24", value)
}

// panr-25 (float32)
//
// Instrument manorama right 25

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr25() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-25")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr25(value float32) error {
	return e.Element.SetProperty("panr-25", value)
}

// panr-26 (float32)
//
// Instrument manorama right 26

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr26() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-26")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr26(value float32) error {
	return e.Element.SetProperty("panr-26", value)
}

// panr-27 (float32)
//
// Instrument manorama right 27

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr27() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-27")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr27(value float32) error {
	return e.Element.SetProperty("panr-27", value)
}

// panr-28 (float32)
//
// Instrument manorama right 28

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr28() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-28")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr28(value float32) error {
	return e.Element.SetProperty("panr-28", value)
}

// panr-29 (float32)
//
// Instrument manorama right 29

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr29() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-29")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr29(value float32) error {
	return e.Element.SetProperty("panr-29", value)
}

// panr-3 (float32)
//
// Instrument manorama right 3

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr3() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr3(value float32) error {
	return e.Element.SetProperty("panr-3", value)
}

// panr-30 (float32)
//
// Instrument manorama right 30

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr30() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-30")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr30(value float32) error {
	return e.Element.SetProperty("panr-30", value)
}

// panr-31 (float32)
//
// Instrument manorama right 31

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr31() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-31")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr31(value float32) error {
	return e.Element.SetProperty("panr-31", value)
}

// panr-32 (float32)
//
// Instrument manorama right 32

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr32() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-32")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr32(value float32) error {
	return e.Element.SetProperty("panr-32", value)
}

// panr-33 (float32)
//
// Instrument manorama right 33

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr33() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-33")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr33(value float32) error {
	return e.Element.SetProperty("panr-33", value)
}

// panr-34 (float32)
//
// Instrument manorama right 34

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr34() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-34")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr34(value float32) error {
	return e.Element.SetProperty("panr-34", value)
}

// panr-35 (float32)
//
// Instrument manorama right 35

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr35() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-35")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr35(value float32) error {
	return e.Element.SetProperty("panr-35", value)
}

// panr-36 (float32)
//
// Instrument manorama right 36

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr36() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-36")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr36(value float32) error {
	return e.Element.SetProperty("panr-36", value)
}

// panr-37 (float32)
//
// Instrument manorama right 37

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr37() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-37")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr37(value float32) error {
	return e.Element.SetProperty("panr-37", value)
}

// panr-38 (float32)
//
// Instrument manorama right 38

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr38() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-38")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr38(value float32) error {
	return e.Element.SetProperty("panr-38", value)
}

// panr-39 (float32)
//
// Instrument manorama right 39

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr39() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-39")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr39(value float32) error {
	return e.Element.SetProperty("panr-39", value)
}

// panr-4 (float32)
//
// Instrument manorama right 4

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr4() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr4(value float32) error {
	return e.Element.SetProperty("panr-4", value)
}

// panr-40 (float32)
//
// Instrument manorama right 40

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr40() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-40")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr40(value float32) error {
	return e.Element.SetProperty("panr-40", value)
}

// panr-41 (float32)
//
// Instrument manorama right 41

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr41() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-41")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr41(value float32) error {
	return e.Element.SetProperty("panr-41", value)
}

// panr-42 (float32)
//
// Instrument manorama right 42

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr42() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-42")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr42(value float32) error {
	return e.Element.SetProperty("panr-42", value)
}

// panr-43 (float32)
//
// Instrument manorama right 43

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr43() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-43")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr43(value float32) error {
	return e.Element.SetProperty("panr-43", value)
}

// panr-44 (float32)
//
// Instrument manorama right 44

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr44() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-44")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr44(value float32) error {
	return e.Element.SetProperty("panr-44", value)
}

// panr-45 (float32)
//
// Instrument manorama right 45

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr45() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-45")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr45(value float32) error {
	return e.Element.SetProperty("panr-45", value)
}

// panr-46 (float32)
//
// Instrument manorama right 46

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr46() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-46")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr46(value float32) error {
	return e.Element.SetProperty("panr-46", value)
}

// panr-47 (float32)
//
// Instrument manorama right 47

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr47() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("panr-47")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr47(value float32) error {
	return e.Element.SetProperty("panr-47", value)
}

// panr-5 (float32)
//
// Instrument manorama right 5

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr5() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr5(value float32) error {
	return e.Element.SetProperty("panr-5", value)
}

// panr-6 (float32)
//
// Instrument manorama right 6

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr6() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr6(value float32) error {
	return e.Element.SetProperty("panr-6", value)
}

// panr-7 (float32)
//
// Instrument manorama right 7

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr7() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr7(value float32) error {
	return e.Element.SetProperty("panr-7", value)
}

// panr-8 (float32)
//
// Instrument manorama right 8

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr8() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr8(value float32) error {
	return e.Element.SetProperty("panr-8", value)
}

// panr-9 (float32)
//
// Instrument manorama right 9

func (e *LspPlugInPluginsLv2MultisamplerX48) GetPanr9() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetPanr9(value float32) error {
	return e.Element.SetProperty("panr-9", value)
}

// wet (float32)
//
// Wet amount

func (e *LspPlugInPluginsLv2MultisamplerX48) GetWet() (float32, error) {
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

func (e *LspPlugInPluginsLv2MultisamplerX48) SetWet(value float32) error {
	return e.Element.SetProperty("wet", value)
}

// ----- Constants -----

type LspPlugInPluginsLv2MultisamplerX48Msel string

const (
	LspPlugInPluginsLv2MultisamplerX48MselInstruments LspPlugInPluginsLv2MultisamplerX48Msel = "Instruments" // Instruments (0) – Instruments
	LspPlugInPluginsLv2MultisamplerX48MselMixer011 LspPlugInPluginsLv2MultisamplerX48Msel = "Mixer 0-11" // Mixer 0-11 (1) – Mixer 0-11
	LspPlugInPluginsLv2MultisamplerX48MselMixer1223 LspPlugInPluginsLv2MultisamplerX48Msel = "Mixer 12-23" // Mixer 12-23 (2) – Mixer 12-23
	LspPlugInPluginsLv2MultisamplerX48MselMixer2435 LspPlugInPluginsLv2MultisamplerX48Msel = "Mixer 24-35" // Mixer 24-35 (3) – Mixer 24-35
	LspPlugInPluginsLv2MultisamplerX48MselMixer3647 LspPlugInPluginsLv2MultisamplerX48Msel = "Mixer 36-47" // Mixer 36-47 (4) – Mixer 36-47
)

