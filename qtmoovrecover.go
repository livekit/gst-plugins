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
)

type Qtmoovrecover struct {
	Element *gst.Element
}

func NewQtmoovrecover() (*Qtmoovrecover, error) {
	e, err := gst.NewElement("qtmoovrecover")
	if err != nil {
		return nil, err
	}
	return &Qtmoovrecover{Element: e}, nil
}

func NewQtmoovrecoverWithName(name string) (*Qtmoovrecover, error) {
	e, err := gst.NewElementWithName("qtmoovrecover", name)
	if err != nil {
		return nil, err
	}
	return &Qtmoovrecover{Element: e}, nil
}

// ----- Properties -----

// broken-input (string)
//
// Path to broken input file. (If qtmux was on faststart mode, this file is the faststart file)

func (e *Qtmoovrecover) GetBrokenInput() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("broken-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Qtmoovrecover) SetBrokenInput(value string) error {
	return e.Element.SetProperty("broken-input", value)
}

// faststart-mode (bool)
//
// If the broken input is from faststart mode

func (e *Qtmoovrecover) GetFaststartMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("faststart-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Qtmoovrecover) SetFaststartMode(value bool) error {
	return e.Element.SetProperty("faststart-mode", value)
}

// fixed-output (string)
//
// Path to write the fixed file to (used as output)

func (e *Qtmoovrecover) GetFixedOutput() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fixed-output")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Qtmoovrecover) SetFixedOutput(value string) error {
	return e.Element.SetProperty("fixed-output", value)
}

// recovery-input (string)
//
// Path to recovery file (used as input)

func (e *Qtmoovrecover) GetRecoveryInput() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("recovery-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Qtmoovrecover) SetRecoveryInput(value string) error {
	return e.Element.SetProperty("recovery-input", value)
}

