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

type Downloadbuffer struct {
	Element *gst.Element
}

func NewDownloadbuffer() (*Downloadbuffer, error) {
	e, err := gst.NewElement("downloadbuffer")
	if err != nil {
		return nil, err
	}
	return &Downloadbuffer{Element: e}, nil
}

func NewDownloadbufferWithName(name string) (*Downloadbuffer, error) {
	e, err := gst.NewElementWithName("downloadbuffer", name)
	if err != nil {
		return nil, err
	}
	return &Downloadbuffer{Element: e}, nil
}

// ----- Properties -----

// high-percent (int)
//
// High threshold for buffering to finish. Emits GST_MESSAGE_BUFFERING with a value of 100%%

func (e *Downloadbuffer) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Downloadbuffer) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

// low-percent (int)
//
// Low threshold for buffering to start. Emits GST_MESSAGE_BUFFERING with a value of 0%%

func (e *Downloadbuffer) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Downloadbuffer) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

// max-size-bytes (uint)
//
// Max. amount of data to buffer (bytes, 0=disable)

func (e *Downloadbuffer) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Downloadbuffer) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

// max-size-time (uint64)
//
// Max. amount of data to buffer (in ns, 0=disable)

func (e *Downloadbuffer) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Downloadbuffer) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

// temp-location (string)
//
// Location to store temporary files in (Only read this property, use temp-template to configure the name template)

func (e *Downloadbuffer) GetTempLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// temp-remove (bool)
//
// When temp-template is set, remove the temporary file when going to READY.

func (e *Downloadbuffer) GetTempRemove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temp-remove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Downloadbuffer) SetTempRemove(value bool) error {
	return e.Element.SetProperty("temp-remove", value)
}

// temp-template (string)
//
// File template to store temporary files in, should contain directory and XXXXXX. (NULL == disabled)

func (e *Downloadbuffer) GetTempTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Downloadbuffer) SetTempTemplate(value string) error {
	return e.Element.SetProperty("temp-template", value)
}

